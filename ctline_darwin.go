package macos

// #import <CoreText/CoreText.h>
import "C"

var (
	CTFontAttributeName                       = C.kCTFontAttributeName
	CTForegroundColorFromContextAttributeName = C.kCTForegroundColorFromContextAttributeName
)

type CTLine = C.CTLineRef

func CTLineCreateWithAttributedString(attrString CFAttributedString) CTLine {
	return C.CTLineCreateWithAttributedString(attrString)
}

func (l CTLine) GetTypographicBounds(ascent, descent, leading *float64) float64 {
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

func (l CTLine) GetStringIndexForPosition(x, y float64) int {
	return int(C.CTLineGetStringIndexForPosition(l, C.CGPointMake(C.CGFloat(x), C.CGFloat(y))))
}

func (l CTLine) GetOffsetForStringIndex(index int, secondaryOffset *float64) float64 {
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

func (l CTLine) Draw(gc CGContext) {
	C.CTLineDraw(l, gc)
}

func (l CTLine) Release() {
	C.CFRelease(CFType(l))
}
