package cg

// #import <CoreGraphics/CoreGraphics.h>
import "C"

type DirectDisplayID = C.CGDirectDisplayID

func DisplayIsMain(id DirectDisplayID) bool {
	return C.CGDisplayIsMain(id) != 0
}

func DisplayBounds(id DirectDisplayID) (x, y, width, height float64) {
	r := C.CGDisplayBounds(id)
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}
