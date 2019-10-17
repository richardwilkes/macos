package cf

import "unsafe"

// #import <CoreFoundation/CoreFoundation.h>
import "C"

type Array = C.CFArrayRef

func (a Array) GetCount() int {
	return int(C.CFArrayGetCount(a))
}

func (a Array) GetValueAtIndex(index int) unsafe.Pointer {
	return C.CFArrayGetValueAtIndex(a, C.CFIndex(index))
}

func (a Array) Release() {
	C.CFRelease(Type(a))
}
