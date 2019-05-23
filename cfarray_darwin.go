package macos

// #import <CoreFoundation/CoreFoundation.h>
import "C"
import "unsafe"

type CFArray = C.CFArrayRef

func (a CFArray) GetCount() int {
	return int(C.CFArrayGetCount(a))
}

func (a CFArray) GetValueAtIndex(index int) unsafe.Pointer {
	return C.CFArrayGetValueAtIndex(a, C.CFIndex(index))
}

func (a CFArray) Release() {
	C.CFRelease(CFType(a))
}
