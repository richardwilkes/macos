// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ns

import (
	"github.com/richardwilkes/macos/cf"
	"github.com/richardwilkes/macos/cg"
)

/*
#import <CoreGraphics/CoreGraphics.h>

typedef void *NSViewPtr;
*/
import "C"

//export viewDrawCallback
func viewDrawCallback(view C.NSViewPtr, gc C.CGContextRef, x, y, width, height C.CGFloat, inLiveResize bool) {
	if d, ok := viewDelegateMap[view]; ok {
		d.ViewDraw(&View{native: view}, cg.Context(gc), float64(x), float64(y), float64(width), float64(height), inLiveResize)
	}
}

//export viewMouseDownCallback
func viewMouseDownCallback(view C.NSViewPtr, x, y C.CGFloat, button, clickCount, mod int) {
	if d, ok := viewDelegateMap[view]; ok {
		d.ViewMouseDownEvent(&View{native: view}, float64(x), float64(y), button, clickCount, mod)
	}
}

//export viewMouseDragCallback
func viewMouseDragCallback(view C.NSViewPtr, x, y C.CGFloat, button, mod int) {
	if d, ok := viewDelegateMap[view]; ok {
		d.ViewMouseDragEvent(&View{native: view}, float64(x), float64(y), button, mod)
	}
}

//export viewMouseUpCallback
func viewMouseUpCallback(view C.NSViewPtr, x, y C.CGFloat, button, mod int) {
	if d, ok := viewDelegateMap[view]; ok {
		d.ViewMouseUpEvent(&View{native: view}, float64(x), float64(y), button, mod)
	}
}

//export viewMouseEnterCallback
func viewMouseEnterCallback(view C.NSViewPtr, x, y C.CGFloat, mod int) {
	if d, ok := viewDelegateMap[view]; ok {
		d.ViewMouseEnterEvent(&View{native: view}, float64(x), float64(y), mod)
	}
}

//export viewMouseMoveCallback
func viewMouseMoveCallback(view C.NSViewPtr, x, y C.CGFloat, mod int) {
	if d, ok := viewDelegateMap[view]; ok {
		d.ViewMouseMoveEvent(&View{native: view}, float64(x), float64(y), mod)
	}
}

//export viewMouseExitCallback
func viewMouseExitCallback(view C.NSViewPtr) {
	if d, ok := viewDelegateMap[view]; ok {
		d.ViewMouseExitEvent(&View{native: view})
	}
}

//export viewMouseWheelCallback
func viewMouseWheelCallback(view C.NSViewPtr, x, y, dx, dy C.CGFloat, mod int) {
	if d, ok := viewDelegateMap[view]; ok {
		d.ViewMouseWheelEvent(&View{native: view}, float64(x), float64(y), float64(dx), float64(dy), mod)
	}
}

//export viewCursorUpdateCallback
func viewCursorUpdateCallback(view C.NSViewPtr, x, y C.CGFloat, mod int) {
	if d, ok := viewDelegateMap[view]; ok {
		d.ViewCursorUpdateEvent(&View{native: view}, float64(x), float64(y), mod)
	}
}

//export viewKeyDownCallback
func viewKeyDownCallback(view C.NSViewPtr, keyCode int, ch C.CFStringRef, mod int, repeat bool) {
	if d, ok := viewDelegateMap[view]; ok {
		var r rune
		if ch != 0 {
			if runes := []rune(cf.String(ch).String()); len(runes) > 0 {
				r = runes[0]
			}
		}
		d.ViewKeyDownEvent(&View{native: view}, keyCode, r, mod, repeat)
	}
}

//export viewKeyUpCallback
func viewKeyUpCallback(view C.NSViewPtr, keyCode, mod int) {
	if d, ok := viewDelegateMap[view]; ok {
		d.ViewKeyUpEvent(&View{native: view}, keyCode, mod)
	}
}
