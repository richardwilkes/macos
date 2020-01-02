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

var AffineTransformIdentity = C.CGAffineTransformIdentity

type AffineTransform = C.CGAffineTransform

func AffineTransformMake(a, b, c, d, tx, ty float64) AffineTransform {
	return C.CGAffineTransform{
		a:  C.CGFloat(a),
		b:  C.CGFloat(b),
		c:  C.CGFloat(c),
		d:  C.CGFloat(d),
		tx: C.CGFloat(tx),
		ty: C.CGFloat(ty),
	}
}
