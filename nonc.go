// This file contains some of the functionality offered by wgpu's Rust API that
// is absent from the C API. It currently only contains the API that is required
// by jello.

package wgpu

func (tf TextureFormat) BlockCopySize(aspect TextureAspect) (uint32, bool) {
	switch tf {
	case TextureFormatR8Unorm,
		TextureFormatR8Snorm,
		TextureFormatR8Uint,
		TextureFormatR8Sint:
		return 1, true

	case TextureFormatRG8Unorm,
		TextureFormatRG8Snorm,
		TextureFormatRG8Uint,
		TextureFormatRG8Sint:
		return 2, true

	case /* TextureFormatR16Unorm,
		TextureFormatR16Snorm,
		*/TextureFormatR16Uint,
		TextureFormatR16Sint,
		TextureFormatR16Float:
		return 2, true

	case TextureFormatRGBA8Unorm,
		TextureFormatRGBA8UnormSrgb,
		TextureFormatRGBA8Snorm,
		TextureFormatRGBA8Uint,
		TextureFormatRGBA8Sint,
		TextureFormatBGRA8Unorm,
		TextureFormatBGRA8UnormSrgb:
		return 4, true

	case /* TextureFormatRG16Unorm,
		TextureFormatRG16Snorm,
		*/TextureFormatRG16Uint,
		TextureFormatRG16Sint,
		TextureFormatRG16Float:
		return 4, true

	case TextureFormatR32Uint,
		TextureFormatR32Sint,
		TextureFormatR32Float:
		return 4, true

	case TextureFormatRGB9E5Ufloat,
		TextureFormatRGB10A2Uint,
		TextureFormatRGB10A2Unorm,
		TextureFormatRG11B10Ufloat:
		return 4, true

	case /* TextureFormatRGBA16Unorm,
		TextureFormatRGBA16Snorm,
		*/TextureFormatRGBA16Uint,
		TextureFormatRGBA16Sint,
		TextureFormatRGBA16Float:
		return 8, true

	case TextureFormatRG32Uint,
		TextureFormatRG32Sint,
		TextureFormatRG32Float:
		return 8, true

	case TextureFormatRGBA32Uint,
		TextureFormatRGBA32Sint,
		TextureFormatRGBA32Float:
		return 16, true

	case TextureFormatStencil8:
		return 1, true

	case TextureFormatDepth16Unorm:
		return 2, true

	case TextureFormatDepth32Float:
		return 4, true

	case TextureFormatDepth24Plus:
		return 0, false

	case TextureFormatDepth24PlusStencil8:
		switch aspect {
		case TextureAspectDepthOnly:
			return 0, false
		case TextureAspectStencilOnly:
			return 1, true
		default:
			return 0, false
		}

	case TextureFormatDepth32FloatStencil8:
		switch aspect {
		case TextureAspectDepthOnly:
			return 4, true
		case TextureAspectStencilOnly:
			return 1, true
		default:
			return 0, false
		}

	// case TextureFormatNV12:
	// 	switch aspect {
	// 	case TextureAspect__Plane0:
	// 		return 1, true
	// 	case TextureAspect__Plane1:
	// 		return 2, true
	// 	default:
	// 		return 0, false
	// 	}

	case TextureFormatBC1RGBAUnorm,
		TextureFormatBC1RGBAUnormSrgb,
		TextureFormatBC4RUnorm,
		TextureFormatBC4RSnorm:
		return 8, true

	case TextureFormatBC2RGBAUnorm,
		TextureFormatBC2RGBAUnormSrgb,
		TextureFormatBC3RGBAUnorm,
		TextureFormatBC3RGBAUnormSrgb,
		TextureFormatBC5RGUnorm,
		TextureFormatBC5RGSnorm,
		TextureFormatBC6HRGBUfloat,
		TextureFormatBC6HRGBFloat,
		TextureFormatBC7RGBAUnorm,
		TextureFormatBC7RGBAUnormSrgb:
		return 16, true

	case TextureFormatETC2RGB8Unorm,
		TextureFormatETC2RGB8UnormSrgb,
		TextureFormatETC2RGB8A1Unorm,
		TextureFormatETC2RGB8A1UnormSrgb,
		TextureFormatEACR11Unorm,
		TextureFormatEACR11Snorm:
		return 8, true

	case TextureFormatETC2RGBA8Unorm,
		TextureFormatETC2RGBA8UnormSrgb,
		TextureFormatEACRG11Unorm,
		TextureFormatEACRG11Snorm:
		return 16, true

	case TextureFormatASTC4x4Unorm,
		TextureFormatASTC4x4UnormSrgb,
		TextureFormatASTC5x4Unorm,
		TextureFormatASTC5x4UnormSrgb,
		TextureFormatASTC5x5Unorm,
		TextureFormatASTC5x5UnormSrgb,
		TextureFormatASTC6x5Unorm,
		TextureFormatASTC6x5UnormSrgb,
		TextureFormatASTC6x6Unorm,
		TextureFormatASTC6x6UnormSrgb,
		TextureFormatASTC8x5Unorm,
		TextureFormatASTC8x5UnormSrgb,
		TextureFormatASTC8x6Unorm,
		TextureFormatASTC8x6UnormSrgb,
		TextureFormatASTC8x8Unorm,
		TextureFormatASTC8x8UnormSrgb,
		TextureFormatASTC10x5Unorm,
		TextureFormatASTC10x5UnormSrgb,
		TextureFormatASTC10x6Unorm,
		TextureFormatASTC10x6UnormSrgb,
		TextureFormatASTC10x8Unorm,
		TextureFormatASTC10x8UnormSrgb,
		TextureFormatASTC10x10Unorm,
		TextureFormatASTC10x10UnormSrgb,
		TextureFormatASTC12x10Unorm,
		TextureFormatASTC12x10UnormSrgb,
		TextureFormatASTC12x12Unorm,
		TextureFormatASTC12x12UnormSrgb:
		return 16, true

	default:
		return 0, false
	}
}
