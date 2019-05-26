package macos

/*
typedef void *NSApplicationPtr;
typedef void *NSNotificationPtr;
*/
import "C"

//export applicationWillFinishLaunchingCallback
func applicationWillFinishLaunchingCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationWillFinishLaunching(&NSNotification{native: aNotification})
	}
}

//export applicationDidFinishLaunchingCallback
func applicationDidFinishLaunchingCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationDidFinishLaunching(&NSNotification{native: aNotification})
	}
}

//export applicationShouldTerminateCallback
func applicationShouldTerminateCallback(sender C.NSApplicationPtr) NSApplicationTerminateReply {
	if currentAppDelegate != nil {
		return currentAppDelegate.ApplicationShouldTerminate(&NSApplication{native: sender})
	}
	return NSTerminateNow
}

//export applicationShouldTerminateAfterLastWindowClosedCallback
func applicationShouldTerminateAfterLastWindowClosedCallback(theApplication C.NSApplicationPtr) bool {
	if currentAppDelegate != nil {
		return currentAppDelegate.ApplicationShouldTerminateAfterLastWindowClosed(&NSApplication{native: theApplication})
	}
	return true
}

//export applicationWillTerminateCallback
func applicationWillTerminateCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationWillTerminate(&NSNotification{native: aNotification})
	}
}

//export applicationWillBecomeActiveCallback
func applicationWillBecomeActiveCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationWillBecomeActive(&NSNotification{native: aNotification})
	}
}

//export applicationDidBecomeActiveCallback
func applicationDidBecomeActiveCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationDidBecomeActive(&NSNotification{native: aNotification})
	}
}

//export applicationWillResignActiveCallback
func applicationWillResignActiveCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationWillResignActive(&NSNotification{native: aNotification})
	}
}

//export applicationDidResignActiveCallback
func applicationDidResignActiveCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationDidResignActive(&NSNotification{native: aNotification})
	}
}

//export themeChangedCallback
func themeChangedCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ThemeChanged(&NSNotification{native: aNotification})
	}
}
