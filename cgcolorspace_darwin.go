package macos

// #import <CoreGraphics/CoreGraphics.h>
import "C"

type CGColorSpace = C.CGColorSpaceRef

func CGColorSpaceCreateDeviceRGB() CGColorSpace {
	return C.CGColorSpaceCreateDeviceRGB()
}

func (cs CGColorSpace) Release() {
	C.CGColorSpaceRelease(cs)
}
