// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ct

import (
	"unsafe"

	"github.com/richardwilkes/macos/cf"
)

/*
#import <CoreText/CoreText.h>

void myRegisterFontDescriptors(CFArrayRef descriptors, CTFontManagerScope scope, bool enabled) {
	CTFontManagerRegisterFontDescriptors(descriptors, scope, enabled, nil);
}
*/
import "C"

type FontManagerScope uint32

const (
	FontManagerScopeNone FontManagerScope = iota
	FontManagerScopeProcess
	FontManagerScopeUser
	FontManagerScopeSession
)

func FontManagerCreateFontDescriptorFromData(data cf.Data) FontDescriptor {
	return C.CTFontManagerCreateFontDescriptorFromData(C.CFDataRef(data))
}

func FontManagerEnableFontDescriptors(enable bool, descriptors ...FontDescriptor) {
	if len(descriptors) > 0 {
		a := cf.MutableArrayCreate(len(descriptors))
		for i := range descriptors {
			a.AppendValue(unsafe.Pointer(descriptors[i])) //nolint:govet
		}
		ma := a.AsCFArray()
		C.CTFontManagerEnableFontDescriptors(C.CFArrayRef(ma), C.bool(enable))
		ma.Release()
	}
}

func FontManagerRegisterFontDescriptors(scope FontManagerScope, enabled bool, descriptors ...FontDescriptor) {
	if len(descriptors) > 0 {
		a := cf.MutableArrayCreate(len(descriptors))
		for i := range descriptors {
			a.AppendValue(unsafe.Pointer(descriptors[i])) //nolint:govet
		}
		ma := a.AsCFArray()
		C.myRegisterFontDescriptors(C.CFArrayRef(ma), C.CTFontManagerScope(scope), C.bool(enabled))
		ma.Release()
	}
}
