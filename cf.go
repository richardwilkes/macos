package macos

// #import <CoreFoundation/CoreFoundation.h>
import "C"

type (
	CFData       = C.CFDataRef
	CFDictionary = C.CFDictionaryRef
)

func CFDataCreate(data []byte) CFData {
	return C.CFDataCreate(0, (*C.UInt8)(&data[0]), C.CFIndex(len(data)))
}

func (d CFData) Release() {
	C.CFRelease(C.CFTypeRef(d))
}
