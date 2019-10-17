package ns

import (
	"github.com/richardwilkes/macos"
	"github.com/richardwilkes/toolbox/softref"
)

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

type Cursor softref.SoftRef

type cursorRef struct {
	key        string
	native     C.NSCursorPtr
	releasable bool
}

func CursorSetHiddenUntilMouseMoves(hide bool) {
	C.nsCursorSetHiddenUntilMouseMoves(C.bool(hide))
}

func CursorInitWithImageHotSpotRetain(img *Image, hotX, hotY float64) *Cursor {
	ref, _ := softref.DefaultPool.NewSoftRef(&cursorRef{
		key:        macos.NextRefKey(),
		native:     C.nsCursorInitWithImageHotSpotRetain(img.native(), C.CGFloat(hotX), C.CGFloat(hotY)),
		releasable: true,
	})
	return (*Cursor)(ref)
}

func cursorInit(native C.NSCursorPtr) *Cursor {
	ref, _ := softref.DefaultPool.NewSoftRef(&cursorRef{
		key:    macos.NextRefKey(),
		native: native,
	})
	return (*Cursor)(ref)
}

func (c *Cursor) native() C.NSCursorPtr {
	return c.Resource.(*cursorRef).native
}

func (c *Cursor) Set() {
	C.nsCursorSet(c.native())
}

func (r *cursorRef) Key() string {
	return r.key
}

func (r *cursorRef) Release() {
	if r.releasable {
		C.nsCursorRelease(r.native)
		r.native = nil
	}
}
