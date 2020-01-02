// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ut

import "github.com/richardwilkes/macos/cf"

// #import <Cocoa/Cocoa.h>
import "C"

func TypeCreatePreferredIdentifierForTagClassMimeType(mimeType string) string {
	tag := cf.StringCreateWithString(mimeType)
	defer tag.Release()
	uti := cf.String(C.UTTypeCreatePreferredIdentifierForTag(C.kUTTagClassMIMEType, C.CFStringRef(tag), 0))
	if uti == 0 {
		return ""
	}
	defer uti.Release()
	return uti.String()
}

func TypeCopyPreferredTagWithClassMimeType(uti string) string {
	tag := cf.StringCreateWithString(uti)
	defer tag.Release()
	mimeType := cf.String(C.UTTypeCopyPreferredTagWithClass(C.CFStringRef(tag), C.kUTTagClassMIMEType))
	if mimeType == 0 {
		return ""
	}
	defer mimeType.Release()
	return mimeType.String()
}
