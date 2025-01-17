package wgpu

// TODO(dh): https://github.com/gfx-rs/wgpu/issues/3794#issuecomment-1954631975

// #include <stdlib.h>
// #include "./webgpu.h"
// #include "./wgpu.h"
//
// void requestAdapterCallback(WGPURequestAdapterStatus status, WGPUAdapter adapter, char *msg, void *data);
// void requestDeviceCallback(WGPURequestDeviceStatus status, WGPUDevice device, char *msg, void *data);
// void deviceLostCallback(WGPUDeviceLostReason reason, char *msg, uintptr_t hnd);
// void popErrorScopeCallback(WGPUErrorType type, char *message, void *userdata);
// void mapCallback(WGPUBufferMapAsyncStatus status, void *userdata);
// void doneCallback(WGPUQueueWorkDoneStatus status, void *userdata);
// void uncapturedErrorCallback(WGPUErrorType type, char *message, void *userdata);
import "C"

import (
	"errors"
	"fmt"
	"runtime"
	"runtime/cgo"
	"structs"
	"sync"
	"unsafe"

	"golang.org/x/exp/constraints"
	"honnef.co/go/safeish"
)

type fp = *[0]byte
type up = unsafe.Pointer

//go:generate stringer -type requestAdapterStatus,requestDeviceStatus,PowerPreference,BackendType,AdapterType,DeviceLostReason,ErrorFilter,PrimitiveTopology,IndexFormat,FrontFace,CullMode,TextureFormat,VertexFormat,CompareFunction,StencilOperation,BlendOperation,BlendFactor,ColorWriteMask,VertexStepMode,PresentMode,CompositeAlphaMode,TextureUsage,TextureViewDimension,TextureAspect,LoadOp,StoreOp,QueryType,BufferUsage,MapMode,BufferMapState,ShaderStage,BufferBindingType,SamplerBindingType,TextureSampleType,StorageTextureAccess,TextureDimension,AddressMode,FilterMode,MipmapFilterMode -linecomment -output webgpu1_string.go
type requestAdapterStatus uint32
type PowerPreference uint32
type BackendType uint32
type AdapterType uint32
type DeviceLostReason uint32
type ErrorFilter uint32
type PrimitiveTopology uint32
type IndexFormat uint32
type FrontFace uint32
type CullMode uint32
type TextureFormat uint32
type VertexFormat uint32
type CompareFunction uint32
type StencilOperation uint32
type BlendOperation uint32
type BlendFactor uint32
type ColorWriteMask uint32
type VertexStepMode uint32
type PresentMode uint32
type CompositeAlphaMode uint32
type TextureUsage uint32
type TextureViewDimension uint32
type TextureAspect uint32
type LoadOp uint32
type StoreOp uint32
type QueryType uint32
type BufferUsage uint32
type MapMode uint32
type BufferMapState uint32
type ShaderStage uint32
type BufferBindingType uint32
type SamplerBindingType uint32
type TextureSampleType uint32
type StorageTextureAccess uint32
type TextureDimension uint32
type AddressMode uint32
type FilterMode uint32
type MipmapFilterMode uint32

//go:generate stringer -type FeatureName -output webgpu2_string.go
type FeatureName uint32

var ErrAdapterUnavailable = errors.New("no adapter available")

type RequestAdapterError struct {
	Message string
}

func (err RequestAdapterError) Error() string { return err.Message }

const (
	requestAdapterStatusSuccess     requestAdapterStatus = C.WGPURequestAdapterStatus_Success     // Success
	requestAdapterStatusUnavailable requestAdapterStatus = C.WGPURequestAdapterStatus_Unavailable // Unavailable
	requestAdapterStatusError       requestAdapterStatus = C.WGPURequestAdapterStatus_Error       // Error
	requestAdapterStatusUnknown     requestAdapterStatus = C.WGPURequestAdapterStatus_Unknown     // Unknown
)

type requestDeviceStatus uint32

type RequestDeviceError struct {
	Message string
}

func (err RequestDeviceError) Error() string { return err.Message }

const (
	requestDeviceStatusSuccess requestDeviceStatus = C.WGPURequestDeviceStatus_Success // Success
	requestDeviceStatusError   requestDeviceStatus = C.WGPURequestDeviceStatus_Error   // Error
	requestDeviceStatusUnknown requestDeviceStatus = C.WGPURequestDeviceStatus_Unknown // Unknown
)

const (
	FeatureNameUndefined               FeatureName = C.WGPUFeatureName_Undefined               // Undefined
	FeatureNameDepthClipControl        FeatureName = C.WGPUFeatureName_DepthClipControl        // DepthClipControl
	FeatureNameDepth32FloatStencil8    FeatureName = C.WGPUFeatureName_Depth32FloatStencil8    // Depth32FloatStencil8
	FeatureNameTimestampQuery          FeatureName = C.WGPUFeatureName_TimestampQuery          // TimestampQuery
	FeatureNameTextureCompressionBC    FeatureName = C.WGPUFeatureName_TextureCompressionBC    // TextureCompressionBC
	FeatureNameTextureCompressionETC2  FeatureName = C.WGPUFeatureName_TextureCompressionETC2  // TextureCompressionETC2
	FeatureNameTextureCompressionASTC  FeatureName = C.WGPUFeatureName_TextureCompressionASTC  // TextureCompressionASTC
	FeatureNameIndirectFirstInstance   FeatureName = C.WGPUFeatureName_IndirectFirstInstance   // IndirectFirstInstance
	FeatureNameShaderF16               FeatureName = C.WGPUFeatureName_ShaderF16               // ShaderF16
	FeatureNameRG11B10UfloatRenderable FeatureName = C.WGPUFeatureName_RG11B10UfloatRenderable // RG11B10UfloatRenderable
	FeatureNameBGRA8UnormStorage       FeatureName = C.WGPUFeatureName_BGRA8UnormStorage       // BGRA8UnormStorage
	FeatureNameFloat32Filterable       FeatureName = C.WGPUFeatureName_Float32Filterable       // Float32Filterable
)

const (
	PowerPreferenceUndefined       PowerPreference = C.WGPUPowerPreference_Undefined       // Undefined
	PowerPreferenceLowPower        PowerPreference = C.WGPUPowerPreference_LowPower        // LowPower
	PowerPreferenceHighPerformance PowerPreference = C.WGPUPowerPreference_HighPerformance // HighPerformance
)

const (
	BackendTypeUndefined BackendType = C.WGPUBackendType_Undefined // Undefined
	BackendTypeNull      BackendType = C.WGPUBackendType_Null      // Null
	BackendTypeWebGPU    BackendType = C.WGPUBackendType_WebGPU    // WebGPU
	BackendTypeD3D11     BackendType = C.WGPUBackendType_D3D11     // D3D11
	BackendTypeD3D12     BackendType = C.WGPUBackendType_D3D12     // D3D12
	BackendTypeMetal     BackendType = C.WGPUBackendType_Metal     // Metal
	BackendTypeVulkan    BackendType = C.WGPUBackendType_Vulkan    // Vulkan
	BackendTypeOpenGL    BackendType = C.WGPUBackendType_OpenGL    // OpenGL
	BackendTypeOpenGLES  BackendType = C.WGPUBackendType_OpenGLES  // OpenGLES
)

const (
	AdapterTypeDiscreteGPU   AdapterType = C.WGPUAdapterType_DiscreteGPU   // DiscreteGPU
	AdapterTypeIntegratedGPU AdapterType = C.WGPUAdapterType_IntegratedGPU // IntegratedGPU
	AdapterTypeCPU           AdapterType = C.WGPUAdapterType_CPU           // CPU
	AdapterTypeUnknown       AdapterType = C.WGPUAdapterType_Unknown       // Unknown
)

const (
	DeviceLostReasonUndefined DeviceLostReason = C.WGPUDeviceLostReason_Undefined // Undefined
	DeviceLostReasonDestroyed DeviceLostReason = C.WGPUDeviceLostReason_Destroyed // Destroyed
)

const (
	ErrorFilterValidation  ErrorFilter = C.WGPUErrorFilter_Validation  // Validation
	ErrorFilterOutOfMemory ErrorFilter = C.WGPUErrorFilter_OutOfMemory // OutOfMemory
	ErrorFilterInternal    ErrorFilter = C.WGPUErrorFilter_Internal    // Internal
)

const (
	PrimitiveTopologyPointList     PrimitiveTopology = C.WGPUPrimitiveTopology_PointList     // PointList
	PrimitiveTopologyLineList      PrimitiveTopology = C.WGPUPrimitiveTopology_LineList      // LineList
	PrimitiveTopologyLineStrip     PrimitiveTopology = C.WGPUPrimitiveTopology_LineStrip     // LineStrip
	PrimitiveTopologyTriangleList  PrimitiveTopology = C.WGPUPrimitiveTopology_TriangleList  // TriangleList
	PrimitiveTopologyTriangleStrip PrimitiveTopology = C.WGPUPrimitiveTopology_TriangleStrip // TriangleStrip
)

const (
	IndexFormatUndefined IndexFormat = C.WGPUIndexFormat_Undefined // Undefined
	IndexFormatUint16    IndexFormat = C.WGPUIndexFormat_Uint16    // Uint16
	IndexFormatUint32    IndexFormat = C.WGPUIndexFormat_Uint32    // Uint32
)

const (
	FrontFaceCCW FrontFace = C.WGPUFrontFace_CCW // CCW
	FrontFaceCW  FrontFace = C.WGPUFrontFace_CW  // CW
)

const (
	CullModeNone  CullMode = C.WGPUCullMode_None  // None
	CullModeFront CullMode = C.WGPUCullMode_Front // Front
	CullModeBack  CullMode = C.WGPUCullMode_Back  // Back
)

const (
	TextureFormatUndefined            TextureFormat = C.WGPUTextureFormat_Undefined            // Undefined
	TextureFormatR8Unorm              TextureFormat = C.WGPUTextureFormat_R8Unorm              // R8Unorm
	TextureFormatR8Snorm              TextureFormat = C.WGPUTextureFormat_R8Snorm              // R8Snorm
	TextureFormatR8Uint               TextureFormat = C.WGPUTextureFormat_R8Uint               // R8Uint
	TextureFormatR8Sint               TextureFormat = C.WGPUTextureFormat_R8Sint               // R8Sint
	TextureFormatR16Uint              TextureFormat = C.WGPUTextureFormat_R16Uint              // R16Uint
	TextureFormatR16Sint              TextureFormat = C.WGPUTextureFormat_R16Sint              // R16Sint
	TextureFormatR16Float             TextureFormat = C.WGPUTextureFormat_R16Float             // R16Float
	TextureFormatRG8Unorm             TextureFormat = C.WGPUTextureFormat_RG8Unorm             // RG8Unorm
	TextureFormatRG8Snorm             TextureFormat = C.WGPUTextureFormat_RG8Snorm             // RG8Snorm
	TextureFormatRG8Uint              TextureFormat = C.WGPUTextureFormat_RG8Uint              // RG8Uint
	TextureFormatRG8Sint              TextureFormat = C.WGPUTextureFormat_RG8Sint              // RG8Sint
	TextureFormatR32Float             TextureFormat = C.WGPUTextureFormat_R32Float             // R32Float
	TextureFormatR32Uint              TextureFormat = C.WGPUTextureFormat_R32Uint              // R32Uint
	TextureFormatR32Sint              TextureFormat = C.WGPUTextureFormat_R32Sint              // R32Sint
	TextureFormatRG16Uint             TextureFormat = C.WGPUTextureFormat_RG16Uint             // RG16Uint
	TextureFormatRG16Sint             TextureFormat = C.WGPUTextureFormat_RG16Sint             // RG16Sint
	TextureFormatRG16Float            TextureFormat = C.WGPUTextureFormat_RG16Float            // RG16Float
	TextureFormatRGBA8Unorm           TextureFormat = C.WGPUTextureFormat_RGBA8Unorm           // RGBA8Unorm
	TextureFormatRGBA8UnormSrgb       TextureFormat = C.WGPUTextureFormat_RGBA8UnormSrgb       // RGBA8UnormSrgb
	TextureFormatRGBA8Snorm           TextureFormat = C.WGPUTextureFormat_RGBA8Snorm           // RGBA8Snorm
	TextureFormatRGBA8Uint            TextureFormat = C.WGPUTextureFormat_RGBA8Uint            // RGBA8Uint
	TextureFormatRGBA8Sint            TextureFormat = C.WGPUTextureFormat_RGBA8Sint            // RGBA8Sint
	TextureFormatBGRA8Unorm           TextureFormat = C.WGPUTextureFormat_BGRA8Unorm           // BGRA8Unorm
	TextureFormatBGRA8UnormSrgb       TextureFormat = C.WGPUTextureFormat_BGRA8UnormSrgb       // BGRA8UnormSrgb
	TextureFormatRGB10A2Uint          TextureFormat = C.WGPUTextureFormat_RGB10A2Uint          // RGB10A2Uint
	TextureFormatRGB10A2Unorm         TextureFormat = C.WGPUTextureFormat_RGB10A2Unorm         // RGB10A2Unorm
	TextureFormatRG11B10Ufloat        TextureFormat = C.WGPUTextureFormat_RG11B10Ufloat        // RG11B10Ufloat
	TextureFormatRGB9E5Ufloat         TextureFormat = C.WGPUTextureFormat_RGB9E5Ufloat         // RGB9E5Ufloat
	TextureFormatRG32Float            TextureFormat = C.WGPUTextureFormat_RG32Float            // RG32Float
	TextureFormatRG32Uint             TextureFormat = C.WGPUTextureFormat_RG32Uint             // RG32Uint
	TextureFormatRG32Sint             TextureFormat = C.WGPUTextureFormat_RG32Sint             // RG32Sint
	TextureFormatRGBA16Uint           TextureFormat = C.WGPUTextureFormat_RGBA16Uint           // RGBA16Uint
	TextureFormatRGBA16Sint           TextureFormat = C.WGPUTextureFormat_RGBA16Sint           // RGBA16Sint
	TextureFormatRGBA16Float          TextureFormat = C.WGPUTextureFormat_RGBA16Float          // RGBA16Float
	TextureFormatRGBA32Float          TextureFormat = C.WGPUTextureFormat_RGBA32Float          // RGBA32Float
	TextureFormatRGBA32Uint           TextureFormat = C.WGPUTextureFormat_RGBA32Uint           // RGBA32Uint
	TextureFormatRGBA32Sint           TextureFormat = C.WGPUTextureFormat_RGBA32Sint           // RGBA32Sint
	TextureFormatStencil8             TextureFormat = C.WGPUTextureFormat_Stencil8             // Stencil8
	TextureFormatDepth16Unorm         TextureFormat = C.WGPUTextureFormat_Depth16Unorm         // Depth16Unorm
	TextureFormatDepth24Plus          TextureFormat = C.WGPUTextureFormat_Depth24Plus          // Depth24Plus
	TextureFormatDepth24PlusStencil8  TextureFormat = C.WGPUTextureFormat_Depth24PlusStencil8  // Depth24PlusStencil8
	TextureFormatDepth32Float         TextureFormat = C.WGPUTextureFormat_Depth32Float         // Depth32Float
	TextureFormatDepth32FloatStencil8 TextureFormat = C.WGPUTextureFormat_Depth32FloatStencil8 // Depth32FloatStencil8
	TextureFormatBC1RGBAUnorm         TextureFormat = C.WGPUTextureFormat_BC1RGBAUnorm         // BC1RGBAUnorm
	TextureFormatBC1RGBAUnormSrgb     TextureFormat = C.WGPUTextureFormat_BC1RGBAUnormSrgb     // BC1RGBAUnormSrgb
	TextureFormatBC2RGBAUnorm         TextureFormat = C.WGPUTextureFormat_BC2RGBAUnorm         // BC2RGBAUnorm
	TextureFormatBC2RGBAUnormSrgb     TextureFormat = C.WGPUTextureFormat_BC2RGBAUnormSrgb     // BC2RGBAUnormSrgb
	TextureFormatBC3RGBAUnorm         TextureFormat = C.WGPUTextureFormat_BC3RGBAUnorm         // BC3RGBAUnorm
	TextureFormatBC3RGBAUnormSrgb     TextureFormat = C.WGPUTextureFormat_BC3RGBAUnormSrgb     // BC3RGBAUnormSrgb
	TextureFormatBC4RUnorm            TextureFormat = C.WGPUTextureFormat_BC4RUnorm            // BC4RUnorm
	TextureFormatBC4RSnorm            TextureFormat = C.WGPUTextureFormat_BC4RSnorm            // BC4RSnorm
	TextureFormatBC5RGUnorm           TextureFormat = C.WGPUTextureFormat_BC5RGUnorm           // BC5RGUnorm
	TextureFormatBC5RGSnorm           TextureFormat = C.WGPUTextureFormat_BC5RGSnorm           // BC5RGSnorm
	TextureFormatBC6HRGBUfloat        TextureFormat = C.WGPUTextureFormat_BC6HRGBUfloat        // BC6HRGBUfloat
	TextureFormatBC6HRGBFloat         TextureFormat = C.WGPUTextureFormat_BC6HRGBFloat         // BC6HRGBFloat
	TextureFormatBC7RGBAUnorm         TextureFormat = C.WGPUTextureFormat_BC7RGBAUnorm         // BC7RGBAUnorm
	TextureFormatBC7RGBAUnormSrgb     TextureFormat = C.WGPUTextureFormat_BC7RGBAUnormSrgb     // BC7RGBAUnormSrgb
	TextureFormatETC2RGB8Unorm        TextureFormat = C.WGPUTextureFormat_ETC2RGB8Unorm        // ETC2RGB8Unorm
	TextureFormatETC2RGB8UnormSrgb    TextureFormat = C.WGPUTextureFormat_ETC2RGB8UnormSrgb    // ETC2RGB8UnormSrgb
	TextureFormatETC2RGB8A1Unorm      TextureFormat = C.WGPUTextureFormat_ETC2RGB8A1Unorm      // ETC2RGB8A1Unorm
	TextureFormatETC2RGB8A1UnormSrgb  TextureFormat = C.WGPUTextureFormat_ETC2RGB8A1UnormSrgb  // ETC2RGB8A1UnormSrgb
	TextureFormatETC2RGBA8Unorm       TextureFormat = C.WGPUTextureFormat_ETC2RGBA8Unorm       // ETC2RGBA8Unorm
	TextureFormatETC2RGBA8UnormSrgb   TextureFormat = C.WGPUTextureFormat_ETC2RGBA8UnormSrgb   // ETC2RGBA8UnormSrgb
	TextureFormatEACR11Unorm          TextureFormat = C.WGPUTextureFormat_EACR11Unorm          // EACR11Unorm
	TextureFormatEACR11Snorm          TextureFormat = C.WGPUTextureFormat_EACR11Snorm          // EACR11Snorm
	TextureFormatEACRG11Unorm         TextureFormat = C.WGPUTextureFormat_EACRG11Unorm         // EACRG11Unorm
	TextureFormatEACRG11Snorm         TextureFormat = C.WGPUTextureFormat_EACRG11Snorm         // EACRG11Snorm
	TextureFormatASTC4x4Unorm         TextureFormat = C.WGPUTextureFormat_ASTC4x4Unorm         // ASTC4x4Unorm
	TextureFormatASTC4x4UnormSrgb     TextureFormat = C.WGPUTextureFormat_ASTC4x4UnormSrgb     // ASTC4x4UnormSrgb
	TextureFormatASTC5x4Unorm         TextureFormat = C.WGPUTextureFormat_ASTC5x4Unorm         // ASTC5x4Unorm
	TextureFormatASTC5x4UnormSrgb     TextureFormat = C.WGPUTextureFormat_ASTC5x4UnormSrgb     // ASTC5x4UnormSrgb
	TextureFormatASTC5x5Unorm         TextureFormat = C.WGPUTextureFormat_ASTC5x5Unorm         // ASTC5x5Unorm
	TextureFormatASTC5x5UnormSrgb     TextureFormat = C.WGPUTextureFormat_ASTC5x5UnormSrgb     // ASTC5x5UnormSrgb
	TextureFormatASTC6x5Unorm         TextureFormat = C.WGPUTextureFormat_ASTC6x5Unorm         // ASTC6x5Unorm
	TextureFormatASTC6x5UnormSrgb     TextureFormat = C.WGPUTextureFormat_ASTC6x5UnormSrgb     // ASTC6x5UnormSrgb
	TextureFormatASTC6x6Unorm         TextureFormat = C.WGPUTextureFormat_ASTC6x6Unorm         // ASTC6x6Unorm
	TextureFormatASTC6x6UnormSrgb     TextureFormat = C.WGPUTextureFormat_ASTC6x6UnormSrgb     // ASTC6x6UnormSrgb
	TextureFormatASTC8x5Unorm         TextureFormat = C.WGPUTextureFormat_ASTC8x5Unorm         // ASTC8x5Unorm
	TextureFormatASTC8x5UnormSrgb     TextureFormat = C.WGPUTextureFormat_ASTC8x5UnormSrgb     // ASTC8x5UnormSrgb
	TextureFormatASTC8x6Unorm         TextureFormat = C.WGPUTextureFormat_ASTC8x6Unorm         // ASTC8x6Unorm
	TextureFormatASTC8x6UnormSrgb     TextureFormat = C.WGPUTextureFormat_ASTC8x6UnormSrgb     // ASTC8x6UnormSrgb
	TextureFormatASTC8x8Unorm         TextureFormat = C.WGPUTextureFormat_ASTC8x8Unorm         // ASTC8x8Unorm
	TextureFormatASTC8x8UnormSrgb     TextureFormat = C.WGPUTextureFormat_ASTC8x8UnormSrgb     // ASTC8x8UnormSrgb
	TextureFormatASTC10x5Unorm        TextureFormat = C.WGPUTextureFormat_ASTC10x5Unorm        // ASTC10x5Unorm
	TextureFormatASTC10x5UnormSrgb    TextureFormat = C.WGPUTextureFormat_ASTC10x5UnormSrgb    // ASTC10x5UnormSrgb
	TextureFormatASTC10x6Unorm        TextureFormat = C.WGPUTextureFormat_ASTC10x6Unorm        // ASTC10x6Unorm
	TextureFormatASTC10x6UnormSrgb    TextureFormat = C.WGPUTextureFormat_ASTC10x6UnormSrgb    // ASTC10x6UnormSrgb
	TextureFormatASTC10x8Unorm        TextureFormat = C.WGPUTextureFormat_ASTC10x8Unorm        // ASTC10x8Unorm
	TextureFormatASTC10x8UnormSrgb    TextureFormat = C.WGPUTextureFormat_ASTC10x8UnormSrgb    // ASTC10x8UnormSrgb
	TextureFormatASTC10x10Unorm       TextureFormat = C.WGPUTextureFormat_ASTC10x10Unorm       // ASTC10x10Unorm
	TextureFormatASTC10x10UnormSrgb   TextureFormat = C.WGPUTextureFormat_ASTC10x10UnormSrgb   // ASTC10x10UnormSrgb
	TextureFormatASTC12x10Unorm       TextureFormat = C.WGPUTextureFormat_ASTC12x10Unorm       // ASTC12x10Unorm
	TextureFormatASTC12x10UnormSrgb   TextureFormat = C.WGPUTextureFormat_ASTC12x10UnormSrgb   // ASTC12x10UnormSrgb
	TextureFormatASTC12x12Unorm       TextureFormat = C.WGPUTextureFormat_ASTC12x12Unorm       // ASTC12x12Unorm
	TextureFormatASTC12x12UnormSrgb   TextureFormat = C.WGPUTextureFormat_ASTC12x12UnormSrgb   // ASTC12x12UnormSrgb
)

const (
	VertexFormatUndefined VertexFormat = C.WGPUVertexFormat_Undefined // Undefined
	VertexFormatUint8x2   VertexFormat = C.WGPUVertexFormat_Uint8x2   // Uint8x2
	VertexFormatUint8x4   VertexFormat = C.WGPUVertexFormat_Uint8x4   // Uint8x4
	VertexFormatSint8x2   VertexFormat = C.WGPUVertexFormat_Sint8x2   // Sint8x2
	VertexFormatSint8x4   VertexFormat = C.WGPUVertexFormat_Sint8x4   // Sint8x4
	VertexFormatUnorm8x2  VertexFormat = C.WGPUVertexFormat_Unorm8x2  // Unorm8x2
	VertexFormatUnorm8x4  VertexFormat = C.WGPUVertexFormat_Unorm8x4  // Unorm8x4
	VertexFormatSnorm8x2  VertexFormat = C.WGPUVertexFormat_Snorm8x2  // Snorm8x2
	VertexFormatSnorm8x4  VertexFormat = C.WGPUVertexFormat_Snorm8x4  // Snorm8x4
	VertexFormatUint16x2  VertexFormat = C.WGPUVertexFormat_Uint16x2  // Uint16x2
	VertexFormatUint16x4  VertexFormat = C.WGPUVertexFormat_Uint16x4  // Uint16x4
	VertexFormatSint16x2  VertexFormat = C.WGPUVertexFormat_Sint16x2  // Sint16x2
	VertexFormatSint16x4  VertexFormat = C.WGPUVertexFormat_Sint16x4  // Sint16x4
	VertexFormatUnorm16x2 VertexFormat = C.WGPUVertexFormat_Unorm16x2 // Unorm16x2
	VertexFormatUnorm16x4 VertexFormat = C.WGPUVertexFormat_Unorm16x4 // Unorm16x4
	VertexFormatSnorm16x2 VertexFormat = C.WGPUVertexFormat_Snorm16x2 // Snorm16x2
	VertexFormatSnorm16x4 VertexFormat = C.WGPUVertexFormat_Snorm16x4 // Snorm16x4
	VertexFormatFloat16x2 VertexFormat = C.WGPUVertexFormat_Float16x2 // Float16x2
	VertexFormatFloat16x4 VertexFormat = C.WGPUVertexFormat_Float16x4 // Float16x4
	VertexFormatFloat32   VertexFormat = C.WGPUVertexFormat_Float32   // Float32
	VertexFormatFloat32x2 VertexFormat = C.WGPUVertexFormat_Float32x2 // Float32x2
	VertexFormatFloat32x3 VertexFormat = C.WGPUVertexFormat_Float32x3 // Float32x3
	VertexFormatFloat32x4 VertexFormat = C.WGPUVertexFormat_Float32x4 // Float32x4
	VertexFormatUint32    VertexFormat = C.WGPUVertexFormat_Uint32    // Uint32
	VertexFormatUint32x2  VertexFormat = C.WGPUVertexFormat_Uint32x2  // Uint32x2
	VertexFormatUint32x3  VertexFormat = C.WGPUVertexFormat_Uint32x3  // Uint32x3
	VertexFormatUint32x4  VertexFormat = C.WGPUVertexFormat_Uint32x4  // Uint32x4
	VertexFormatSint32    VertexFormat = C.WGPUVertexFormat_Sint32    // Sint32
	VertexFormatSint32x2  VertexFormat = C.WGPUVertexFormat_Sint32x2  // Sint32x2
	VertexFormatSint32x3  VertexFormat = C.WGPUVertexFormat_Sint32x3  // Sint32x3
	VertexFormatSint32x4  VertexFormat = C.WGPUVertexFormat_Sint32x4  // Sint32x4
)

const (
	CompareFunctionUndefined    CompareFunction = C.WGPUCompareFunction_Undefined    // Undefined
	CompareFunctionNever        CompareFunction = C.WGPUCompareFunction_Never        // Never
	CompareFunctionLess         CompareFunction = C.WGPUCompareFunction_Less         // Less
	CompareFunctionLessEqual    CompareFunction = C.WGPUCompareFunction_LessEqual    // LessEqual
	CompareFunctionGreater      CompareFunction = C.WGPUCompareFunction_Greater      // Greater
	CompareFunctionGreaterEqual CompareFunction = C.WGPUCompareFunction_GreaterEqual // GreaterEqual
	CompareFunctionEqual        CompareFunction = C.WGPUCompareFunction_Equal        // Equal
	CompareFunctionNotEqual     CompareFunction = C.WGPUCompareFunction_NotEqual     // NotEqual
	CompareFunctionAlways       CompareFunction = C.WGPUCompareFunction_Always       // Always
)

const (
	StencilOperationKeep           StencilOperation = C.WGPUStencilOperation_Keep           // Keep
	StencilOperationZero           StencilOperation = C.WGPUStencilOperation_Zero           // Zero
	StencilOperationReplace        StencilOperation = C.WGPUStencilOperation_Replace        // Replace
	StencilOperationInvert         StencilOperation = C.WGPUStencilOperation_Invert         // Invert
	StencilOperationIncrementClamp StencilOperation = C.WGPUStencilOperation_IncrementClamp // IncrementClamp
	StencilOperationDecrementClamp StencilOperation = C.WGPUStencilOperation_DecrementClamp // DecrementClamp
	StencilOperationIncrementWrap  StencilOperation = C.WGPUStencilOperation_IncrementWrap  // IncrementWrap
	StencilOperationDecrementWrap  StencilOperation = C.WGPUStencilOperation_DecrementWrap  // DecrementWrap
)

const (
	BlendOperationAdd             BlendOperation = C.WGPUBlendOperation_Add             // Add
	BlendOperationSubtract        BlendOperation = C.WGPUBlendOperation_Subtract        // Subtract
	BlendOperationReverseSubtract BlendOperation = C.WGPUBlendOperation_ReverseSubtract // ReverseSubtract
	BlendOperationMin             BlendOperation = C.WGPUBlendOperation_Min             // Min
	BlendOperationMax             BlendOperation = C.WGPUBlendOperation_Max             // Max
)

const (
	BlendFactorZero              BlendFactor = C.WGPUBlendFactor_Zero              // Zero
	BlendFactorOne               BlendFactor = C.WGPUBlendFactor_One               // One
	BlendFactorSrc               BlendFactor = C.WGPUBlendFactor_Src               // Src
	BlendFactorOneMinusSrc       BlendFactor = C.WGPUBlendFactor_OneMinusSrc       // OneMinusSrc
	BlendFactorSrcAlpha          BlendFactor = C.WGPUBlendFactor_SrcAlpha          // SrcAlpha
	BlendFactorOneMinusSrcAlpha  BlendFactor = C.WGPUBlendFactor_OneMinusSrcAlpha  // OneMinusSrcAlpha
	BlendFactorDst               BlendFactor = C.WGPUBlendFactor_Dst               // Dst
	BlendFactorOneMinusDst       BlendFactor = C.WGPUBlendFactor_OneMinusDst       // OneMinusDst
	BlendFactorDstAlpha          BlendFactor = C.WGPUBlendFactor_DstAlpha          // DstAlpha
	BlendFactorOneMinusDstAlpha  BlendFactor = C.WGPUBlendFactor_OneMinusDstAlpha  // OneMinusDstAlpha
	BlendFactorSrcAlphaSaturated BlendFactor = C.WGPUBlendFactor_SrcAlphaSaturated // SrcAlphaSaturated
	BlendFactorConstant          BlendFactor = C.WGPUBlendFactor_Constant          // Constant
	BlendFactorOneMinusConstant  BlendFactor = C.WGPUBlendFactor_OneMinusConstant  // OneMinusConstant
)

const (
	ColorWriteMaskNone  ColorWriteMask = C.WGPUColorWriteMask_None  // None
	ColorWriteMaskRed   ColorWriteMask = C.WGPUColorWriteMask_Red   // Red
	ColorWriteMaskGreen ColorWriteMask = C.WGPUColorWriteMask_Green // Green
	ColorWriteMaskBlue  ColorWriteMask = C.WGPUColorWriteMask_Blue  // Blue
	ColorWriteMaskAlpha ColorWriteMask = C.WGPUColorWriteMask_Alpha // Alpha
	ColorWriteMaskAll   ColorWriteMask = C.WGPUColorWriteMask_All   // All
)

const (
	VertexStepModeVertex              VertexStepMode = C.WGPUVertexStepMode_Vertex              // Vertex
	VertexStepModeInstance            VertexStepMode = C.WGPUVertexStepMode_Instance            // Instance
	VertexStepModeVertexBufferNotUsed VertexStepMode = C.WGPUVertexStepMode_VertexBufferNotUsed // VertexBufferNotUsed
)

const (
	PresentModeFifo        PresentMode = C.WGPUPresentMode_Fifo        // Fifo
	PresentModeFifoRelaxed PresentMode = C.WGPUPresentMode_FifoRelaxed // FifoRelaxed
	PresentModeImmediate   PresentMode = C.WGPUPresentMode_Immediate   // Immediate
	PresentModeMailbox     PresentMode = C.WGPUPresentMode_Mailbox     // Mailbox
)

const (
	CompositeAlphaModeAuto            CompositeAlphaMode = C.WGPUCompositeAlphaMode_Auto            // Auto
	CompositeAlphaModeOpaque          CompositeAlphaMode = C.WGPUCompositeAlphaMode_Opaque          // Opaque
	CompositeAlphaModePremultiplied   CompositeAlphaMode = C.WGPUCompositeAlphaMode_Premultiplied   // Premultiplied
	CompositeAlphaModeUnpremultiplied CompositeAlphaMode = C.WGPUCompositeAlphaMode_Unpremultiplied // Unpremultiplied
	CompositeAlphaModeInherit         CompositeAlphaMode = C.WGPUCompositeAlphaMode_Inherit         // Inherit
)

const (
	TextureUsageNone             TextureUsage = C.WGPUTextureUsage_None             // None
	TextureUsageCopySrc          TextureUsage = C.WGPUTextureUsage_CopySrc          // CopySrc
	TextureUsageCopyDst          TextureUsage = C.WGPUTextureUsage_CopyDst          // CopyDst
	TextureUsageTextureBinding   TextureUsage = C.WGPUTextureUsage_TextureBinding   // TextureBinding
	TextureUsageStorageBinding   TextureUsage = C.WGPUTextureUsage_StorageBinding   // StorageBinding
	TextureUsageRenderAttachment TextureUsage = C.WGPUTextureUsage_RenderAttachment // RenderAttachment
)

const (
	TextureViewDimensionUndefined TextureViewDimension = C.WGPUTextureViewDimension_Undefined // undefined
	TextureViewDimension1D        TextureViewDimension = C.WGPUTextureViewDimension_1D        // 1D
	TextureViewDimension2D        TextureViewDimension = C.WGPUTextureViewDimension_2D        // 2D
	TextureViewDimension2DArray   TextureViewDimension = C.WGPUTextureViewDimension_2DArray   // 2DArray
	TextureViewDimensionCube      TextureViewDimension = C.WGPUTextureViewDimension_Cube      // Cube
	TextureViewDimensionCubeArray TextureViewDimension = C.WGPUTextureViewDimension_CubeArray // CubeArray
	TextureViewDimension3D        TextureViewDimension = C.WGPUTextureViewDimension_3D        // 3D
)

const (
	TextureAspectAll         TextureAspect = C.WGPUTextureAspect_All         // All
	TextureAspectStencilOnly TextureAspect = C.WGPUTextureAspect_StencilOnly // StencilOnly
	TextureAspectDepthOnly   TextureAspect = C.WGPUTextureAspect_DepthOnly   // DepthOnly
)

const (
	LoadOpUndefined LoadOp = C.WGPULoadOp_Undefined // Undefined
	LoadOpClear     LoadOp = C.WGPULoadOp_Clear     // Clear
	LoadOpLoad      LoadOp = C.WGPULoadOp_Load      // Load
)

const (
	StoreOpUndefined StoreOp = C.WGPUStoreOp_Undefined // Undefined
	StoreOpStore     StoreOp = C.WGPUStoreOp_Store     // Store
	StoreOpDiscard   StoreOp = C.WGPUStoreOp_Discard   // Discard
)

const (
	QueryTypeOcclusion QueryType = C.WGPUQueryType_Occlusion // Occlusion
	QueryTypeTimestamp QueryType = C.WGPUQueryType_Timestamp // Timestamp
)

const (
	BufferUsageNone         BufferUsage = C.WGPUBufferUsage_None         // None
	BufferUsageMapRead      BufferUsage = C.WGPUBufferUsage_MapRead      // MapRead
	BufferUsageMapWrite     BufferUsage = C.WGPUBufferUsage_MapWrite     // MapWrite
	BufferUsageCopySrc      BufferUsage = C.WGPUBufferUsage_CopySrc      // CopySrc
	BufferUsageCopyDst      BufferUsage = C.WGPUBufferUsage_CopyDst      // CopyDst
	BufferUsageIndex        BufferUsage = C.WGPUBufferUsage_Index        // Index
	BufferUsageVertex       BufferUsage = C.WGPUBufferUsage_Vertex       // Vertex
	BufferUsageUniform      BufferUsage = C.WGPUBufferUsage_Uniform      // Uniform
	BufferUsageStorage      BufferUsage = C.WGPUBufferUsage_Storage      // Storage
	BufferUsageIndirect     BufferUsage = C.WGPUBufferUsage_Indirect     // Indirect
	BufferUsageQueryResolve BufferUsage = C.WGPUBufferUsage_QueryResolve // QueryResolve
)

const (
	MapModeNone  MapMode = C.WGPUMapMode_None  // None
	MapModeRead  MapMode = C.WGPUMapMode_Read  // Read
	MapModeWrite MapMode = C.WGPUMapMode_Write // Write
)

const (
	BufferMapStateUnmapped BufferMapState = C.WGPUBufferMapState_Unmapped // Unmapped
	BufferMapStatePending  BufferMapState = C.WGPUBufferMapState_Pending  // Pending
	BufferMapStateMapped   BufferMapState = C.WGPUBufferMapState_Mapped   // Mapped
)

const (
	ShaderStageNone     ShaderStage = C.WGPUShaderStage_None     // None
	ShaderStageVertex   ShaderStage = C.WGPUShaderStage_Vertex   // Vertex
	ShaderStageFragment ShaderStage = C.WGPUShaderStage_Fragment // Fragment
	ShaderStageCompute  ShaderStage = C.WGPUShaderStage_Compute  // Compute
)

const (
	BufferBindingTypeUndefined       BufferBindingType = C.WGPUBufferBindingType_Undefined       // Undefined
	BufferBindingTypeUniform         BufferBindingType = C.WGPUBufferBindingType_Uniform         // Uniform
	BufferBindingTypeStorage         BufferBindingType = C.WGPUBufferBindingType_Storage         // Storage
	BufferBindingTypeReadOnlyStorage BufferBindingType = C.WGPUBufferBindingType_ReadOnlyStorage // ReadOnlyStorage
)

const (
	SamplerBindingTypeUndefined    SamplerBindingType = C.WGPUSamplerBindingType_Undefined    // Undefined
	SamplerBindingTypeFiltering    SamplerBindingType = C.WGPUSamplerBindingType_Filtering    // Filtering
	SamplerBindingTypeNonFiltering SamplerBindingType = C.WGPUSamplerBindingType_NonFiltering // NonFiltering
	SamplerBindingTypeComparison   SamplerBindingType = C.WGPUSamplerBindingType_Comparison   // Comparison
)

const (
	TextureSampleTypeUndefined         TextureSampleType = C.WGPUTextureSampleType_Undefined         // Undefined
	TextureSampleTypeFloat             TextureSampleType = C.WGPUTextureSampleType_Float             // Float
	TextureSampleTypeUnfilterableFloat TextureSampleType = C.WGPUTextureSampleType_UnfilterableFloat // UnfilterableFloat
	TextureSampleTypeDepth             TextureSampleType = C.WGPUTextureSampleType_Depth             // Depth
	TextureSampleTypeSint              TextureSampleType = C.WGPUTextureSampleType_Sint              // Sint
	TextureSampleTypeUint              TextureSampleType = C.WGPUTextureSampleType_Uint              // Uint
)

const (
	StorageTextureAccessUndefined StorageTextureAccess = C.WGPUStorageTextureAccess_Undefined // Undefined
	StorageTextureAccessWriteOnly StorageTextureAccess = C.WGPUStorageTextureAccess_WriteOnly // WriteOnly
	StorageTextureAccessReadOnly  StorageTextureAccess = C.WGPUStorageTextureAccess_ReadOnly  // ReadOnly
	StorageTextureAccessReadWrite StorageTextureAccess = C.WGPUStorageTextureAccess_ReadWrite // ReadWrite
)

const (
	TextureDimension1D TextureDimension = C.WGPUTextureDimension_1D // 1D
	TextureDimension2D TextureDimension = C.WGPUTextureDimension_2D // 2D
	TextureDimension3D TextureDimension = C.WGPUTextureDimension_3D // 3D
)

const (
	AddressModeRepeat       AddressMode = C.WGPUAddressMode_Repeat       // Repeat
	AddressModeMirrorRepeat AddressMode = C.WGPUAddressMode_MirrorRepeat // MirrorRepeat
	AddressModeClampToEdge  AddressMode = C.WGPUAddressMode_ClampToEdge  // ClampToEdge
)

const (
	FilterModeNearest FilterMode = C.WGPUFilterMode_Nearest // Nearest
	FilterModeLinear  FilterMode = C.WGPUFilterMode_Linear  // Linear
)

const (
	MipmapFilterModeNearest MipmapFilterMode = C.WGPUMipmapFilterMode_Nearest // Nearest
	MipmapFilterModeLinear  MipmapFilterMode = C.WGPUMipmapFilterMode_Linear  // Linear
)

type Adapter C.struct_WGPUAdapterImpl
type BindGroupLayout C.struct_WGPUBindGroupLayoutImpl
type Instance C.struct_WGPUInstanceImpl
type PipelineLayout C.struct_WGPUPipelineLayoutImpl
type Queue C.struct_WGPUQueueImpl
type RenderPipeline C.struct_WGPURenderPipelineImpl
type Surface C.struct_WGPUSurfaceImpl
type Texture C.struct_WGPUTextureImpl
type TextureView C.struct_WGPUTextureViewImpl
type CommandBuffer C.struct_WGPUCommandBufferImpl
type RenderPassEncoder C.struct_WGPURenderPassEncoderImpl
type QuerySet C.struct_WGPUQuerySetImpl
type Buffer C.struct_WGPUBufferImpl
type CommandEncoder C.struct_WGPUCommandEncoderImpl
type ShaderModule C.struct_WGPUShaderModuleImpl
type Sampler C.struct_WGPUSamplerImpl
type BindGroup C.struct_WGPUBindGroupImpl
type ComputePassEncoder C.struct_WGPUComputePassEncoderImpl
type ComputePipeline C.struct_WGPUComputePipelineImpl
type RenderBundle C.struct_WGPURenderBundleImpl
type RenderBundleEncoder C.struct_WGPURenderBundleEncoderImpl

var (
	ErrDeviceLost = errors.New("device lost")
	ErrUnknown    = errors.New("unknown error")
)

var (
	ErrCurrentTextureTimeout     = errors.New("timeout")
	ErrCurrentTextureOutdated    = errors.New("outdated")
	ErrCurrentTextureLost        = errors.New("lost")
	ErrCurrentTextureOutOfMemory = errors.New("out of memory")
)

var (
	ErrMapValidationError         = errors.New("validation error")
	ErrMapDestroyedBeforeCallback = errors.New("destroyed before callback")
	ErrMapUnmappedBeforeCallback  = errors.New("unmapped before callback")
	ErrMapMappingAlreadyPending   = errors.New("mapping already pending")
	ErrMapOffsetOutOfRange        = errors.New("offset out of range")
	ErrMapSizeOutOfRange          = errors.New("size out of range")
)

var cstrings sync.Map

func getString(s string) *C.char {
	if c, ok := cstrings.Load(s); ok {
		return c.(*C.char)
	}
	c := C.CString(s)
	cstrings.Store(s, c)
	return c
}

type Device struct {
	hnd  C.WGPUDevice
	lost chan DeviceLost

	mapMu sync.Mutex
	// This may roll over on 32 bit, but if that results in a collision we'll
	// have other problems (such as a handle having stuck around for months).
	mapCounter uintptr
	mapHandles map[uintptr]struct {
		pinner runtime.Pinner
		ch     chan<- error
	}
}

type DeviceLost struct {
	Device  *Device
	Reason  DeviceLostReason
	Message string
}

type SurfaceCapabilities struct {
	Formats      []TextureFormat
	PresentModes []PresentMode
	AlphaModes   []CompositeAlphaMode
}

func (s *Surface) Capabilities(adapter *Adapter) SurfaceCapabilities {
	var caps C.WGPUSurfaceCapabilities
	C.wgpuSurfaceGetCapabilities(s.c(), adapter.c(), &caps)

	formats := cloneCArray[TextureFormat](caps.formats, caps.formatCount)
	presentModes := cloneCArray[PresentMode](caps.presentModes, caps.presentModeCount)
	alphaModes := cloneCArray[CompositeAlphaMode](caps.alphaModes, caps.alphaModeCount)
	C.wgpuSurfaceCapabilitiesFreeMembers(caps)
	return SurfaceCapabilities{
		Formats:      formats,
		PresentModes: presentModes,
		AlphaModes:   alphaModes,
	}
}

type instanceRequestAdapterResult struct {
	status  requestAdapterStatus
	adapter C.WGPUAdapter
	msg     string
	called  bool
}

type InstanceDescriptor struct {
	Extras *InstanceExtras
}

type AndroidNativeWindow struct {
	Window unsafe.Pointer
}

type MetalLayer struct {
	Layer unsafe.Pointer
}

type WaylandSurface struct {
	Label   string
	Display unsafe.Pointer
	Surface unsafe.Pointer
}

type WindowsHWND struct {
	HINSTANCE unsafe.Pointer
	HWND      unsafe.Pointer
}

type XCBWindow struct {
	Connection unsafe.Pointer
	Window     uint32
}

type XlibWindow struct {
	Display unsafe.Pointer
	Window  uint64
}

var _ NativeSurface = AndroidNativeWindow{}
var _ NativeSurface = MetalLayer{}
var _ NativeSurface = WaylandSurface{}
var _ NativeSurface = WindowsHWND{}
var _ NativeSurface = XCBWindow{}
var _ NativeSurface = XlibWindow{}

func (AndroidNativeWindow) nativeSurface() {}
func (MetalLayer) nativeSurface()          {}
func (WaylandSurface) nativeSurface()      {}
func (WindowsHWND) nativeSurface()         {}
func (XCBWindow) nativeSurface()           {}
func (XlibWindow) nativeSurface()          {}

// CloneCArray copies the elements of a dynamic C array (i.e. a pointer + a
// number of elements) into a newly allocated Go slice. It will convert from TC
// to TG in the process, requiring them to have the same size and memory layout.
func cloneCArray[TG, TC any, Int constraints.Integer](ptr *TC, n Int) []TG {
	if ptr == nil {
		return nil
	}
	if n == 0 {
		return []TG{}
	}
	if unsafe.Sizeof(*new(TC)) != unsafe.Sizeof(*new(TG)) {
		panic("TC and TG have different sizes")
	}

	out := make([]TG, n)
	in := unsafe.Slice(safeish.Cast[*TG](ptr), n)
	copy(out, in)
	return out
}

func calloc[T any]() *T {
	return (*T)(C.calloc(C.size_t(unsafe.Sizeof(*new(T))), 1))
}

func callocn[T any](n int) []T {
	ptr := (*T)(C.calloc(C.size_t(unsafe.Sizeof(*new(T))), C.size_t(n)))
	return unsafe.Slice(ptr, n)
}

func valloc[T any](v T) *T {
	ptr := (*T)(C.malloc(C.size_t(unsafe.Sizeof(*new(T)))))
	*ptr = v
	return ptr
}

func free[T any](ptr *T) {
	C.free(up(ptr))
}

func freen[E any, T ~[]E](ptr T) {
	free(unsafe.SliceData(ptr))
}

func maybeCall(fn func()) {
	if fn != nil {
		fn()
	}
}

func cstring(s string) *C.char {
	if s == "" {
		return nil
	}
	return C.CString(s)
}

func CreateInstance(desc InstanceDescriptor) *Instance {
	cdesc := calloc[C.WGPUInstanceDescriptor]()
	defer free(cdesc)
	if desc.Extras != nil {
		chain := calloc[C.WGPUInstanceExtras]()
		defer free(chain)
		chain.chain.sType = C.WGPUSType_InstanceExtras
		chain.backends = C.WGPUInstanceBackendFlags(desc.Extras.Backends)
		chain.flags = C.WGPUInstanceFlags(desc.Extras.Flags)
		chain.dx12ShaderCompiler = C.WGPUDx12Compiler(desc.Extras.DX12ShaderCompiler)
		chain.gles3MinorVersion = C.WGPUGles3MinorVersion(desc.Extras.GLES3MinorVersion)

		chain.dxilPath = cstring(desc.Extras.DXILPath)
		chain.dxcPath = cstring(desc.Extras.DXCPath)
		defer free(chain.dxilPath)
		defer free(chain.dxcPath)

		cdesc.nextInChain = &chain.chain
	}
	hnd := C.wgpuCreateInstance(cdesc)
	return safeish.Cast[*Instance](hnd)
}

// One of AndroidNativeWindow, MetalLayer, WaylandSurface, WindowsHWND,
// XCBWindow, or XlibWindow.
type NativeSurface interface {
	nativeSurface()
}

type SurfaceDescriptor struct {
	Label  string
	Native NativeSurface
}

func (ins *Instance) CreateSurface(desc SurfaceDescriptor) *Surface {
	cdesc := &C.WGPUSurfaceDescriptor{
		label: getString(desc.Label),
	}
	native := desc.Native
	switch native := native.(type) {
	case AndroidNativeWindow:
		chain := calloc[C.WGPUSurfaceDescriptorFromAndroidNativeWindow]()
		defer free(chain)
		chain.chain.sType = C.WGPUSType_SurfaceDescriptorFromAndroidNativeWindow
		chain.window = native.Window
		cdesc.nextInChain = &chain.chain
	case MetalLayer:
		chain := calloc[C.WGPUSurfaceDescriptorFromMetalLayer]()
		defer free(chain)
		chain.chain.sType = C.WGPUSType_SurfaceDescriptorFromMetalLayer
		chain.layer = native.Layer
		cdesc.nextInChain = &chain.chain
	case WaylandSurface:
		chain := calloc[C.WGPUSurfaceDescriptorFromWaylandSurface]()
		defer free(chain)
		chain.chain.sType = C.WGPUSType_SurfaceDescriptorFromWaylandSurface
		chain.display = native.Display
		chain.surface = native.Surface
		cdesc.nextInChain = &chain.chain
	case WindowsHWND:
		chain := calloc[C.WGPUSurfaceDescriptorFromWindowsHWND]()
		defer free(chain)
		chain.chain.sType = C.WGPUSType_SurfaceDescriptorFromWindowsHWND
		chain.hinstance = native.HINSTANCE
		chain.hwnd = native.HWND
		cdesc.nextInChain = &chain.chain
	case XCBWindow:
		chain := calloc[C.WGPUSurfaceDescriptorFromXcbWindow]()
		defer free(chain)
		chain.chain.sType = C.WGPUSType_SurfaceDescriptorFromXcbWindow
		chain.connection = native.Connection
		chain.window = C.uint32_t(native.Window)
		cdesc.nextInChain = &chain.chain
	case XlibWindow:
		chain := calloc[C.WGPUSurfaceDescriptorFromXlibWindow]()
		defer free(chain)
		chain.chain.sType = C.WGPUSType_SurfaceDescriptorFromXlibWindow
		chain.display = native.Display
		chain.window = C.uint64_t(native.Window)
		cdesc.nextInChain = &chain.chain
	case nil:
	default:
		panic(fmt.Sprintf("unhandled type %T", native))
	}
	hnd := C.wgpuInstanceCreateSurface(ins.c(), cdesc)
	return safeish.Cast[*Surface](hnd)
}

type RequestAdapterOptions struct {
	CompatibleSurface    *Surface
	PowerPreference      PowerPreference
	ForceFallbackAdapter bool

	// wgpu doesn't support BackendType and requires the use of InstanceExtras.Backends
	// with InstanceDescriptor instead.
}

func toBool(b bool) C.WGPUBool {
	if b {
		return 1
	} else {
		return 0
	}
}

func (ins *Instance) RequestAdapter(desc RequestAdapterOptions) (*Adapter, error) {
	var ret instanceRequestAdapterResult
	cdesc := valloc(C.WGPURequestAdapterOptions{
		compatibleSurface:    desc.CompatibleSurface.c(),
		powerPreference:      C.WGPUPowerPreference(desc.PowerPreference),
		backendType:          C.WGPUBackendType_Undefined,
		forceFallbackAdapter: toBool(desc.ForceFallbackAdapter),
	})
	defer free(cdesc)

	C.wgpuInstanceRequestAdapter(ins.c(), cdesc, fp(C.requestAdapterCallback), up(&ret))
	if !ret.called {
		panic("callback wasn't called")
	}
	switch ret.status {
	case requestAdapterStatusSuccess:
		return safeish.Cast[*Adapter](ret.adapter), nil
	case requestAdapterStatusUnavailable:
		return nil, ErrAdapterUnavailable
	default:
		return nil, RequestAdapterError{Message: ret.msg}
	}
}

//export requestAdapterCallback
func requestAdapterCallback(
	status C.WGPURequestAdapterStatus,
	adapter C.WGPUAdapter,
	msg *C.char,
	data up,
) {
	ret := (*instanceRequestAdapterResult)(data)
	ret.status = requestAdapterStatus(status)
	ret.adapter = adapter
	ret.msg = C.GoString(msg)
	ret.called = true
}

//export requestDeviceCallback
func requestDeviceCallback(
	status C.WGPURequestDeviceStatus,
	device C.WGPUDevice,
	msg *C.char,
	data up,
) {
	ret := (*adapterRequestDeviceResult)(data)
	ret.status = requestDeviceStatus(status)
	ret.device = device
	ret.msg = C.GoString(msg)
	ret.called = true
}

func (adp *Adapter) Features() []FeatureName {
	n := C.wgpuAdapterEnumerateFeatures(adp.c(), nil)
	if n <= 0 {
		return nil
	}
	names := make([]C.WGPUFeatureName, n)
	C.wgpuAdapterEnumerateFeatures(adp.c(), &names[0])
	return *(*[]FeatureName)(up(&names))
}

type AdapterProperties struct {
	VendorID          uint32
	VendorName        string
	Architecture      string
	DeviceID          uint32
	Name              string
	DriverDescription string
	AdapterType       AdapterType
	BackendType       BackendType
}

func (adp *Adapter) Properties() AdapterProperties {
	var cprops C.struct_WGPUAdapterProperties
	C.wgpuAdapterGetProperties(adp.c(), &cprops)
	return AdapterProperties{
		VendorID:          uint32(cprops.vendorID),
		VendorName:        C.GoString(cprops.vendorName),
		Architecture:      C.GoString(cprops.architecture),
		DeviceID:          uint32(cprops.deviceID),
		Name:              C.GoString(cprops.name),
		DriverDescription: C.GoString(cprops.driverDescription),
		AdapterType:       AdapterType(cprops.adapterType),
		BackendType:       BackendType(cprops.backendType),
	}
}

var DefaultLimits = Limits{
	// From https://www.w3.org/TR/webgpu/#limit-default

	MaxTextureDimension1D:                     8192,
	MaxTextureDimension2D:                     8192,
	MaxTextureDimension3D:                     2048,
	MaxTextureArrayLayers:                     256,
	MaxBindGroups:                             4,
	MaxBindGroupsPlusVertexBuffers:            24,
	MaxBindingsPerBindGroup:                   1000,
	MaxDynamicUniformBuffersPerPipelineLayout: 8,
	MaxDynamicStorageBuffersPerPipelineLayout: 4,
	MaxSampledTexturesPerShaderStage:          16,
	MaxSamplersPerShaderStage:                 16,
	MaxStorageBuffersPerShaderStage:           8,
	MaxStorageTexturesPerShaderStage:          4,
	MaxUniformBuffersPerShaderStage:           12,
	MaxUniformBufferBindingSize:               65536,
	MaxStorageBufferBindingSize:               134217728,
	MinUniformBufferOffsetAlignment:           256,
	MinStorageBufferOffsetAlignment:           256,
	MaxVertexBuffers:                          8,
	MaxBufferSize:                             268435456,
	MaxVertexAttributes:                       16,
	MaxVertexBufferArrayStride:                2048,
	MaxInterStageShaderComponents:             64,
	MaxInterStageShaderVariables:              16,
	MaxColorAttachments:                       8,
	MaxColorAttachmentBytesPerSample:          32,
	MaxComputeWorkgroupStorageSize:            16384,
	MaxComputeInvocationsPerWorkgroup:         256,
	MaxComputeWorkgroupSizeX:                  256,
	MaxComputeWorkgroupSizeY:                  256,
	MaxComputeWorkgroupSizeZ:                  64,
	MaxComputeWorkgroupsPerDimension:          65535,
}

type Limits struct {
	_ structs.HostLayout

	MaxTextureDimension1D                     uint32
	MaxTextureDimension2D                     uint32
	MaxTextureDimension3D                     uint32
	MaxTextureArrayLayers                     uint32
	MaxBindGroups                             uint32
	MaxBindGroupsPlusVertexBuffers            uint32
	MaxBindingsPerBindGroup                   uint32
	MaxDynamicUniformBuffersPerPipelineLayout uint32
	MaxDynamicStorageBuffersPerPipelineLayout uint32
	MaxSampledTexturesPerShaderStage          uint32
	MaxSamplersPerShaderStage                 uint32
	MaxStorageBuffersPerShaderStage           uint32
	MaxStorageTexturesPerShaderStage          uint32
	MaxUniformBuffersPerShaderStage           uint32
	MaxUniformBufferBindingSize               uint64
	MaxStorageBufferBindingSize               uint64
	MinUniformBufferOffsetAlignment           uint32
	MinStorageBufferOffsetAlignment           uint32
	MaxVertexBuffers                          uint32
	MaxBufferSize                             uint64
	MaxVertexAttributes                       uint32
	MaxVertexBufferArrayStride                uint32
	MaxInterStageShaderComponents             uint32
	MaxInterStageShaderVariables              uint32
	MaxColorAttachments                       uint32
	MaxColorAttachmentBytesPerSample          uint32
	MaxComputeWorkgroupStorageSize            uint32
	MaxComputeInvocationsPerWorkgroup         uint32
	MaxComputeWorkgroupSizeX                  uint32
	MaxComputeWorkgroupSizeY                  uint32
	MaxComputeWorkgroupSizeZ                  uint32
	MaxComputeWorkgroupsPerDimension          uint32
}

var DefaultNativeLimits = NativeLimits{
	MaxPushConstantSize:   ^uint32(0),
	MaxNonSamplerBindings: ^uint32(0),
}

type NativeLimits struct {
	_ structs.HostLayout

	MaxPushConstantSize   uint32
	MaxNonSamplerBindings uint32
}

type AllLimits struct {
	Limits
	NativeLimits
}

func (adp *Adapter) Limits() AllLimits {
	var climits C.struct_WGPUSupportedLimits
	chain := calloc[C.WGPUSupportedLimitsExtras]()
	defer free(chain)
	chain.chain.sType = C.WGPUSType_SupportedLimitsExtras
	climits.nextInChain = &chain.chain
	if C.wgpuAdapterGetLimits(adp.c(), &climits) == 0 {
		panic("wgpuAdapterGetLimits call failed unexpectedly")
	}
	return AllLimits{
		Limits:       safeish.Cast[Limits](climits.limits),
		NativeLimits: safeish.Cast[NativeLimits](chain.limits),
	}
}

type DeviceDescriptor struct {
	Label            string
	RequiredFeatures []FeatureName
	RequiredLimits   *RequiredLimits
	DefaultQueue     QueueDescriptor
}

type adapterRequestDeviceResult struct {
	status requestDeviceStatus
	device C.WGPUDevice
	msg    string
	called bool
}

//export deviceLostCallback
func deviceLostCallback(reason C.WGPUDeviceLostReason, message *C.char, hnd uintptr) {
	handle := cgo.Handle(hnd)
	dev := handle.Value().(*Device)

	// The device lost callback only gets called once, and it gets called even
	// if the user explicitly destroys the device. Ergo we can be sure we won't
	// see this handle again.
	handle.Delete()
	dev.lost <- DeviceLost{
		Device:  dev,
		Reason:  DeviceLostReason(reason),
		Message: C.GoString(message),
	}
	close(dev.lost)
}

type RequiredLimits struct {
	Limits       Limits
	NativeLimits NativeLimits
}

type QueueDescriptor struct {
	Label string
}

func (adp *Adapter) RequestDevice(desc *DeviceDescriptor) (*Device, error) {
	dev := &Device{
		mapHandles: make(map[uintptr]struct {
			pinner runtime.Pinner
			ch     chan<- error
		}),
		lost: make(chan DeviceLost, 1),
	}
	var ret adapterRequestDeviceResult
	cdesc := &C.WGPUDeviceDescriptor{
		deviceLostUserdata: C.uintptr_t(cgo.NewHandle(dev)),
		deviceLostCallback: safeish.Cast[C.WGPUDeviceLostCallback](deviceLostCallback),
	}
	if desc != nil {
		cdesc.requiredFeatureCount = C.size_t(len(desc.RequiredFeatures))
		features := unsafe.SliceData(desc.RequiredFeatures)
		var pinner runtime.Pinner
		defer pinner.Unpin()
		pinner.Pin(features)
		cdesc.requiredFeatures = safeish.Cast[*C.WGPUFeatureName](features)
		if desc.Label != "" {
			cdesc.label = getString(desc.Label)
		}
		if desc.RequiredLimits != nil {
			climits := calloc[C.WGPURequiredLimits]()
			defer free(climits)
			cdesc.requiredLimits = climits

			chain := calloc[C.WGPURequiredLimitsExtras]()
			defer free(chain)
			chain.chain.sType = C.WGPUSType_RequiredLimitsExtras
			chain.limits = safeish.Cast[C.WGPUNativeLimits](desc.RequiredLimits.NativeLimits)
			climits.nextInChain = &chain.chain
			climits.limits = safeish.Cast[C.WGPULimits](desc.RequiredLimits.Limits)
		}

		// TODO(dh): we would support WGPUDeviceExtras.tracePath here, but the
		// release builds of wgpu-native don't have the trace feature enabled.
	}

	C.wgpuAdapterRequestDevice(adp.c(), cdesc, fp(C.requestDeviceCallback), up(&ret))
	if !ret.called {
		panic("callback wasn't called")
	}
	if ret.status != requestDeviceStatusSuccess {
		return nil, RequestDeviceError{Message: ret.msg}
	}
	dev.hnd = ret.device
	C.wgpuDeviceSetUncapturedErrorCallback(ret.device, fp(C.uncapturedErrorCallback), nil)
	return dev, nil
}

// Destroy forcibly destroys the device handle, invalidating any resources that
// reference it, such as bind groups.
//
// Note: it doesn't currently do anything because wgpu doesn't implement it yet
func (dev *Device) Destroy() {
	C.wgpuDeviceDestroy(dev.hnd)
}

// Lost returns a channel that will be signaled when the device is lost.
// Repeated calls to Lost return the same channel.
func (dev *Device) Lost() <-chan DeviceLost {
	return dev.lost
}

// Can be chained in WGPUShaderModuleDescriptor
// typedef struct WGPUShaderModuleSPIRVDescriptor {
//     WGPUChainedStruct chain;
//     uint32_t codeSize;
//     uint32_t const * code;
// } WGPUShaderModuleSPIRVDescriptor WGPU_STRUCTURE_ATTRIBUTE;

// typedef struct WGPUShaderModuleGLSLDescriptor {
//     WGPUChainedStruct chain;
//     WGPUShaderStage stage;
//     char const * code;
//     uint32_t defineCount;
//     WGPUShaderDefine * defines;
// } WGPUShaderModuleGLSLDescriptor;

var _ ShaderSource = ShaderSourceWGSL("")

func (ShaderSourceWGSL) shaderSource() {}

type ShaderSourceWGSL []byte

type ShaderModuleCompilationHint struct {
	EntryPoint string
	Layout     *PipelineLayout
}

// One of ShaderSourceWGSL.
type ShaderSource interface {
	shaderSource()
}

type ShaderModuleDescriptor struct {
	Label  string
	Hints  []ShaderModuleCompilationHint
	Source ShaderSource
}

func (dev *Device) PushErrorScope(filter ErrorFilter) {
	C.wgpuDevicePushErrorScope(dev.hnd, C.WGPUErrorFilter(filter))
}

func (dev *Device) PopErrorScope() error {
	ret := calloc[popErrorScopeResult]()
	defer free(ret)

	C.wgpuDevicePopErrorScope(dev.hnd, fp(C.popErrorScopeCallback), up(ret))
	// wgpu generates quite verbose errors that span multiple lines and that
	// begin with the kind of error. Therefore we don't have to do anything but
	// return the messages.
	switch ret.typ {
	case C.WGPUErrorType_NoError:
		return nil
	case C.WGPUErrorType_Validation:
		return ValidationError{C.GoString(ret.msg)}
	case C.WGPUErrorType_OutOfMemory:
		return OutOfMemoryError{C.GoString(ret.msg)}
	case C.WGPUErrorType_Internal:
		return InternalError{C.GoString(ret.msg)}
	case C.WGPUErrorType_Unknown:
		return UnknownError{C.GoString(ret.msg)}
	case C.WGPUErrorType_DeviceLost:
		// As we understand it, device lost should never be returned by
		// PopErrorScope and should only trigger the device lost callback.
		panic("unexpected 'device lost' error")
	default:
		return UnknownError{C.GoString(ret.msg)}
	}
}

type popErrorScopeResult struct {
	typ C.WGPUErrorType
	msg *C.char
}

//export popErrorScopeCallback
func popErrorScopeCallback(typ C.WGPUErrorType, msg *C.char, data unsafe.Pointer) {
	res := (*popErrorScopeResult)(data)
	res.typ = typ
	res.msg = msg
}

type PipelineLayoutDescriptor struct {
	Label            string
	BindGroupLayouts []*BindGroupLayout
}

func (dev *Device) CreatePipelineLayout(desc *PipelineLayoutDescriptor) *PipelineLayout {
	cdesc := calloc[C.WGPUPipelineLayoutDescriptor]()
	defer free(cdesc)
	if desc != nil {
		var pinner runtime.Pinner
		defer pinner.Unpin()
		cdesc.label = getString(desc.Label)
		cdesc.bindGroupLayoutCount = C.size_t(len(desc.BindGroupLayouts))
		layouts := unsafe.SliceData(desc.BindGroupLayouts)
		pinner.Pin(layouts)
		cdesc.bindGroupLayouts = safeish.Cast[*C.WGPUBindGroupLayout](layouts)
	}
	hnd := C.wgpuDeviceCreatePipelineLayout(dev.hnd, cdesc)
	return safeish.Cast[*PipelineLayout](hnd)
}

func (dev *Device) CreateShaderModule(desc ShaderModuleDescriptor) *ShaderModule {
	cdesc := &C.WGPUShaderModuleDescriptor{
		label: getString(desc.Label),
	}

	if len(desc.Hints) > 0 {
		cdesc.hintCount = C.size_t(len(desc.Hints))
		chints := make([]C.WGPUShaderModuleCompilationHint, len(desc.Hints))
		for i, hint := range desc.Hints {
			chints[i].entryPoint = getString(hint.EntryPoint)
			if hint.Layout != nil {
				chints[i].layout = hint.Layout.c()
			}
		}
	}

	switch source := desc.Source.(type) {
	case ShaderSourceWGSL:
		chain := calloc[C.WGPUShaderModuleWGSLDescriptor]()
		defer free(chain)
		chain.chain.sType = C.WGPUSType_ShaderModuleWGSLDescriptor
		code := C.CString(string(source))
		defer free(code)
		chain.code = code
		cdesc.nextInChain = &chain.chain
	default:
		panic(fmt.Sprintf("unhandled type %T", source))
	}

	hnd := C.wgpuDeviceCreateShaderModule(dev.hnd, cdesc)
	return safeish.Cast[*ShaderModule](hnd)
}

func (dev *Device) Features() []FeatureName {
	n := C.wgpuDeviceEnumerateFeatures(dev.hnd, nil)
	if n <= 0 {
		return nil
	}
	names := make([]C.WGPUFeatureName, n)
	C.wgpuDeviceEnumerateFeatures(dev.hnd, &names[0])
	return safeish.SliceCast[[]FeatureName](names)
}

func (dev *Device) Limits() AllLimits {
	var climits C.struct_WGPUSupportedLimits
	chain := calloc[C.WGPUSupportedLimitsExtras]()
	defer free(chain)
	chain.chain.sType = C.WGPUSType_SupportedLimitsExtras
	climits.nextInChain = &chain.chain
	if C.wgpuDeviceGetLimits(dev.hnd, &climits) == 0 {
		panic("wgpuDeviceGetLimits call failed unexpectedly")
	}
	return AllLimits{
		Limits:       safeish.Cast[Limits](climits.limits),
		NativeLimits: safeish.Cast[NativeLimits](chain.limits),
	}
}

func (dev *Device) HasFeature(name FeatureName) bool {
	return C.wgpuDeviceHasFeature(dev.hnd, C.WGPUFeatureName(name)) != 0
}

func (adp *Adapter) HasFeature(name FeatureName) bool {
	return C.wgpuAdapterHasFeature(adp.c(), C.WGPUFeatureName(name)) != 0
}

func (dev *Device) Queue() *Queue {
	hnd := C.wgpuDeviceGetQueue(dev.hnd)
	return safeish.Cast[*Queue](hnd)
}

type CommandEncoderDescriptor struct {
	Label string
}

func (dev *Device) CreateCommandEncoder(desc *CommandEncoderDescriptor) *CommandEncoder {
	cdesc := calloc[C.WGPUCommandEncoderDescriptor]()
	defer free(cdesc)
	if desc != nil {
		cdesc.label = getString(desc.Label)
	}
	hnd := C.wgpuDeviceCreateCommandEncoder(dev.hnd, cdesc)
	return safeish.Cast[*CommandEncoder](hnd)
}

type ConstantEntry struct {
	Key   string
	Value float64
}

type VertexAttribute struct {
	Format         VertexFormat
	Offset         uint64
	ShaderLocation uint32
}

type VertexBufferLayout struct {
	ArrayStride uint64
	StepMode    VertexStepMode
	Attributes  []VertexAttribute
}

type VertexState struct {
	Module     *ShaderModule
	EntryPoint string
	Constants  []ConstantEntry
	Buffers    []VertexBufferLayout
}

type StencilFaceState struct {
	_ structs.HostLayout

	Compare     CompareFunction
	FailOp      StencilOperation
	DepthFailOp StencilOperation
	PassOp      StencilOperation
}

type DepthStencilState struct {
	Format              TextureFormat
	DepthWriteEnabled   bool
	DepthCompare        CompareFunction
	StencilFront        StencilFaceState
	StencilBack         StencilFaceState
	StencilReadMask     uint32
	StencilWriteMask    uint32
	DepthBias           int32
	DepthBiasSlopeScale float32
	DepthBiasClamp      float32
}

var DefaultMultisampleState = &MultisampleState{
	Count:                  1,
	Mask:                   0xFFFFFFFF,
	AlphaToCoverageEnabled: false,
}

type MultisampleState struct {
	Count                  uint32
	Mask                   uint32
	AlphaToCoverageEnabled bool
}

type BlendComponent struct {
	Operation BlendOperation
	SrcFactor BlendFactor
	DstFactor BlendFactor
}

type BlendState struct {
	_ structs.HostLayout

	Color BlendComponent
	Alpha BlendComponent
}

type ColorTargetState struct {
	Format    TextureFormat
	Blend     *BlendState
	WriteMask ColorWriteMask
}

type FragmentState struct {
	Module     *ShaderModule
	EntryPoint string
	Constants  []ConstantEntry
	Targets    []ColorTargetState
}

type PrimitiveDepthClipControl struct {
	UnclippedDepth bool
}

var DefaultPrimitiveState = &PrimitiveState{
	Topology:         PrimitiveTopologyTriangleList,
	StripIndexFormat: IndexFormatUndefined,
	FrontFace:        FrontFaceCCW,
	CullMode:         CullModeNone,
}

type PrimitiveState struct {
	Topology         PrimitiveTopology
	StripIndexFormat IndexFormat
	FrontFace        FrontFace
	CullMode         CullMode

	DepthClipControl *PrimitiveDepthClipControl
}

type RenderPipelineDescriptor struct {
	Label        string          // optional
	Layout       *PipelineLayout // optional
	Vertex       *VertexState
	Primitive    *PrimitiveState
	DepthStencil *DepthStencilState // optional
	Multisample  *MultisampleState
	Fragment     *FragmentState // optional
}

func (dev *Device) CreateRenderPipeline(desc *RenderPipelineDescriptor) *RenderPipeline {
	var pinner runtime.Pinner
	defer pinner.Unpin()

	cdesc := calloc[C.WGPURenderPipelineDescriptor]()
	defer free(cdesc)
	cdesc.label = getString(desc.Label)
	if desc.Layout != nil {
		cdesc.layout = desc.Layout.c()
	}

	{
		cvstate := &cdesc.vertex
		cvstate.module = desc.Vertex.Module.c()
		cvstate.entryPoint = getString(desc.Vertex.EntryPoint)
		cvstate.constantCount = C.size_t(len(desc.Vertex.Constants))
		constants := callocn[C.WGPUConstantEntry](len(desc.Vertex.Constants))
		defer freen(constants)
		cvstate.constants = unsafe.SliceData(constants)
		for i, c := range desc.Vertex.Constants {
			constants[i] = C.WGPUConstantEntry{
				key:   getString(c.Key),
				value: C.double(c.Value),
			}
		}
		cvstate.bufferCount = C.size_t(len(desc.Vertex.Buffers))
		buffers := callocn[C.WGPUVertexBufferLayout](len(desc.Vertex.Buffers))
		defer freen(buffers)
		cvstate.buffers = unsafe.SliceData(buffers)
		for i, b := range desc.Vertex.Buffers {
			attributes := unsafe.SliceData(b.Attributes)
			pinner.Pin(attributes)
			buffers[i] = C.WGPUVertexBufferLayout{
				arrayStride:    C.uint64_t(b.ArrayStride),
				stepMode:       C.WGPUVertexStepMode(b.StepMode),
				attributeCount: C.size_t(len(b.Attributes)),
				attributes:     safeish.Cast[*C.WGPUVertexAttribute](attributes),
			}
		}
	}

	primitive := desc.Primitive
	if primitive == nil {
		primitive = DefaultPrimitiveState
	}
	cdesc.primitive = C.WGPUPrimitiveState{
		topology:         C.WGPUPrimitiveTopology(primitive.Topology),
		stripIndexFormat: C.WGPUIndexFormat(primitive.StripIndexFormat),
		frontFace:        C.WGPUFrontFace(primitive.FrontFace),
		cullMode:         C.WGPUCullMode(primitive.CullMode),
	}
	if primitive.DepthClipControl != nil {
		chain := calloc[C.WGPUPrimitiveDepthClipControl]()
		defer free(chain)
		chain.chain.sType = C.WGPUSType_PrimitiveDepthClipControl
		chain.unclippedDepth = toBool(primitive.DepthClipControl.UnclippedDepth)
		cdesc.nextInChain = &chain.chain
	}

	if desc.DepthStencil != nil {
		cdepth := calloc[C.WGPUDepthStencilState]()
		defer free(cdepth)
		cdesc.depthStencil = cdepth
		*cdepth = C.WGPUDepthStencilState{
			format:              C.WGPUTextureFormat(desc.DepthStencil.Format),
			depthWriteEnabled:   toBool(desc.DepthStencil.DepthWriteEnabled),
			depthCompare:        C.WGPUCompareFunction(desc.DepthStencil.DepthCompare),
			stencilFront:        safeish.Cast[C.WGPUStencilFaceState](desc.DepthStencil.StencilFront),
			stencilBack:         safeish.Cast[C.WGPUStencilFaceState](desc.DepthStencil.StencilBack),
			stencilReadMask:     C.uint32_t(desc.DepthStencil.StencilReadMask),
			stencilWriteMask:    C.uint32_t(desc.DepthStencil.StencilWriteMask),
			depthBias:           C.int32_t(desc.DepthStencil.DepthBias),
			depthBiasSlopeScale: C.float(desc.DepthStencil.DepthBiasSlopeScale),
			depthBiasClamp:      C.float(desc.DepthStencil.DepthBiasClamp),
		}
	}

	multisample := desc.Multisample
	if multisample == nil {
		multisample = DefaultMultisampleState
	}
	cdesc.multisample = C.WGPUMultisampleState{
		count:                  C.uint32_t(multisample.Count),
		mask:                   C.uint32_t(multisample.Mask),
		alphaToCoverageEnabled: toBool(multisample.AlphaToCoverageEnabled),
	}

	if desc.Fragment != nil {
		cfrag := calloc[C.WGPUFragmentState]()
		defer free(cfrag)
		cdesc.fragment = cfrag

		cfrag.module = desc.Fragment.Module.c()
		cfrag.entryPoint = getString(desc.Fragment.EntryPoint)
		cfrag.constantCount = C.size_t(len(desc.Fragment.Constants))

		constants := callocn[C.WGPUConstantEntry](len(desc.Fragment.Constants))
		defer freen(constants)
		cfrag.constants = unsafe.SliceData(constants)
		for i, c := range desc.Fragment.Constants {
			constants[i] = C.WGPUConstantEntry{
				key:   getString(c.Key),
				value: C.double(c.Value),
			}
		}

		cfrag.targetCount = C.size_t(len(desc.Fragment.Targets))
		ctargets := callocn[C.WGPUColorTargetState](len(desc.Fragment.Targets))
		defer freen(ctargets)
		cfrag.targets = unsafe.SliceData(ctargets)
		for i, t := range desc.Fragment.Targets {
			ctargets[i].format = C.WGPUTextureFormat(t.Format)
			ctargets[i].writeMask = C.WGPUColorWriteMaskFlags(t.WriteMask)
			if t.Blend != nil {
				cblend := calloc[C.WGPUBlendState]()
				defer free(cblend)
				ctargets[i].blend = cblend
				*cblend = *safeish.Cast[*C.WGPUBlendState](t.Blend)
			}
		}
	}

	hnd := C.wgpuDeviceCreateRenderPipeline(dev.hnd, cdesc)
	return safeish.Cast[*RenderPipeline](hnd)
}

type QuerySetDescriptor struct {
	Label string
	Type  QueryType
	Count uint32
}

func (dev *Device) CreateQuerySet(desc *QuerySetDescriptor) *QuerySet {
	cdesc := valloc(C.WGPUQuerySetDescriptor{
		label: getString(desc.Label),
		_type: C.WGPUQueryType(desc.Type),
		count: C.uint32_t(desc.Count),
	})
	defer free(cdesc)
	hnd := C.wgpuDeviceCreateQuerySet(dev.hnd, cdesc)
	return safeish.Cast[*QuerySet](hnd)
}

type SurfaceConfiguration struct {
	Format      TextureFormat
	Usage       TextureUsage
	ViewFormats []TextureFormat
	AlphaMode   CompositeAlphaMode
	Width       uint32
	Height      uint32
	PresentMode PresentMode

	DesiredMaximumFrameLatency uint32
}

func (s *Surface) Configure(dev *Device, config *SurfaceConfiguration) {
	var pinner runtime.Pinner
	defer pinner.Unpin()
	viewFormats := unsafe.SliceData(config.ViewFormats)
	pinner.Pin(viewFormats)
	cdesc := valloc(C.WGPUSurfaceConfiguration{
		device:          dev.hnd,
		format:          C.WGPUTextureFormat(config.Format),
		usage:           C.WGPUTextureUsageFlags(config.Usage),
		viewFormatCount: C.size_t(len(config.ViewFormats)),
		viewFormats:     safeish.Cast[*C.WGPUTextureFormat](viewFormats),
		alphaMode:       C.WGPUCompositeAlphaMode(config.AlphaMode),
		width:           C.uint32_t(config.Width),
		height:          C.uint32_t(config.Height),
		presentMode:     C.WGPUPresentMode(config.PresentMode),
	})
	defer free(cdesc)

	if config.DesiredMaximumFrameLatency != 0 {
		cextra := calloc[C.WGPUSurfaceConfigurationExtras]()
		defer free(cextra)
		cextra.chain.sType = C.WGPUSType_SurfaceConfigurationExtras
		cextra.desiredMaximumFrameLatency = C.WGPUBool(config.DesiredMaximumFrameLatency)
		cdesc.nextInChain = &cextra.chain
	}

	C.wgpuSurfaceConfigure(s.c(), cdesc)
}

func (s *Surface) Unconfigure() {
	C.wgpuSurfaceUnconfigure(s.c())
}

func (s *Surface) PreferredFormat(adapter *Adapter) TextureFormat {
	return TextureFormat(C.wgpuSurfaceGetPreferredFormat(s.c(), adapter.c()))
}

type SurfaceTexture struct {
	Texture    *Texture
	Suboptimal bool
}

func (s *Surface) CurrentTexture() (SurfaceTexture, error) {
	ret := calloc[C.WGPUSurfaceTexture]()
	defer free(ret)
	// FIXME(dh): there can be validation errors that aren't returned in the
	// status, but we don't have access to the device to push an error scope.
	C.wgpuSurfaceGetCurrentTexture(s.c(), ret)
	switch ret.status {
	case C.WGPUSurfaceGetCurrentTextureStatus_Success:
		return SurfaceTexture{
			Texture:    safeish.Cast[*Texture](ret.texture),
			Suboptimal: ret.suboptimal != 0,
		}, nil
	case C.WGPUSurfaceGetCurrentTextureStatus_Timeout:
		return SurfaceTexture{}, ErrCurrentTextureTimeout
	case C.WGPUSurfaceGetCurrentTextureStatus_Outdated:
		return SurfaceTexture{}, ErrCurrentTextureOutdated
	case C.WGPUSurfaceGetCurrentTextureStatus_Lost:
		return SurfaceTexture{}, ErrCurrentTextureLost
	case C.WGPUSurfaceGetCurrentTextureStatus_OutOfMemory:
		return SurfaceTexture{}, ErrCurrentTextureOutOfMemory
	case C.WGPUSurfaceGetCurrentTextureStatus_DeviceLost:
		return SurfaceTexture{}, ErrDeviceLost
	default:
		panic(fmt.Sprintf("invalid status %d", ret.status))
	}
}

func (s *Surface) Present() {
	C.wgpuSurfacePresent(s.c())
}

type TextureViewDescriptor struct {
	Label           string
	Format          TextureFormat
	Dimension       TextureViewDimension
	BaseMipLevel    uint32
	MipLevelCount   uint32
	BaseArrayLayer  uint32
	ArrayLayerCount uint32
	Aspect          TextureAspect
}

var DefaultTextureViewDescriptor = &TextureViewDescriptor{
	ArrayLayerCount: ^uint32(0),
	MipLevelCount:   ^uint32(0),
}

func (tex *Texture) CreateView(desc *TextureViewDescriptor) *TextureView {
	if desc == nil {
		desc = DefaultTextureViewDescriptor
	}
	cdesc := valloc(C.WGPUTextureViewDescriptor{
		label:           getString(desc.Label),
		format:          C.WGPUTextureFormat(desc.Format),
		dimension:       C.WGPUTextureViewDimension(desc.Dimension),
		baseMipLevel:    C.uint32_t(desc.BaseMipLevel),
		mipLevelCount:   C.uint32_t(desc.MipLevelCount),
		baseArrayLayer:  C.uint32_t(desc.BaseArrayLayer),
		arrayLayerCount: C.uint32_t(desc.ArrayLayerCount),
		aspect:          C.WGPUTextureAspect(desc.Aspect),
	})
	defer free(cdesc)
	hnd := C.wgpuTextureCreateView(tex.c(), cdesc)
	return safeish.Cast[*TextureView](hnd)
}

func (tex *Texture) Format() TextureFormat {
	return TextureFormat(C.wgpuTextureGetFormat(tex.c()))
}

func (enc *CommandEncoder) InsertDebugMarker(label string) {
	clabel := getString(label)
	C.wgpuCommandEncoderInsertDebugMarker(enc.c(), clabel)
}

type CommandBufferDescriptor struct {
	Label string
}

func (enc *CommandEncoder) Finish(desc *CommandBufferDescriptor) *CommandBuffer {
	cdesc := valloc(C.WGPUCommandBufferDescriptor{})
	defer free(cdesc)
	if desc != nil {
		cdesc.label = getString(desc.Label)
	}
	hnd := C.wgpuCommandEncoderFinish(enc.c(), cdesc)
	return safeish.Cast[*CommandBuffer](hnd)
}

type RenderPassDescriptor struct {
	Label                  string
	ColorAttachments       []RenderPassColorAttachment
	DepthStencilAttachment *RenderPassDepthStencilAttachment // nullable
	OcclusionQuerySet      *QuerySet                         // nullable
	TimestampWrites        *RenderPassTimestampWrites        //nullable
	// TODO(dh): support WGPURenderPassDescriptorMaxDrawCount
}

type RenderPassTimestampWrites struct {
	QuerySet                  *QuerySet
	BeginningOfPassWriteIndex uint32
	EndOfPassWriteIndex       uint32
}

type RenderPassDepthStencilAttachment struct {
	View              *TextureView
	DepthLoadOp       LoadOp
	DepthStoreOp      StoreOp
	DepthClearValue   float32
	DepthReadOnly     bool
	StencilLoadOp     LoadOp
	StencilStoreOp    StoreOp
	StencilClearValue uint32
	StencilReadOnly   bool
}

type RenderPassColorAttachment struct {
	View          *TextureView // nullable
	ResolveTarget *TextureView // nullable
	LoadOp        LoadOp
	StoreOp       StoreOp
	ClearValue    Color
}

func (enc *CommandEncoder) BeginRenderPass(desc *RenderPassDescriptor) *RenderPassEncoder {
	cattachments := callocn[C.WGPURenderPassColorAttachment](len(desc.ColorAttachments))
	defer freen(cattachments)
	cdesc := valloc(C.WGPURenderPassDescriptor{
		label:                getString(desc.Label),
		colorAttachmentCount: C.size_t(len(desc.ColorAttachments)),
		colorAttachments:     unsafe.SliceData(cattachments),
	})
	defer free(cdesc)

	for i, a := range desc.ColorAttachments {
		var viewHnd, resolveHnd C.WGPUTextureView
		if a.View != nil {
			viewHnd = a.View.c()
		}
		if a.ResolveTarget != nil {
			resolveHnd = a.ResolveTarget.c()
		}
		cattachments[i] = C.WGPURenderPassColorAttachment{
			view:          viewHnd,
			resolveTarget: resolveHnd,
			loadOp:        C.WGPULoadOp(a.LoadOp),
			storeOp:       C.WGPUStoreOp(a.StoreOp),
			clearValue:    safeish.Cast[C.WGPUColor](a.ClearValue),
		}
	}

	if desc.DepthStencilAttachment != nil {
		cdepth := calloc[C.WGPURenderPassDepthStencilAttachment]()
		defer free(cdepth)
		cdesc.depthStencilAttachment = cdepth
		*cdepth = C.WGPURenderPassDepthStencilAttachment{
			view:              desc.DepthStencilAttachment.View.c(),
			depthLoadOp:       C.WGPULoadOp(desc.DepthStencilAttachment.DepthLoadOp),
			depthStoreOp:      C.WGPUStoreOp(desc.DepthStencilAttachment.DepthStoreOp),
			depthClearValue:   C.float(desc.DepthStencilAttachment.DepthClearValue),
			depthReadOnly:     toBool(desc.DepthStencilAttachment.DepthReadOnly),
			stencilLoadOp:     C.WGPULoadOp(desc.DepthStencilAttachment.StencilLoadOp),
			stencilStoreOp:    C.WGPUStoreOp(desc.DepthStencilAttachment.StencilStoreOp),
			stencilClearValue: C.uint32_t(desc.DepthStencilAttachment.StencilClearValue),
			stencilReadOnly:   toBool(desc.DepthStencilAttachment.StencilReadOnly),
		}
	}
	if desc.OcclusionQuerySet != nil {
		cdesc.occlusionQuerySet = desc.OcclusionQuerySet.c()
	}
	if desc.TimestampWrites != nil {
		cts := calloc[C.WGPURenderPassTimestampWrites]()
		defer free(cts)
		cdesc.timestampWrites = cts
		*cts = C.WGPURenderPassTimestampWrites{
			querySet:                  desc.TimestampWrites.QuerySet.c(),
			beginningOfPassWriteIndex: C.uint32_t(desc.TimestampWrites.BeginningOfPassWriteIndex),
			endOfPassWriteIndex:       C.uint32_t(desc.TimestampWrites.EndOfPassWriteIndex),
		}
	}

	hnd := C.wgpuCommandEncoderBeginRenderPass(enc.c(), cdesc)
	return safeish.Cast[*RenderPassEncoder](hnd)
}

func (enc *RenderPassEncoder) SetPipeline(p *RenderPipeline) {
	C.wgpuRenderPassEncoderSetPipeline(enc.c(), p.c())
}

func (enc *RenderPassEncoder) Draw(vertexCount, instanceCount, firstVertex, firstInstance uint32) {
	C.wgpuRenderPassEncoderDraw(
		enc.c(),
		C.uint32_t(vertexCount),
		C.uint32_t(instanceCount),
		C.uint32_t(firstVertex),
		C.uint32_t(firstInstance),
	)
}

func (enc *RenderPassEncoder) SetVertexBuffer(slot uint32, buffer *Buffer, offset, size uint64) {
	C.wgpuRenderPassEncoderSetVertexBuffer(
		enc.c(),
		C.uint32_t(slot),
		buffer.c(),
		C.uint64_t(offset),
		C.uint64_t(size),
	)
}

func (enc *RenderPassEncoder) SetIndexBuffer(buffer *Buffer, format IndexFormat, offset, size uint64) {
	C.wgpuRenderPassEncoderSetIndexBuffer(
		enc.c(),
		buffer.c(),
		C.WGPUIndexFormat(format),
		C.uint64_t(offset),
		C.uint64_t(size),
	)
}

func (enc *RenderPassEncoder) DrawIndexed(indexCount, instanceCount, firstIndex uint32, baseVertex int32, firstInstance uint32) {
	C.wgpuRenderPassEncoderDrawIndexed(
		enc.c(),
		C.uint32_t(indexCount),
		C.uint32_t(instanceCount),
		C.uint32_t(firstIndex),
		C.int32_t(baseVertex),
		C.uint32_t(firstInstance),
	)
}

func (queue *Queue) Submit(commands ...*CommandBuffer) {
	ccmds := callocn[C.WGPUCommandBuffer](len(commands))
	for i, cmd := range commands {
		ccmds[i] = cmd.c()
	}
	C.wgpuQueueSubmit(queue.c(), C.size_t(len(commands)), unsafe.SliceData(ccmds))
}

type Color struct {
	_ structs.HostLayout

	R, G, B, A float64
}

func (enc *RenderPassEncoder) End() {
	C.wgpuRenderPassEncoderEnd(enc.c())
}

func (enc *CommandEncoder) ResolveQuerySet(querySet *QuerySet, firstQuery uint32, queryCount uint32, destination *Buffer, destinationOffset uint64) {
	C.wgpuCommandEncoderResolveQuerySet(
		enc.c(),
		querySet.c(),
		C.uint32_t(firstQuery),
		C.uint32_t(queryCount),
		destination.c(),
		C.uint64_t(destinationOffset),
	)
}

type BufferDescriptor struct {
	Label            string
	Usage            BufferUsage
	Size             uint64
	MappedAtCreation bool
}

func (dev *Device) CreateBuffer(desc *BufferDescriptor) *Buffer {
	cdesc := valloc(C.WGPUBufferDescriptor{
		label:            getString(desc.Label),
		usage:            C.WGPUBufferUsageFlags(desc.Usage),
		size:             C.uint64_t(desc.Size),
		mappedAtCreation: toBool(desc.MappedAtCreation),
	})
	defer free(cdesc)
	hnd := C.wgpuDeviceCreateBuffer(dev.hnd, cdesc)
	return safeish.Cast[*Buffer](hnd)
}

func (buf *Buffer) Destroy() {
	C.wgpuBufferDestroy(buf.c())
}

func (q *Queue) WriteBuffer(buffer *Buffer, offset uint64, data []byte) {
	C.wgpuQueueWriteBuffer(q.c(), buffer.c(), C.uint64_t(offset), up(unsafe.SliceData(data)), C.size_t(len(data)))
}

func (enc *CommandEncoder) CopyBufferToBuffer(
	source *Buffer,
	sourceOffset uint64,
	destination *Buffer,
	destinationOffset uint64,
	size uint64,
) {
	C.wgpuCommandEncoderCopyBufferToBuffer(
		enc.c(),
		source.c(),
		C.uint64_t(sourceOffset),
		destination.c(),
		C.uint64_t(destinationOffset),
		C.uint64_t(size),
	)
}

func (enc *CommandEncoder) WriteTimestamp(q *QuerySet, queryIndex uint32) {
	C.wgpuCommandEncoderWriteTimestamp(enc.c(), q.c(), C.uint32_t(queryIndex))
}

func (enc *CommandEncoder) ClearBuffer(buf *Buffer, offset, size uint64) {
	C.wgpuCommandEncoderClearBuffer(enc.c(), buf.c(), C.uint64_t(offset), C.uint64_t(size))
}

func (q *Queue) OnSubmittedWorkDone(dev *Device) <-chan error {
	ch := make(chan error, 1)
	dev.mapMu.Lock()
	dev.mapCounter++
	id := dev.mapCounter
	var pinner runtime.Pinner
	pinner.Pin(dev)
	dev.mapHandles[id] = struct {
		pinner runtime.Pinner
		ch     chan<- error
	}{
		pinner: pinner,
		ch:     ch,
	}
	data := calloc[callbackData]()
	*data = callbackData{
		dev: dev,
		id:  id,
	}
	dev.mapMu.Unlock()
	C.wgpuQueueOnSubmittedWorkDone(
		q.c(),
		fp(C.doneCallback),
		up(data),
	)
	return ch
}

func (buf *Buffer) Map(dev *Device, mode MapMode, offset int, size int) <-chan error {
	ch := make(chan error, 1)
	dev.mapMu.Lock()
	dev.mapCounter++
	id := dev.mapCounter
	var pinner runtime.Pinner
	pinner.Pin(dev)
	dev.mapHandles[id] = struct {
		pinner runtime.Pinner
		ch     chan<- error
	}{
		pinner: pinner,
		ch:     ch,
	}
	data := calloc[callbackData]()
	*data = callbackData{
		dev: dev,
		id:  id,
	}
	dev.mapMu.Unlock()
	C.wgpuBufferMapAsync(
		buf.c(),
		C.WGPUMapModeFlags(mode),
		C.size_t(offset),
		C.size_t(size),
		fp(C.mapCallback),
		up(data),
	)
	return ch
}

type callbackData struct {
	dev *Device
	id  uintptr
}

//export mapCallback
func mapCallback(status C.WGPUBufferMapAsyncStatus, data unsafe.Pointer) {
	d := *(*callbackData)(data)
	C.free(data)

	d.dev.mapMu.Lock()
	hnd, ok := d.dev.mapHandles[d.id]
	delete(d.dev.mapHandles, d.id)
	d.dev.mapMu.Unlock()
	if !ok {
		panic(fmt.Sprintf("internal error: missing handle for id %d", d.id))
	}

	hnd.pinner.Unpin()
	var err error
	switch status {
	case C.WGPUBufferMapAsyncStatus_Success:
	case C.WGPUBufferMapAsyncStatus_ValidationError:
		err = ErrMapValidationError
	case C.WGPUBufferMapAsyncStatus_Unknown:
		err = ErrUnknown
	case C.WGPUBufferMapAsyncStatus_DeviceLost:
		err = ErrDeviceLost
	case C.WGPUBufferMapAsyncStatus_DestroyedBeforeCallback:
		err = ErrMapDestroyedBeforeCallback
	case C.WGPUBufferMapAsyncStatus_UnmappedBeforeCallback:
		err = ErrMapUnmappedBeforeCallback
	case C.WGPUBufferMapAsyncStatus_MappingAlreadyPending:
		err = ErrMapMappingAlreadyPending
	case C.WGPUBufferMapAsyncStatus_OffsetOutOfRange:
		err = ErrMapOffsetOutOfRange
	case C.WGPUBufferMapAsyncStatus_SizeOutOfRange:
		err = ErrMapSizeOutOfRange
	default:
		panic(fmt.Sprintf("invalid status %d", status))
	}
	hnd.ch <- err
}

//export doneCallback
func doneCallback(status C.WGPUQueueWorkDoneStatus, data unsafe.Pointer) {
	d := *(*callbackData)(data)
	C.free(data)

	d.dev.mapMu.Lock()
	hnd, ok := d.dev.mapHandles[d.id]
	delete(d.dev.mapHandles, d.id)
	d.dev.mapMu.Unlock()
	if !ok {
		panic(fmt.Sprintf("internal error: missing handle for id %d", d.id))
	}

	hnd.pinner.Unpin()
	var err error
	switch status {
	case C.WGPUQueueWorkDoneStatus_Success:
	case C.WGPUQueueWorkDoneStatus_Error:
		// XXX proper error message
		err = errors.New("some kind of error")
	case C.WGPUQueueWorkDoneStatus_Unknown:
		err = ErrUnknown
	case C.WGPUQueueWorkDoneStatus_DeviceLost:
		err = ErrDeviceLost
	default:
		panic(fmt.Sprintf("invalid status %d", status))
	}
	hnd.ch <- err
}

func (buf *Buffer) Unmap() {
	C.wgpuBufferUnmap(buf.c())
}

func (buf *Buffer) ReadOnlyMappedRange(offset, size int) []byte {
	ptr := C.wgpuBufferGetConstMappedRange(buf.c(), C.size_t(offset), C.size_t(size))
	return unsafe.Slice((*byte)(ptr), size)
}

func (buf *Buffer) MappedRange(offset, size int) []byte {
	ptr := C.wgpuBufferGetMappedRange(buf.c(), C.size_t(offset), C.size_t(size))
	return unsafe.Slice((*byte)(ptr), size)
}

func (buf *Buffer) Size() uint64 {
	return uint64(C.wgpuBufferGetSize(buf.c()))
}

func (buf *Buffer) Usage() BufferUsage {
	return BufferUsage(C.wgpuBufferGetUsage(buf.c()))
}

func (q *QuerySet) Destroy() {
	C.wgpuQuerySetDestroy(q.c())
}

func (q *QuerySet) Count() uint32 {
	return uint32(C.wgpuQuerySetGetCount(q.c()))
}

func (q *QuerySet) Type() QueryType {
	return QueryType(C.wgpuQuerySetGetType(q.c()))
}

type BindGroupLayoutEntry struct {
	Binding        uint32
	Visibility     ShaderStage
	Buffer         *BufferBindingLayout
	Sampler        *SamplerBindingLayout
	Texture        *TextureBindingLayout
	StorageTexture *StorageTextureBindingLayout

	Count uint32
}

type BufferBindingLayout struct {
	Type             BufferBindingType
	HasDynamicOffset bool
	MinBindingSize   uint64
}

type SamplerBindingLayout struct {
	Type SamplerBindingType
}

type TextureBindingLayout struct {
	SampleType    TextureSampleType
	ViewDimension TextureViewDimension
	Multisampled  bool
}

type StorageTextureBindingLayout struct {
	Access        StorageTextureAccess
	Format        TextureFormat
	ViewDimension TextureViewDimension
}

type BindGroupLayoutDescriptor struct {
	Label   string
	Entries []BindGroupLayoutEntry
}

func (dev *Device) CreateBindGroupLayout(desc *BindGroupLayoutDescriptor) *BindGroupLayout {
	centries := callocn[C.WGPUBindGroupLayoutEntry](len(desc.Entries))
	defer freen(centries)
	cextras := callocn[C.WGPUBindGroupLayoutEntryExtras](len(desc.Entries))
	defer freen(cextras)
	cdesc := valloc(C.WGPUBindGroupLayoutDescriptor{
		label:      getString(desc.Label),
		entryCount: C.size_t(len(desc.Entries)),
		entries:    unsafe.SliceData(centries),
	})
	defer free(cdesc)

	for i, e := range desc.Entries {
		ce := &centries[i]
		ce.binding = C.uint32_t(e.Binding)
		ce.visibility = C.WGPUShaderStageFlags(e.Visibility)
		if e.Buffer != nil {
			ce.buffer = C.WGPUBufferBindingLayout{
				_type:            C.WGPUBufferBindingType(e.Buffer.Type),
				hasDynamicOffset: toBool(e.Buffer.HasDynamicOffset),
				minBindingSize:   C.uint64_t(e.Buffer.MinBindingSize),
			}
		}
		if e.Sampler != nil {
			ce.sampler = C.WGPUSamplerBindingLayout{
				_type: C.WGPUSamplerBindingType(e.Sampler.Type),
			}
		}
		if e.StorageTexture != nil {
			ce.storageTexture = C.WGPUStorageTextureBindingLayout{
				access:        C.WGPUStorageTextureAccess(e.StorageTexture.Access),
				format:        C.WGPUTextureFormat(e.StorageTexture.Format),
				viewDimension: C.WGPUTextureViewDimension(e.StorageTexture.ViewDimension),
			}
		}
		if e.Texture != nil {
			ce.texture = C.WGPUTextureBindingLayout{
				sampleType:    C.WGPUTextureSampleType(e.Texture.SampleType),
				viewDimension: C.WGPUTextureViewDimension(e.Texture.ViewDimension),
				multisampled:  toBool(e.Texture.Multisampled),
			}
		}
		cextra := &cextras[i]
		cextra.chain.sType = C.WGPUSType_BindGroupLayoutEntryExtras
		cextra.count = C.uint32_t(e.Count)
		ce.nextInChain = &cextra.chain
	}

	return safeish.Cast[*BindGroupLayout](C.wgpuDeviceCreateBindGroupLayout(dev.hnd, cdesc))
}

type BindGroupEntry struct {
	Binding     uint32
	Buffer      *Buffer // nullable
	Offset      uint64
	Size        uint64
	Sampler     *Sampler     // nullable
	TextureView *TextureView // nullable

	Buffers      []*Buffer
	Samplers     []*Sampler
	TextureViews []*TextureView
}

type BindGroupDescriptor struct {
	Label   string
	Layout  *BindGroupLayout
	Entries []BindGroupEntry
}

func (dev *Device) CreateBindGroup(desc *BindGroupDescriptor) *BindGroup {
	centries := callocn[C.WGPUBindGroupEntry](len(desc.Entries))
	defer freen(centries)
	cextras := callocn[C.WGPUBindGroupEntryExtras](len(desc.Entries))
	defer freen(cextras)
	cdesc := valloc(C.WGPUBindGroupDescriptor{
		label:      getString(desc.Label),
		layout:     (*C.struct_WGPUBindGroupLayoutImpl)(desc.Layout),
		entryCount: C.size_t(len(desc.Entries)),
		entries:    unsafe.SliceData(centries),
	})
	defer free(cdesc)

	for i, e := range desc.Entries {
		cextra := &cextras[i]
		cextra.chain.sType = C.WGPUSType_BindGroupEntryExtras
		cextra.bufferCount = C.size_t(len(e.Buffers))
		cextra.samplerCount = C.size_t(len(e.Samplers))
		cextra.textureViewCount = C.size_t(len(e.TextureViews))
		if len(e.Buffers) > 0 {
			a := callocn[C.WGPUBuffer](len(e.Buffers))
			defer freen(a)
			copy(a, safeish.SliceCast[[]C.WGPUBuffer](e.Buffers))
			cextra.buffers = unsafe.SliceData(a)
		}
		if len(e.Samplers) > 0 {
			a := callocn[C.WGPUSampler](len(e.Samplers))
			defer freen(a)
			copy(a, safeish.SliceCast[[]C.WGPUSampler](e.Samplers))
			cextra.samplers = unsafe.SliceData(a)
		}
		if len(e.TextureViews) > 0 {
			a := callocn[C.WGPUTextureView](len(e.TextureViews))
			defer freen(a)
			copy(a, safeish.SliceCast[[]C.WGPUTextureView](e.TextureViews))
			cextra.textureViews = unsafe.SliceData(a)
		}
		centries[i] = C.WGPUBindGroupEntry{
			binding:     C.uint32_t(e.Binding),
			buffer:      e.Buffer.c(),
			offset:      C.uint64_t(e.Offset),
			size:        C.uint64_t(e.Size),
			sampler:     e.Sampler.c(),
			textureView: e.TextureView.c(),
		}
		centries[i].nextInChain = &cextra.chain
	}

	return safeish.Cast[*BindGroup](C.wgpuDeviceCreateBindGroup(dev.hnd, cdesc))
}

func (enc *RenderPassEncoder) SetBindGroup(
	groupIndex uint32,
	group *BindGroup,
	dynamicOffsets []uint32,
) {
	C.wgpuRenderPassEncoderSetBindGroup(
		enc.c(),
		C.uint32_t(groupIndex),
		group.c(),
		C.size_t(len(dynamicOffsets)),
		safeish.Cast[*C.uint32_t](unsafe.SliceData(dynamicOffsets)),
	)
}

func (dev *Device) CreateTexture(desc *TextureDescriptor) *Texture {
	viewFormats := callocn[C.WGPUTextureFormat](len(desc.ViewFormats))
	defer freen(viewFormats)
	copy(viewFormats, safeish.SliceCast[[]C.WGPUTextureFormat](desc.ViewFormats))
	cdesc := valloc(C.WGPUTextureDescriptor{
		label:           getString(desc.Label),
		usage:           C.WGPUTextureUsageFlags(desc.Usage),
		dimension:       C.WGPUTextureDimension(desc.Dimension),
		size:            safeish.Cast[C.WGPUExtent3D](desc.Size),
		format:          C.WGPUTextureFormat(desc.Format),
		mipLevelCount:   C.uint32_t(desc.MipLevelCount),
		sampleCount:     C.uint32_t(desc.SampleCount),
		viewFormatCount: C.size_t(len(desc.ViewFormats)),
		viewFormats:     unsafe.SliceData(viewFormats),
	})
	defer free(cdesc)
	return safeish.Cast[*Texture](C.wgpuDeviceCreateTexture(dev.hnd, cdesc))
}

type TextureDescriptor struct {
	Label         string
	Usage         TextureUsage
	Dimension     TextureDimension
	Size          Extent3D
	Format        TextureFormat
	MipLevelCount uint32
	SampleCount   uint32
	ViewFormats   []TextureFormat
}

type Extent3D struct {
	_ structs.HostLayout

	Width              uint32
	Height             uint32
	DepthOrArrayLayers uint32
}

func (q *Queue) WriteTexture(
	destination *ImageCopyTexture,
	data []byte,
	dataLayout *TextureDataLayout,
	writeSize *Extent3D,
) {
	C.wgpuQueueWriteTexture(
		q.c(),
		safeish.Cast[*C.WGPUImageCopyTexture](destination),
		up(unsafe.SliceData(data)),
		C.size_t(len(data)),
		safeish.Cast[*C.WGPUTextureDataLayout](dataLayout),
		safeish.Cast[*C.WGPUExtent3D](writeSize),
	)
}

type ImageCopyTexture struct {
	_ structs.HostLayout
	_ uintptr

	Texture  *Texture
	MipLevel uint32
	Origin   Origin3D
	Aspect   TextureAspect
}

type Origin3D struct {
	_ structs.HostLayout

	X uint32
	Y uint32
	Z uint32
}

type TextureDataLayout struct {
	_ structs.HostLayout
	_ uintptr

	Offset       uint64
	BytesPerRow  uint32
	RowsPerImage uint32
}

func (dev *Device) CreateSampler(desc *SamplerDescriptor) *Sampler {
	if desc == nil {
		return safeish.Cast[*Sampler](C.wgpuDeviceCreateSampler(dev.hnd, nil))
	}

	cdesc := valloc(C.WGPUSamplerDescriptor{
		label:         getString(desc.Label),
		addressModeU:  C.WGPUAddressMode(desc.AddressModeU),
		addressModeV:  C.WGPUAddressMode(desc.AddressModeV),
		addressModeW:  C.WGPUAddressMode(desc.AddressModeW),
		magFilter:     C.WGPUFilterMode(desc.MagFilter),
		minFilter:     C.WGPUFilterMode(desc.MinFilter),
		mipmapFilter:  C.WGPUMipmapFilterMode(desc.MipmapFilter),
		lodMinClamp:   C.float(desc.LODMinClamp),
		lodMaxClamp:   C.float(desc.LODMaxClamp),
		compare:       C.WGPUCompareFunction(desc.Compare),
		maxAnisotropy: C.uint16_t(desc.MaxAnisotropy),
	})
	defer free(cdesc)
	return safeish.Cast[*Sampler](C.wgpuDeviceCreateSampler(dev.hnd, cdesc))
}

type SamplerDescriptor struct {
	Label         string
	AddressModeU  AddressMode
	AddressModeV  AddressMode
	AddressModeW  AddressMode
	MagFilter     FilterMode
	MinFilter     FilterMode
	MipmapFilter  MipmapFilterMode
	LODMinClamp   float32
	LODMaxClamp   float32
	Compare       CompareFunction
	MaxAnisotropy uint16
}

func (enc *RenderPassEncoder) BeginOcclusionQuery(queryIndex uint32) {
	C.wgpuRenderPassEncoderBeginOcclusionQuery(enc.c(), C.uint32_t(queryIndex))
}

func (enc *RenderPassEncoder) DrawIndexedIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuRenderPassEncoderDrawIndexedIndirect(enc.c(), indirectBuffer.c(), C.uint64_t(indirectOffset))
}

func (enc *RenderPassEncoder) DrawIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuRenderPassEncoderDrawIndirect(enc.c(), indirectBuffer.c(), C.uint64_t(indirectOffset))
}

func (enc *RenderPassEncoder) EndOcclusionQuery() {
	C.wgpuRenderPassEncoderEndOcclusionQuery(enc.c())
}

func (enc *RenderPassEncoder) ExecuteBundles(bundles ...*RenderBundle) {
	C.wgpuRenderPassEncoderExecuteBundles(enc.c(), C.size_t(len(bundles)), safeish.Cast[*C.WGPURenderBundle](unsafe.SliceData(bundles)))
}

func (enc *RenderPassEncoder) InsertDebugMarker(markerLabel string) {
	C.wgpuRenderPassEncoderInsertDebugMarker(enc.c(), getString(markerLabel))
}

func (enc *RenderPassEncoder) PopDebugGroup() {
	C.wgpuRenderPassEncoderPopDebugGroup(enc.c())
}

func (enc *RenderPassEncoder) PushDebugGroup(groupLabel string) {
	C.wgpuRenderPassEncoderPushDebugGroup(enc.c(), getString(groupLabel))
}

func (enc *RenderPassEncoder) SetBlendConstant(color *Color) {
	C.wgpuRenderPassEncoderSetBlendConstant(enc.c(), safeish.Cast[*C.WGPUColor](color))
}

func (enc *RenderPassEncoder) SetScissorRect(x, y, width, height uint32) {
	C.wgpuRenderPassEncoderSetScissorRect(enc.c(), C.uint32_t(x), C.uint32_t(y), C.uint32_t(width), C.uint32_t(height))
}

func (enc *RenderPassEncoder) SetStencilReference(reference uint32) {
	C.wgpuRenderPassEncoderSetStencilReference(enc.c(), C.uint32_t(reference))
}

func (enc *RenderPassEncoder) SetViewport(x, y, width, height, minDepth, maxDepth float32) {
	C.wgpuRenderPassEncoderSetViewport(enc.c(), C.float(x), C.float(y), C.float(width), C.float(height), C.float(minDepth), C.float(maxDepth))
}

func (tex *Texture) Destroy() {
	C.wgpuTextureDestroy(tex.c())
}

func (tex *Texture) DepthOrArrayLayers() uint32 {
	return uint32(C.wgpuTextureGetDepthOrArrayLayers(tex.c()))
}

func (tex *Texture) Dimension() TextureDimension {
	return TextureDimension(C.wgpuTextureGetDimension(tex.c()))
}

func (tex *Texture) Height() uint32 {
	return uint32(C.wgpuTextureGetHeight(tex.c()))
}

func (tex *Texture) MipLevelCount() uint32 {
	return uint32(C.wgpuTextureGetMipLevelCount(tex.c()))
}

func (tex *Texture) SampleCount() uint32 {
	return uint32(C.wgpuTextureGetSampleCount(tex.c()))
}

func (tex *Texture) Usage() TextureUsage {
	return TextureUsage(C.wgpuTextureGetUsage(tex.c()))
}

func (tex *Texture) Width() uint32 {
	return uint32(C.wgpuTextureGetWidth(tex.c()))
}

func (enc *CommandEncoder) BeginComputePass(desc *ComputePassDescriptor) *ComputePassEncoder {
	cdesc := calloc[C.WGPUComputePassDescriptor]()
	defer free(cdesc)
	if desc != nil {
		cdesc.label = getString(desc.Label)
		if desc.TimestampWrites != nil {
			cts := calloc[C.WGPUComputePassTimestampWrites]()
			defer free(cts)
			cdesc.timestampWrites = cts
			cts.querySet = desc.TimestampWrites.QuerySet.c()
			cts.beginningOfPassWriteIndex = C.uint32_t(desc.TimestampWrites.BeginningOfPassWriteIndex)
			cts.endOfPassWriteIndex = C.uint32_t(desc.TimestampWrites.EndOfPassWriteIndex)
		}
	}
	hnd := C.wgpuCommandEncoderBeginComputePass(enc.c(), cdesc)
	return safeish.Cast[*ComputePassEncoder](hnd)
}

type ComputePassDescriptor struct {
	Label           string
	TimestampWrites *ComputePassTimestampWrites // nullable
}

type ComputePassTimestampWrites struct {
	QuerySet                  *QuerySet
	BeginningOfPassWriteIndex uint32
	EndOfPassWriteIndex       uint32
}

func (enc *CommandEncoder) CopyBufferToTexture(source *ImageCopyBuffer, destination *ImageCopyTexture, copySize *Extent3D) {
	C.wgpuCommandEncoderCopyBufferToTexture(
		enc.c(),
		safeish.Cast[*C.WGPUImageCopyBuffer](source),
		safeish.Cast[*C.WGPUImageCopyTexture](destination),
		safeish.Cast[*C.WGPUExtent3D](copySize),
	)
}

func (enc *CommandEncoder) CopyTextureToBuffer(source *ImageCopyTexture, destination *ImageCopyBuffer, copySize *Extent3D) {
	C.wgpuCommandEncoderCopyTextureToBuffer(
		enc.c(),
		safeish.Cast[*C.WGPUImageCopyTexture](source),
		safeish.Cast[*C.WGPUImageCopyBuffer](destination),
		safeish.Cast[*C.WGPUExtent3D](copySize),
	)
}

func (enc *CommandEncoder) CopyTextureToTexture(source *ImageCopyTexture, destination *ImageCopyTexture, copySize *Extent3D) {
	C.wgpuCommandEncoderCopyTextureToTexture(
		enc.c(),
		safeish.Cast[*C.WGPUImageCopyTexture](source),
		safeish.Cast[*C.WGPUImageCopyTexture](destination),
		safeish.Cast[*C.WGPUExtent3D](copySize),
	)
}

func (enc *CommandEncoder) PopDebugGroup() {
	C.wgpuCommandEncoderPopDebugGroup(enc.c())
}

func (enc *CommandEncoder) PushDebugGroup(label string) {
	C.wgpuCommandEncoderPushDebugGroup(enc.c(), getString(label))
}

type ImageCopyBuffer struct {
	_ structs.HostLayout
	_ uintptr

	Layout TextureDataLayout
	Buffer *Buffer
}

func (dev *Device) CreateComputePipeline(desc *ComputePipelineDescriptor) *ComputePipeline {
	cdesc := calloc[C.WGPUComputePipelineDescriptor]()
	defer free(cdesc)
	cdesc.label = getString(desc.Label)
	cdesc.layout = desc.Layout.c()
	cdesc.compute.module = desc.Compute.Module.c()
	cdesc.compute.entryPoint = getString(desc.Compute.EntryPoint)
	cdesc.compute.constantCount = C.size_t(len(desc.Compute.Constants))
	constants := callocn[C.WGPUConstantEntry](len(desc.Compute.Constants))
	defer freen(constants)
	cdesc.compute.constants = unsafe.SliceData(constants)
	for i, c := range desc.Compute.Constants {
		constants[i] = C.WGPUConstantEntry{
			key:   getString(c.Key),
			value: C.double(c.Value),
		}
	}

	hnd := C.wgpuDeviceCreateComputePipeline(dev.hnd, cdesc)
	return safeish.Cast[*ComputePipeline](hnd)
}

type ComputePipelineDescriptor struct {
	Label   string
	Layout  *PipelineLayout // nullable
	Compute ProgrammableStageDescriptor
}

type ProgrammableStageDescriptor struct {
	Module     *ShaderModule
	EntryPoint string
	Constants  []ConstantEntry
}

func (p *RenderPipeline) BindGroupLayout(groupIndex uint32) *BindGroupLayout {
	return safeish.Cast[*BindGroupLayout](C.wgpuRenderPipelineGetBindGroupLayout(p.c(), C.uint32_t(groupIndex)))
}

func (enc *ComputePassEncoder) DispatchWorkgroups(countX, countY, countZ uint32) {
	C.wgpuComputePassEncoderDispatchWorkgroups(enc.c(), C.uint32_t(countX), C.uint32_t(countY), C.uint32_t(countZ))
}

func (enc *ComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer *Buffer, indirectOffset uint64) {
	C.wgpuComputePassEncoderDispatchWorkgroupsIndirect(enc.c(), indirectBuffer.c(), C.uint64_t(indirectOffset))
}

func (enc *ComputePassEncoder) End() {
	C.wgpuComputePassEncoderEnd(enc.c())
}

func (enc *ComputePassEncoder) InsertDebugMarker(label string) {
	C.wgpuComputePassEncoderInsertDebugMarker(enc.c(), getString(label))
}

func (enc *ComputePassEncoder) PopDebugGroup() {
	C.wgpuComputePassEncoderPopDebugGroup(enc.c())
}

func (enc *ComputePassEncoder) PushDebugGroup(label string) {
	C.wgpuComputePassEncoderPushDebugGroup(enc.c(), getString(label))
}

func (enc *ComputePassEncoder) SetBindGroup(groupIndex uint32, group *BindGroup, dynamicOffsets []uint32) {
	C.wgpuComputePassEncoderSetBindGroup(
		enc.c(),
		C.uint32_t(groupIndex),
		group.c(),
		C.size_t(len(dynamicOffsets)),
		safeish.Cast[*C.uint32_t](unsafe.SliceData(dynamicOffsets)),
	)
}

func (enc *ComputePassEncoder) SetPipeline(p *ComputePipeline) {
	C.wgpuComputePassEncoderSetPipeline(enc.c(), p.c())
}

func (p *ComputePipeline) BindGroupLayout(groupIndex uint32) *BindGroupLayout {
	hnd := C.wgpuComputePipelineGetBindGroupLayout(p.c(), C.uint32_t(groupIndex))
	return safeish.Cast[*BindGroupLayout](hnd)
}

//export uncapturedErrorCallback
func uncapturedErrorCallback(typ C.WGPUErrorType, msg *C.char, data unsafe.Pointer) {
	switch typ {
	case C.WGPUErrorType_Validation:
		panic(ValidationError{C.GoString(msg)})
	case C.WGPUErrorType_OutOfMemory:
		panic(OutOfMemoryError{C.GoString(msg)})
	case C.WGPUErrorType_Internal:
		panic(InternalError{C.GoString(msg)})
	case C.WGPUErrorType_Unknown:
		panic(UnknownError{C.GoString(msg)})
	case C.WGPUErrorType_DeviceLost:
		panic(DeviceLostError{C.GoString(msg)})
	default:
		panic(fmt.Sprintf("unhandled error type %d", typ))
	}
}

type ValidationError struct{ message string }
type OutOfMemoryError struct{ message string }
type InternalError struct{ message string }
type UnknownError struct{ message string }
type DeviceLostError struct{ message string }

func (err ValidationError) Error() string  { return err.message }
func (err OutOfMemoryError) Error() string { return err.message }
func (err InternalError) Error() string    { return err.message }
func (err UnknownError) Error() string     { return err.message }
func (err DeviceLostError) Error() string  { return err.message }

func (ptr *Adapter) Release()             { C.wgpuAdapterRelease(ptr.c()) }
func (ptr *BindGroup) Release()           { C.wgpuBindGroupRelease(ptr.c()) }
func (ptr *BindGroupLayout) Release()     { C.wgpuBindGroupLayoutRelease(ptr.c()) }
func (ptr *Buffer) Release()              { C.wgpuBufferRelease(ptr.c()) }
func (ptr *CommandBuffer) Release()       { C.wgpuCommandBufferRelease(ptr.c()) }
func (ptr *CommandEncoder) Release()      { C.wgpuCommandEncoderRelease(ptr.c()) }
func (ptr *ComputePassEncoder) Release()  { C.wgpuComputePassEncoderRelease(ptr.c()) }
func (ptr *ComputePipeline) Release()     { C.wgpuComputePipelineRelease(ptr.c()) }
func (ptr *Device) Release()              { C.wgpuDeviceRelease(ptr.hnd) }
func (ptr *Instance) Release()            { C.wgpuInstanceRelease(ptr.c()) }
func (ptr *PipelineLayout) Release()      { C.wgpuPipelineLayoutRelease(ptr.c()) }
func (ptr *QuerySet) Release()            { C.wgpuQuerySetRelease(ptr.c()) }
func (ptr *Queue) Release()               { C.wgpuQueueRelease(ptr.c()) }
func (ptr *RenderBundle) Release()        { C.wgpuRenderBundleRelease(ptr.c()) }
func (ptr *RenderBundleEncoder) Release() { C.wgpuRenderBundleEncoderRelease(ptr.c()) }
func (ptr *RenderPassEncoder) Release()   { C.wgpuRenderPassEncoderRelease(ptr.c()) }
func (ptr *RenderPipeline) Release()      { C.wgpuRenderPipelineRelease(ptr.c()) }
func (ptr *Sampler) Release()             { C.wgpuSamplerRelease(ptr.c()) }
func (ptr *ShaderModule) Release()        { C.wgpuShaderModuleRelease(ptr.c()) }
func (ptr *Surface) Release()             { C.wgpuSurfaceRelease(ptr.c()) }
func (ptr *Texture) Release()             { C.wgpuTextureRelease(ptr.c()) }
func (ptr *TextureView) Release()         { C.wgpuTextureViewRelease(ptr.c()) }

func (ptr *Adapter) c() *C.struct_WGPUAdapterImpl {
	return (*C.struct_WGPUAdapterImpl)(ptr)
}
func (ptr *BindGroup) c() *C.struct_WGPUBindGroupImpl {
	return (*C.struct_WGPUBindGroupImpl)(ptr)
}
func (ptr *BindGroupLayout) c() *C.struct_WGPUBindGroupLayoutImpl {
	return (*C.struct_WGPUBindGroupLayoutImpl)(ptr)
}
func (ptr *Buffer) c() *C.struct_WGPUBufferImpl {
	return (*C.struct_WGPUBufferImpl)(ptr)
}
func (ptr *CommandBuffer) c() *C.struct_WGPUCommandBufferImpl {
	return (*C.struct_WGPUCommandBufferImpl)(ptr)
}
func (ptr *CommandEncoder) c() *C.struct_WGPUCommandEncoderImpl {
	return (*C.struct_WGPUCommandEncoderImpl)(ptr)
}
func (ptr *ComputePassEncoder) c() *C.struct_WGPUComputePassEncoderImpl {
	return (*C.struct_WGPUComputePassEncoderImpl)(ptr)
}
func (ptr *ComputePipeline) c() *C.struct_WGPUComputePipelineImpl {
	return (*C.struct_WGPUComputePipelineImpl)(ptr)
}
func (ptr *Instance) c() *C.struct_WGPUInstanceImpl {
	return (*C.struct_WGPUInstanceImpl)(ptr)
}
func (ptr *PipelineLayout) c() *C.struct_WGPUPipelineLayoutImpl {
	return (*C.struct_WGPUPipelineLayoutImpl)(ptr)
}
func (ptr *QuerySet) c() *C.struct_WGPUQuerySetImpl {
	return (*C.struct_WGPUQuerySetImpl)(ptr)
}
func (ptr *Queue) c() *C.struct_WGPUQueueImpl {
	return (*C.struct_WGPUQueueImpl)(ptr)
}
func (ptr *RenderBundle) c() *C.struct_WGPURenderBundleImpl {
	return (*C.struct_WGPURenderBundleImpl)(ptr)
}
func (ptr *RenderBundleEncoder) c() *C.struct_WGPURenderBundleEncoderImpl {
	return (*C.struct_WGPURenderBundleEncoderImpl)(ptr)
}
func (ptr *RenderPassEncoder) c() *C.struct_WGPURenderPassEncoderImpl {
	return (*C.struct_WGPURenderPassEncoderImpl)(ptr)
}
func (ptr *RenderPipeline) c() *C.struct_WGPURenderPipelineImpl {
	return (*C.struct_WGPURenderPipelineImpl)(ptr)
}
func (ptr *Sampler) c() *C.struct_WGPUSamplerImpl {
	return (*C.struct_WGPUSamplerImpl)(ptr)
}
func (ptr *ShaderModule) c() *C.struct_WGPUShaderModuleImpl {
	return (*C.struct_WGPUShaderModuleImpl)(ptr)
}
func (ptr *Surface) c() *C.struct_WGPUSurfaceImpl {
	return (*C.struct_WGPUSurfaceImpl)(ptr)
}
func (ptr *Texture) c() *C.struct_WGPUTextureImpl {
	return (*C.struct_WGPUTextureImpl)(ptr)
}
func (ptr *TextureView) c() *C.struct_WGPUTextureViewImpl {
	return (*C.struct_WGPUTextureViewImpl)(ptr)
}

// Currently not implemented by wgpu-native:
// WGPU_EXPORT void wgpuShaderModuleGetCompilationInfo(WGPUShaderModule shaderModule, WGPUCompilationInfoCallback callback, void * userdata) WGPU_FUNCTION_ATTRIBUTE;
// WGPU_EXPORT void wgpuDeviceCreateRenderPipelineAsync(WGPUDevice device, WGPURenderPipelineDescriptor const * descriptor, WGPUCreateRenderPipelineAsyncCallback callback, void * userdata) WGPU_FUNCTION_ATTRIBUTE;
// WGPU_EXPORT void wgpuDeviceCreateComputePipelineAsync(WGPUDevice device, WGPUComputePipelineDescriptor const * descriptor, WGPUCreateComputePipelineAsyncCallback callback, void * userdata) WGPU_FUNCTION_ATTRIBUTE;
// WGPU_EXPORT void wgpuInstanceProcessEvents(WGPUInstance instance) WGPU_FUNCTION_ATTRIBUTE;
// func (ptr *BindGroup) SetLabel(label string) { C.wgpuBindGroupSetLabel(ptr.c(), getString(label)) }
// func (ptr *BindGroupLayout) SetLabel(label string) {C.wgpuBindGroupLayoutSetLabel(ptr.c(), getString(label))}
// func (buf *Buffer) MapState() BufferMapState {return BufferMapState(C.wgpuBufferGetMapState(buf.c()))}
// func (ptr *Buffer) SetLabel(label string) {C.wgpuBufferSetLabel(ptr.c(), getString(label))}
// func (ptr *CommandBuffer) SetLabel(label string) {C.wgpuCommandBufferSetLabel(ptr.c(), getString(label))}
// func (ptr *CommandEncoder) SetLabel(label string) {C.wgpuCommandEncoderSetLabel(ptr.c(), getString(label))}
// func (ptr *ComputePassEncoder) SetLabel(label string) {C.wgpuComputePassEncoderSetLabel(ptr.c(), getString(label))}
// func (ptr *ComputePipeline) SetLabel(label string) {C.wgpuComputePipelineSetLabel(ptr.c(), getString(label))}
// func (ptr *Device) SetLabel(label string) {C.wgpuDeviceSetLabel(ptr.hnd, getString(label))}
// func (ptr *PipelineLayout) SetLabel(label string) {C.wgpuPipelineLayoutSetLabel(ptr.c(), getString(label))}
// func (ptr *QuerySet) SetLabel(label string) {C.wgpuQuerySetSetLabel(ptr.c(), getString(label))}
// func (ptr *Queue) SetLabel(label string) {C.wgpuQueueSetLabel(ptr.c(), getString(label))}
// func (ptr *RenderBundle) SetLabel(label string) {C.wgpuRenderBundleSetLabel(ptr.c(), getString(label))}
// func (ptr *RenderBundleEncoder) SetLabel(label string) {C.wgpuRenderBundleEncoderSetLabel(ptr.c(), getString(label))}
// func (ptr *RenderPassEncoder) SetLabel(label string) {C.wgpuRenderPassEncoderSetLabel(ptr.c(), getString(label))}
// func (ptr *RenderPipeline) SetLabel(label string) {C.wgpuRenderPipelineSetLabel(ptr.c(), getString(label))}
// func (ptr *Sampler) SetLabel(label string) {C.wgpuSamplerSetLabel(ptr.c(), getString(label))}
// func (ptr *ShaderModule) SetLabel(label string) {C.wgpuShaderModuleSetLabel(ptr.c(), getString(label))}
// func (ptr *Texture) SetLabel(label string) {C.wgpuTextureSetLabel(ptr.c(), getString(label))}
// func (ptr *TextureView) SetLabel(label string) {C.wgpuTextureViewSetLabel(ptr.c(), getString(label))}
