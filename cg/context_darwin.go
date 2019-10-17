package cg

import "unsafe"

// #import <CoreGraphics/CoreGraphics.h>
import "C"

const (
	LineCapButt LineCap = iota
	LineCapRound
	LineCapSquare
)

const (
	LineJoinMiter LineJoin = iota
	LineJoinRound
	LineJoinBevel
)

const (
	InterpolationQualityDefault InterpolationQuality = iota
	InterpolationQualityNone
	InterpolationQualityLow
	InterpolationQualityHigh
	InterpolationQualityMedium
)

type (
	Context              = C.CGContextRef
	LineCap              int
	LineJoin             int
	InterpolationQuality int
)

func BitmapContextCreate(data unsafe.Pointer, width, height, bitsPerComponent, bytesPerRow int, space ColorSpace, bitmapInfo BitmapInfo) Context {
	return C.CGBitmapContextCreate(data, C.size_t(width), C.size_t(height), C.size_t(bitsPerComponent), C.size_t(bytesPerRow), space, C.uint32_t(bitmapInfo))
}

func (c Context) BitmapContextCreateImage() Image {
	return C.CGBitmapContextCreateImage(c)
}

func (c Context) SaveGState() {
	C.CGContextSaveGState(c)
}

func (c Context) RestoreGState() {
	C.CGContextRestoreGState(c)
}

func (c Context) SetAlpha(opacity float64) {
	C.CGContextSetAlpha(c, C.CGFloat(opacity))
}

func (c Context) SetRGBFillColor(r, g, b, a float64) {
	C.CGContextSetRGBFillColor(c, C.CGFloat(r), C.CGFloat(g), C.CGFloat(b), C.CGFloat(a))
}

func (c Context) SetFillColorSpace(space ColorSpace) {
	C.CGContextSetFillColorSpace(c, space)
}

func (c Context) SetFillPattern(pattern Pattern, components ...float64) {
	C.CGContextSetFillPattern(c, pattern, floatSliceToCGFloatPtr(components))
}

func (c Context) SetRGBStrokeColor(r, g, b, a float64) {
	C.CGContextSetRGBStrokeColor(c, C.CGFloat(r), C.CGFloat(g), C.CGFloat(b), C.CGFloat(a))
}

func (c Context) SetStrokePattern(pattern Pattern, components ...float64) {
	C.CGContextSetStrokePattern(c, pattern, floatSliceToCGFloatPtr(components))
}

func (c Context) SetPatternPhase(x, y float64) {
	C.CGContextSetPatternPhase(c, C.CGSizeMake(C.CGFloat(x), C.CGFloat(y)))
}

func (c Context) SetLineWidth(width float64) {
	C.CGContextSetLineWidth(c, C.CGFloat(width))
}

func (c Context) SetLineCap(lineCap LineCap) {
	C.CGContextSetLineCap(c, C.CGLineCap(lineCap))
}

func (c Context) SetLineJoin(lineJoin LineJoin) {
	C.CGContextSetLineJoin(c, C.CGLineJoin(lineJoin))
}

func (c Context) SetMiterLimit(limit float64) {
	C.CGContextSetMiterLimit(c, C.CGFloat(limit))
}

func (c Context) SetLineDash(phase float64, segments ...float64) {
	C.CGContextSetLineDash(c, C.CGFloat(phase), floatSliceToCGFloatPtr(segments), C.size_t(len(segments)))
}

func (c Context) SetInterpolationQuality(quality InterpolationQuality) {
	C.CGContextSetInterpolationQuality(c, C.CGInterpolationQuality(quality))
}

func (c Context) TranslateCTM(x, y float64) {
	C.CGContextTranslateCTM(c, C.CGFloat(x), C.CGFloat(y))
}

func (c Context) ScaleCTM(x, y float64) {
	C.CGContextScaleCTM(c, C.CGFloat(x), C.CGFloat(y))
}

func (c Context) RotateCTM(angleInRadians float64) {
	C.CGContextRotateCTM(c, C.CGFloat(angleInRadians))
}

func (c Context) SetTextMatrix(matrix AffineTransform) {
	C.CGContextSetTextMatrix(c, matrix)
}

func (c Context) GetClipBoundingBox() (x, y, width, height float64) {
	r := C.CGContextGetClipBoundingBox(c)
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func (c Context) Clip() {
	C.CGContextClip(c)
}

func (c Context) EOClip() {
	C.CGContextEOClip(c)
}
func (c Context) BeginPath() {
	C.CGContextBeginPath(c)
}

func (c Context) MoveToPoint(x, y float64) {
	C.CGContextMoveToPoint(c, C.CGFloat(x), C.CGFloat(y))
}

func (c Context) AddLineToPoint(x, y float64) {
	C.CGContextAddLineToPoint(c, C.CGFloat(x), C.CGFloat(y))
}

func (c Context) AddQuadCurveToPoint(cpx, cpy, x, y float64) {
	C.CGContextAddQuadCurveToPoint(c, C.CGFloat(cpx), C.CGFloat(cpy), C.CGFloat(x), C.CGFloat(y))
}

func (c Context) AddCurveToPoint(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	C.CGContextAddCurveToPoint(c, C.CGFloat(cp1x), C.CGFloat(cp1y), C.CGFloat(cp2x), C.CGFloat(cp2y), C.CGFloat(x), C.CGFloat(y))
}

func (c Context) AddRect(x, y, width, height float64) {
	C.CGContextAddRect(c, C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)))
}

func (c Context) AddEllipseInRect(x, y, width, height float64) {
	C.CGContextAddEllipseInRect(c, C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)))
}

func (c Context) ClosePath() {
	C.CGContextClosePath(c)
}

func (c Context) FillPath() {
	C.CGContextFillPath(c)
}

func (c Context) EOFillPath() {
	C.CGContextEOFillPath(c)
}

func (c Context) ReplacePathWithStrokedPath() {
	C.CGContextReplacePathWithStrokedPath(c)
}

func (c Context) StrokePath() {
	C.CGContextStrokePath(c)
}

func (c Context) DrawImage(x, y, width, height float64, img Image) {
	C.CGContextDrawImage(c, C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)), img)
}

func (c Context) DrawLayer(x, y, width, height float64, layer Layer) {
	C.CGContextDrawLayerInRect(c, C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)), layer)
}

func (c Context) DrawLinearGradient(gradient Gradient, sx, sy, ex, ey float64, options GradientDrawingOptions) {
	C.CGContextDrawLinearGradient(c, gradient, C.CGPointMake(C.CGFloat(sx), C.CGFloat(sy)), C.CGPointMake(C.CGFloat(ex), C.CGFloat(ey)), C.CGGradientDrawingOptions(options))
}

func (c Context) DrawRadialGradient(gradient Gradient, sx, sy, sr, ex, ey, er float64, options GradientDrawingOptions) {
	C.CGContextDrawRadialGradient(c, gradient, C.CGPointMake(C.CGFloat(sx), C.CGFloat(sy)), C.CGFloat(sr), C.CGPointMake(C.CGFloat(ex), C.CGFloat(ey)), C.CGFloat(er), C.CGGradientDrawingOptions(options))
}

func (c Context) Release() {
	C.CGContextRelease(c)
}

func floatSliceToCGFloatPtr(in []float64) *C.CGFloat {
	if len(in) == 0 {
		return nil
	}
	out := make([]C.CGFloat, len(in))
	for i := range out {
		out[i] = C.CGFloat(in[i])
	}
	return &out[0]
}
