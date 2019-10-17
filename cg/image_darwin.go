package cg

// #import <ImageIO/ImageIO.h>
import "C"

const (
	RenderingIntentDefault ColorRenderingIntent = iota
	RenderingIntentAbsoluteColorimetric
	RenderingIntentRelativeColorimetric
	RenderingIntentPerceptual
	RenderingIntentSaturation
)

const (
	ImageAlphaNone ImageAlphaInfo = iota
	ImageAlphaPremultipliedLast
	ImageAlphaPremultipliedFirst
	ImageAlphaLast
	ImageAlphaFirst
	ImageAlphaNoneSkipLast
	ImageAlphaNoneSkipFirst
	ImageAlphaOnly
)

const (
	ImageByteOrderDefault ImageByteOrderInfo = iota << 12
	ImageByteOrder16Little
	ImageByteOrder32Little
	ImageByteOrder16Big
	ImageByteOrder32Big
	ImageByteOrderMask ImageByteOrderInfo = 0x7000
)

const (
	BitmapAlphaInfoMask           BitmapInfo = 0x1F
	BitmapAlphaNone                          = BitmapInfo(ImageAlphaNone)
	BitmapAlphaPremultipliedLast             = BitmapInfo(ImageAlphaPremultipliedLast)
	BitmapAlphaPremultipliedFirst            = BitmapInfo(ImageAlphaPremultipliedFirst)
	BitmapAlphaLast                          = BitmapInfo(ImageAlphaLast)
	BitmapAlphaFirst                         = BitmapInfo(ImageAlphaFirst)
	BitmapAlphaNoneSkipLast                  = BitmapInfo(ImageAlphaNoneSkipLast)
	BitmapAlphaNoneSkipFirst                 = BitmapInfo(ImageAlphaNoneSkipFirst)
	BitmapAlphaOnly                          = BitmapInfo(ImageAlphaOnly)
	BitmapFloatInfoMask           BitmapInfo = 0xF00
	BitmapFloatComponents         BitmapInfo = 1 << 8
	BitmapByteOrderMask                      = BitmapInfo(ImageByteOrderMask)
	BitmapByteOrderDefault                   = BitmapInfo(ImageByteOrderDefault)
	BitmapByteOrder16Little                  = BitmapInfo(ImageByteOrder16Little)
	BitmapByteOrder32Little                  = BitmapInfo(ImageByteOrder32Little)
	BitmapByteOrder16Big                     = BitmapInfo(ImageByteOrder16Big)
	BitmapByteOrder32Big                     = BitmapInfo(ImageByteOrder32Big)
	BitmapByteOrder16Host                    = BitmapInfo(C.kCGBitmapByteOrder16Host)
	BitmapByteOrder32Host                    = BitmapInfo(C.kCGBitmapByteOrder32Host)
)

type (
	Image                = C.CGImageRef
	ColorRenderingIntent int32
	ImageAlphaInfo       uint32
	ImageByteOrderInfo   uint32
	BitmapInfo           uint32
)

func ImageCreate(width, height, bitsPerComponent, bitsPerPixel, bytesPerRow int, space ColorSpace, bitmapInfo BitmapInfo, provider DataProvider, decode []float64, shouldInterpolate bool, intent ColorRenderingIntent) Image {
	return C.CGImageCreate(C.size_t(width), C.size_t(height), C.size_t(bitsPerComponent), C.size_t(bitsPerPixel), C.size_t(bytesPerRow), space, C.CGBitmapInfo(bitmapInfo), provider, floatSliceToCGFloatPtr(decode), C.bool(shouldInterpolate), C.CGColorRenderingIntent(intent))
}

func (img Image) CreateWithImageInRect(x, y, width, height float64) Image {
	return C.CGImageCreateWithImageInRect(img, C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)))
}

func (img Image) GetWidth() int {
	return int(C.CGImageGetWidth(img))
}

func (img Image) GetHeight() int {
	return int(C.CGImageGetHeight(img))
}

func (img Image) Release() {
	C.CGImageRelease(img)
}
