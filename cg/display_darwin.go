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

type DirectDisplayID = C.CGDirectDisplayID

func DisplayIsMain(id DirectDisplayID) bool {
	return C.CGDisplayIsMain(id) != 0
}

func DisplayBounds(id DirectDisplayID) (x, y, width, height float64) {
	r := C.CGDisplayBounds(id)
	return float64(r.origin.x), float64(r.origin.y), float64(r.size.width), float64(r.size.height)
}
