package cg

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
	PatternTilingNoDistortion PatternTiling = iota
	PatternTilingConstantSpacingMinimalDistortion
	PatternTilingConstantSpacing
)

var (
	nextPatternID      int32 = 1
	patternLock        sync.Mutex
	patternCallbackMap = make(map[int32]PatternCallback)
)

type (
	Pattern       = C.CGPatternRef
	PatternTiling int
)

type PatternCallback interface {
	PatternDraw(gc Context)
	PatternRelease()
}

func PatternCreate(x, y, width, height float64, matrix AffineTransform, xStep, yStep float64, tiling PatternTiling, isColored bool, callbacks PatternCallback) Pattern {
	patternLock.Lock()
	id := nextPatternID
	nextPatternID++
	patternCallbackMap[id] = callbacks
	patternLock.Unlock()
	return C.CGPatternCreate(unsafe.Pointer(&id), C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)), matrix, C.CGFloat(xStep), C.CGFloat(yStep), C.CGPatternTiling(tiling), C.bool(isColored), C.patternCallbacksPtr) //nolint:gocritic,staticcheck
}

func (p Pattern) Release() {
	C.CGPatternRelease(p)
}
