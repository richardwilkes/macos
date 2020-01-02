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

const (
	GradientDrawsBeforeStartLocation GradientDrawingOptions = 1 << iota
	GradientDrawsAfterEndLocation
)

type (
	Gradient               = C.CGGradientRef
	GradientDrawingOptions int
)

func GradientCreateWithColorComponents(space ColorSpace, components, locations []float64) Gradient {
	return C.CGGradientCreateWithColorComponents(space, floatSliceToCGFloatPtr(components), floatSliceToCGFloatPtr(locations), C.size_t(len(locations)))
}

func (g Gradient) Retain() {
	C.CGGradientRetain(g)
}

func (g Gradient) Release() {
	C.CGGradientRelease(g)
}
