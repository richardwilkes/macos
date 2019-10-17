package ct

import (
	"github.com/richardwilkes/macos/cf"
	"github.com/richardwilkes/macos/cg"
)

// #import <CoreText/CoreText.h>
import "C"

var (
	FontAttributeName                       = cf.String(C.kCTFontAttributeName)
	ForegroundColorFromContextAttributeName = cf.String(C.kCTForegroundColorFromContextAttributeName)
)

type Line = C.CTLineRef

func LineCreateWithAttributedString(attrString cf.AttributedString) Line {
	return C.CTLineCreateWithAttributedString(C.CFAttributedStringRef(attrString))
}

func (l Line) GetTypographicBounds(ascent, descent, leading *float64) float64 {
	var af, df, lf *C.CGFloat
	if ascent != nil {
		var aa C.CGFloat
		af = &aa
	}
	if descent != nil {
		var dd C.CGFloat
		df = &dd
	}
	if leading != nil {
		var ll C.CGFloat
		lf = &ll
	}
	width := float64(C.CTLineGetTypographicBounds(l, af, df, lf))
	if ascent != nil {
		*ascent = float64(*af)
	}
	if descent != nil {
		*descent = float64(*df)
	}
	if leading != nil {
		*leading = float64(*lf)
	}
	return width
}

func (l Line) GetStringIndexForPosition(x, y float64) int {
	return int(C.CTLineGetStringIndexForPosition(l, C.CGPointMake(C.CGFloat(x), C.CGFloat(y))))
}

func (l Line) GetOffsetForStringIndex(index int, secondaryOffset *float64) float64 {
	var offset *C.CGFloat
	if secondaryOffset != nil {
		var o C.CGFloat
		offset = &o
	}
	result := float64(C.CTLineGetOffsetForStringIndex(l, C.CFIndex(index), offset))
	if secondaryOffset != nil {
		*secondaryOffset = float64(*offset)
	}
	return result
}

func (l Line) Draw(gc cg.Context) {
	C.CTLineDraw(l, C.CGContextRef(gc))
}

func (l Line) Release() {
	C.CFRelease(C.CFTypeRef(l))
}
