// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ct

import "github.com/richardwilkes/macos/cf"

// #import <CoreText/CoreText.h>
import "C"

type FontCollection = C.CTFontCollectionRef

func FontCollectionCreateFromAvailableFonts(options cf.Dictionary) FontCollection {
	return C.CTFontCollectionCreateFromAvailableFonts(C.CFDictionaryRef(options))
}

func (fc FontCollection) CreateMatchingFontDescriptors() cf.Array {
	return cf.Array(C.CTFontCollectionCreateMatchingFontDescriptors(fc))
}

func (fc FontCollection) Release() {
	C.CFRelease(C.CFTypeRef(fc))
}
