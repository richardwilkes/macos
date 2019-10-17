package cg

// #import <CoreGraphics/CoreGraphics.h>
import "C"

type ColorSpace = C.CGColorSpaceRef

func ColorSpaceCreateDeviceRGB() ColorSpace {
	return C.CGColorSpaceCreateDeviceRGB()
}

func ColorSpaceCreatePattern(baseSpace ColorSpace) ColorSpace {
	return C.CGColorSpaceCreatePattern(baseSpace)
}

func (cs ColorSpace) Release() {
	C.CGColorSpaceRelease(cs)
}
