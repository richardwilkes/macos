// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ns

/*
typedef void *NSWindowPtr;
typedef void *NSDraggingInfoPtr;
*/
import "C"

//export windowDidResizeCallback
func windowDidResizeCallback(w C.NSWindowPtr) {
	if d, ok := windowDelegateMap[w]; ok {
		d.WindowDidResize(&Window{native: w})
	}
}

//export windowDidBecomeKeyCallback
func windowDidBecomeKeyCallback(w C.NSWindowPtr) {
	if d, ok := windowDelegateMap[w]; ok {
		d.WindowDidBecomeKey(&Window{native: w})
	}
}

//export windowDidResignKeyCallback
func windowDidResignKeyCallback(w C.NSWindowPtr) {
	if d, ok := windowDelegateMap[w]; ok {
		d.WindowDidResignKey(&Window{native: w})
	}
}

//export windowShouldCloseCallback
func windowShouldCloseCallback(w C.NSWindowPtr) bool {
	if d, ok := windowDelegateMap[w]; ok {
		return d.WindowShouldClose(&Window{native: w})
	}
	return true
}

//export windowWillCloseCallback
func windowWillCloseCallback(w C.NSWindowPtr) {
	if d, ok := windowDelegateMap[w]; ok {
		d.WindowWillClose(&Window{native: w})
	}
}

//export windowDraggingEnteredCallback
func windowDraggingEnteredCallback(diNative C.NSDraggingInfoPtr) DragOperation {
	di := &DraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := windowDelegateMap[wnd.native]; ok {
		return d.WindowDragEntered(wnd, di)
	}
	return DragOperationNone
}

//export windowDraggingUpdatedCallback
func windowDraggingUpdatedCallback(diNative C.NSDraggingInfoPtr) DragOperation {
	di := &DraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := windowDelegateMap[wnd.native]; ok {
		return d.WindowDragUpdated(wnd, di)
	}
	return DragOperationNone
}

//export windowDraggingExitedCallback
func windowDraggingExitedCallback(diNative C.NSDraggingInfoPtr) {
	di := &DraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := windowDelegateMap[wnd.native]; ok {
		d.WindowDragExited(wnd)
	}
}

//export windowDraggingEndedCallback
func windowDraggingEndedCallback(diNative C.NSDraggingInfoPtr) {
	di := &DraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := windowDelegateMap[wnd.native]; ok {
		d.WindowDragEnded(wnd)
	}
}

//export windowPrepareForDragCallback
func windowPrepareForDragCallback(diNative C.NSDraggingInfoPtr) bool {
	di := &DraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := windowDelegateMap[wnd.native]; ok {
		return d.WindowDropIsAcceptable(wnd, di)
	}
	return false
}

//export windowPerformDragCallback
func windowPerformDragCallback(diNative C.NSDraggingInfoPtr) bool {
	di := &DraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := windowDelegateMap[wnd.native]; ok {
		return d.WindowDrop(wnd, di)
	}
	return false
}

//export windowConcludeDragCallback
func windowConcludeDragCallback(diNative C.NSDraggingInfoPtr) {
	di := &DraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := windowDelegateMap[wnd.native]; ok {
		d.WindowDropFinished(wnd, di)
	}
}
