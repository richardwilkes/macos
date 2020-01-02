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

type Data = C.CFDataRef

func DataCreate(data []byte) Data {
	return C.CFDataCreate(0, (*C.UInt8)(&data[0]), C.CFIndex(len(data)))
}

func (d Data) GetLength() int {
	return int(C.CFDataGetLength(d))
}

func (d Data) GetBytes(location, length int) []byte {
	buffer := make([]byte, length)
	C.CFDataGetBytes(d, C.CFRangeMake(C.CFIndex(location), C.CFIndex(length)), (*C.UInt8)(unsafe.Pointer(&buffer[0])))
	return buffer
}

func (d Data) Release() {
	C.CFRelease(Type(d))
}
