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

type Array = C.CFArrayRef

func (a Array) GetCount() int {
	return int(C.CFArrayGetCount(a))
}

func (a Array) GetValueAtIndex(index int) unsafe.Pointer {
	return C.CFArrayGetValueAtIndex(a, C.CFIndex(index))
}

func (a Array) Release() {
	C.CFRelease(Type(a))
}
