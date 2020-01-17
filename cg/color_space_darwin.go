// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package cg

// #import <CoreGraphics/CoreGraphics.h>
import "C"

type ColorSpace = C.CGColorSpaceRef

func ColorSpaceCreateDeviceRGB() ColorSpace {
	return C.CGColorSpaceCreateDeviceRGB()
}

func ColorSpaceCreatePattern(baseSpace ColorSpace) ColorSpace {
	return C.CGColorSpaceCreatePattern(baseSpace)
}

func ColorSpaceCreateSRGB() ColorSpace {
	return C.CGColorSpaceCreateWithName(C.kCGColorSpaceSRGB)
}

func (cs ColorSpace) Release() {
	C.CGColorSpaceRelease(cs)
}
