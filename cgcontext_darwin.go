package macos

import "unsafe"

// #import <CoreGraphics/CoreGraphics.h>
import "C"

const (
	CGLineCapButt CGLineCap = iota
	CGLineCapRound
	CGLineCapSquare
)

const (
	CGLineJoinMiter CGLineJoin = iota
	CGLineJoinRound
	CGLineJoinBevel
)

const (
	CGInterpolationQualityDefault CGInterpolationQuality = iota
	CGInterpolationQualityNone
	CGInterpolationQualityLow
	CGInterpolationQualityHigh
	CGInterpolationQualityMedium
)

type (
	CGContext              = C.CGContextRef
	CGLineCap              int
	CGLineJoin             int
	CGInterpolationQuality int
)

func CGBitmapContextCreate(data unsafe.Pointer, width, height, bitsPerComponent, bytesPerRow int, space CGColorSpace, bitmapInfo CGBitmapInfo) CGContext {
	return C.CGBitmapContextCreate(data, C.size_t(width), C.size_t(height), C.size_t(bitsPerComponent), C.size_t(bytesPerRow), space, C.uint32_t(bitmapInfo))
}

func (c CGContext) BitmapContextCreateImage() CGImage {
	return C.CGBitmapContextCreateImage(c)
}

func (c CGContext) SaveGState() {
	C.CGContextSaveGState(c)
}

func (c CGContext) RestoreGState() {
	C.CGContextRestoreGState(c)
}

func (c CGContext) SetAlpha(opacity float64) {
	C.CGContextSetAlpha(c, C.CGFloat(opacity))
}

func (c CGContext) SetRGBFillColor(r, g, b, a float64) {
	C.CGContextSetRGBFillColor(c, C.CGFloat(r), C.CGFloat(g), C.CGFloat(b), C.CGFloat(a))
}

func (c CGContext) SetFillColorSpace(space CGColorSpace) {
	C.CGContextSetFillColorSpace(c, space)
}

func (c CGContext) SetFillPattern(pattern CGPattern, components ...float64) {
	C.CGContextSetFillPattern(c, pattern, floatSliceToCGFloatPtr(components))
}

func (c CGContext) SetRGBStrokeColor(r, g, b, a float64) {
	C.CGContextSetRGBStrokeColor(c, C.CGFloat(r), C.CGFloat(g), C.CGFloat(b), C.CGFloat(a))
}

func (c CGContext) SetStrokePattern(pattern CGPattern, components ...float64) {
	C.CGContextSetStrokePattern(c, pattern, floatSliceToCGFloatPtr(components))
}

func (c CGContext) SetPatternPhase(x, y float64) {
	C.CGContextSetPatternPhase(c, C.CGSizeMake(C.CGFloat(x), C.CGFloat(y)))
}

func (c CGContext) SetLineWidth(width float64) {
	C.CGContextSetLineWidth(c, C.CGFloat(width))
}

func (c CGContext) SetLineCap(lineCap CGLineCap) {
	C.CGContextSetLineCap(c, C.CGLineCap(lineCap))
}

func (c CGContext) SetLineJoin(lineJoin CGLineJoin) {
	C.CGContextSetLineJoin(c, C.CGLineJoin(lineJoin))
}

func (c CGContext) SetMiterLimit(limit float64) {
	C.CGContextSetMiterLimit(c, C.CGFloat(limit))
}

func (c CGContext) SetLineDash(phase float64, segments ...float64) {
	C.CGContextSetLineDash(c, C.CGFloat(phase), floatSliceToCGFloatPtr(segments), C.size_t(len(segments)))
}

func (c CGContext) SetInterpolationQuality(quality CGInterpolationQuality) {
	C.CGContextSetInterpolationQuality(c, C.CGInterpolationQuality(quality))
}

func (c CGContext) TranslateCTM(x, y float64) {
	C.CGContextTranslateCTM(c, C.CGFloat(x), C.CGFloat(y))
}

func (c CGContext) ScaleCTM(x, y float64) {
	C.CGContextScaleCTM(c, C.CGFloat(x), C.CGFloat(y))
}

func (c CGContext) RotateCTM(angleInRadians float64) {
	C.CGContextRotateCTM(c, C.CGFloat(angleInRadians))
}

func (c CGContext) SetTextMatrix(matrix CGAffineTransform) {
	C.CGContextSetTextMatrix(c, matrix)
}

func (c CGContext) GetClipBoundingBox() (x, y, width, height float64) {
	r := C.CGContextGetClipBoundingBox(c)
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func (c CGContext) Clip() {
	C.CGContextClip(c)
}

func (c CGContext) EOClip() {
	C.CGContextEOClip(c)
}
func (c CGContext) BeginPath() {
	C.CGContextBeginPath(c)
}

func (c CGContext) MoveToPoint(x, y float64) {
	C.CGContextMoveToPoint(c, C.CGFloat(x), C.CGFloat(y))
}

func (c CGContext) AddLineToPoint(x, y float64) {
	C.CGContextAddLineToPoint(c, C.CGFloat(x), C.CGFloat(y))
}

func (c CGContext) AddQuadCurveToPoint(cpx, cpy, x, y float64) {
	C.CGContextAddQuadCurveToPoint(c, C.CGFloat(cpx), C.CGFloat(cpy), C.CGFloat(x), C.CGFloat(y))
}

func (c CGContext) AddCurveToPoint(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	C.CGContextAddCurveToPoint(c, C.CGFloat(cp1x), C.CGFloat(cp1y), C.CGFloat(cp2x), C.CGFloat(cp2y), C.CGFloat(x), C.CGFloat(y))
}

func (c CGContext) AddRect(x, y, width, height float64) {
	C.CGContextAddRect(c, C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)))
}

func (c CGContext) AddEllipseInRect(x, y, width, height float64) {
	C.CGContextAddEllipseInRect(c, C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)))
}

func (c CGContext) ClosePath() {
	C.CGContextClosePath(c)
}

func (c CGContext) FillPath() {
	C.CGContextFillPath(c)
}

func (c CGContext) EOFillPath() {
	C.CGContextEOFillPath(c)
}

func (c CGContext) ReplacePathWithStrokedPath() {
	C.CGContextReplacePathWithStrokedPath(c)
}

func (c CGContext) StrokePath() {
	C.CGContextStrokePath(c)
}

func (c CGContext) DrawImage(x, y, width, height float64, img CGImage) {
	C.CGContextDrawImage(c, C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)), img)
}

func (c CGContext) DrawLayer(x, y, width, height float64, layer CGLayer) {
	C.CGContextDrawLayerInRect(c, C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)), layer)
}

func (c CGContext) DrawLinearGradient(gradient CGGradient, sx, sy, ex, ey float64, options CGGradientDrawingOptions) {
	C.CGContextDrawLinearGradient(c, gradient, C.CGPointMake(C.CGFloat(sx), C.CGFloat(sy)), C.CGPointMake(C.CGFloat(ex), C.CGFloat(ey)), C.CGGradientDrawingOptions(options))
}

func (c CGContext) DrawRadialGradient(gradient CGGradient, sx, sy, sr, ex, ey, er float64, options CGGradientDrawingOptions) {
	C.CGContextDrawRadialGradient(c, gradient, C.CGPointMake(C.CGFloat(sx), C.CGFloat(sy)), C.CGFloat(sr), C.CGPointMake(C.CGFloat(ex), C.CGFloat(ey)), C.CGFloat(er), C.CGGradientDrawingOptions(options))
}

func (c CGContext) Release() {
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
