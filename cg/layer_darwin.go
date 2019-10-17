package cg

// #import <CoreGraphics/CoreGraphics.h>
import "C"

type (
	Layer = C.CGLayerRef
)

func LayerCreateWithContext(ctx Context, width, height float64) Layer {
	return C.CGLayerCreateWithContext(ctx, C.CGSizeMake(C.CGFloat(width), C.CGFloat(height)), 0)
}

func (layer Layer) Size() (width, height float64) {
	size := C.CGLayerGetSize(layer)
	return float64(size.width), float64(size.height)
}

func (layer Layer) Context() Context {
	return C.CGLayerGetContext(layer)
}

func (layer Layer) Retain() {
	C.CGLayerRetain(layer)
}

func (layer Layer) Release() {
	C.CGLayerRelease(layer)
}
