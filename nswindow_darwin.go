package macos

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
import (
	"math"
)

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

type NSDragOperation = uint32

const (
	NSDragOperationCopy NSDragOperation = 1 << iota
	NSDragOperationLink
	NSDragOperationGeneric
	NSDragOperationPrivate
	NSDragOperationMove
	NSDragOperationDelete
	NSDragOperationNone  NSDragOperation = 0
	NSDragOperationEvery NSDragOperation = math.MaxUint32
)

var nsWindowDelegateMap = make(map[C.NSWindowPtr]NSWindowDelegate)

type (
	NSWindowStyleMask int
	NSWindowNative    = C.NSWindowPtr
)

type NSWindow struct {
	native C.NSWindowPtr
}

type NSWindowDelegate interface {
	WindowDidResize(wnd *NSWindow)
	WindowDidBecomeKey(wnd *NSWindow)
	WindowDidResignKey(wnd *NSWindow)
	WindowShouldClose(wnd *NSWindow) bool
	WindowWillClose(wnd *NSWindow)
	WindowDragEntered(wnd *NSWindow, di *NSDraggingInfo) NSDragOperation
	WindowDragUpdated(wnd *NSWindow, di *NSDraggingInfo) NSDragOperation
	WindowDragExited(wnd *NSWindow)
	WindowDragEnded(wnd *NSWindow)
	WindowDropIsAcceptable(wnd *NSWindow, di *NSDraggingInfo) bool
	WindowDrop(wnd *NSWindow, di *NSDraggingInfo) bool
	WindowDropFinished(wnd *NSWindow, di *NSDraggingInfo)
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

func (w *NSWindow) SetDelegate(delegate NSWindowDelegate) {
	nsWindowDelegateMap[w.native] = delegate
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

func (w *NSWindow) ContentView() *NSView {
	return &NSView{native: C.nsWindowContentView(w.native)}
}

func (w *NSWindow) SetContentView(view *NSView) {
	C.nsWindowSetContentView(w.native, view.native)
}

func (w *NSWindow) Close() {
	C.nsWindowClose(w.native)
	delete(nsWindowDelegateMap, w.native)
}

func (w *NSWindow) MouseLocationOutsideOfEventStream() (x, y float64) {
	p := C.nsWindowMouseLocationOutsideOfEventStream(w.native)
	return float64(p.x), float64(p.y)
}

func (w *NSWindow) RegisterForDraggedTypes(dataTypes ...string) {
	a := CFMutableArrayCreate(len(dataTypes))
	for _, dt := range dataTypes {
		a.AppendStringValue(dt)
	}
	C.nsWindowRegisterForDraggedTypes(w.native, a.AsCFArray())
}
