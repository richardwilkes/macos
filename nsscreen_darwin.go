package macos

/*
#import <Cocoa/Cocoa.h>
#import <CoreGraphics/CoreGraphics.h>

typedef void *NSScreenPtr;

CFArrayRef nsScreenScreens() {
	return (CFArrayRef)[NSScreen screens];
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

type CGDirectDisplayID = C.CGDirectDisplayID

type NSScreen struct {
	native C.NSScreenPtr
}

func Screens() []*NSScreen {
	s := C.nsScreenScreens()
	screens := make([]*NSScreen, s.GetCount())
	for i := range screens {
		screens[i] = &NSScreen{native: C.NSScreenPtr(s.GetValueAtIndex(i))}
	}
	return screens
}

func (s *NSScreen) Frame() (x, y, width, height float64) {
	r := C.nsScreenFrame(s.native)
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func (s *NSScreen) VisibleFrame() (x, y, width, height float64) {
	r := C.nsScreenVisibleFrame(s.native)
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func (s *NSScreen) BackingScaleFactor() float64 {
	return float64(C.nsScreenBackingScaleFactor(s.native))
}

func (s *NSScreen) DisplayID() CGDirectDisplayID {
	return C.nsScreenDisplayID(s.native)
}

func CGDisplayIsMain(id CGDirectDisplayID) bool {
	return C.CGDisplayIsMain(id) != 0
}

func CGDisplayBounds(id CGDirectDisplayID) (x, y, width, height float64) {
	r := C.CGDisplayBounds(id)
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}
