package wgpu

// #include <stdlib.h>
// #include "./wgpu.h"
//
// void logCallback(WGPULogLevel level, char * mnessage, void * userdata);
import "C"
import (
	"log"
	"strconv"
	"strings"
	"structs"
	"unsafe"

	"honnef.co/go/safeish"
)

//go:generate stringer -type LogLevel,InstanceBackend,InstanceFlag,DX12Compiler,GLES3MinorVersion -output wgpu_string.go -linecomment
type LogLevel uint32
type InstanceBackend uint32
type InstanceBackendFlags = InstanceBackend
type InstanceFlag uint32
type InstanceFlags = InstanceFlag
type DX12Compiler uint32
type GLES3MinorVersion uint32

const (
	LogLevelOff   LogLevel = C.WGPULogLevel_Off   // Off
	LogLevelError LogLevel = C.WGPULogLevel_Error // Error
	LogLevelWarn  LogLevel = C.WGPULogLevel_Warn  // Warn
	LogLevelInfo  LogLevel = C.WGPULogLevel_Info  // Info
	LogLevelDebug LogLevel = C.WGPULogLevel_Debug // Debug
	LogLevelTrace LogLevel = C.WGPULogLevel_Trace // Trace
)

const (
	InstanceBackendAll           InstanceBackend = C.WGPUInstanceBackend_All           // All
	InstanceBackendVulkan        InstanceBackend = C.WGPUInstanceBackend_Vulkan        // Vulkan
	InstanceBackendGL            InstanceBackend = C.WGPUInstanceBackend_GL            // GL
	InstanceBackendMetal         InstanceBackend = C.WGPUInstanceBackend_Metal         // Metal
	InstanceBackendDX12          InstanceBackend = C.WGPUInstanceBackend_DX12          // DX12
	InstanceBackendDX11          InstanceBackend = C.WGPUInstanceBackend_DX11          // DX11
	InstanceBackendBrowserWebGPU InstanceBackend = C.WGPUInstanceBackend_BrowserWebGPU // BrowserWebGPU
	InstanceBackendPrimary       InstanceBackend = C.WGPUInstanceBackend_Primary       // Primary
	InstanceBackendSecondary     InstanceBackend = C.WGPUInstanceBackend_Secondary     // Secondary
)

const (
	InstanceFlagDefault          InstanceFlag = C.WGPUInstanceFlag_Default          // Default
	InstanceFlagDebug            InstanceFlag = C.WGPUInstanceFlag_Debug            // Debug
	InstanceFlagValidation       InstanceFlag = C.WGPUInstanceFlag_Validation       // Validation
	InstanceFlagDiscardHalLabels InstanceFlag = C.WGPUInstanceFlag_DiscardHalLabels // DiscardHalLabels
)

const (
	DX12CompilerUndefined DX12Compiler = C.WGPUDx12Compiler_Undefined // Undefined
	DX12CompilerFxc       DX12Compiler = C.WGPUDx12Compiler_Fxc       // Fxc
	DX12CompilerDxc       DX12Compiler = C.WGPUDx12Compiler_Dxc       // Dxc
)

const (
	GLES3MinorVersionAutomatic GLES3MinorVersion = C.WGPUGles3MinorVersion_Automatic // Automatic
	GLES3MinorVersionVersion0  GLES3MinorVersion = C.WGPUGles3MinorVersion_Version0  // Version0
	GLES3MinorVersionVersion1  GLES3MinorVersion = C.WGPUGles3MinorVersion_Version1  // Version1
	GLES3MinorVersionVersion2  GLES3MinorVersion = C.WGPUGles3MinorVersion_Version2  // Version2
)

const (
	NativeFeatureNamePushConstants                                         FeatureName = C.WGPUNativeFeature_PushConstants                                         // PushConstants
	NativeFeatureNameTextureAdapterSpecificFormatFeatures                  FeatureName = C.WGPUNativeFeature_TextureAdapterSpecificFormatFeatures                  // TextureAdapterSpecificFormatFeatures
	NativeFeatureNameMultiDrawIndirect                                     FeatureName = C.WGPUNativeFeature_MultiDrawIndirect                                     // MultiDrawIndirect
	NativeFeatureNameMultiDrawIndirectCount                                FeatureName = C.WGPUNativeFeature_MultiDrawIndirectCount                                // MultiDrawIndirectCount
	NativeFeatureNameVertexWritableStorage                                 FeatureName = C.WGPUNativeFeature_VertexWritableStorage                                 // VertexWritableStorage
	NativeFeatureNameTextureBindingArray                                   FeatureName = C.WGPUNativeFeature_TextureBindingArray                                   // TextureBindingArray
	NativeFeatureNameSampledTextureAndStorageBufferArrayNonUniformIndexing FeatureName = C.WGPUNativeFeature_SampledTextureAndStorageBufferArrayNonUniformIndexing // SampledTextureAndStorageBufferArrayNonUniformIndexing
	NativeFeatureNamePipelineStatisticsQuery                               FeatureName = C.WGPUNativeFeature_PipelineStatisticsQuery                               // PipelineStatisticsQuery
	NativeFeatureNameStorageResourceBindingArray                           FeatureName = C.WGPUNativeFeature_StorageResourceBindingArray                           // StorageResourceBindingArray
	NativeFeatureNamePartiallyBoundBindingArray                            FeatureName = C.WGPUNativeFeature_PartiallyBoundBindingArray                            // PartiallyBoundBindingArray
)

func init() {
	C.wgpuSetLogCallback(fp(C.logCallback), nil)
}

//export logCallback
func logCallback(level C.WGPULogLevel, msg *C.char, data up) {
	log.Println(LogLevel(level), C.GoString(msg))
}

func Version() string {
	v := C.wgpuGetVersion()
	pieces := []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
	}
	for len(pieces) > 1 && pieces[0] == 0 {
		pieces = pieces[1:]
	}
	var sb strings.Builder
	sb.WriteByte('v')
	for i := len(pieces) - 1; i >= 0; i-- {
		sb.WriteString(strconv.Itoa(int(pieces[i])))
		if i > 0 {
			sb.WriteByte('.')
		}
	}
	return sb.String()
}

func SetLogLevel(level LogLevel) {
	C.wgpuSetLogLevel(C.WGPULogLevel(level))
}

type InstanceExtras struct {
	Backends           InstanceBackendFlags
	Flags              InstanceFlags
	DX12ShaderCompiler DX12Compiler
	GLES3MinorVersion  GLES3MinorVersion
	DXILPath           string
	DXCPath            string
}

func (dev *Device) Poll(wait bool) bool {
	// Returns true if the queue is empty, or false if there are more queue submissions still in flight.

	// TODO(dh): support WGPUUnwrappedSubmissionIndex
	return C.wgpuDevicePoll(dev.hnd, toBool(wait), nil) != 0
}

type InstanceEnumerateAdapterOptions struct {
	Backends InstanceBackendFlags
}

// Adapters returns all adapters available on this instance. The returned
// adapters are ready to use. You should call [Adapter.Release] on unneeded
// adapters.
func (ins *Instance) Adapters(opts InstanceEnumerateAdapterOptions) []*Adapter {
	copts := valloc(C.WGPUInstanceEnumerateAdapterOptions{
		backends: C.WGPUInstanceBackendFlags(opts.Backends),
	})
	defer free(copts)

	for {
		n := C.wgpuInstanceEnumerateAdapters(ins.c(), copts, nil)
		out := make([]C.WGPUAdapter, n)
		nn := C.wgpuInstanceEnumerateAdapters(ins.c(), copts, unsafe.SliceData(out))
		if nn <= n {
			return safeish.SliceCast[[]*Adapter](out[:nn])
		}
	}
}

type RegistryReport struct {
	_ structs.HostLayout

	NumAllocated        int
	NumKeptFromUser     int
	NumReleasedFromUser int
	NumError            int
	ElementSize         int
}

type GlobalReport struct {
	_ structs.HostLayout

	Surfaces    RegistryReport
	BackendType BackendType
	Vulkan      HubReport
	Metal       HubReport
	DX12        HubReport
	GL          HubReport
}

type HubReport struct {
	_ structs.HostLayout

	Adapters         RegistryReport
	Devices          RegistryReport
	Queues           RegistryReport
	PipelineLayouts  RegistryReport
	ShaderModules    RegistryReport
	BindGroupLayouts RegistryReport
	BindGroups       RegistryReport
	CommandBuffers   RegistryReport
	RenderBundles    RegistryReport
	RenderPipelines  RegistryReport
	ComputePipelines RegistryReport
	QuerySets        RegistryReport
	Buffers          RegistryReport
	Textures         RegistryReport
	TextureViews     RegistryReport
	Samplers         RegistryReport
}

// Check that sizeof(GlobalReport) == sizeof(C.WGPUGlobalReport)
const _ = -uint(unsafe.Sizeof(GlobalReport{}) - unsafe.Sizeof(C.WGPUGlobalReport{}))

func (ins *Instance) Report(out *GlobalReport) {
	C.wgpuGenerateReport(ins.c(), safeish.Cast[*C.WGPUGlobalReport](out))
}
