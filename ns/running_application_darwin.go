package ns

/*
#import <Cocoa/Cocoa.h>

typedef void *NSRunningApplicationPtr;

NSRunningApplicationPtr nsRunningApplicationCurrent() {
	return (NSRunningApplicationPtr)[NSRunningApplication currentApplication];
}

void nsRunningApplicationActivateWithOptions(NSRunningApplicationPtr app, NSApplicationActivationOptions options) {
	[(NSRunningApplication *)app activateWithOptions:options];
}

bool nsRunningApplicationHide(NSRunningApplicationPtr app) {
	return [(NSRunningApplication *)app hide] != 0;
}
*/
import "C"

type RunningApplication struct {
	native C.NSRunningApplicationPtr
}

func RunningApplicationCurrent() *RunningApplication {
	return &RunningApplication{native: C.nsRunningApplicationCurrent()}
}

func (app *RunningApplication) ActivateWithOptions(options ApplicationActivationOptions) {
	C.nsRunningApplicationActivateWithOptions(app.native, C.NSApplicationActivationOptions(options))
}

func (app *RunningApplication) Hide() bool {
	return bool(C.nsRunningApplicationHide(app.native))
}
