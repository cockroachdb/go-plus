// Copyright 2021 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package proc provides more fine-grained control over the Go runtime.
//
// It also creates a new abstration: the inheritable goroutine ID, or
// IGID, which is automatically copied from parent to children
// goroutines.
//
// Programs should use the Supported() API function to determine whether
// the Go runtime extension is available.
package proc

// Supported returns true iff the Go ID extension is available
// in the Go runtime.
func Supported() bool {
	return internalSupported()
}

// GetGID retrieves the goroutine (G's) ID.
//
// This works even when the extension is not supported,
// by relying on github.com/petermattis/goid.
func GetGID() int64 {
	return internalGetGID()
}

// GetIGID retrieves the inheritable goroutine ID (IGID).
// This is inherited from the goroutine's parent.
// Top-level goroutines are assigned their own ID as group ID.
//
// Returns 0 if the extension is not supported.
func GetIGID() int64 {
	return internalGetIGID()
}

// SetIGID sets the current goroutine's inheritable
// goroutine ID.
// This value is inherited to children goroutines.
//
// A possible good value to use is the value of GetGID(),
// to obtain a behavior similar to process group IDs on unix.
//
// No-op if the extension is not supported.
func SetIGID(igid int64) {
	internalSetIGID(igid)
}

// GetPID retrieves the P (virtual CPU) ID of the current goroutine.
// These are by design set between 0 and gomaxprocs - 1.
// This can be used to create data structures with CPU affinity.
//
// Note: because of preemption, there is no guarantee that the
// goroutine remains on the same P after the call to GetPID()
// completes. See Pin/UnpinGoroutine().
//
// Returns 0 if the extension is not supported.
func GetPID() int32 {
	return internalGetPID()
}

// PinGoroutine pins the current G to its P and disables preemption.
// The caller is responsible for calling Unpin. The GC is not running
// between Pin and Unpin.
// Returns the ID of the P that the G has been pinned to.
//
// Returns -1 if the extension is not supported.
func PinGoroutine() int {
	return internalPinGoroutine()
}

// UnpinGoroutine unpints the current G from its P.
// The caller is responsible for ensuring that PinGoroutine()
// has been called first. The behavior is undefined
// otherwise.
//
// No-op if the extension is not supported.
func UnpinGoroutine() {
	internalUnpinGoroutine()
}

// SetDefaultPanicOnFault sets the default value of the "panic on
// fault" that is otherwise customizable with debug.SetPanicOnFault.
//
// The default value is used for "top level" goroutines that are
// started during the init process / runtime. The flag is inherited to
// children goroutines.
// When not customized, the go runtime uses "false" as default.
//
// Returns false if the extension is not supported.
func SetDefaultPanicOnFault(enabled bool) bool {
	return internalSetDefaultPanicOnFault(enabled)
}

// GetDefaultPanicOnFault retrieves the default value of the "panic on
// fault" flag.
//
// Returns false if the extension is not supported.
func GetDefaultPanicOnFault() bool {
	return internalGetDefaultPanicOnFault()
}

// GetOSThreadID retrieves the OS-level thread ID, which can be used to e.g.
// set scheduling policies or retrieve CPU usage statistics.
//
// Note: because of preemption, there is no guarantee that the current
// goroutine remains on the same OS thread after a call to this
// function completes. See Pin/UnpinGoroutine.
//
// Returns 0 if the extension is not supported.
func GetOSThreadID() uint64 {
	return internalGetOSThreadID()
}

// GetGoroutineStackSize retrieves an approximation of the current goroutine (G)'s
// stack size.
//
// Returns 0 if the extension is not supported.
func GetGoroutineStackSize() uintptr {
	return internalGetGoroutineStackSize()
}

// GetExternalErrorFd retrieves the file descriptor used to emit
// panic messages and other internal errors by the go runtime.
//
// Returns -1 if the extension is not supported.
func GetExternalErrorFd() int {
	return internalGetExternalErrorFd()
}

// SetExternalErrorFd changes the file descriptor used to emit panic
// messages and other internal errors by the go runtime.
//
// No-op if the extension is not supported.
func SetExternalErrorFd(fd int) {
	internalSetExternalErrorFd(fd)
}
