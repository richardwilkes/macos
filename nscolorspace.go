package macos

/*
#import <Cocoa/Cocoa.h>

typedef void *NSColorSpacePtr;

NSColorSpacePtr nsColorSpaceDeviceRGBColorSpace() {
	return [NSColorSpace deviceRGBColorSpace];
}
*/
import "C"

type NSColorSpace struct {
	native C.NSColorSpacePtr
}

func NSColorSpaceDeviceRGBColorSpace() *NSColorSpace {
	return &NSColorSpace{native: C.nsColorSpaceDeviceRGBColorSpace()}
}
