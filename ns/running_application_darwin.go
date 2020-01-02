// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

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
