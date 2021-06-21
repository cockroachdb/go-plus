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

// +build goplus

package proc

import "runtime"

func internalSupported() bool { return true }

func internalGetGID() int64 {
	return runtime.GetGID()
}

func internalGetIGID() int64 {
	return runtime.GetLogicalTaskGroupID()
}

func internalSetIGID(igid int64) {
	runtime.SetLogicalTaskGroupID(igid)
}

func internalGetPID() int32 {
	return runtime.GetPID()
}

func internalPinGoroutine() int {
	return runtime.PinGoroutine()
}

func internalUnpinGoroutine() {
	runtime.UnpinGoroutine()
}

func internalSetDefaultPanicOnFault(enabled bool) bool {
	return runtime.SetDefaultPanicOnFault(enabled)
}

func internalGetDefaultPanicOnFault() bool {
	return runtime.GetDefaultPanicOnFault()
}

func internalGetOSThreadID() uint64 {
	return runtime.GetOSThreadID()
}

func internalGetGoroutineStackSize() uintptr {
	return runtime.GetGoroutineStackSize()
}

func internalGetExternalErrorFd() int {
	return runtime.GetExternalErrorFd()
}

func internalSetExternalErrorFd(fd int) {
	runtime.SetExternalErrorFd(fd)
}
