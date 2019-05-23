package macos

// #import <CoreGraphics/CoreGraphics.h>
import "C"

const (
	CGGradientDrawsBeforeStartLocation CGGradientDrawingOptions = 1 << iota
	CGGradientDrawsAfterEndLocation
)

type (
	CGGradient               = C.CGGradientRef
	CGGradientDrawingOptions int
)

func CGGradientCreateWithColorComponents(space CGColorSpace, components, locations []float64) CGGradient {
	return C.CGGradientCreateWithColorComponents(space, floatSliceToCGFloatPtr(components), floatSliceToCGFloatPtr(locations), C.size_t(len(locations)))
}

func (g CGGradient) Retain() {
	C.CGGradientRetain(g)
}

func (g CGGradient) Release() {
	C.CGGradientRelease(g)
}
