package ns

import "C"

//export menuItemValidateCallback
func menuItemValidateCallback(tag int) bool {
	if validator, ok := menuItemValidators[tag]; ok && validator != nil {
		return validator(tag)
	}
	return true
}

//export menuItemHandleCallback
func menuItemHandleCallback(tag int) {
	if handler, ok := menuItemHandlers[tag]; ok && handler != nil {
		handler(tag)
	}
}
