package macos

// #import <CoreGraphics/CoreGraphics.h>
import "C"

var CGAffineTransformIdentity = C.CGAffineTransformIdentity

type CGAffineTransform = C.CGAffineTransform

func CGAffineTransformMake(a, b, c, d, tx, ty float64) CGAffineTransform {
	return C.CGAffineTransform{
		a:  C.CGFloat(a),
		b:  C.CGFloat(b),
		c:  C.CGFloat(c),
		d:  C.CGFloat(d),
		tx: C.CGFloat(tx),
		ty: C.CGFloat(ty),
	}
}
