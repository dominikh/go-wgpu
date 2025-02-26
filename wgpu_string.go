// Code generated by "stringer -type LogLevel,InstanceBackend,InstanceFlag,DX12Compiler,GLES3MinorVersion -output wgpu_string.go -linecomment"; DO NOT EDIT.

package wgpu

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LogLevelOff-0]
	_ = x[LogLevelError-1]
	_ = x[LogLevelWarn-2]
	_ = x[LogLevelInfo-3]
	_ = x[LogLevelDebug-4]
	_ = x[LogLevelTrace-5]
}

const _LogLevel_name = "OffErrorWarnInfoDebugTrace"

var _LogLevel_index = [...]uint8{0, 3, 8, 12, 16, 21, 26}

func (i LogLevel) String() string {
	if i >= LogLevel(len(_LogLevel_index)-1) {
		return "LogLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LogLevel_name[_LogLevel_index[i]:_LogLevel_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InstanceBackendAll-0]
	_ = x[InstanceBackendVulkan-1]
	_ = x[InstanceBackendGL-2]
	_ = x[InstanceBackendMetal-4]
	_ = x[InstanceBackendDX12-8]
	_ = x[InstanceBackendDX11-16]
	_ = x[InstanceBackendBrowserWebGPU-32]
	_ = x[InstanceBackendPrimary-45]
	_ = x[InstanceBackendSecondary-18]
}

const (
	_InstanceBackend_name_0 = "AllVulkanGL"
	_InstanceBackend_name_1 = "Metal"
	_InstanceBackend_name_2 = "DX12"
	_InstanceBackend_name_3 = "DX11"
	_InstanceBackend_name_4 = "Secondary"
	_InstanceBackend_name_5 = "BrowserWebGPU"
	_InstanceBackend_name_6 = "Primary"
)

var (
	_InstanceBackend_index_0 = [...]uint8{0, 3, 9, 11}
)

func (i InstanceBackend) String() string {
	switch {
	case i <= 2:
		return _InstanceBackend_name_0[_InstanceBackend_index_0[i]:_InstanceBackend_index_0[i+1]]
	case i == 4:
		return _InstanceBackend_name_1
	case i == 8:
		return _InstanceBackend_name_2
	case i == 16:
		return _InstanceBackend_name_3
	case i == 18:
		return _InstanceBackend_name_4
	case i == 32:
		return _InstanceBackend_name_5
	case i == 45:
		return _InstanceBackend_name_6
	default:
		return "InstanceBackend(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InstanceFlagDefault-0]
	_ = x[InstanceFlagDebug-1]
	_ = x[InstanceFlagValidation-2]
	_ = x[InstanceFlagDiscardHalLabels-4]
}

const (
	_InstanceFlag_name_0 = "DefaultDebugValidation"
	_InstanceFlag_name_1 = "DiscardHalLabels"
)

var (
	_InstanceFlag_index_0 = [...]uint8{0, 7, 12, 22}
)

func (i InstanceFlag) String() string {
	switch {
	case i <= 2:
		return _InstanceFlag_name_0[_InstanceFlag_index_0[i]:_InstanceFlag_index_0[i+1]]
	case i == 4:
		return _InstanceFlag_name_1
	default:
		return "InstanceFlag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DX12CompilerUndefined-0]
	_ = x[DX12CompilerFxc-1]
	_ = x[DX12CompilerDxc-2]
}

const _DX12Compiler_name = "UndefinedFxcDxc"

var _DX12Compiler_index = [...]uint8{0, 9, 12, 15}

func (i DX12Compiler) String() string {
	if i >= DX12Compiler(len(_DX12Compiler_index)-1) {
		return "DX12Compiler(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DX12Compiler_name[_DX12Compiler_index[i]:_DX12Compiler_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GLES3MinorVersionAutomatic-0]
	_ = x[GLES3MinorVersionVersion0-1]
	_ = x[GLES3MinorVersionVersion1-2]
	_ = x[GLES3MinorVersionVersion2-3]
}

const _GLES3MinorVersion_name = "AutomaticVersion0Version1Version2"

var _GLES3MinorVersion_index = [...]uint8{0, 9, 17, 25, 33}

func (i GLES3MinorVersion) String() string {
	if i >= GLES3MinorVersion(len(_GLES3MinorVersion_index)-1) {
		return "GLES3MinorVersion(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GLES3MinorVersion_name[_GLES3MinorVersion_index[i]:_GLES3MinorVersion_index[i+1]]
}
