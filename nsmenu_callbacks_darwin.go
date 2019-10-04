package macos

// typedef void *NSMenuPtr;
import "C"

//export updateMenuCallback
func updateMenuCallback(menu C.NSMenuPtr) {
	if updater, ok := nsMenuUpdaters[menu]; ok && updater != nil {
		updater(&NSMenu{native: menu})
	}
}
