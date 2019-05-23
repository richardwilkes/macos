package macos

// #import <CoreGraphics/CoreGraphics.h>
import "C"

type CGColorSpace = C.CGColorSpaceRef

func CGColorSpaceCreateDeviceRGB() CGColorSpace {
	return C.CGColorSpaceCreateDeviceRGB()
}

func CGColorSpaceCreatePattern(baseSpace CGColorSpace) CGColorSpace {
	return C.CGColorSpaceCreatePattern(baseSpace)
}

func (cs CGColorSpace) Release() {
	C.CGColorSpaceRelease(cs)
}
