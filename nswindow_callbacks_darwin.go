package macos

/*
typedef void *NSWindowPtr;
typedef void *NSDraggingInfoPtr;
*/
import "C"

//export windowDidResizeCallback
func windowDidResizeCallback(w C.NSWindowPtr) {
	if d, ok := nsWindowDelegateMap[w]; ok {
		d.WindowDidResize(&NSWindow{native: w})
	}
}

//export windowDidBecomeKeyCallback
func windowDidBecomeKeyCallback(w C.NSWindowPtr) {
	if d, ok := nsWindowDelegateMap[w]; ok {
		d.WindowDidBecomeKey(&NSWindow{native: w})
	}
}

//export windowDidResignKeyCallback
func windowDidResignKeyCallback(w C.NSWindowPtr) {
	if d, ok := nsWindowDelegateMap[w]; ok {
		d.WindowDidResignKey(&NSWindow{native: w})
	}
}

//export windowShouldCloseCallback
func windowShouldCloseCallback(w C.NSWindowPtr) bool {
	if d, ok := nsWindowDelegateMap[w]; ok {
		return d.WindowShouldClose(&NSWindow{native: w})
	}
	return true
}

//export windowWillCloseCallback
func windowWillCloseCallback(w C.NSWindowPtr) {
	if d, ok := nsWindowDelegateMap[w]; ok {
		d.WindowWillClose(&NSWindow{native: w})
	}
}

//export windowDraggingEnteredCallback
func windowDraggingEnteredCallback(diNative C.NSDraggingInfoPtr) NSDragOperation {
	di := &NSDraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := nsWindowDelegateMap[wnd.native]; ok {
		return d.WindowDragEntered(wnd, di)
	}
	return NSDragOperationNone
}

//export windowDraggingUpdatedCallback
func windowDraggingUpdatedCallback(diNative C.NSDraggingInfoPtr) NSDragOperation {
	di := &NSDraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := nsWindowDelegateMap[wnd.native]; ok {
		return d.WindowDragUpdated(wnd, di)
	}
	return NSDragOperationNone
}

//export windowDraggingExitedCallback
func windowDraggingExitedCallback(diNative C.NSDraggingInfoPtr) {
	di := &NSDraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := nsWindowDelegateMap[wnd.native]; ok {
		d.WindowDragExited(wnd)
	}
}

//export windowDraggingEndedCallback
func windowDraggingEndedCallback(diNative C.NSDraggingInfoPtr) {
	di := &NSDraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := nsWindowDelegateMap[wnd.native]; ok {
		d.WindowDragEnded(wnd)
	}
}

//export windowPrepareForDragCallback
func windowPrepareForDragCallback(diNative C.NSDraggingInfoPtr) bool {
	di := &NSDraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := nsWindowDelegateMap[wnd.native]; ok {
		return d.WindowDropIsAcceptable(wnd, di)
	}
	return false
}

//export windowPerformDragCallback
func windowPerformDragCallback(diNative C.NSDraggingInfoPtr) bool {
	di := &NSDraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := nsWindowDelegateMap[wnd.native]; ok {
		return d.WindowDrop(wnd, di)
	}
	return false
}

//export windowConcludeDragCallback
func windowConcludeDragCallback(diNative C.NSDraggingInfoPtr) {
	di := &NSDraggingInfo{native: diNative}
	wnd := di.DestinationWindow()
	if d, ok := nsWindowDelegateMap[wnd.native]; ok {
		d.WindowDropFinished(wnd, di)
	}
}
