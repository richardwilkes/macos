package wk

import (
	"github.com/richardwilkes/macos/cf"
	"github.com/richardwilkes/macos/ns"
)

/*
#import <WebKit/WebKit.h>

typedef void *WKWebViewPtr;
typedef void *WKNavigationPtr;
typedef void *WKNavigationActionPtr;
typedef void *WKNavigationResponsePtr;
typedef void *NSURLAuthenticationChallengePtr;
typedef void *NSURLCredentialPtr;
*/
import "C"

//export webViewDidCommitNavigation
func webViewDidCommitNavigation(webView C.WKWebViewPtr, nav C.WKNavigationPtr) {
	if d, ok := webViewDelegateMap[webView]; ok {
		d.delegate.WebViewDidCommitNavigation(d.webView, &Navigation{native: nav})
	}
}

//export webViewDidStartProvisionalNavigation
func webViewDidStartProvisionalNavigation(webView C.WKWebViewPtr, nav C.WKNavigationPtr) {
	if d, ok := webViewDelegateMap[webView]; ok {
		d.delegate.WebViewDidStartProvisionalNavigation(d.webView, &Navigation{native: nav})
	}
}

//export webViewDidReceiveServerRedirectForProvisionNavigation
func webViewDidReceiveServerRedirectForProvisionNavigation(webView C.WKWebViewPtr, nav C.WKNavigationPtr) {
	if d, ok := webViewDelegateMap[webView]; ok {
		d.delegate.WebViewDidReceiveServerRedirectForProvisionNavigation(d.webView, &Navigation{native: nav})
	}
}

//export webViewDidReceiveAuthenticationChallenge
func webViewDidReceiveAuthenticationChallenge(webView C.WKWebViewPtr, challenge C.NSURLAuthenticationChallengePtr, disposition *C.NSURLSessionAuthChallengeDisposition) C.NSURLCredentialPtr {
	*disposition = C.NSURLSessionAuthChallengePerformDefaultHandling
	if d, ok := webViewDelegateMap[webView]; ok {
		disp, credential := d.delegate.WebViewDidReceiveAuthenticationChallenge(d.webView, ns.NewURLAuthenticationChallengeFromNative(ns.URLAuthenticationChallengeNative(challenge)))
		*disposition = C.NSURLSessionAuthChallengeDisposition(disp)
		return C.NSURLCredentialPtr(credential.Native())
	}
	return nil
}

//export webViewDidFailNavigationWithError
func webViewDidFailNavigationWithError(webView C.WKWebViewPtr, nav C.WKNavigationPtr, msg C.CFStringRef) {
	if d, ok := webViewDelegateMap[webView]; ok {
		d.delegate.WebViewDidFailNavigationWithError(d.webView, &Navigation{native: nav}, cf.String(msg).String())
	}
}

//export webViewDidFailProvisionalNavigationWithError
func webViewDidFailProvisionalNavigationWithError(webView C.WKWebViewPtr, nav C.WKNavigationPtr, msg C.CFStringRef) {
	if d, ok := webViewDelegateMap[webView]; ok {
		d.delegate.WebViewDidFailProvisionalNavigationWithError(d.webView, &Navigation{native: nav}, cf.String(msg).String())
	}
}

//export webViewDidFinishNavigation
func webViewDidFinishNavigation(webView C.WKWebViewPtr, nav C.WKNavigationPtr) {
	if d, ok := webViewDelegateMap[webView]; ok {
		d.delegate.WebViewDidFinishNavigation(d.webView, &Navigation{native: nav})
	}
}

//export webViewWebContentProcessDidTerminate
func webViewWebContentProcessDidTerminate(webView C.WKWebViewPtr) {
	if d, ok := webViewDelegateMap[webView]; ok {
		d.delegate.WebViewWebContentProcessDidTerminate(d.webView)
	}
}

//export webViewDecidePolicyForNavigationAction
func webViewDecidePolicyForNavigationAction(webView C.WKWebViewPtr, action C.WKNavigationActionPtr) C.WKNavigationActionPolicy {
	if d, ok := webViewDelegateMap[webView]; ok {
		return C.WKNavigationActionPolicy(d.delegate.WebViewDecidePolicyForNavigationAction(d.webView, &NavigationAction{native: action}))
	}
	return C.WKNavigationActionPolicy(NavigationActionPolicyCancel)
}

//export webViewDecidePolicyForNavigationResponse
func webViewDecidePolicyForNavigationResponse(webView C.WKWebViewPtr, response C.WKNavigationResponsePtr) C.WKNavigationResponsePolicy {
	if d, ok := webViewDelegateMap[webView]; ok {
		return C.WKNavigationResponsePolicy(d.delegate.WebViewDecidePolicyForNavigationResponse(d.webView, &NavigationResponse{native: response}))
	}
	return C.WKNavigationResponsePolicy(NavigationResponsePolicyCancel)
}
