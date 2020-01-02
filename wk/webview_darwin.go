// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package wk

import (
	"github.com/richardwilkes/macos/ns"
)

/*
#import <WebKit/WebKit.h>

typedef void *NSViewPtr;
typedef void *WKWebViewPtr;
typedef void *WKWebViewConfigurationPtr;
typedef void *WKNavigationPtr;
typedef void *WKNavigationActionPtr;
typedef void *WKNavigationResponsePtr;
typedef void *NSURLAuthenticationChallengePtr;
typedef void *NSURLRequestPtr;
typedef void *NSURLCredentialPtr;

void webViewDidCommitNavigation(WKWebViewPtr webView, WKNavigationPtr nav);
void webViewDidStartProvisionalNavigation(WKWebViewPtr webView, WKNavigationPtr nav);
void webViewDidReceiveServerRedirectForProvisionNavigation(WKWebViewPtr webView, WKNavigationPtr nav);
NSURLCredentialPtr *webViewDidReceiveAuthenticationChallenge(WKWebViewPtr webView, NSURLAuthenticationChallenge *challenge, NSURLSessionAuthChallengeDisposition *disposition);
void webViewDidFailNavigationWithError(WKWebViewPtr webView, WKNavigationPtr nav, CFStringRef msg);
void webViewDidFailProvisionalNavigationWithError(WKWebViewPtr webView, WKNavigationPtr nav, CFStringRef msg);
void webViewDidFinishNavigation(WKWebViewPtr webView, WKNavigationPtr nav);
void webViewWebContentProcessDidTerminate(WKWebViewPtr webView);
WKNavigationActionPolicy webViewDecidePolicyForNavigationAction(WKWebViewPtr webView, WKNavigationPtr nav);
WKNavigationResponsePolicy webViewDecidePolicyForNavigationResponse(WKWebViewPtr webView, WKNavigationPtr nav);

@interface WebViewNavDelegate : NSObject<WKNavigationDelegate>
@end

@implementation WebViewNavDelegate
- (void)webView:(WKWebView *)webView didCommitNavigation:(WKNavigation *)navigation {
	webViewDidCommitNavigation(webView, navigation);
}

- (void)webView:(WKWebView *)webView didStartProvisionalNavigation:(WKNavigation *)navigation {
	webViewDidStartProvisionalNavigation(webView, navigation);
}

- (void)webView:(WKWebView *)webView didReceiveServerRedirectForProvisionalNavigation:(WKNavigation *)navigation {
	webViewDidReceiveServerRedirectForProvisionNavigation(webView, navigation);
}

- (void)webView:(WKWebView *)webView didReceiveAuthenticationChallenge:(NSURLAuthenticationChallenge *)challenge completionHandler:(void (^)(NSURLSessionAuthChallengeDisposition disposition, NSURLCredential *credential))completionHandler {
	NSURLSessionAuthChallengeDisposition disposition;
	NSURLCredentialPtr credential = webViewDidReceiveAuthenticationChallenge(webView, challenge, &disposition);
	completionHandler(disposition, (NSURLCredential *)credential);
}

- (void)webView:(WKWebView *)webView didFailNavigation:(WKNavigation *)navigation withError:(NSError *)error {
	webViewDidFailNavigationWithError(webView, navigation, (CFStringRef)[error localizedDescription]);
}

- (void)webView:(WKWebView *)webView didFailProvisionalNavigation:(WKNavigation *)navigation withError:(NSError *)error {
	webViewDidFailProvisionalNavigationWithError(webView, navigation, (CFStringRef)[error localizedDescription]);
}

- (void)webView:(WKWebView *)webView didFinishNavigation:(WKNavigation *)navigation {
	webViewDidFinishNavigation(webView, navigation);
}

- (void)webViewWebContentProcessDidTerminate:(WKWebView *)webView {
	webViewWebContentProcessDidTerminate(webView);
}

- (void)webView:(WKWebView *)webView decidePolicyForNavigationAction:(WKNavigationAction *)navigationAction decisionHandler:(void (^)(WKNavigationActionPolicy))decisionHandler {
	decisionHandler(webViewDecidePolicyForNavigationAction(webView, navigationAction));
}

- (void)webView:(WKWebView *)webView decidePolicyForNavigationResponse:(WKNavigationResponse *)navigationResponse decisionHandler:(void (^)(WKNavigationResponsePolicy))decisionHandler {
	decisionHandler(webViewDecidePolicyForNavigationResponse(webView, navigationResponse));
}
@end

static WebViewNavDelegate *webViewNavDelegate = nil;

WKWebViewPtr webViewInitWithFrameConfiguration(CGFloat x, CGFloat y, CGFloat width, CGFloat height, WKWebViewConfigurationPtr config) {
	WKWebView *view = [[WKWebView alloc] initWithFrame:NSMakeRect(x, y, width, height) configuration:(WKWebViewConfiguration *)config];
	if (!webViewNavDelegate) {
		webViewNavDelegate = [WebViewNavDelegate new];
	}
	[view setNavigationDelegate:webViewNavDelegate];
	return (WKWebViewPtr)view;
}

void webViewLoadRequest(WKWebViewPtr view, NSURLRequestPtr request) {
	[(WKWebView *)view loadRequest:(NSURLRequest *)request];
}

void webViewRelease(WKWebViewPtr view) {
	[(WKWebView *)view release];
}
*/
import "C"

const (
	NavigationActionPolicyCancel NavigationActionPolicy = iota
	NavigationActionPolicyAllow
)

const (
	NavigationResponsePolicyCancel NavigationResponsePolicy = iota
	NavigationResponsePolicyAllow
)

var webViewDelegateMap = make(map[C.WKWebViewPtr]*webViewDelegateInfo)

type (
	NavigationActionPolicy   int
	NavigationResponsePolicy int
	NavigationPtr            = C.WKNavigationPtr
)

type webViewDelegateInfo struct {
	webView  *WebView
	delegate NavigationDelegate
}

type WebView struct {
	ns.View
}

type Navigation struct {
	native C.WKNavigationPtr
}

func (a *Navigation) Native() NavigationPtr {
	return a.native
}

type NavigationAction struct {
	native C.WKNavigationActionPtr
}

type NavigationResponse struct {
	native C.WKNavigationResponsePtr
}

type NavigationDelegate interface {
	WebViewDidCommitNavigation(webView *WebView, nav *Navigation)
	WebViewDidStartProvisionalNavigation(webView *WebView, nav *Navigation)
	WebViewDidReceiveServerRedirectForProvisionNavigation(webView *WebView, nav *Navigation)
	WebViewDidReceiveAuthenticationChallenge(webView *WebView, challenge *ns.URLAuthenticationChallenge) (disposition ns.URLSessionAuthChallengeDisposition, credential *ns.URLCredential)
	WebViewDidFailNavigationWithError(webView *WebView, nav *Navigation, errorMsg string)
	WebViewDidFailProvisionalNavigationWithError(webView *WebView, nav *Navigation, errorMsg string)
	WebViewDidFinishNavigation(webView *WebView, nav *Navigation)
	WebViewWebContentProcessDidTerminate(webView *WebView)
	WebViewDecidePolicyForNavigationAction(webView *WebView, action *NavigationAction) NavigationActionPolicy
	WebViewDecidePolicyForNavigationResponse(webView *WebView, response *NavigationResponse) NavigationResponsePolicy
}

func WebViewInitWithFrameConfiguration(x, y, width, height float64, config *WebViewConfiguration, delegate NavigationDelegate) *WebView {
	webView := &WebView{View: *ns.NewViewFromNative(ns.ViewNative(C.webViewInitWithFrameConfiguration(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), config.native)))}
	webViewDelegateMap[C.WKWebViewPtr(webView.Native())] = &webViewDelegateInfo{
		webView:  webView,
		delegate: delegate,
	}
	return webView
}

func (v *WebView) LoadRequest(request *ns.URLRequest) {
	C.webViewLoadRequest(C.WKWebViewPtr(v.Native()), C.NSURLRequestPtr(request.Native()))
}

func (v *WebView) Release() {
	C.webViewRelease(C.WKWebViewPtr(v.Native()))
	delete(webViewDelegateMap, C.WKWebViewPtr(v.Native()))
}
