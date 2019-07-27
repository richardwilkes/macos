package macos

// #import <CoreFoundation/CoreFoundation.h>
import "C"
import "unsafe"

type CFMutableArray struct {
	native C.CFMutableArrayRef
}

func CFMutableArrayCreate(capacity int) *CFMutableArray {
	return &CFMutableArray{native: C.CFArrayCreateMutable(0, C.CFIndex(capacity), &C.kCFTypeArrayCallBacks)}
}

func CFMutableArrayCreateNoCap() *CFMutableArray {
	return CFMutableArrayCreate(0)
}

func (a *CFMutableArray) AppendValue(value unsafe.Pointer) {
	C.CFArrayAppendValue(a.native, value)
}

func (a *CFMutableArray) AppendStringValue(value string) {
	C.CFArrayAppendValue(a.native, unsafe.Pointer(CFStringCreateWithString(value)))
}

func (a *CFMutableArray) AsCFArray() CFArray {
	return CFArray(unsafe.Pointer(a.native))
}
