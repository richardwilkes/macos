package macos

/*
#import <CoreGraphics/CoreGraphics.h>

typedef void *NSViewPtr;
*/
import "C"

//export viewDrawCallback
func viewDrawCallback(view C.NSViewPtr, gc CGContext, x, y, width, height C.CGFloat, inLiveResize bool) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewDraw(&NSView{native: view}, gc, float64(x), float64(y), float64(width), float64(height), inLiveResize)
	}
}

//export viewMouseDownCallback
func viewMouseDownCallback(view C.NSViewPtr, x, y C.CGFloat, button, clickCount, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseDownEvent(&NSView{native: view}, float64(x), float64(y), button, clickCount, mod)
	}
}

//export viewMouseDragCallback
func viewMouseDragCallback(view C.NSViewPtr, x, y C.CGFloat, button, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseDragEvent(&NSView{native: view}, float64(x), float64(y), button, mod)
	}
}

//export viewMouseUpCallback
func viewMouseUpCallback(view C.NSViewPtr, x, y C.CGFloat, button, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseUpEvent(&NSView{native: view}, float64(x), float64(y), button, mod)
	}
}

//export viewMouseEnterCallback
func viewMouseEnterCallback(view C.NSViewPtr, x, y C.CGFloat, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseEnterEvent(&NSView{native: view}, float64(x), float64(y), mod)
	}
}

//export viewMouseMoveCallback
func viewMouseMoveCallback(view C.NSViewPtr, x, y C.CGFloat, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseMoveEvent(&NSView{native: view}, float64(x), float64(y), mod)
	}
}

//export viewMouseExitCallback
func viewMouseExitCallback(view C.NSViewPtr) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseExitEvent(&NSView{native: view})
	}
}

//export viewMouseWheelCallback
func viewMouseWheelCallback(view C.NSViewPtr, x, y, dx, dy C.CGFloat, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseWheelEvent(&NSView{native: view}, float64(x), float64(y), float64(dx), float64(dy), mod)
	}
}

//export viewCursorUpdateCallback
func viewCursorUpdateCallback(view C.NSViewPtr, x, y C.CGFloat, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewCursorUpdateEvent(&NSView{native: view}, float64(x), float64(y), mod)
	}
}

//export viewKeyDownCallback
func viewKeyDownCallback(view C.NSViewPtr, keyCode int, ch CFString, mod int, repeat bool) {
	if d, ok := nsViewDelegateMap[view]; ok {
		var r rune
		if ch != 0 {
			if runes := []rune(ch.String()); len(runes) > 0 {
				r = runes[0]
			}
		}
		d.ViewKeyDownEvent(&NSView{native: view}, keyCode, r, mod, repeat)
	}
}

//export viewKeyUpCallback
func viewKeyUpCallback(view C.NSViewPtr, keyCode, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewKeyUpEvent(&NSView{native: view}, keyCode, mod)
	}
}
