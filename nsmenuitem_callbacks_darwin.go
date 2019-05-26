package macos

import "C"

//export menuItemValidateCallback
func menuItemValidateCallback(tag int) bool {
	if validator, ok := nsMenuItemValidators[tag]; ok && validator != nil {
		return validator()
	}
	return true
}

//export menuItemHandleCallback
func menuItemHandleCallback(tag int) {
	if handler, ok := nsMenuItemHandlers[tag]; ok && handler != nil {
		handler()
	}
}
