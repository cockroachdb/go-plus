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

// +build !crlext

package proc

import "github.com/petermattis/goid"

func internalSupported() bool { return false }

func internalGetGID() int64 {
	return goid.Get()
}

func internalGetIGID() int64 {
	return 0
}

func internalSetIGID(igid int64) {
}

func internalGetPID() int32 {
	return 0
}

func internalPinGoroutine() int {
	return -1
}

func internalUnpinGoroutine() {
}

func internalSetDefaultPanicOnFault(enabled bool) bool {
	return false
}

func internalGetDefaultPanicOnFault() bool {
	return false
}

func internalGetOSThreadID() uint64 {
	return 0
}

func internalGetGoroutineStackSize() uintptr {
	return 0
}

func internalGetExternalErrorFd() int {
	return -1
}

func internalSetExternalErrorFd(fd int) {
}
