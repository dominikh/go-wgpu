// # API stability
//
// WebGPU, webgpu.h, and wgpu are still in development and change regularly. Our
// bindings works with specific versions of webgpu.h and wgpu only, and will
// regularly be updated to track changes, which means our API might change as well.
//
// # Error handling
//
// Error handling in WebGPU, webgpu.h and wgpu is still in flux and incomplete,
// making it impossible to handle some errors cleanly, as they will panic in
// Rust and cause SIGABRT. The errors that _can_ be handled fall into one of the
// following categories:
//
// Failure to acquire an adapter or device will return a Go error.
//
// By default, all validation errors, out-of-memory errors, and internal errors
// will panic in Go.
//
// Using [Device.PushErrorScope] and [Device.PopErrorScope], validation errors,
// out-of-memory errors, and internal errors can be caught and turned into Go
// errors. This allows, for example, to handle fallible memory allocations.
//
// The alternative to explicit use of error scopes would've been to return Go
// errors from most functions in this package. However, the vast majority of
// validation errors are programmer mistakes, which should panic, not return
// errors. Similarly, a lot of out of memory situations are fatal. Furthermore,
// instead of checking the error for every function call, it often suffices to
// check them for a sequence of function calls, for example to check that
// rendering a full frame succeeded. This can be accomplished much easier with
// error scopes. Finally, fine-grained use of error scopes may lead to
// unnecessary synchronisation between the CPU and GPU.
//
// Some functions are internally implemented using callbacks, such as
// [Surface.CurrentTexture]. Errors that occur in such functions do not use
// error scopes and are instead passed to the callbacks, and we turn them into
// Go errors.
//
// Device loss will invoke a callback. WebGPU intends that most API continues to
// work, by becoming no-ops, after a device has been lost, so that device loss
// can be handled reliably via the callback. This hasn't been fully implemented
// yet, and handling device loss reliably is error-prone, no pun intended.
//
// Due to a combination of the designs of webgpu.h and wgpu, some errors related
// to resources that do not belong to a device cannot be handled, as they cause
// panics inside Rust. This applies to surface creation and configuration, for
// example.
package wgpu

// https://github.com/gfx-rs/wgpu/issues/5132 -  Mark the device as lost when hal produces a device lost error
// https://github.com/gfx-rs/wgpu/issues/1624 - Figure out how to deal with device loss
// https://github.com/gpuweb/gpuweb/issues/1629 -  Make lost devices appear to function as much as possible
// https://github.com/gfx-rs/wgpu-native/issues/113 -  When will panics become errors that can be handled?
