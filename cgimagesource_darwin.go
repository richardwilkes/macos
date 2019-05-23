package macos

// #import <ImageIO/ImageIO.h>
import "C"

type CGImageSource = C.CGImageSourceRef

func CGImageSourceCreateWithData(data CFData, options CFDictionary) CGImageSource {
	return C.CGImageSourceCreateWithData(data, options)
}

func (is CGImageSource) CreateImageAtIndex(index int, options CFDictionary) CGImage {
	return C.CGImageSourceCreateImageAtIndex(is, C.size_t(index), options)
}

func (is CGImageSource) Release() {
	C.CFRelease(C.CFTypeRef(is))
}
