// Code generated by "stringer -type requestAdapterStatus,requestDeviceStatus,PowerPreference,BackendType,AdapterType,DeviceLostReason,ErrorFilter,ErrorType,PrimitiveTopology,IndexFormat,FrontFace,CullMode,TextureFormat,VertexFormat,CompareFunction,StencilOperation,BlendOperation,BlendFactor,ColorWriteMask,VertexStepMode,PresentMode,CompositeAlphaMode,TextureUsage,TextureViewDimension,TextureAspect,LoadOp,StoreOp,QueryType,BufferUsage,MapMode,BufferMapState,ShaderStage,BufferBindingType,SamplerBindingType,TextureSampleType,StorageTextureAccess,TextureDimension,AddressMode,FilterMode,MipmapFilterMode -linecomment -output webgpu1_string.go"; DO NOT EDIT.

package wgpu

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[requestAdapterStatusSuccess-0]
	_ = x[requestAdapterStatusUnavailable-1]
	_ = x[requestAdapterStatusError-2]
	_ = x[requestAdapterStatusUnknown-3]
}

const _requestAdapterStatus_name = "SuccessUnavailableErrorUnknown"

var _requestAdapterStatus_index = [...]uint8{0, 7, 18, 23, 30}

func (i requestAdapterStatus) String() string {
	if i >= requestAdapterStatus(len(_requestAdapterStatus_index)-1) {
		return "requestAdapterStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _requestAdapterStatus_name[_requestAdapterStatus_index[i]:_requestAdapterStatus_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[requestDeviceStatusSuccess-0]
	_ = x[requestDeviceStatusError-1]
	_ = x[requestDeviceStatusUnknown-2]
}

const _requestDeviceStatus_name = "SuccessErrorUnknown"

var _requestDeviceStatus_index = [...]uint8{0, 7, 12, 19}

func (i requestDeviceStatus) String() string {
	if i >= requestDeviceStatus(len(_requestDeviceStatus_index)-1) {
		return "requestDeviceStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _requestDeviceStatus_name[_requestDeviceStatus_index[i]:_requestDeviceStatus_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PowerPreferenceUndefined-0]
	_ = x[PowerPreferenceLowPower-1]
	_ = x[PowerPreferenceHighPerformance-2]
}

const _PowerPreference_name = "UndefinedLowPowerHighPerformance"

var _PowerPreference_index = [...]uint8{0, 9, 17, 32}

func (i PowerPreference) String() string {
	if i >= PowerPreference(len(_PowerPreference_index)-1) {
		return "PowerPreference(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PowerPreference_name[_PowerPreference_index[i]:_PowerPreference_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BackendTypeUndefined-0]
	_ = x[BackendTypeNull-1]
	_ = x[BackendTypeWebGPU-2]
	_ = x[BackendTypeD3D11-3]
	_ = x[BackendTypeD3D12-4]
	_ = x[BackendTypeMetal-5]
	_ = x[BackendTypeVulkan-6]
	_ = x[BackendTypeOpenGL-7]
	_ = x[BackendTypeOpenGLES-8]
}

const _BackendType_name = "UndefinedNullWebGPUD3D11D3D12MetalVulkanOpenGLOpenGLES"

var _BackendType_index = [...]uint8{0, 9, 13, 19, 24, 29, 34, 40, 46, 54}

func (i BackendType) String() string {
	if i >= BackendType(len(_BackendType_index)-1) {
		return "BackendType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BackendType_name[_BackendType_index[i]:_BackendType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AdapterTypeDiscreteGPU-0]
	_ = x[AdapterTypeIntegratedGPU-1]
	_ = x[AdapterTypeCPU-2]
	_ = x[AdapterTypeUnknown-3]
}

const _AdapterType_name = "DiscreteGPUIntegratedGPUCPUUnknown"

var _AdapterType_index = [...]uint8{0, 11, 24, 27, 34}

func (i AdapterType) String() string {
	if i >= AdapterType(len(_AdapterType_index)-1) {
		return "AdapterType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AdapterType_name[_AdapterType_index[i]:_AdapterType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DeviceLostReasonUndefined-0]
	_ = x[DeviceLostReasonDestroyed-1]
}

const _DeviceLostReason_name = "UndefinedDestroyed"

var _DeviceLostReason_index = [...]uint8{0, 9, 18}

func (i DeviceLostReason) String() string {
	if i >= DeviceLostReason(len(_DeviceLostReason_index)-1) {
		return "DeviceLostReason(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DeviceLostReason_name[_DeviceLostReason_index[i]:_DeviceLostReason_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrorFilterValidation-0]
	_ = x[ErrorFilterOutOfMemory-1]
	_ = x[ErrorFilterInternal-2]
}

const _ErrorFilter_name = "ValidationOutOfMemoryInternal"

var _ErrorFilter_index = [...]uint8{0, 10, 21, 29}

func (i ErrorFilter) String() string {
	if i >= ErrorFilter(len(_ErrorFilter_index)-1) {
		return "ErrorFilter(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorFilter_name[_ErrorFilter_index[i]:_ErrorFilter_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrorTypeNoError-0]
	_ = x[ErrorTypeValidation-1]
	_ = x[ErrorTypeOutOfMemory-2]
	_ = x[ErrorTypeInternal-3]
	_ = x[ErrorTypeUnknown-4]
	_ = x[ErrorTypeDeviceLost-5]
}

const _ErrorType_name = "NoErrorValidationOutOfMemoryInternalUnknownDeviceLost"

var _ErrorType_index = [...]uint8{0, 7, 17, 28, 36, 43, 53}

func (i ErrorType) String() string {
	if i >= ErrorType(len(_ErrorType_index)-1) {
		return "ErrorType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorType_name[_ErrorType_index[i]:_ErrorType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PrimitiveTopologyPointList-0]
	_ = x[PrimitiveTopologyLineList-1]
	_ = x[PrimitiveTopologyLineStrip-2]
	_ = x[PrimitiveTopologyTriangleList-3]
	_ = x[PrimitiveTopologyTriangleStrip-4]
}

const _PrimitiveTopology_name = "PointListLineListLineStripTriangleListTriangleStrip"

var _PrimitiveTopology_index = [...]uint8{0, 9, 17, 26, 38, 51}

func (i PrimitiveTopology) String() string {
	if i >= PrimitiveTopology(len(_PrimitiveTopology_index)-1) {
		return "PrimitiveTopology(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PrimitiveTopology_name[_PrimitiveTopology_index[i]:_PrimitiveTopology_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IndexFormatUndefined-0]
	_ = x[IndexFormatUint16-1]
	_ = x[IndexFormatUint32-2]
}

const _IndexFormat_name = "UndefinedUint16Uint32"

var _IndexFormat_index = [...]uint8{0, 9, 15, 21}

func (i IndexFormat) String() string {
	if i >= IndexFormat(len(_IndexFormat_index)-1) {
		return "IndexFormat(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IndexFormat_name[_IndexFormat_index[i]:_IndexFormat_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FrontFaceCCW-0]
	_ = x[FrontFaceCW-1]
}

const _FrontFace_name = "CCWCW"

var _FrontFace_index = [...]uint8{0, 3, 5}

func (i FrontFace) String() string {
	if i >= FrontFace(len(_FrontFace_index)-1) {
		return "FrontFace(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FrontFace_name[_FrontFace_index[i]:_FrontFace_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CullModeNone-0]
	_ = x[CullModeFront-1]
	_ = x[CullModeBack-2]
}

const _CullMode_name = "NoneFrontBack"

var _CullMode_index = [...]uint8{0, 4, 9, 13}

func (i CullMode) String() string {
	if i >= CullMode(len(_CullMode_index)-1) {
		return "CullMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CullMode_name[_CullMode_index[i]:_CullMode_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TextureFormatUndefined-0]
	_ = x[TextureFormatR8Unorm-1]
	_ = x[TextureFormatR8Snorm-2]
	_ = x[TextureFormatR8Uint-3]
	_ = x[TextureFormatR8Sint-4]
	_ = x[TextureFormatR16Uint-5]
	_ = x[TextureFormatR16Sint-6]
	_ = x[TextureFormatR16Float-7]
	_ = x[TextureFormatRG8Unorm-8]
	_ = x[TextureFormatRG8Snorm-9]
	_ = x[TextureFormatRG8Uint-10]
	_ = x[TextureFormatRG8Sint-11]
	_ = x[TextureFormatR32Float-12]
	_ = x[TextureFormatR32Uint-13]
	_ = x[TextureFormatR32Sint-14]
	_ = x[TextureFormatRG16Uint-15]
	_ = x[TextureFormatRG16Sint-16]
	_ = x[TextureFormatRG16Float-17]
	_ = x[TextureFormatRGBA8Unorm-18]
	_ = x[TextureFormatRGBA8UnormSrgb-19]
	_ = x[TextureFormatRGBA8Snorm-20]
	_ = x[TextureFormatRGBA8Uint-21]
	_ = x[TextureFormatRGBA8Sint-22]
	_ = x[TextureFormatBGRA8Unorm-23]
	_ = x[TextureFormatBGRA8UnormSrgb-24]
	_ = x[TextureFormatRGB10A2Uint-25]
	_ = x[TextureFormatRGB10A2Unorm-26]
	_ = x[TextureFormatRG11B10Ufloat-27]
	_ = x[TextureFormatRGB9E5Ufloat-28]
	_ = x[TextureFormatRG32Float-29]
	_ = x[TextureFormatRG32Uint-30]
	_ = x[TextureFormatRG32Sint-31]
	_ = x[TextureFormatRGBA16Uint-32]
	_ = x[TextureFormatRGBA16Sint-33]
	_ = x[TextureFormatRGBA16Float-34]
	_ = x[TextureFormatRGBA32Float-35]
	_ = x[TextureFormatRGBA32Uint-36]
	_ = x[TextureFormatRGBA32Sint-37]
	_ = x[TextureFormatStencil8-38]
	_ = x[TextureFormatDepth16Unorm-39]
	_ = x[TextureFormatDepth24Plus-40]
	_ = x[TextureFormatDepth24PlusStencil8-41]
	_ = x[TextureFormatDepth32Float-42]
	_ = x[TextureFormatDepth32FloatStencil8-43]
	_ = x[TextureFormatBC1RGBAUnorm-44]
	_ = x[TextureFormatBC1RGBAUnormSrgb-45]
	_ = x[TextureFormatBC2RGBAUnorm-46]
	_ = x[TextureFormatBC2RGBAUnormSrgb-47]
	_ = x[TextureFormatBC3RGBAUnorm-48]
	_ = x[TextureFormatBC3RGBAUnormSrgb-49]
	_ = x[TextureFormatBC4RUnorm-50]
	_ = x[TextureFormatBC4RSnorm-51]
	_ = x[TextureFormatBC5RGUnorm-52]
	_ = x[TextureFormatBC5RGSnorm-53]
	_ = x[TextureFormatBC6HRGBUfloat-54]
	_ = x[TextureFormatBC6HRGBFloat-55]
	_ = x[TextureFormatBC7RGBAUnorm-56]
	_ = x[TextureFormatBC7RGBAUnormSrgb-57]
	_ = x[TextureFormatETC2RGB8Unorm-58]
	_ = x[TextureFormatETC2RGB8UnormSrgb-59]
	_ = x[TextureFormatETC2RGB8A1Unorm-60]
	_ = x[TextureFormatETC2RGB8A1UnormSrgb-61]
	_ = x[TextureFormatETC2RGBA8Unorm-62]
	_ = x[TextureFormatETC2RGBA8UnormSrgb-63]
	_ = x[TextureFormatEACR11Unorm-64]
	_ = x[TextureFormatEACR11Snorm-65]
	_ = x[TextureFormatEACRG11Unorm-66]
	_ = x[TextureFormatEACRG11Snorm-67]
	_ = x[TextureFormatASTC4x4Unorm-68]
	_ = x[TextureFormatASTC4x4UnormSrgb-69]
	_ = x[TextureFormatASTC5x4Unorm-70]
	_ = x[TextureFormatASTC5x4UnormSrgb-71]
	_ = x[TextureFormatASTC5x5Unorm-72]
	_ = x[TextureFormatASTC5x5UnormSrgb-73]
	_ = x[TextureFormatASTC6x5Unorm-74]
	_ = x[TextureFormatASTC6x5UnormSrgb-75]
	_ = x[TextureFormatASTC6x6Unorm-76]
	_ = x[TextureFormatASTC6x6UnormSrgb-77]
	_ = x[TextureFormatASTC8x5Unorm-78]
	_ = x[TextureFormatASTC8x5UnormSrgb-79]
	_ = x[TextureFormatASTC8x6Unorm-80]
	_ = x[TextureFormatASTC8x6UnormSrgb-81]
	_ = x[TextureFormatASTC8x8Unorm-82]
	_ = x[TextureFormatASTC8x8UnormSrgb-83]
	_ = x[TextureFormatASTC10x5Unorm-84]
	_ = x[TextureFormatASTC10x5UnormSrgb-85]
	_ = x[TextureFormatASTC10x6Unorm-86]
	_ = x[TextureFormatASTC10x6UnormSrgb-87]
	_ = x[TextureFormatASTC10x8Unorm-88]
	_ = x[TextureFormatASTC10x8UnormSrgb-89]
	_ = x[TextureFormatASTC10x10Unorm-90]
	_ = x[TextureFormatASTC10x10UnormSrgb-91]
	_ = x[TextureFormatASTC12x10Unorm-92]
	_ = x[TextureFormatASTC12x10UnormSrgb-93]
	_ = x[TextureFormatASTC12x12Unorm-94]
	_ = x[TextureFormatASTC12x12UnormSrgb-95]
}

const _TextureFormat_name = "UndefinedR8UnormR8SnormR8UintR8SintR16UintR16SintR16FloatRG8UnormRG8SnormRG8UintRG8SintR32FloatR32UintR32SintRG16UintRG16SintRG16FloatRGBA8UnormRGBA8UnormSrgbRGBA8SnormRGBA8UintRGBA8SintBGRA8UnormBGRA8UnormSrgbRGB10A2UintRGB10A2UnormRG11B10UfloatRGB9E5UfloatRG32FloatRG32UintRG32SintRGBA16UintRGBA16SintRGBA16FloatRGBA32FloatRGBA32UintRGBA32SintStencil8Depth16UnormDepth24PlusDepth24PlusStencil8Depth32FloatDepth32FloatStencil8BC1RGBAUnormBC1RGBAUnormSrgbBC2RGBAUnormBC2RGBAUnormSrgbBC3RGBAUnormBC3RGBAUnormSrgbBC4RUnormBC4RSnormBC5RGUnormBC5RGSnormBC6HRGBUfloatBC6HRGBFloatBC7RGBAUnormBC7RGBAUnormSrgbETC2RGB8UnormETC2RGB8UnormSrgbETC2RGB8A1UnormETC2RGB8A1UnormSrgbETC2RGBA8UnormETC2RGBA8UnormSrgbEACR11UnormEACR11SnormEACRG11UnormEACRG11SnormASTC4x4UnormASTC4x4UnormSrgbASTC5x4UnormASTC5x4UnormSrgbASTC5x5UnormASTC5x5UnormSrgbASTC6x5UnormASTC6x5UnormSrgbASTC6x6UnormASTC6x6UnormSrgbASTC8x5UnormASTC8x5UnormSrgbASTC8x6UnormASTC8x6UnormSrgbASTC8x8UnormASTC8x8UnormSrgbASTC10x5UnormASTC10x5UnormSrgbASTC10x6UnormASTC10x6UnormSrgbASTC10x8UnormASTC10x8UnormSrgbASTC10x10UnormASTC10x10UnormSrgbASTC12x10UnormASTC12x10UnormSrgbASTC12x12UnormASTC12x12UnormSrgb"

var _TextureFormat_index = [...]uint16{0, 9, 16, 23, 29, 35, 42, 49, 57, 65, 73, 80, 87, 95, 102, 109, 117, 125, 134, 144, 158, 168, 177, 186, 196, 210, 221, 233, 246, 258, 267, 275, 283, 293, 303, 314, 325, 335, 345, 353, 365, 376, 395, 407, 427, 439, 455, 467, 483, 495, 511, 520, 529, 539, 549, 562, 574, 586, 602, 615, 632, 647, 666, 680, 698, 709, 720, 732, 744, 756, 772, 784, 800, 812, 828, 840, 856, 868, 884, 896, 912, 924, 940, 952, 968, 981, 998, 1011, 1028, 1041, 1058, 1072, 1090, 1104, 1122, 1136, 1154}

func (i TextureFormat) String() string {
	if i >= TextureFormat(len(_TextureFormat_index)-1) {
		return "TextureFormat(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TextureFormat_name[_TextureFormat_index[i]:_TextureFormat_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VertexFormatUndefined-0]
	_ = x[VertexFormatUint8x2-1]
	_ = x[VertexFormatUint8x4-2]
	_ = x[VertexFormatSint8x2-3]
	_ = x[VertexFormatSint8x4-4]
	_ = x[VertexFormatUnorm8x2-5]
	_ = x[VertexFormatUnorm8x4-6]
	_ = x[VertexFormatSnorm8x2-7]
	_ = x[VertexFormatSnorm8x4-8]
	_ = x[VertexFormatUint16x2-9]
	_ = x[VertexFormatUint16x4-10]
	_ = x[VertexFormatSint16x2-11]
	_ = x[VertexFormatSint16x4-12]
	_ = x[VertexFormatUnorm16x2-13]
	_ = x[VertexFormatUnorm16x4-14]
	_ = x[VertexFormatSnorm16x2-15]
	_ = x[VertexFormatSnorm16x4-16]
	_ = x[VertexFormatFloat16x2-17]
	_ = x[VertexFormatFloat16x4-18]
	_ = x[VertexFormatFloat32-19]
	_ = x[VertexFormatFloat32x2-20]
	_ = x[VertexFormatFloat32x3-21]
	_ = x[VertexFormatFloat32x4-22]
	_ = x[VertexFormatUint32-23]
	_ = x[VertexFormatUint32x2-24]
	_ = x[VertexFormatUint32x3-25]
	_ = x[VertexFormatUint32x4-26]
	_ = x[VertexFormatSint32-27]
	_ = x[VertexFormatSint32x2-28]
	_ = x[VertexFormatSint32x3-29]
	_ = x[VertexFormatSint32x4-30]
}

const _VertexFormat_name = "UndefinedUint8x2Uint8x4Sint8x2Sint8x4Unorm8x2Unorm8x4Snorm8x2Snorm8x4Uint16x2Uint16x4Sint16x2Sint16x4Unorm16x2Unorm16x4Snorm16x2Snorm16x4Float16x2Float16x4Float32Float32x2Float32x3Float32x4Uint32Uint32x2Uint32x3Uint32x4Sint32Sint32x2Sint32x3Sint32x4"

var _VertexFormat_index = [...]uint8{0, 9, 16, 23, 30, 37, 45, 53, 61, 69, 77, 85, 93, 101, 110, 119, 128, 137, 146, 155, 162, 171, 180, 189, 195, 203, 211, 219, 225, 233, 241, 249}

func (i VertexFormat) String() string {
	if i >= VertexFormat(len(_VertexFormat_index)-1) {
		return "VertexFormat(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VertexFormat_name[_VertexFormat_index[i]:_VertexFormat_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CompareFunctionUndefined-0]
	_ = x[CompareFunctionNever-1]
	_ = x[CompareFunctionLess-2]
	_ = x[CompareFunctionLessEqual-3]
	_ = x[CompareFunctionGreater-4]
	_ = x[CompareFunctionGreaterEqual-5]
	_ = x[CompareFunctionEqual-6]
	_ = x[CompareFunctionNotEqual-7]
	_ = x[CompareFunctionAlways-8]
}

const _CompareFunction_name = "UndefinedNeverLessLessEqualGreaterGreaterEqualEqualNotEqualAlways"

var _CompareFunction_index = [...]uint8{0, 9, 14, 18, 27, 34, 46, 51, 59, 65}

func (i CompareFunction) String() string {
	if i >= CompareFunction(len(_CompareFunction_index)-1) {
		return "CompareFunction(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CompareFunction_name[_CompareFunction_index[i]:_CompareFunction_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StencilOperationKeep-0]
	_ = x[StencilOperationZero-1]
	_ = x[StencilOperationReplace-2]
	_ = x[StencilOperationInvert-3]
	_ = x[StencilOperationIncrementClamp-4]
	_ = x[StencilOperationDecrementClamp-5]
	_ = x[StencilOperationIncrementWrap-6]
	_ = x[StencilOperationDecrementWrap-7]
}

const _StencilOperation_name = "KeepZeroReplaceInvertIncrementClampDecrementClampIncrementWrapDecrementWrap"

var _StencilOperation_index = [...]uint8{0, 4, 8, 15, 21, 35, 49, 62, 75}

func (i StencilOperation) String() string {
	if i >= StencilOperation(len(_StencilOperation_index)-1) {
		return "StencilOperation(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StencilOperation_name[_StencilOperation_index[i]:_StencilOperation_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BlendOperationAdd-0]
	_ = x[BlendOperationSubtract-1]
	_ = x[BlendOperationReverseSubtract-2]
	_ = x[BlendOperationMin-3]
	_ = x[BlendOperationMax-4]
}

const _BlendOperation_name = "AddSubtractReverseSubtractMinMax"

var _BlendOperation_index = [...]uint8{0, 3, 11, 26, 29, 32}

func (i BlendOperation) String() string {
	if i >= BlendOperation(len(_BlendOperation_index)-1) {
		return "BlendOperation(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BlendOperation_name[_BlendOperation_index[i]:_BlendOperation_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BlendFactorZero-0]
	_ = x[BlendFactorOne-1]
	_ = x[BlendFactorSrc-2]
	_ = x[BlendFactorOneMinusSrc-3]
	_ = x[BlendFactorSrcAlpha-4]
	_ = x[BlendFactorOneMinusSrcAlpha-5]
	_ = x[BlendFactorDst-6]
	_ = x[BlendFactorOneMinusDst-7]
	_ = x[BlendFactorDstAlpha-8]
	_ = x[BlendFactorOneMinusDstAlpha-9]
	_ = x[BlendFactorSrcAlphaSaturated-10]
	_ = x[BlendFactorConstant-11]
	_ = x[BlendFactorOneMinusConstant-12]
}

const _BlendFactor_name = "ZeroOneSrcOneMinusSrcSrcAlphaOneMinusSrcAlphaDstOneMinusDstDstAlphaOneMinusDstAlphaSrcAlphaSaturatedConstantOneMinusConstant"

var _BlendFactor_index = [...]uint8{0, 4, 7, 10, 21, 29, 45, 48, 59, 67, 83, 100, 108, 124}

func (i BlendFactor) String() string {
	if i >= BlendFactor(len(_BlendFactor_index)-1) {
		return "BlendFactor(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BlendFactor_name[_BlendFactor_index[i]:_BlendFactor_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ColorWriteMaskNone-0]
	_ = x[ColorWriteMaskRed-1]
	_ = x[ColorWriteMaskGreen-2]
	_ = x[ColorWriteMaskBlue-4]
	_ = x[ColorWriteMaskAlpha-8]
	_ = x[ColorWriteMaskAll-15]
}

const (
	_ColorWriteMask_name_0 = "NoneRedGreen"
	_ColorWriteMask_name_1 = "Blue"
	_ColorWriteMask_name_2 = "Alpha"
	_ColorWriteMask_name_3 = "All"
)

var (
	_ColorWriteMask_index_0 = [...]uint8{0, 4, 7, 12}
)

func (i ColorWriteMask) String() string {
	switch {
	case i <= 2:
		return _ColorWriteMask_name_0[_ColorWriteMask_index_0[i]:_ColorWriteMask_index_0[i+1]]
	case i == 4:
		return _ColorWriteMask_name_1
	case i == 8:
		return _ColorWriteMask_name_2
	case i == 15:
		return _ColorWriteMask_name_3
	default:
		return "ColorWriteMask(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VertexStepModeVertex-0]
	_ = x[VertexStepModeInstance-1]
	_ = x[VertexStepModeVertexBufferNotUsed-2]
}

const _VertexStepMode_name = "VertexInstanceVertexBufferNotUsed"

var _VertexStepMode_index = [...]uint8{0, 6, 14, 33}

func (i VertexStepMode) String() string {
	if i >= VertexStepMode(len(_VertexStepMode_index)-1) {
		return "VertexStepMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VertexStepMode_name[_VertexStepMode_index[i]:_VertexStepMode_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PresentModeFifo-0]
	_ = x[PresentModeFifoRelaxed-1]
	_ = x[PresentModeImmediate-2]
	_ = x[PresentModeMailbox-3]
}

const _PresentMode_name = "FifoFifoRelaxedImmediateMailbox"

var _PresentMode_index = [...]uint8{0, 4, 15, 24, 31}

func (i PresentMode) String() string {
	if i >= PresentMode(len(_PresentMode_index)-1) {
		return "PresentMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PresentMode_name[_PresentMode_index[i]:_PresentMode_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CompositeAlphaModeAuto-0]
	_ = x[CompositeAlphaModeOpaque-1]
	_ = x[CompositeAlphaModePremultiplied-2]
	_ = x[CompositeAlphaModeUnpremultiplied-3]
	_ = x[CompositeAlphaModeInherit-4]
}

const _CompositeAlphaMode_name = "AutoOpaquePremultipliedUnpremultipliedInherit"

var _CompositeAlphaMode_index = [...]uint8{0, 4, 10, 23, 38, 45}

func (i CompositeAlphaMode) String() string {
	if i >= CompositeAlphaMode(len(_CompositeAlphaMode_index)-1) {
		return "CompositeAlphaMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CompositeAlphaMode_name[_CompositeAlphaMode_index[i]:_CompositeAlphaMode_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TextureUsageNone-0]
	_ = x[TextureUsageCopySrc-1]
	_ = x[TextureUsageCopyDst-2]
	_ = x[TextureUsageTextureBinding-4]
	_ = x[TextureUsageStorageBinding-8]
	_ = x[TextureUsageRenderAttachment-16]
}

const (
	_TextureUsage_name_0 = "NoneCopySrcCopyDst"
	_TextureUsage_name_1 = "TextureBinding"
	_TextureUsage_name_2 = "StorageBinding"
	_TextureUsage_name_3 = "RenderAttachment"
)

var (
	_TextureUsage_index_0 = [...]uint8{0, 4, 11, 18}
)

func (i TextureUsage) String() string {
	switch {
	case i <= 2:
		return _TextureUsage_name_0[_TextureUsage_index_0[i]:_TextureUsage_index_0[i+1]]
	case i == 4:
		return _TextureUsage_name_1
	case i == 8:
		return _TextureUsage_name_2
	case i == 16:
		return _TextureUsage_name_3
	default:
		return "TextureUsage(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TextureViewDimensionUndefined-0]
	_ = x[TextureViewDimension1D-1]
	_ = x[TextureViewDimension2D-2]
	_ = x[TextureViewDimension2DArray-3]
	_ = x[TextureViewDimensionCube-4]
	_ = x[TextureViewDimensionCubeArray-5]
	_ = x[TextureViewDimension3D-6]
}

const _TextureViewDimension_name = "undefined1D2D2DArrayCubeCubeArray3D"

var _TextureViewDimension_index = [...]uint8{0, 9, 11, 13, 20, 24, 33, 35}

func (i TextureViewDimension) String() string {
	if i >= TextureViewDimension(len(_TextureViewDimension_index)-1) {
		return "TextureViewDimension(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TextureViewDimension_name[_TextureViewDimension_index[i]:_TextureViewDimension_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TextureAspectAll-0]
	_ = x[TextureAspectStencilOnly-1]
	_ = x[TextureAspectDepthOnly-2]
}

const _TextureAspect_name = "AllStencilOnlyDepthOnly"

var _TextureAspect_index = [...]uint8{0, 3, 14, 23}

func (i TextureAspect) String() string {
	if i >= TextureAspect(len(_TextureAspect_index)-1) {
		return "TextureAspect(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TextureAspect_name[_TextureAspect_index[i]:_TextureAspect_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LoadOpUndefined-0]
	_ = x[LoadOpClear-1]
	_ = x[LoadOpLoad-2]
}

const _LoadOp_name = "UndefinedClearLoad"

var _LoadOp_index = [...]uint8{0, 9, 14, 18}

func (i LoadOp) String() string {
	if i >= LoadOp(len(_LoadOp_index)-1) {
		return "LoadOp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LoadOp_name[_LoadOp_index[i]:_LoadOp_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StoreOpUndefined-0]
	_ = x[StoreOpStore-1]
	_ = x[StoreOpDiscard-2]
}

const _StoreOp_name = "UndefinedStoreDiscard"

var _StoreOp_index = [...]uint8{0, 9, 14, 21}

func (i StoreOp) String() string {
	if i >= StoreOp(len(_StoreOp_index)-1) {
		return "StoreOp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StoreOp_name[_StoreOp_index[i]:_StoreOp_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[QueryTypeOcclusion-0]
	_ = x[QueryTypeTimestamp-1]
}

const _QueryType_name = "OcclusionTimestamp"

var _QueryType_index = [...]uint8{0, 9, 18}

func (i QueryType) String() string {
	if i >= QueryType(len(_QueryType_index)-1) {
		return "QueryType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _QueryType_name[_QueryType_index[i]:_QueryType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BufferUsageNone-0]
	_ = x[BufferUsageMapRead-1]
	_ = x[BufferUsageMapWrite-2]
	_ = x[BufferUsageCopySrc-4]
	_ = x[BufferUsageCopyDst-8]
	_ = x[BufferUsageIndex-16]
	_ = x[BufferUsageVertex-32]
	_ = x[BufferUsageUniform-64]
	_ = x[BufferUsageStorage-128]
	_ = x[BufferUsageIndirect-256]
	_ = x[BufferUsageQueryResolve-512]
}

const (
	_BufferUsage_name_0 = "NoneMapReadMapWrite"
	_BufferUsage_name_1 = "CopySrc"
	_BufferUsage_name_2 = "CopyDst"
	_BufferUsage_name_3 = "Index"
	_BufferUsage_name_4 = "Vertex"
	_BufferUsage_name_5 = "Uniform"
	_BufferUsage_name_6 = "Storage"
	_BufferUsage_name_7 = "Indirect"
	_BufferUsage_name_8 = "QueryResolve"
)

var (
	_BufferUsage_index_0 = [...]uint8{0, 4, 11, 19}
)

func (i BufferUsage) String() string {
	switch {
	case i <= 2:
		return _BufferUsage_name_0[_BufferUsage_index_0[i]:_BufferUsage_index_0[i+1]]
	case i == 4:
		return _BufferUsage_name_1
	case i == 8:
		return _BufferUsage_name_2
	case i == 16:
		return _BufferUsage_name_3
	case i == 32:
		return _BufferUsage_name_4
	case i == 64:
		return _BufferUsage_name_5
	case i == 128:
		return _BufferUsage_name_6
	case i == 256:
		return _BufferUsage_name_7
	case i == 512:
		return _BufferUsage_name_8
	default:
		return "BufferUsage(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MapModeNone-0]
	_ = x[MapModeRead-1]
	_ = x[MapModeWrite-2]
}

const _MapMode_name = "NoneReadWrite"

var _MapMode_index = [...]uint8{0, 4, 8, 13}

func (i MapMode) String() string {
	if i >= MapMode(len(_MapMode_index)-1) {
		return "MapMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MapMode_name[_MapMode_index[i]:_MapMode_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BufferMapStateUnmapped-0]
	_ = x[BufferMapStatePending-1]
	_ = x[BufferMapStateMapped-2]
}

const _BufferMapState_name = "UnmappedPendingMapped"

var _BufferMapState_index = [...]uint8{0, 8, 15, 21}

func (i BufferMapState) String() string {
	if i >= BufferMapState(len(_BufferMapState_index)-1) {
		return "BufferMapState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BufferMapState_name[_BufferMapState_index[i]:_BufferMapState_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ShaderStageNone-0]
	_ = x[ShaderStageVertex-1]
	_ = x[ShaderStageFragment-2]
	_ = x[ShaderStageCompute-4]
}

const (
	_ShaderStage_name_0 = "NoneVertexFragment"
	_ShaderStage_name_1 = "Compute"
)

var (
	_ShaderStage_index_0 = [...]uint8{0, 4, 10, 18}
)

func (i ShaderStage) String() string {
	switch {
	case i <= 2:
		return _ShaderStage_name_0[_ShaderStage_index_0[i]:_ShaderStage_index_0[i+1]]
	case i == 4:
		return _ShaderStage_name_1
	default:
		return "ShaderStage(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BufferBindingTypeUndefined-0]
	_ = x[BufferBindingTypeUniform-1]
	_ = x[BufferBindingTypeStorage-2]
	_ = x[BufferBindingTypeReadOnlyStorage-3]
}

const _BufferBindingType_name = "UndefinedUniformStorageReadOnlyStorage"

var _BufferBindingType_index = [...]uint8{0, 9, 16, 23, 38}

func (i BufferBindingType) String() string {
	if i >= BufferBindingType(len(_BufferBindingType_index)-1) {
		return "BufferBindingType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BufferBindingType_name[_BufferBindingType_index[i]:_BufferBindingType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SamplerBindingTypeUndefined-0]
	_ = x[SamplerBindingTypeFiltering-1]
	_ = x[SamplerBindingTypeNonFiltering-2]
	_ = x[SamplerBindingTypeComparison-3]
}

const _SamplerBindingType_name = "UndefinedFilteringNonFilteringComparison"

var _SamplerBindingType_index = [...]uint8{0, 9, 18, 30, 40}

func (i SamplerBindingType) String() string {
	if i >= SamplerBindingType(len(_SamplerBindingType_index)-1) {
		return "SamplerBindingType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SamplerBindingType_name[_SamplerBindingType_index[i]:_SamplerBindingType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TextureSampleTypeUndefined-0]
	_ = x[TextureSampleTypeFloat-1]
	_ = x[TextureSampleTypeUnfilterableFloat-2]
	_ = x[TextureSampleTypeDepth-3]
	_ = x[TextureSampleTypeSint-4]
	_ = x[TextureSampleTypeUint-5]
}

const _TextureSampleType_name = "UndefinedFloatUnfilterableFloatDepthSintUint"

var _TextureSampleType_index = [...]uint8{0, 9, 14, 31, 36, 40, 44}

func (i TextureSampleType) String() string {
	if i >= TextureSampleType(len(_TextureSampleType_index)-1) {
		return "TextureSampleType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TextureSampleType_name[_TextureSampleType_index[i]:_TextureSampleType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StorageTextureAccessUndefined-0]
	_ = x[StorageTextureAccessWriteOnly-1]
	_ = x[StorageTextureAccessReadOnly-2]
	_ = x[StorageTextureAccessReadWrite-3]
}

const _StorageTextureAccess_name = "UndefinedWriteOnlyReadOnlyReadWrite"

var _StorageTextureAccess_index = [...]uint8{0, 9, 18, 26, 35}

func (i StorageTextureAccess) String() string {
	if i >= StorageTextureAccess(len(_StorageTextureAccess_index)-1) {
		return "StorageTextureAccess(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StorageTextureAccess_name[_StorageTextureAccess_index[i]:_StorageTextureAccess_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TextureDimension1D-0]
	_ = x[TextureDimension2D-1]
	_ = x[TextureDimension3D-2]
}

const _TextureDimension_name = "1D2D3D"

var _TextureDimension_index = [...]uint8{0, 2, 4, 6}

func (i TextureDimension) String() string {
	if i >= TextureDimension(len(_TextureDimension_index)-1) {
		return "TextureDimension(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TextureDimension_name[_TextureDimension_index[i]:_TextureDimension_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AddressModeRepeat-0]
	_ = x[AddressModeMirrorRepeat-1]
	_ = x[AddressModeClampToEdge-2]
}

const _AddressMode_name = "RepeatMirrorRepeatClampToEdge"

var _AddressMode_index = [...]uint8{0, 6, 18, 29}

func (i AddressMode) String() string {
	if i >= AddressMode(len(_AddressMode_index)-1) {
		return "AddressMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AddressMode_name[_AddressMode_index[i]:_AddressMode_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FilterModeNearest-0]
	_ = x[FilterModeLinear-1]
}

const _FilterMode_name = "NearestLinear"

var _FilterMode_index = [...]uint8{0, 7, 13}

func (i FilterMode) String() string {
	if i >= FilterMode(len(_FilterMode_index)-1) {
		return "FilterMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FilterMode_name[_FilterMode_index[i]:_FilterMode_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MipmapFilterModeNearest-0]
	_ = x[MipmapFilterModeLinear-1]
}

const _MipmapFilterMode_name = "NearestLinear"

var _MipmapFilterMode_index = [...]uint8{0, 7, 13}

func (i MipmapFilterMode) String() string {
	if i >= MipmapFilterMode(len(_MipmapFilterMode_index)-1) {
		return "MipmapFilterMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MipmapFilterMode_name[_MipmapFilterMode_index[i]:_MipmapFilterMode_index[i+1]]
}