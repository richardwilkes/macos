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

type (
	Layer = C.CGLayerRef
)

func LayerCreateWithContext(ctx Context, width, height float64) Layer {
	return C.CGLayerCreateWithContext(ctx, C.CGSizeMake(C.CGFloat(width), C.CGFloat(height)), 0)
}

func (layer Layer) Size() (width, height float64) {
	size := C.CGLayerGetSize(layer)
	return float64(size.width), float64(size.height)
}

func (layer Layer) Context() Context {
	return C.CGLayerGetContext(layer)
}

func (layer Layer) Retain() {
	C.CGLayerRetain(layer)
}

func (layer Layer) Release() {
	C.CGLayerRelease(layer)
}
