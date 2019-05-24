package macos

/*
#import <Cocoa/Cocoa.h>
#import <CoreGraphics/CoreGraphics.h>

typedef void *NSImagePtr;

NSImagePtr nsImageInitWithCGImageSizeRetain(CGImageRef img, CGFloat width, CGFloat height) {
	return [[[NSImage alloc] initWithCGImage:img size:NSMakeSize(width, height)] retain];
}

void nsImageRetain(NSImagePtr img) {
	[((NSImage *)img) retain];
}

void nsImageRelease(NSImagePtr img) {
	[((NSImage *)img) release];
}
*/
import "C"

type NSImage struct {
	native     C.NSImagePtr
	refcnt     int
	releasable bool
}

func NSImageInitWithCGImageSizeRetain(img CGImage, width, height float64) *NSImage {
	return &NSImage{
		native: C.nsImageInitWithCGImageSizeRetain(img, C.CGFloat(width), C.CGFloat(height)),
		refcnt: 1,
	}
}

func (img *NSImage) Retain() {
	if img.releasable && img.native != nil {
		img.refcnt++
		C.nsImageRetain(img.native)
	}
}

func (img *NSImage) Release() {
	if img.releasable && img.native != nil {
		img.refcnt--
		if img.refcnt <= 0 {
			C.nsImageRelease(img.native)
			img.native = nil
		}
	}
}
