// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package cf

// #import <CoreFoundation/CoreFoundation.h>
import "C"

type (
	AttributedString        = C.CFAttributedStringRef
	MutableAttributedString = C.CFMutableAttributedStringRef
)

func AttributedStringCreateMutable(maxLength int) MutableAttributedString {
	return C.CFAttributedStringCreateMutable(0, C.CFIndex(maxLength))
}

func (s MutableAttributedString) BeginEditing() {
	C.CFAttributedStringBeginEditing(s)
}

func (s MutableAttributedString) ReplaceString(location, length int, replacement String) {
	C.CFAttributedStringReplaceString(s, C.CFRangeMake(C.CFIndex(location), C.CFIndex(length)), replacement)
}

func (s MutableAttributedString) SetAttribute(location, length int, attrName String, value Type) {
	C.CFAttributedStringSetAttribute(s, C.CFRangeMake(C.CFIndex(location), C.CFIndex(length)), attrName, value)
}

func (s MutableAttributedString) EndEditing() {
	C.CFAttributedStringEndEditing(s)
}

func (s MutableAttributedString) Release() {
	C.CFRelease(Type(s))
}
