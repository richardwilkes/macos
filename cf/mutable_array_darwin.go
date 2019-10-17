package cf

import "unsafe"

// #import <CoreFoundation/CoreFoundation.h>
import "C"

type MutableArray struct {
	native C.CFMutableArrayRef
}

func MutableArrayCreate(capacity int) *MutableArray {
	return &MutableArray{native: C.CFArrayCreateMutable(0, C.CFIndex(capacity), &C.kCFTypeArrayCallBacks)} //nolint:gocritic,staticcheck
}

func MutableArrayCreateNoCap() *MutableArray {
	return MutableArrayCreate(0)
}

func (a *MutableArray) AppendValue(value unsafe.Pointer) {
	C.CFArrayAppendValue(a.native, value)
}

func (a *MutableArray) AppendStringValue(value string) {
	C.CFArrayAppendValue(a.native, unsafe.Pointer(StringCreateWithString(value))) //nolint:govet
}

func (a *MutableArray) AsCFArray() Array {
	return Array(unsafe.Pointer(a.native))
}
