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
