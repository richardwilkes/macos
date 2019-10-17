package ns

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

type AutoreleasePool struct {
	native C.NSAutoreleasePoolPtr
}

func NewAutoreleasePool() *AutoreleasePool {
	return &AutoreleasePool{native: C.nsNSAutoreleasePool()}
}

func (p *AutoreleasePool) Release() {
	if p.native != nil {
		C.nsAutoreleasePoolRelease(p.native)
		p.native = nil
	}
}
