// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package cf

import "unsafe"

// #import <CoreFoundation/CoreFoundation.h>
import "C"

type String = C.CFStringRef

func StringCreateWithString(str string) String {
	return StringCreateWithBytes([]byte(str))
}

func StringCreateWithBytes(bytes []byte) String {
	var ptr *C.UInt8
	if len(bytes) > 0 {
		ptr = (*C.UInt8)(unsafe.Pointer(&bytes[0]))
	}
	return C.CFStringCreateWithBytes(0, ptr, C.CFIndex(len(bytes)), C.kCFStringEncodingUTF8, 0)
}

func (s String) GetLength() int {
	return int(C.CFStringGetLength(s))
}

func (s String) String() string {
	var allocedPtr *C.char
	strPtr := C.CFStringGetCStringPtr(s, C.kCFStringEncodingUTF8)
	if strPtr == nil {
		stringLength := s.GetLength()
		maxBytes := 4*stringLength + 1
		allocedPtr = (*C.char)(C.malloc(C.size_t(maxBytes)))
		C.CFStringGetCString(s, allocedPtr, C.CFIndex(maxBytes), C.kCFStringEncodingUTF8)
		strPtr = allocedPtr
	}
	str := C.GoString(strPtr)
	if allocedPtr != nil {
		C.free(unsafe.Pointer(allocedPtr))
	}
	return str
}

func (s String) Release() {
	C.CFRelease(Type(s))
}
