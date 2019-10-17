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
