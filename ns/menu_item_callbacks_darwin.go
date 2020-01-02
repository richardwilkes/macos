// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ns

// typedef void *NSMenuItemPtr;
import "C"

//export menuItemValidateCallback
func menuItemValidateCallback(item C.NSMenuItemPtr) bool {
	if validator, ok := menuItemValidators[item]; ok && validator != nil {
		return validator(&MenuItem{native: item})
	}
	return true
}

//export menuItemHandleCallback
func menuItemHandleCallback(item C.NSMenuItemPtr) {
	if handler, ok := menuItemHandlers[item]; ok && handler != nil {
		handler(&MenuItem{native: item})
	}
}
