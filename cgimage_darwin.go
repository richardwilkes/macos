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
	CGImageAlphaNone CGImageAlphaInfo = iota
	CGImageAlphaPremultipliedLast
	CGImageAlphaPremultipliedFirst
	CGImageAlphaLast
	CGImageAlphaFirst
	CGImageAlphaNoneSkipLast
	CGImageAlphaNoneSkipFirst
	CGImageAlphaOnly
)

const (
	CGImageByteOrderDefault CGImageByteOrderInfo = iota << 12
	CGImageByteOrder16Little
	CGImageByteOrder32Little
	CGImageByteOrder16Big
	CGImageByteOrder32Big
	CGImageByteOrderMask CGImageByteOrderInfo = 0x7000
)

const (
	CGBitmapAlphaInfoMask           CGBitmapInfo = 0x1F
	CGBitmapAlphaNone                            = CGBitmapInfo(CGImageAlphaNone)
	CGBitmapAlphaPremultipliedLast               = CGBitmapInfo(CGImageAlphaPremultipliedLast)
	CGBitmapAlphaPremultipliedFirst              = CGBitmapInfo(CGImageAlphaPremultipliedFirst)
	CGBitmapAlphaLast                            = CGBitmapInfo(CGImageAlphaLast)
	CGBitmapAlphaFirst                           = CGBitmapInfo(CGImageAlphaFirst)
	CGBitmapAlphaNoneSkipLast                    = CGBitmapInfo(CGImageAlphaNoneSkipLast)
	CGBitmapAlphaNoneSkipFirst                   = CGBitmapInfo(CGImageAlphaNoneSkipFirst)
	CGBitmapAlphaOnly                            = CGBitmapInfo(CGImageAlphaOnly)
	CGBitmapFloatInfoMask           CGBitmapInfo = 0xF00
	CGBitmapFloatComponents         CGBitmapInfo = 1 << 8
	CGBitmapByteOrderMask                        = CGBitmapInfo(CGImageByteOrderMask)
	CGBitmapByteOrderDefault                     = CGBitmapInfo(CGImageByteOrderDefault)
	CGBitmapByteOrder16Little                    = CGBitmapInfo(CGImageByteOrder16Little)
	CGBitmapByteOrder32Little                    = CGBitmapInfo(CGImageByteOrder32Little)
	CGBitmapByteOrder16Big                       = CGBitmapInfo(CGImageByteOrder16Big)
	CGBitmapByteOrder32Big                       = CGBitmapInfo(CGImageByteOrder32Big)
	CGBitmapByteOrder16Host                      = CGBitmapInfo(C.kCGBitmapByteOrder16Host)
	CGBitmapByteOrder32Host                      = CGBitmapInfo(C.kCGBitmapByteOrder32Host)
)

type (
	CGImage                = C.CGImageRef
	CGColorRenderingIntent int32
	CGImageAlphaInfo       uint32
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
