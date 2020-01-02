// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

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
