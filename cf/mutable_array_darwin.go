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

type MutableArray struct {
	native C.CFMutableArrayRef
}

func MutableArrayCreate(capacity int) *MutableArray {
	return &MutableArray{native: C.CFArrayCreateMutable(0, C.CFIndex(capacity), &C.kCFTypeArrayCallBacks)} //nolint:gocritic,staticcheck
}

func MutableArrayCreateNoCap() *MutableArray {
	return MutableArrayCreate(0)
}

func (a *MutableArray) AppendValue(value unsafe.Pointer) {
	C.CFArrayAppendValue(a.native, value)
}

func (a *MutableArray) AppendStringValue(value string) {
	C.CFArrayAppendValue(a.native, unsafe.Pointer(StringCreateWithString(value))) //nolint:govet
}

func (a *MutableArray) AsCFArray() Array {
	return Array(unsafe.Pointer(a.native))
}
