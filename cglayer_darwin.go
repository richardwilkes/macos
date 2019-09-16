package macos

// #import <CoreGraphics/CoreGraphics.h>
import "C"

type (
	CGLayer = C.CGLayerRef
)

func CGLayerCreateWithContext(ctx CGContext, width, height float64) CGLayer {
	return C.CGLayerCreateWithContext(ctx, C.CGSizeMake(C.CGFloat(width), C.CGFloat(height)), 0)
}

func (layer CGLayer) Size() (width, height float64) {
	size := C.CGLayerGetSize(layer)
	return float64(size.width), float64(size.height)
}

func (layer CGLayer) Context() CGContext {
	return C.CGLayerGetContext(layer)
}

func (layer CGLayer) Retain() {
	C.CGLayerRetain(layer)
}

func (layer CGLayer) Release() {
	C.CGLayerRelease(layer)
}
