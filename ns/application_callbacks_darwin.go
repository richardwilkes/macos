package ns

/*
#import <CoreFoundation/CoreFoundation.h>

typedef void *NSApplicationPtr;
typedef void *NSNotificationPtr;
typedef void *NSURLPtr;
*/
import "C"
import "github.com/richardwilkes/macos/cf"

//export applicationWillFinishLaunchingCallback
func applicationWillFinishLaunchingCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationWillFinishLaunching(&Notification{native: aNotification})
	}
}

//export applicationDidFinishLaunchingCallback
func applicationDidFinishLaunchingCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationDidFinishLaunching(&Notification{native: aNotification})
	}
}

//export applicationShouldTerminateCallback
func applicationShouldTerminateCallback(sender C.NSApplicationPtr) ApplicationTerminateReply {
	if currentAppDelegate != nil {
		return currentAppDelegate.ApplicationShouldTerminate(&Application{native: sender})
	}
	return TerminateNow
}

//export applicationShouldTerminateAfterLastWindowClosedCallback
func applicationShouldTerminateAfterLastWindowClosedCallback(theApplication C.NSApplicationPtr) bool {
	if currentAppDelegate != nil {
		return currentAppDelegate.ApplicationShouldTerminateAfterLastWindowClosed(&Application{native: theApplication})
	}
	return true
}

//export applicationWillTerminateCallback
func applicationWillTerminateCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationWillTerminate(&Notification{native: aNotification})
	}
}

//export applicationWillBecomeActiveCallback
func applicationWillBecomeActiveCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationWillBecomeActive(&Notification{native: aNotification})
	}
}

//export applicationDidBecomeActiveCallback
func applicationDidBecomeActiveCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationDidBecomeActive(&Notification{native: aNotification})
	}
}

//export applicationWillResignActiveCallback
func applicationWillResignActiveCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationWillResignActive(&Notification{native: aNotification})
	}
}

//export applicationDidResignActiveCallback
func applicationDidResignActiveCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ApplicationDidResignActive(&Notification{native: aNotification})
	}
}

//export applicationOpenURLsCallback
func applicationOpenURLsCallback(theApplication C.NSApplicationPtr, cUrls C.CFArrayRef) {
	if currentAppDelegate != nil {
		urls := cf.Array(cUrls)
		count := urls.GetCount()
		urlList := make([]string, 0, count)
		for i := 0; i < count; i++ {
			one := URL{native: C.NSURLPtr(urls.GetValueAtIndex(i))}
			urlList = append(urlList, one.ResolveFilePath())
		}
		currentAppDelegate.ApplicationOpenURLs(&Application{native: theApplication}, urlList)
	}
}

//export themeChangedCallback
func themeChangedCallback(aNotification C.NSNotificationPtr) {
	if currentAppDelegate != nil {
		currentAppDelegate.ThemeChanged(&Notification{native: aNotification})
	}
}
