package macos

import "unsafe"

// #import <CoreFoundation/CoreFoundation.h>
import "C"

type CFData = C.CFDataRef

func CFDataCreate(data []byte) CFData {
	return C.CFDataCreate(0, (*C.UInt8)(&data[0]), C.CFIndex(len(data)))
}

func (d CFData) GetLength() int {
	return int(C.CFDataGetLength(d))
}

func (d CFData) GetBytes(location, length int) []byte {
	buffer := make([]byte, length)
	C.CFDataGetBytes(d, C.CFRangeMake(C.CFIndex(location), C.CFIndex(length)), (*C.UInt8)(unsafe.Pointer(&buffer[0])))
	return buffer
}

func (d CFData) Release() {
	C.CFRelease(CFType(d))
}
