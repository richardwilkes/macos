// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ns

import "github.com/richardwilkes/macos/cf"

/*
#import <Cocoa/Cocoa.h>

typedef void *NSViewPtr;
typedef void *NSTrackingAreaPtr;

NSTrackingAreaPtr nsTrackingAreaInitWithRectOptionsOwnerUserInfo(CGFloat x, CGFloat y, CGFloat width, CGFloat height, NSTrackingAreaOptions options, NSViewPtr owner, CFDictionaryRef userInfo) {
	return (NSTrackingAreaPtr)[[NSTrackingArea alloc] initWithRect:NSMakeRect(x, y, width, height) options:options owner:(NSView *)owner userInfo:(NSDictionary *)userInfo];
}
*/
import "C"

const (
	TrackingMouseEnteredAndExited    TrackingAreaOptions = 0x01
	TrackingMouseMoved               TrackingAreaOptions = 0x02
	TrackingCursorUpdate             TrackingAreaOptions = 0x04
	TrackingActiveWhenFirstResponder TrackingAreaOptions = 0x10
	TrackingActiveInKeyWindow        TrackingAreaOptions = 0x20
	TrackingActiveInActiveApp        TrackingAreaOptions = 0x40
	TrackingActiveAlways             TrackingAreaOptions = 0x80
	TrackingAssumeInside             TrackingAreaOptions = 0x100
	TrackingInVisibleRect            TrackingAreaOptions = 0x200
	TrackingEnabledDuringMouseDrag   TrackingAreaOptions = 0x400
)

type TrackingAreaOptions int

type TrackingArea struct {
	native C.NSTrackingAreaPtr
}

func TrackingAreaInitWithRectOptionsOwnerUserInfo(x, y, width, height float64, options TrackingAreaOptions, owner *View, userInfo cf.Dictionary) *TrackingArea {
	return &TrackingArea{
		native: C.nsTrackingAreaInitWithRectOptionsOwnerUserInfo(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), C.NSTrackingAreaOptions(options), owner.native, C.CFDictionaryRef(userInfo)),
	}
}
