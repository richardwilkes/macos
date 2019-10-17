package cf

import "unsafe"

// #import <CoreFoundation/CoreFoundation.h>
import "C"

type Data = C.CFDataRef

func DataCreate(data []byte) Data {
	return C.CFDataCreate(0, (*C.UInt8)(&data[0]), C.CFIndex(len(data)))
}

func (d Data) GetLength() int {
	return int(C.CFDataGetLength(d))
}

func (d Data) GetBytes(location, length int) []byte {
	buffer := make([]byte, length)
	C.CFDataGetBytes(d, C.CFRangeMake(C.CFIndex(location), C.CFIndex(length)), (*C.UInt8)(unsafe.Pointer(&buffer[0])))
	return buffer
}

func (d Data) Release() {
	C.CFRelease(Type(d))
}
