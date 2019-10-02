package macos

/*
#import <Cocoa/Cocoa.h>

typedef void *NSCursorPtr;
typedef void *NSImagePtr;

NSCursorPtr nsCursorInitWithImageHotSpotRetain(NSImagePtr img, CGFloat hotX, CGFloat hotY) {
	return [[[NSCursor alloc] initWithImage:img hotSpot:NSMakePoint(hotX,hotY)] retain];
}

void nsCursorSet(NSCursorPtr cursor) {
	[((NSCursor *)cursor) set];
}

void nsCursorSetHiddenUntilMouseMoves(bool hide) {
	[NSCursor setHiddenUntilMouseMoves:hide ? YES : NO];
}

void nsCursorRelease(NSCursorPtr cursor) {
	[((NSCursor *)cursor) release];
}
*/
import "C"
import "github.com/richardwilkes/toolbox/softref"

type NSCursor softref.SoftRef

type nsCursorRef struct {
	key        string
	native     C.NSCursorPtr
	releasable bool
}

func NSCursorSetHiddenUntilMouseMoves(hide bool) {
	C.nsCursorSetHiddenUntilMouseMoves(C.bool(hide))
}

func NSCursorInitWithImageHotSpotRetain(img *NSImage, hotX, hotY float64) *NSCursor {
	ref, _ := softref.DefaultPool.NewSoftRef(&nsCursorRef{
		key:        NextRefKey(),
		native:     C.nsCursorInitWithImageHotSpotRetain(img.native(), C.CGFloat(hotX), C.CGFloat(hotY)),
		releasable: true,
	})
	return (*NSCursor)(ref)
}

func nsCursorInit(native C.NSCursorPtr) *NSCursor {
	ref, _ := softref.DefaultPool.NewSoftRef(&nsCursorRef{
		key:    NextRefKey(),
		native: native,
	})
	return (*NSCursor)(ref)
}

func (c *NSCursor) native() C.NSCursorPtr {
	return c.Resource.(*nsCursorRef).native
}

func (c *NSCursor) Set() {
	C.nsCursorSet(c.native())
}

func (r *nsCursorRef) Key() string {
	return r.key
}

func (r *nsCursorRef) Release() {
	if r.releasable {
		C.nsCursorRelease(r.native)
		r.native = nil
	}
}
