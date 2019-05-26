package macos

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
func wkWebViewDidReceiveAuthenticationChallenge(webView C.WKWebViewPtr, challenge C.NSURLAuthenticationChallengePtr, disposition *C.NSURLSessionAuthChallengeDisposition) C.NSURLCredentialPtr {
	*disposition = C.NSURLSessionAuthChallengePerformDefaultHandling
	if d, ok := wkWebViewDelegateMap[webView]; ok {
		disp, credential := d.delegate.WebViewDidReceiveAuthenticationChallenge(d.webView, &NSURLAuthenticationChallenge{native: challenge})
		*disposition = C.NSURLSessionAuthChallengeDisposition(disp)
		return credential.native
	}
	return nil
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
