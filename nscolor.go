package macos

/*
#import <Cocoa/Cocoa.h>

typedef void *NSColorPtr;
typedef void *NSColorSpacePtr;

CFArrayRef nsColorAlternatingContentBackgroundColors() {
	return (CFArrayRef)[NSColor alternatingContentBackgroundColors];
}

NSColorPtr nsColorUsingColorSpace(NSColorPtr color, NSColorSpacePtr space) {
	return [(NSColor *)color colorUsingColorSpace:(NSColorSpace *)space];
}

void nsColorGetRedGreenBlueAlpha(NSColorPtr color, CGFloat *r, CGFloat *g, CGFloat *b, CGFloat *a) {
	[(NSColor *)color getRed:r green:g blue:b alpha:a];
}

void nsColorDispose(NSColorPtr color) {
	[((NSColor *)color) release];
}
*/
import "C"

type NSColor struct {
	native C.NSColorPtr
}

func AlternatingContentBackgroundColors() []*NSColor {
	data := C.nsColorAlternatingContentBackgroundColors()
	colors := make([]*NSColor, data.GetCount())
	for i := range colors {
		colors[i] = &NSColor{native: C.NSColorPtr(data.GetValueAtIndex(i))}
	}
	return colors
}

func (c *NSColor) ColorUsingColorSpace(space *NSColorSpace) *NSColor {
	return &NSColor{native: C.nsColorUsingColorSpace(c.native, space.native)}
}

func (c *NSColor) GetRedGreenBlueAlpha() (r, g, b, a float64) {
	var rr, gg, bb, aa C.CGFloat
	C.nsColorGetRedGreenBlueAlpha(c.native, &rr, &gg, &bb, &aa)
	return float64(rr), float64(gg), float64(bb), float64(aa)
}

func (c *NSColor) Release() {
	if c.native != nil {
		C.nsColorDispose(c.native)
		c.native = nil
	}
}
