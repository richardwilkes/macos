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

void nsCursorRetain(NSCursorPtr cursor) {
	[((NSCursor *)cursor) retain];
}

void nsCursorRelease(NSCursorPtr cursor) {
	[((NSCursor *)cursor) release];
}
*/
import "C"

type NSCursor struct {
	native     C.NSCursorPtr
	refcnt     int
	releasable bool
}

func NSCursorSetHiddenUntilMouseMoves(hide bool) {
	C.nsCursorSetHiddenUntilMouseMoves(C.bool(hide))
}

func NSCursorInitWithImageHotSpotRetain(img *NSImage, hotX, hotY float64) *NSCursor {
	return &NSCursor{
		native:     C.nsCursorInitWithImageHotSpotRetain(img.native, C.CGFloat(hotX), C.CGFloat(hotY)),
		refcnt:     1,
		releasable: true,
	}
}

func (c *NSCursor) Set() {
	C.nsCursorSet(c.native)
}

func (c *NSCursor) Retain() {
	if c.releasable && c.native != nil {
		c.refcnt++
		C.nsCursorRetain(c.native)
	}
}

func (c *NSCursor) Release() {
	if c.releasable && c.native != nil {
		c.refcnt--
		if c.refcnt <= 0 {
			C.nsCursorRelease(c.native)
			c.native = nil
		}
	}
}
