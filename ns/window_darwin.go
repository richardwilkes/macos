package ns

import (
	"math"

	"github.com/richardwilkes/macos/cf"
)

/*
#import <Cocoa/Cocoa.h>

typedef void *NSWindowPtr;
typedef void *NSViewPtr;
typedef void *NSDraggingInfoPtr;

void windowDidResizeCallback(NSWindowPtr wnd);
void windowDidBecomeKeyCallback(NSWindowPtr wnd);
void windowDidResignKeyCallback(NSWindowPtr wnd);
bool windowShouldCloseCallback(NSWindowPtr wnd);
void windowWillCloseCallback(NSWindowPtr wnd);
NSDragOperation windowDraggingEnteredCallback(NSDraggingInfoPtr di);
NSDragOperation windowDraggingUpdatedCallback(NSDraggingInfoPtr di);
void windowDraggingExitedCallback(NSDraggingInfoPtr di);
void windowDraggingEndedCallback(NSDraggingInfoPtr di);
bool windowPrepareForDragCallback(NSDraggingInfoPtr di);
bool windowPerformDragCallback(NSDraggingInfoPtr di);
void windowConcludeDragCallback(NSDraggingInfoPtr di);

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

@interface WindowDelegate : NSObject<NSWindowDelegate>
@end

@implementation WindowDelegate
-(void)windowDidResize:(NSNotification *)notification {
	windowDidResizeCallback((NSWindowPtr)[notification object]);
}

-(void)windowDidBecomeKey:(NSNotification *)notification {
	windowDidBecomeKeyCallback((NSWindowPtr)[notification object]);
}

-(void)windowDidResignKey:(NSNotification *)notification {
	windowDidResignKeyCallback((NSWindowPtr)[notification object]);
}

-(BOOL)windowShouldClose:(id)sender {
	return windowShouldCloseCallback((NSWindowPtr)sender) ? YES : NO;
}

-(void)windowWillClose:(NSNotification *)notification {
	windowWillCloseCallback((NSWindowPtr)[notification object]);
}

-(NSDragOperation)draggingEntered:(id<NSDraggingInfo>)sender {
	return windowDraggingEnteredCallback((NSDraggingInfoPtr)sender);
}

-(NSDragOperation)draggingUpdated:(id<NSDraggingInfo>)sender {
	return windowDraggingUpdatedCallback((NSDraggingInfoPtr)sender);
}

-(void)draggingExited:(id<NSDraggingInfo>)sender {
	windowDraggingExitedCallback((NSDraggingInfoPtr)sender);
}

-(void)draggingEnded:(id<NSDraggingInfo>)sender {
	windowDraggingEndedCallback((NSDraggingInfoPtr)sender);
}

-(BOOL)prepareForDragOperation:(id<NSDraggingInfo>)sender {
	return windowPrepareForDragCallback((NSDraggingInfoPtr)sender) ? YES : NO;
}

-(BOOL)performDragOperation:(id<NSDraggingInfo>)sender {
	return windowPerformDragCallback((NSDraggingInfoPtr)sender) ? YES : NO;
}

-(void)concludeDragOperation:(id<NSDraggingInfo>)sender {
	windowConcludeDragCallback((NSDraggingInfoPtr)sender);
}
@end

NSRect nsWindowContentRectForFrameRectStyleMask(CGFloat x, CGFloat y, CGFloat width, CGFloat height, NSWindowStyleMask styleMask) {
	return [NSWindow contentRectForFrameRect:NSMakeRect(x, y, width, height) styleMask:styleMask];
}

NSRect nsWindowFrameRectForContentRectStyleMask(CGFloat x, CGFloat y, CGFloat width, CGFloat height, NSWindowStyleMask styleMask) {
	return [NSWindow frameRectForContentRect:NSMakeRect(x, y, width, height) styleMask:styleMask];
}

NSWindowPtr nsWindowInitWithContentRectStyleMask(CGFloat x, CGFloat y, CGFloat width, CGFloat height, NSWindowStyleMask styleMask) {
	KeyableWindow *wnd = [[KeyableWindow alloc] initWithContentRect:NSMakeRect(x, y, width, height) styleMask:styleMask backing:NSBackingStoreBuffered defer:YES];
	[wnd setDelegate: [WindowDelegate new]];
	return (NSWindowPtr)wnd;
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

NSViewPtr nsWindowContentView(NSWindowPtr w) {
	return (NSViewPtr)[(NSWindow *)w contentView];
}

void nsWindowSetContentView(NSWindowPtr w, NSViewPtr v) {
	[(NSWindow *)w setContentView:(NSView *)v];
}

void nsWindowClose(NSWindowPtr w) {
	[(NSWindow *)w close];
}

NSPoint nsWindowMouseLocationOutsideOfEventStream(NSWindowPtr w) {
	return [(NSWindow *)w mouseLocationOutsideOfEventStream];
}

void nsWindowRegisterForDraggedTypes(NSWindowPtr w, CFArrayRef types) {
	[(NSWindow *)w registerForDraggedTypes:(NSArray<NSPasteboardType> *)types];
}
*/
import "C"

const (
	WindowStyleMaskBorderless             WindowStyleMask = 0
	WindowStyleMaskTitled                 WindowStyleMask = 1 << 0
	WindowStyleMaskClosable               WindowStyleMask = 1 << 1
	WindowStyleMaskMiniaturizable         WindowStyleMask = 1 << 2
	WindowStyleMaskResizable              WindowStyleMask = 1 << 3
	WindowStyleMaskUtilityWindow          WindowStyleMask = 1 << 4
	WindowStyleMaskDocModalWindow         WindowStyleMask = 1 << 6
	WindowStyleMaskNonactivatingPanel     WindowStyleMask = 1 << 7
	WindowStyleMaskUnifiedTitleAndToolbar WindowStyleMask = 1 << 12
	WindowStyleMaskHUDWindow              WindowStyleMask = 1 << 13
	WindowStyleMaskFullScreen             WindowStyleMask = 1 << 14
	WindowStyleMaskFullSizeContentView    WindowStyleMask = 1 << 15
)

type DragOperation = uint32

const (
	DragOperationCopy DragOperation = 1 << iota
	DragOperationLink
	DragOperationGeneric
	DragOperationPrivate
	DragOperationMove
	DragOperationDelete
	DragOperationNone  DragOperation = 0
	DragOperationEvery DragOperation = math.MaxUint32
)

var windowDelegateMap = make(map[C.NSWindowPtr]WindowDelegate)

type (
	WindowStyleMask int
	WindowNative    = C.NSWindowPtr
)

type Window struct {
	native C.NSWindowPtr
}

type WindowDelegate interface {
	WindowDidResize(wnd *Window)
	WindowDidBecomeKey(wnd *Window)
	WindowDidResignKey(wnd *Window)
	WindowShouldClose(wnd *Window) bool
	WindowWillClose(wnd *Window)
	WindowDragEntered(wnd *Window, di *DraggingInfo) DragOperation
	WindowDragUpdated(wnd *Window, di *DraggingInfo) DragOperation
	WindowDragExited(wnd *Window)
	WindowDragEnded(wnd *Window)
	WindowDropIsAcceptable(wnd *Window, di *DraggingInfo) bool
	WindowDrop(wnd *Window, di *DraggingInfo) bool
	WindowDropFinished(wnd *Window, di *DraggingInfo)
}

func WindowContentRectForFrameRectStyleMask(x, y, width, height float64, styleMask WindowStyleMask) (cx, cy, cwidth, cheight float64) {
	r := C.nsWindowContentRectForFrameRectStyleMask(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), C.NSWindowStyleMask(styleMask))
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func WindowFrameRectForContentRectStyleMask(x, y, width, height float64, styleMask WindowStyleMask) (cx, cy, cwidth, cheight float64) {
	r := C.nsWindowFrameRectForContentRectStyleMask(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), C.NSWindowStyleMask(styleMask))
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func WindowInitWithContentRectStyleMask(x, y, width, height float64, styleMask WindowStyleMask) *Window {
	return &Window{native: C.nsWindowInitWithContentRectStyleMask(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), C.NSWindowStyleMask(styleMask))}
}

func (w *Window) Native() WindowNative {
	return w.native
}

func (w *Window) SetDelegate(delegate WindowDelegate) {
	windowDelegateMap[w.native] = delegate
}

func (w *Window) SetIsKeyable(keyable bool) {
	C.nsWindowSetIsKeyable(w.native, C.bool(keyable))
}

func (w *Window) DisableCursorRects() {
	C.nsWindowDisableCursorRects(w.native)
}

func (w *Window) SetTitle(title string) {
	str := cf.StringCreateWithString(title)
	defer str.Release()
	C.nsWindowSetTitle(w.native, C.CFStringRef(str))
}

func (w *Window) Frame() (x, y, width, height float64) {
	r := C.nsWindowFrame(w.native)
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}

func (w *Window) SetFrame(x, y, width, height float64, display bool) {
	C.nsWindowSetFrame(w.native, C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), C.bool(display))
}

func (w *Window) MakeKeyAndOrderFront() {
	C.nsWindowMakeKeyAndOrderFront(w.native)
}

func (w *Window) PerformMiniaturize() {
	C.nsWindowPerformMiniaturize(w.native)
}

func (w *Window) PerformZoom() {
	C.nsWindowPerformZoom(w.native)
}

func (w *Window) ContentView() *View {
	return &View{native: C.nsWindowContentView(w.native)}
}

func (w *Window) SetContentView(view *View) {
	C.nsWindowSetContentView(w.native, view.native)
}

func (w *Window) Close() {
	C.nsWindowClose(w.native)
	delete(windowDelegateMap, w.native)
}

func (w *Window) MouseLocationOutsideOfEventStream() (x, y float64) {
	p := C.nsWindowMouseLocationOutsideOfEventStream(w.native)
	return float64(p.x), float64(p.y)
}

func (w *Window) RegisterForDraggedTypes(dataTypes ...string) {
	a := cf.MutableArrayCreate(len(dataTypes))
	for _, dt := range dataTypes {
		a.AppendStringValue(dt)
	}
	C.nsWindowRegisterForDraggedTypes(w.native, C.CFArrayRef(a.AsCFArray()))
}
