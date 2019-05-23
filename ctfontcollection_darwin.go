package macos

// #import <CoreText/CoreText.h>
import "C"

type CTFontCollection = C.CTFontCollectionRef

func CTFontCollectionCreateFromAvailableFonts(options CFDictionary) CTFontCollection {
	return C.CTFontCollectionCreateFromAvailableFonts(options)
}

func (fc CTFontCollection) CreateMatchingFontDescriptors() CFArray {
	return C.CTFontCollectionCreateMatchingFontDescriptors(fc)
}

func (fc CTFontCollection) Release() {
	C.CFRelease(CFType(fc))
}
