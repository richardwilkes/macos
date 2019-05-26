package macos

// typedef void *NSWindowPtr;
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
