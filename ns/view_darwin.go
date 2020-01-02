// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ns

import "github.com/richardwilkes/macos/cg"

/*
#import <Cocoa/Cocoa.h>

typedef void *NSWindowPtr;
typedef void *NSViewPtr;
typedef void *NSTrackingAreaPtr;

void viewDrawCallback(NSViewPtr view, CGContextRef gc, CGFloat x, CGFloat y, CGFloat width, CGFloat height, bool inLiveResize);
void viewMouseDownCallback(NSViewPtr view, CGFloat x, CGFloat y, int button, int clickCount, int mod);
void viewMouseDragCallback(NSViewPtr view, CGFloat x, CGFloat y, int button, int mod);
void viewMouseUpCallback(NSViewPtr view, CGFloat x, CGFloat y, int button, int mod);
void viewMouseEnterCallback(NSViewPtr view, CGFloat x, CGFloat y, int mod);
void viewMouseMoveCallback(NSViewPtr view, CGFloat x, CGFloat y, int mod);
void viewMouseExitCallback(NSViewPtr view);
void viewMouseWheelCallback(NSViewPtr view, CGFloat x, CGFloat y, CGFloat dx, CGFloat dy, int mod);
void viewCursorUpdateCallback(NSViewPtr view, CGFloat x, CGFloat y, int mod);
void viewKeyDownCallback(NSViewPtr view, int keyCode, CFStringRef ch, int mod, bool repeat);
void viewKeyUpCallback(NSViewPtr view, int keyCode, int mod);

@interface DrawingView : NSView
@end

@implementation DrawingView
-(void)drawRect:(NSRect)dirtyRect {
	viewDrawCallback((NSViewPtr)self, [[NSGraphicsContext currentContext] CGContext], dirtyRect.origin.x, dirtyRect.origin.y, dirtyRect.size.width, dirtyRect.size.height, [self inLiveResize] != 0);
}

-(void)mouseDown:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	viewMouseDownCallback((NSViewPtr)self, where.x, where.y, theEvent.buttonNumber, theEvent.clickCount, theEvent.modifierFlags);
}

-(void)rightMouseDown:(NSEvent *)theEvent {
	[self mouseDown:theEvent];
}

-(void)otherMouseDown:(NSEvent *)theEvent {
	[self mouseDown:theEvent];
}

-(void)mouseDragged:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	viewMouseDragCallback((NSViewPtr)self, where.x, where.y, theEvent.buttonNumber, theEvent.modifierFlags);
}

-(void)rightMouseDragged:(NSEvent *)theEvent {
	[self mouseDragged:theEvent];
}

-(void)otherMouseDragged:(NSEvent *)theEvent {
	[self mouseDragged:theEvent];
}

-(void)mouseUp:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	viewMouseUpCallback((NSViewPtr)self, where.x, where.y, theEvent.buttonNumber, theEvent.modifierFlags);
}

-(void)rightMouseUp:(NSEvent *)theEvent {
	[self mouseUp:theEvent];
}

-(void)otherMouseUp:(NSEvent *)theEvent {
	[self mouseUp:theEvent];
}

-(void)mouseEntered:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	viewMouseEnterCallback((NSViewPtr)self, where.x, where.y, theEvent.modifierFlags);
}

-(void)mouseMoved:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	viewMouseMoveCallback((NSViewPtr)self, where.x, where.y, theEvent.modifierFlags);
}

-(void)mouseExited:(NSEvent *)theEvent {
	viewMouseExitCallback((NSViewPtr)self);
}

-(void)scrollWheel:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	viewMouseWheelCallback((NSViewPtr)self, where.x, where.y, theEvent.scrollingDeltaX, theEvent.scrollingDeltaY, theEvent.modifierFlags);
}

-(void)cursorUpdate:(NSEvent *)theEvent {
	NSPoint where = [self convertPoint:theEvent.locationInWindow fromView:nil];
	viewCursorUpdateCallback((NSViewPtr)self, where.x, where.y, theEvent.modifierFlags);
}

-(void)flagsChanged:(NSEvent *)theEvent {
	BOOL down;
	switch (theEvent.keyCode) {
		case 57:	// Caps Lock
			down = (theEvent.modifierFlags & NSEventModifierFlagCapsLock) != 0;
			break;
		case 56:	// Left Shift
		case 60:	// Right Shift
			down = (theEvent.modifierFlags & NSEventModifierFlagShift) != 0;
			break;
		case 59:	// Left Control
		case 62:	// Right Control
			down = (theEvent.modifierFlags & NSEventModifierFlagControl) != 0;
			break;
		case 58:	// Left Option
		case 61:	// Right Option
			down = (theEvent.modifierFlags & NSEventModifierFlagOption) != 0;
			break;
		case 54:	// Right Cmd
		case 55:	// Left Cmd
			down = (theEvent.modifierFlags & NSEventModifierFlagCommand) != 0;
			break;
		default:
			down = true;
			break;
	}
	if (down) {
		viewKeyDownCallback((NSViewPtr)self, theEvent.keyCode, nil, theEvent.modifierFlags, false);
	} else {
		viewKeyUpCallback((NSViewPtr)self, theEvent.keyCode, theEvent.modifierFlags);
	}
}

-(void)keyDown:(NSEvent *)theEvent {
	viewKeyDownCallback((NSViewPtr)self, theEvent.keyCode, (CFStringRef)theEvent.characters, theEvent.modifierFlags, theEvent.ARepeat);
}

-(void)keyUp:(NSEvent *)theEvent {
	viewKeyUpCallback((NSViewPtr)self, theEvent.keyCode, theEvent.modifierFlags);
}

-(void)viewDidEndLiveResize {
	[self setNeedsDisplayInRect:[self bounds]];
}

-(BOOL)acceptsFirstResponder {
	return YES;
}

-(BOOL)isFlipped {
	return YES;
}
@end

NSViewPtr nsNewView() {
	return (NSViewPtr)[DrawingView new];
}

NSWindowPtr nsViewWindow(NSViewPtr v) {
	return (NSWindowPtr)[(NSView *)v window];
}

void nsViewAddTrackingArea(NSViewPtr v, NSTrackingAreaPtr ta) {
	[(NSView *)v addTrackingArea:ta];
}

NSRect nsViewFrame(NSViewPtr v) {
	return [(NSView *)v frame];
}

void nsViewSetFrame(NSViewPtr v, CGFloat x, CGFloat y, CGFloat width, CGFloat height) {
	[(NSView *)v setFrame:NSMakeRect(x, y, width, height)];
}

void nsViewSetNeedsLayout(NSViewPtr v, bool needsLayout) {
	[(NSView *)v setNeedsLayout:needsLayout ? YES : NO];
}

void nsViewSetNeedsDisplay(NSViewPtr v, bool needsDisplay) {
	[(NSView *)v setNeedsDisplay:needsDisplay ? YES : NO];
}

void nsViewSetNeedsDisplayInRect(NSViewPtr v, CGFloat x, CGFloat y, CGFloat width, CGFloat height) {
	[(NSView *)v setNeedsDisplayInRect:NSMakeRect(x, y, width, height)];
}

bool nsViewInLiveResize(NSViewPtr v) {
	return [(NSView *)v inLiveResize] != 0;
}

void nsViewAddSubview(NSViewPtr v, NSViewPtr sub) {
	return [(NSView *)v addSubview:(NSView *)sub];
}

void nsViewRemoveFromSuperview(NSViewPtr v) {
	[(NSView *)v removeFromSuperview];
}

void nsViewRelease(NSViewPtr v) {
	[(NSView *)v release];
}
*/
import "C"

const (
	EventModifierFlagCapsLock = 1 << (iota + 16)
	EventModifierFlagShift
	EventModifierFlagControl
	EventModifierFlagOption
	EventModifierFlagCommand
	EventModifierFlagNumericPad
	EventModifierFlagHelp
	EventModifierFlagFunction
)

var viewDelegateMap = make(map[C.NSViewPtr]ViewDelegate)

type ViewNative = C.NSViewPtr

type View struct {
	native ViewNative
}

type ViewDelegate interface {
	ViewDraw(view *View, gc cg.Context, x, y, width, height float64, inLiveResize bool)
	ViewMouseDownEvent(view *View, x, y float64, button, clickCount, mod int)
	ViewMouseDragEvent(view *View, x, y float64, button, mod int)
	ViewMouseUpEvent(view *View, x, y float64, button, mod int)
	ViewMouseEnterEvent(view *View, x, y float64, mod int)
	ViewMouseMoveEvent(view *View, x, y float64, mod int)
	ViewMouseExitEvent(view *View)
	ViewMouseWheelEvent(view *View, x, y, dx, dy float64, mod int)
	ViewCursorUpdateEvent(view *View, x, y float64, mod int)
	ViewKeyDownEvent(view *View, keyCode int, ch rune, mod int, repeat bool)
	ViewKeyUpEvent(view *View, keyCode int, mod int)
}

func NewView(delegate ViewDelegate) *View {
	view := &View{native: C.nsNewView()}
	viewDelegateMap[view.native] = delegate
	return view
}

func NewViewFromNative(native ViewNative) *View {
	return &View{native: native}
}

func (v *View) Native() ViewNative {
	return v.native
}

func (v *View) Window() *Window {
	if wnd := C.nsViewWindow(v.native); wnd != nil {
		return &Window{native: wnd}
	}
	return nil
}

func (v *View) AddTrackingArea(trackingArea *TrackingArea) {
	C.nsViewAddTrackingArea(v.native, trackingArea.native)
}

func (v *View) Frame() (x, y, width, height float64) {
	r := C.nsViewFrame(v.native)
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func (v *View) SetFrame(x, y, width, height float64) {
	C.nsViewSetFrame(v.native, C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height))
}

func (v *View) SetNeedsLayout(needsLayout bool) {
	C.nsViewSetNeedsLayout(v.native, C.bool(needsLayout))
}

func (v *View) SetNeedsDisplay(needsDisplay bool) {
	C.nsViewSetNeedsDisplay(v.native, C.bool(needsDisplay))
}

func (v *View) SetNeedsDisplayInRect(x, y, width, height float64) {
	C.nsViewSetNeedsDisplayInRect(v.native, C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height))
}

func (v *View) InLiveResize() bool {
	return bool(C.nsViewInLiveResize(v.native))
}

func (v *View) AddSubview(view *View) {
	C.nsViewAddSubview(v.native, view.native)
}

func (v *View) RemoveFromSuperview() {
	C.nsViewRemoveFromSuperview(v.native)
}

func (v *View) Release() {
	C.nsViewRelease(v.native)
	v.ReleaseDelegate()
}

func (v *View) ReleaseDelegate() {
	delete(viewDelegateMap, v.native)
}
