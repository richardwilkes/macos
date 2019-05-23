package macos

// #import <CoreGraphics/CoreGraphics.h>
import "C"

//export patternDrawCallback
func patternDrawCallback(id *int32, gc C.CGContextRef) {
	patternLock.Lock()
	callback, ok := patternCallbackMap[*id]
	patternLock.Unlock()
	if ok {
		callback.PatternDraw(gc)
	}
}

//export patternReleaseCallback
func patternReleaseCallback(id *int32) {
	patternLock.Lock()
	callback, ok := patternCallbackMap[*id]
	delete(patternCallbackMap, *id)
	patternLock.Unlock()
	if ok {
		callback.PatternRelease()
	}
}
