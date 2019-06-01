package macos

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
	NSTrackingMouseEnteredAndExited    NSTrackingAreaOptions = 0x01
	NSTrackingMouseMoved               NSTrackingAreaOptions = 0x02
	NSTrackingCursorUpdate             NSTrackingAreaOptions = 0x04
	NSTrackingActiveWhenFirstResponder NSTrackingAreaOptions = 0x10
	NSTrackingActiveInKeyWindow        NSTrackingAreaOptions = 0x20
	NSTrackingActiveInActiveApp        NSTrackingAreaOptions = 0x40
	NSTrackingActiveAlways             NSTrackingAreaOptions = 0x80
	NSTrackingAssumeInside             NSTrackingAreaOptions = 0x100
	NSTrackingInVisibleRect            NSTrackingAreaOptions = 0x200
	NSTrackingEnabledDuringMouseDrag   NSTrackingAreaOptions = 0x400
)

type NSTrackingAreaOptions int

type NSTrackingArea struct {
	native C.NSTrackingAreaPtr
}

func NSTrackingAreaInitWithRectOptionsOwnerUserInfo(x, y, width, height float64, options NSTrackingAreaOptions, owner *NSView, userInfo CFDictionary) *NSTrackingArea {
	return &NSTrackingArea{
		native: C.nsTrackingAreaInitWithRectOptionsOwnerUserInfo(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), C.NSTrackingAreaOptions(options), owner.native, userInfo),
	}
}
