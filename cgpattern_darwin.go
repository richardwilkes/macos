package macos

import (
	"sync"
	"unsafe"
)

/*
#import <CoreGraphics/CoreGraphics.h>

void patternDrawCallback(void *info, CGContextRef context);
void patternReleaseCallback(void *info);

CGPatternCallbacks patternCallbacks = { 0, patternDrawCallback, patternReleaseCallback };
CGPatternCallbacks *patternCallbacksPtr = &patternCallbacks;
*/
import "C"

const (
	CGPatternTilingNoDistortion CGPatternTiling = iota
	CGPatternTilingConstantSpacingMinimalDistortion
	CGPatternTilingConstantSpacing
)

var (
	nextPatternID      int32 = 1
	patternLock        sync.Mutex
	patternCallbackMap = make(map[int32]CGPatternCallback)
)

type (
	CGPattern       = C.CGPatternRef
	CGPatternTiling int
)

type CGPatternCallback interface {
	PatternDraw(gc CGContext)
	PatternRelease()
}

func CGPatternCreate(x, y, width, height float64, matrix CGAffineTransform, xStep, yStep float64, tiling CGPatternTiling, isColored bool, callbacks CGPatternCallback) CGPattern {
	patternLock.Lock()
	id := nextPatternID
	nextPatternID++
	patternCallbackMap[id] = callbacks
	patternLock.Unlock()
	return C.CGPatternCreate(unsafe.Pointer(&id), C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)), matrix, C.CGFloat(xStep), C.CGFloat(yStep), C.CGPatternTiling(tiling), C.bool(isColored), C.patternCallbacksPtr) //nolint:gocritic,staticcheck
}

func (p CGPattern) Release() {
	C.CGPatternRelease(p)
}
