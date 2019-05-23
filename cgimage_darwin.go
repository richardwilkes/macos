package macos

// #import <ImageIO/ImageIO.h>
import "C"

const (
	CGRenderingIntentDefault CGColorRenderingIntent = iota
	CGRenderingIntentAbsoluteColorimetric
	CGRenderingIntentRelativeColorimetric
	CGRenderingIntentPerceptual
	CGRenderingIntentSaturation
)

const (
	CGImageByteOrderMask     CGImageByteOrderInfo = 0x7000
	CGImageByteOrderDefault                       = 0 << 12
	CGImageByteOrder16Little                      = 1 << 12
	CGImageByteOrder32Little                      = 2 << 12
	CGImageByteOrder16Big                         = 3 << 12
	CGImageByteOrder32Big                         = 4 << 12
)

const (
	CGBitmapAlphaInfoMask     CGBitmapInfo = 0x1F
	CGBitmapFloatInfoMask                  = 0xF00
	CGBitmapFloatComponents                = 1 << 8
	CGBitmapByteOrderMask                  = CGBitmapInfo(CGImageByteOrderMask)
	CGBitmapByteOrderDefault               = CGImageByteOrderDefault
	CGBitmapByteOrder16Little              = CGImageByteOrder16Little
	CGBitmapByteOrder32Little              = CGImageByteOrder32Little
	CGBitmapByteOrder16Big                 = CGImageByteOrder16Big
	CGBitmapByteOrder32Big                 = CGImageByteOrder32Big
)

type (
	CGImage                = C.CGImageRef
	CGColorRenderingIntent int32
	CGImageByteOrderInfo   uint32
	CGBitmapInfo           uint32
)

func CGImageCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow int, space CGColorSpace, bitmapInfo CGBitmapInfo, provider CGDataProvider, decode []float64, shouldInterpolate bool, intent CGColorRenderingIntent) CGImage {
	return C.CGImageCreate(C.size_t(width), C.size_t(height), C.size_t(bitsPerComponent), C.size_t(bitsPerPixel), C.size_t(bytesPerRow), space, C.CGBitmapInfo(bitmapInfo), provider, floatSliceToCGFloatPtr(decode), C.bool(shouldInterpolate), C.CGColorRenderingIntent(intent))
}

func (img CGImage) CreateWithImageInRect(x, y, width, height float64) CGImage {
	return C.CGImageCreateWithImageInRect(img, C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)))
}

func (img CGImage) GetWidth() int {
	return int(C.CGImageGetWidth(img))
}

func (img CGImage) GetHeight() int {
	return int(C.CGImageGetHeight(img))
}

func (img CGImage) Release() {
	C.CGImageRelease(img)
}
