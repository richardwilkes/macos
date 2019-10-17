package ns

import (
	"github.com/richardwilkes/macos"
	"github.com/richardwilkes/macos/cg"
	"github.com/richardwilkes/toolbox/softref"
)

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

type Image softref.SoftRef

type imageRef struct {
	key    string
	native C.NSImagePtr
}

func ImageInitWithCGImageSizeRetain(img cg.Image, width, height float64) *Image {
	ref, _ := softref.DefaultPool.NewSoftRef(&imageRef{
		key:    macos.NextRefKey(),
		native: C.nsImageInitWithCGImageSizeRetain(C.CGImageRef(img), C.CGFloat(width), C.CGFloat(height)),
	})
	return (*Image)(ref)
}

func (img *Image) native() C.NSImagePtr {
	return img.Resource.(*imageRef).native
}

func (r *imageRef) Key() string {
	return r.key
}

func (r *imageRef) Release() {
	C.nsImageRelease(r.native)
	r.native = nil
}
