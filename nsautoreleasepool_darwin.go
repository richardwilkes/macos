package macos

/*
#import <Cocoa/Cocoa.h>

typedef void *NSAutoreleasePoolPtr;

NSAutoreleasePoolPtr nsNSAutoreleasePool() {
	return (NSAutoreleasePoolPtr)[[NSAutoreleasePool alloc] init];
}

void nsAutoreleasePoolRelease(NSAutoreleasePoolPtr p) {
	[(NSAutoreleasePool *)p release];
}
*/
import "C"

type NSAutoreleasePool struct {
	native C.NSAutoreleasePoolPtr
}

func NewNSAutoreleasePool() *NSAutoreleasePool {
	return &NSAutoreleasePool{native: C.nsNSAutoreleasePool()}
}

func (p *NSAutoreleasePool) Release() {
	if p.native != nil {
		C.nsAutoreleasePoolRelease(p.native)
		p.native = nil
	}
}
