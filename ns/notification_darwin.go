package ns

/*
#import <Cocoa/Cocoa.h>

typedef void *NSNotificationPtr;
*/
import "C"

type Notification struct {
	native C.NSNotificationPtr
}
