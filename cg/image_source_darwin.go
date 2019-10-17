package cg

import "github.com/richardwilkes/macos/cf"

// #import <ImageIO/ImageIO.h>
import "C"

type ImageSource = C.CGImageSourceRef

func ImageSourceCreateWithData(data cf.Data, options cf.Dictionary) ImageSource {
	return C.CGImageSourceCreateWithData(C.CFDataRef(data), C.CFDictionaryRef(options))
}

func (is ImageSource) CreateImageAtIndex(index int, options cf.Dictionary) Image {
	return C.CGImageSourceCreateImageAtIndex(is, C.size_t(index), C.CFDictionaryRef(options))
}

func (is ImageSource) Release() {
	C.CFRelease(C.CFTypeRef(is))
}
