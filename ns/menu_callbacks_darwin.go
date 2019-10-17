package ns

// typedef void *NSMenuPtr;
import "C"

//export updateMenuCallback
func updateMenuCallback(menu C.NSMenuPtr) {
	if updater, ok := menuUpdaters[menu]; ok && updater != nil {
		updater(&Menu{native: menu})
	}
}
