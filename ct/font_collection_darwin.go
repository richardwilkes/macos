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
