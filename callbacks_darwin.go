package macos

import "C"

/*
#import <CoreGraphics/CoreGraphics.h>
#import <WebKit/WebKit.h>
#import "authentication_challenge_response.h"

typedef void *NSApplicationPtr;
typedef void *NSNotificationPtr;
typedef void *NSWindowPtr;
typedef void *NSViewPtr;
typedef void *NSMenuItemPtr;
typedef void *WKWebViewPtr;
typedef void *WKNavigationPtr;
typedef void *WKNavigationActionPtr;
typedef void *WKNavigationResponsePtr;
typedef void *NSURLAuthenticationChallengePtr;
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

//export windowDidResizeCallback
func windowDidResizeCallback(w C.NSWindowPtr) {
	if d, ok := nsWindowDelegateMap[w]; ok {
		d.WindowDidResize(&NSWindow{native: w})
	}
}

//export windowDidBecomeKeyCallback
func windowDidBecomeKeyCallback(w C.NSWindowPtr) {
	if d, ok := nsWindowDelegateMap[w]; ok {
		d.WindowDidBecomeKey(&NSWindow{native: w})
	}
}

//export windowDidResignKeyCallback
func windowDidResignKeyCallback(w C.NSWindowPtr) {
	if d, ok := nsWindowDelegateMap[w]; ok {
		d.WindowDidResignKey(&NSWindow{native: w})
	}
}

//export windowShouldCloseCallback
func windowShouldCloseCallback(w C.NSWindowPtr) bool {
	if d, ok := nsWindowDelegateMap[w]; ok {
		return d.WindowShouldClose(&NSWindow{native: w})
	}
	return true
}

//export windowWillCloseCallback
func windowWillCloseCallback(w C.NSWindowPtr) {
	if d, ok := nsWindowDelegateMap[w]; ok {
		d.WindowWillClose(&NSWindow{native: w})
	}
}

//export viewDrawCallback
func viewDrawCallback(view C.NSViewPtr, gc CGContext, x, y, width, height C.CGFloat, inLiveResize bool) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewDraw(&NSView{native: view}, gc, float64(x), float64(y), float64(width), float64(height), inLiveResize)
	}
}

//export viewMouseDownCallback
func viewMouseDownCallback(view C.NSViewPtr, x, y C.CGFloat, button, clickCount, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseDownEvent(&NSView{native: view}, float64(x), float64(y), button, clickCount, mod)
	}
}

//export viewMouseDragCallback
func viewMouseDragCallback(view C.NSViewPtr, x, y C.CGFloat, button, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseDragEvent(&NSView{native: view}, float64(x), float64(y), button, mod)
	}
}

//export viewMouseUpCallback
func viewMouseUpCallback(view C.NSViewPtr, x, y C.CGFloat, button, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseUpEvent(&NSView{native: view}, float64(x), float64(y), button, mod)
	}
}

//export viewMouseEnterCallback
func viewMouseEnterCallback(view C.NSViewPtr, x, y C.CGFloat, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseEnterEvent(&NSView{native: view}, float64(x), float64(y), mod)
	}
}

//export viewMouseMoveCallback
func viewMouseMoveCallback(view C.NSViewPtr, x, y C.CGFloat, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseMoveEvent(&NSView{native: view}, float64(x), float64(y), mod)
	}
}

//export viewMouseExitCallback
func viewMouseExitCallback(view C.NSViewPtr) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseExitEvent(&NSView{native: view})
	}
}

//export viewMouseWheelCallback
func viewMouseWheelCallback(view C.NSViewPtr, x, y, dx, dy C.CGFloat, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewMouseWheelEvent(&NSView{native: view}, float64(x), float64(y), float64(dx), float64(dy), mod)
	}
}

//export viewCursorUpdateCallback
func viewCursorUpdateCallback(view C.NSViewPtr, x, y C.CGFloat, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewCursorUpdateEvent(&NSView{native: view}, float64(x), float64(y), mod)
	}
}

//export viewKeyDownCallback
func viewKeyDownCallback(view C.NSViewPtr, keyCode int, ch CFString, mod int, repeat bool) {
	if d, ok := nsViewDelegateMap[view]; ok {
		var r rune
		if ch != 0 {
			if runes := []rune(ch.String()); len(runes) > 0 {
				r = runes[0]
			}
		}
		d.ViewKeyDownEvent(&NSView{native: view}, keyCode, r, mod, repeat)
	}
}

//export viewKeyUpCallback
func viewKeyUpCallback(view C.NSViewPtr, keyCode, mod int) {
	if d, ok := nsViewDelegateMap[view]; ok {
		d.ViewKeyUpEvent(&NSView{native: view}, keyCode, mod)
	}
}

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

//export patternDrawCallback
func patternDrawCallback(id *int32, gc C.CGContextRef) {
	patternLock.Lock()
	callback, ok := patternCallbackMap[*id]
	patternLock.Unlock()
	if ok {
		callback.PatternDraw(gc)
	}
}

//export patternReleaseCallback
func patternReleaseCallback(id *int32) {
	patternLock.Lock()
	callback, ok := patternCallbackMap[*id]
	delete(patternCallbackMap, *id)
	patternLock.Unlock()
	if ok {
		callback.PatternRelease()
	}
}

//export wkWebViewDidCommitNavigation
func wkWebViewDidCommitNavigation(webView C.WKWebViewPtr, nav C.WKNavigationPtr) {
	if d, ok := wkWebViewDelegateMap[webView]; ok {
		d.delegate.WebViewDidCommitNavigation(d.webView, &WKNavigation{native: nav})
	}
}

//export wkWebViewDidStartProvisionalNavigation
func wkWebViewDidStartProvisionalNavigation(webView C.WKWebViewPtr, nav C.WKNavigationPtr) {
	if d, ok := wkWebViewDelegateMap[webView]; ok {
		d.delegate.WebViewDidStartProvisionalNavigation(d.webView, &WKNavigation{native: nav})
	}
}

//export wkWebViewDidReceiveServerRedirectForProvisionNavigation
func wkWebViewDidReceiveServerRedirectForProvisionNavigation(webView C.WKWebViewPtr, nav C.WKNavigationPtr) {
	if d, ok := wkWebViewDelegateMap[webView]; ok {
		d.delegate.WebViewDidReceiveServerRedirectForProvisionNavigation(d.webView, &WKNavigation{native: nav})
	}
}

//export wkWebViewDidReceiveAuthenticationChallenge
func wkWebViewDidReceiveAuthenticationChallenge(webView C.WKWebViewPtr, challenge C.NSURLAuthenticationChallengePtr) *authenticationChallengeResponse {
	if d, ok := wkWebViewDelegateMap[webView]; ok {
		disposition, credential := d.delegate.WebViewDidReceiveAuthenticationChallenge(d.webView, &NSURLAuthenticationChallenge{native: challenge})
		return &authenticationChallengeResponse{
			disposition: C.NSURLSessionAuthChallengeDisposition(disposition),
			credential:  credential.native,
		}
	}
	return &authenticationChallengeResponse{disposition: C.NSURLSessionAuthChallengeDisposition(NSURLSessionAuthChallengePerformDefaultHandling)}
}

//export wkWebViewDidFailNavigationWithError
func wkWebViewDidFailNavigationWithError(webView C.WKWebViewPtr, nav C.WKNavigationPtr, msg CFString) {
	if d, ok := wkWebViewDelegateMap[webView]; ok {
		d.delegate.WebViewDidFailNavigationWithError(d.webView, &WKNavigation{native: nav}, msg.String())
	}
}

//export wkWebViewDidFailProvisionalNavigationWithError
func wkWebViewDidFailProvisionalNavigationWithError(webView C.WKWebViewPtr, nav C.WKNavigationPtr, msg CFString) {
	if d, ok := wkWebViewDelegateMap[webView]; ok {
		d.delegate.WebViewDidFailProvisionalNavigationWithError(d.webView, &WKNavigation{native: nav}, msg.String())
	}
}

//export wkWebViewDidFinishNavigation
func wkWebViewDidFinishNavigation(webView C.WKWebViewPtr, nav C.WKNavigationPtr) {
	if d, ok := wkWebViewDelegateMap[webView]; ok {
		d.delegate.WebViewDidFinishNavigation(d.webView, &WKNavigation{native: nav})
	}
}

//export wkWebViewWebContentProcessDidTerminate
func wkWebViewWebContentProcessDidTerminate(webView C.WKWebViewPtr) {
	if d, ok := wkWebViewDelegateMap[webView]; ok {
		d.delegate.WebViewWebContentProcessDidTerminate(d.webView)
	}
}

//export wkWebViewDecidePolicyForNavigationAction
func wkWebViewDecidePolicyForNavigationAction(webView C.WKWebViewPtr, action C.WKNavigationActionPtr) C.WKNavigationActionPolicy {
	if d, ok := wkWebViewDelegateMap[webView]; ok {
		return C.WKNavigationActionPolicy(d.delegate.WebViewDecidePolicyForNavigationAction(d.webView, &WKNavigationAction{native: action}))
	}
	return C.WKNavigationActionPolicy(WKNavigationActionPolicyCancel)
}

//export wkWebViewDecidePolicyForNavigationResponse
func wkWebViewDecidePolicyForNavigationResponse(webView C.WKWebViewPtr, response C.WKNavigationResponsePtr) C.WKNavigationResponsePolicy {
	if d, ok := wkWebViewDelegateMap[webView]; ok {
		return C.WKNavigationResponsePolicy(d.delegate.WebViewDecidePolicyForNavigationResponse(d.webView, &WKNavigationResponse{native: response}))
	}
	return C.WKNavigationResponsePolicy(WKNavigationResponsePolicyCancel)
}
