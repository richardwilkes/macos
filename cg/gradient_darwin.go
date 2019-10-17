package cg

// #import <CoreGraphics/CoreGraphics.h>
import "C"

const (
	GradientDrawsBeforeStartLocation GradientDrawingOptions = 1 << iota
	GradientDrawsAfterEndLocation
)

type (
	Gradient               = C.CGGradientRef
	GradientDrawingOptions int
)

func GradientCreateWithColorComponents(space ColorSpace, components, locations []float64) Gradient {
	return C.CGGradientCreateWithColorComponents(space, floatSliceToCGFloatPtr(components), floatSliceToCGFloatPtr(locations), C.size_t(len(locations)))
}

func (g Gradient) Retain() {
	C.CGGradientRetain(g)
}

func (g Gradient) Release() {
	C.CGGradientRelease(g)
}
