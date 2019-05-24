package macos

/*
#import <Cocoa/Cocoa.h>

typedef void *NSWindowPtr;

@interface KeyableWindow : NSWindow
{
@protected
	bool _isKeyable;
}
@property(assign, readwrite) bool isKeyable;
@end

@implementation KeyableWindow
@synthesize isKeyable = _isKeyable;
-(BOOL)canBecomeKeyWindow {
	return _isKeyable;
}
@end

NSRect nsWindowContentRectForFrameRectStyleMask(CGFloat x, CGFloat y, CGFloat width, CGFloat height, NSWindowStyleMask styleMask) {
	return [NSWindow contentRectForFrameRect:NSMakeRect(x, y, width, height) styleMask:styleMask];
}

NSRect nsWindowFrameRectForContentRectStyleMask(CGFloat x, CGFloat y, CGFloat width, CGFloat height, NSWindowStyleMask styleMask) {
	return [NSWindow frameRectForContentRect:NSMakeRect(x, y, width, height) styleMask:styleMask];
}

NSWindowPtr nsWindowInitWithContentRectStyleMask(CGFloat x, CGFloat y, CGFloat width, CGFloat height, NSWindowStyleMask styleMask) {
	return (NSWindowPtr)[[KeyableWindow alloc] initWithContentRect:NSMakeRect(x, y, width, height) styleMask:styleMask backing:NSBackingStoreBuffered defer:YES];
}

void nsWindowSetIsKeyable(NSWindowPtr w, bool keyable) {
	[(KeyableWindow *)w setIsKeyable:keyable];
}

void nsWindowDisableCursorRects(NSWindowPtr w) {
	[(NSWindow *)w disableCursorRects];
}

void nsWindowSetTitle(NSWindowPtr w, CFStringRef title) {
	[(NSWindow *)w setTitle:(NSString *)title];
}

NSRect nsWindowFrame(NSWindowPtr w) {
	return [(NSWindow *)w frame];
}

void nsWindowSetFrame(NSWindowPtr w, CGFloat x, CGFloat y, CGFloat width, CGFloat height, bool display) {
	[(NSWindow *)w setFrame:NSMakeRect(x, y, width, height) display:display ? YES : NO];
}

void nsWindowMakeKeyAndOrderFront(NSWindowPtr w) {
	[(NSWindow *)w makeKeyAndOrderFront:nil];
}

void nsWindowPerformMiniaturize(NSWindowPtr w) {
	[(NSWindow *)w performMiniaturize:nil];
}

void nsWindowPerformZoom(NSWindowPtr w) {
	[(NSWindow *)w performZoom:nil];
}

void nsWindowClose(NSWindowPtr w) {
	[(NSWindow *)w close];
}
*/
import "C"

const (
	NSWindowStyleMaskBorderless             NSWindowStyleMask = 0
	NSWindowStyleMaskTitled                 NSWindowStyleMask = 1 << 0
	NSWindowStyleMaskClosable               NSWindowStyleMask = 1 << 1
	NSWindowStyleMaskMiniaturizable         NSWindowStyleMask = 1 << 2
	NSWindowStyleMaskResizable              NSWindowStyleMask = 1 << 3
	NSWindowStyleMaskUtilityWindow          NSWindowStyleMask = 1 << 4
	NSWindowStyleMaskDocModalWindow         NSWindowStyleMask = 1 << 6
	NSWindowStyleMaskNonactivatingPanel     NSWindowStyleMask = 1 << 7
	NSWindowStyleMaskUnifiedTitleAndToolbar NSWindowStyleMask = 1 << 12
	NSWindowStyleMaskHUDWindow              NSWindowStyleMask = 1 << 13
	NSWindowStyleMaskFullScreen             NSWindowStyleMask = 1 << 14
	NSWindowStyleMaskFullSizeContentView    NSWindowStyleMask = 1 << 15
)

type (
	NSWindowStyleMask int
	NSWindowNative    = C.NSWindowPtr
)

type NSWindow struct {
	native NSWindowNative
}

func NSWindowContentRectForFrameRectStyleMask(x, y, width, height float64, styleMask NSWindowStyleMask) (cx, cy, cwidth, cheight float64) {
	r := C.nsWindowContentRectForFrameRectStyleMask(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), C.NSWindowStyleMask(styleMask))
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func NSWindowFrameRectForContentRectStyleMask(x, y, width, height float64, styleMask NSWindowStyleMask) (cx, cy, cwidth, cheight float64) {
	r := C.nsWindowFrameRectForContentRectStyleMask(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), C.NSWindowStyleMask(styleMask))
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func NSWindowInitWithContentRectStyleMask(x, y, width, height float64, styleMask NSWindowStyleMask) *NSWindow {
	return &NSWindow{native: C.nsWindowInitWithContentRectStyleMask(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), C.NSWindowStyleMask(styleMask))}
}

func (w *NSWindow) Native() NSWindowNative {
	return w.native
}

func (w *NSWindow) SetIsKeyable(keyable bool) {
	C.nsWindowSetIsKeyable(w.native, C.bool(keyable))
}

func (w *NSWindow) DisableCursorRects() {
	C.nsWindowDisableCursorRects(w.native)
}

func (w *NSWindow) SetTitle(title string) {
	str := CFStringCreateWithString(title)
	defer str.Release()
	C.nsWindowSetTitle(w.native, str)
}

func (w *NSWindow) Frame() (x, y, width, height float64) {
	r := C.nsWindowFrame(w.native)
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func (w *NSWindow) SetFrame(x, y, width, height float64, display bool) {
	C.nsWindowSetFrame(w.native, C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), C.bool(display))
}

func (w *NSWindow) MakeKeyAndOrderFront() {
	C.nsWindowMakeKeyAndOrderFront(w.native)
}

func (w *NSWindow) PerformMiniaturize() {
	C.nsWindowPerformMiniaturize(w.native)
}

func (w *NSWindow) PerformZoom() {
	C.nsWindowPerformZoom(w.native)
}

func (w *NSWindow) Close() {
	C.nsWindowClose(w.native)
}
