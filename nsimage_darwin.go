package macos

/*
#import <Cocoa/Cocoa.h>
#import <CoreGraphics/CoreGraphics.h>

typedef void *NSImagePtr;

NSImagePtr nsImageInitWithCGImageSizeRetain(CGImageRef img, CGFloat width, CGFloat height) {
	return [[[NSImage alloc] initWithCGImage:img size:NSMakeSize(width, height)] retain];
}

void nsImageRelease(NSImagePtr img) {
	[((NSImage *)img) release];
}
*/
import "C"

import "github.com/richardwilkes/toolbox/softref"

type NSImage softref.SoftRef

type nsImageRef struct {
	key    string
	native C.NSImagePtr
}

func NSImageInitWithCGImageSizeRetain(img CGImage, width, height float64) *NSImage {
	ref, _ := softref.DefaultPool.NewSoftRef(&nsImageRef{
		key:    NextRefKey(),
		native: C.nsImageInitWithCGImageSizeRetain(img, C.CGFloat(width), C.CGFloat(height)),
	})
	return (*NSImage)(ref)
}

func (img *NSImage) native() C.NSImagePtr {
	return img.Resource.(*nsImageRef).native
}

func (r *nsImageRef) Key() string {
	return r.key
}

func (r *nsImageRef) Release() {
	C.nsImageRelease(r.native)
	r.native = nil
}
