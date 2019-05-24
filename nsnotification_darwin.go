package macos

/*
#import <Cocoa/Cocoa.h>

typedef void *NSNotificationPtr;
*/
import "C"

type NSNotification struct {
	native C.NSNotificationPtr
}
