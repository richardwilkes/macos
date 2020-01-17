// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package cg

import "github.com/richardwilkes/macos/cf"

// #import <CoreGraphics/CoreGraphics.h>
import "C"

type ColorSpace = C.CGColorSpaceRef

const ColorSpaceSRGBName = "CGColorSpaceSRGB"

func ColorSpaceCreateDeviceRGB() ColorSpace {
	return C.CGColorSpaceCreateDeviceRGB()
}

func ColorSpaceCreatePattern(baseSpace ColorSpace) ColorSpace {
	return C.CGColorSpaceCreatePattern(baseSpace)
}

func ColorSpaceCreateWithName(name string) ColorSpace {
	str := cf.StringCreateWithString(name)
	defer str.Release()
	return C.CGColorSpaceCreateWithName(C.CFStringRef(str))
}

func (cs ColorSpace) Release() {
	C.CGColorSpaceRelease(cs)
}
