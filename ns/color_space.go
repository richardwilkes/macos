package ns

/*
#import <Cocoa/Cocoa.h>

typedef void *NSColorSpacePtr;

NSColorSpacePtr nsColorSpaceDeviceRGBColorSpace() {
	return [NSColorSpace deviceRGBColorSpace];
}
*/
import "C"

type ColorSpace struct {
	native C.NSColorSpacePtr
}

func ColorSpaceDeviceRGBColorSpace() *ColorSpace {
	return &ColorSpace{native: C.nsColorSpaceDeviceRGBColorSpace()}
}
