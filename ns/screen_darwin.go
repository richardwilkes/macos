package ns

import (
	"github.com/richardwilkes/macos/cf"
	"github.com/richardwilkes/macos/cg"
)

/*
#import <Cocoa/Cocoa.h>
#import <CoreGraphics/CoreGraphics.h>

typedef void *NSScreenPtr;

CFArrayRef nsScreenScreens() {
	return (CFArrayRef)[NSScreen screens];
}

NSScreenPtr nsScreenMainScreen() {
	return (NSScreenPtr)[NSScreen mainScreen];
}

NSRect nsScreenFrame(NSScreenPtr s) {
	return [(NSScreen *)s frame];
}

NSRect nsScreenVisibleFrame(NSScreenPtr s) {
	return [(NSScreen *)s visibleFrame];
}

CGFloat nsScreenBackingScaleFactor(NSScreenPtr s) {
	return [(NSScreen *)s backingScaleFactor];
}

CGDirectDisplayID nsScreenDisplayID(NSScreenPtr s) {
	return (CGDirectDisplayID)[[[(NSScreen *)s deviceDescription] objectForKey:@"NSScreenNumber"] unsignedIntValue];
}
*/
import "C"

type Screen struct {
	native C.NSScreenPtr
}

func Screens() []*Screen {
	s := cf.Array(C.nsScreenScreens())
	screens := make([]*Screen, s.GetCount())
	for i := range screens {
		screens[i] = &Screen{native: C.NSScreenPtr(s.GetValueAtIndex(i))}
	}
	return screens
}

func MainScreen() *Screen {
	return &Screen{native: C.nsScreenMainScreen()}
}

func (s *Screen) Frame() (x, y, width, height float64) {
	r := C.nsScreenFrame(s.native)
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func (s *Screen) VisibleFrame() (x, y, width, height float64) {
	r := C.nsScreenVisibleFrame(s.native)
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func (s *Screen) BackingScaleFactor() float64 {
	return float64(C.nsScreenBackingScaleFactor(s.native))
}

func (s *Screen) DisplayID() cg.DirectDisplayID {
	return cg.DirectDisplayID(C.nsScreenDisplayID(s.native))
}
