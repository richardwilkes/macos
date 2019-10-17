package cg

import "github.com/richardwilkes/macos/cf"

// #import <ImageIO/ImageIO.h>
import "C"

type DataProvider = C.CGDataProviderRef

func DataProviderCreateWithCFData(data cf.Data) DataProvider {
	return C.CGDataProviderCreateWithCFData(C.CFDataRef(data))
}

func (dp DataProvider) Release() {
	C.CGDataProviderRelease(dp)
}
