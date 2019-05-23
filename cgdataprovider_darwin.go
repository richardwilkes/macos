package macos

// #import <ImageIO/ImageIO.h>
import "C"

type CGDataProvider = C.CGDataProviderRef

func CGDataProviderCreateWithCFData(data CFData) CGDataProvider {
	return C.CGDataProviderCreateWithCFData(data)
}

func (dp CGDataProvider) Release() {
	C.CGDataProviderRelease(dp)
}
