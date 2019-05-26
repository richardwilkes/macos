package macos

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

void wkWebViewDidCommitNavigation(WKWebViewPtr webView, WKNavigationPtr nav);
void wkWebViewDidStartProvisionalNavigation(WKWebViewPtr webView, WKNavigationPtr nav);
void wkWebViewDidReceiveServerRedirectForProvisionNavigation(WKWebViewPtr webView, WKNavigationPtr nav);
NSURLCredentialPtr *wkWebViewDidReceiveAuthenticationChallenge(WKWebViewPtr webView, NSURLAuthenticationChallenge *challenge, NSURLSessionAuthChallengeDisposition *disposition);
void wkWebViewDidFailNavigationWithError(WKWebViewPtr webView, WKNavigationPtr nav, CFStringRef msg);
void wkWebViewDidFailProvisionalNavigationWithError(WKWebViewPtr webView, WKNavigationPtr nav, CFStringRef msg);
void wkWebViewDidFinishNavigation(WKWebViewPtr webView, WKNavigationPtr nav);
void wkWebViewWebContentProcessDidTerminate(WKWebViewPtr webView);
WKNavigationActionPolicy wkWebViewDecidePolicyForNavigationAction(WKWebViewPtr webView, WKNavigationPtr nav);
WKNavigationResponsePolicy wkWebViewDecidePolicyForNavigationResponse(WKWebViewPtr webView, WKNavigationPtr nav);

@interface WebViewNavDelegate : NSObject<WKNavigationDelegate>
@end

@implementation WebViewNavDelegate
- (void)webView:(WKWebView *)webView didCommitNavigation:(WKNavigation *)navigation {
	wkWebViewDidCommitNavigation(webView, navigation);
}

- (void)webView:(WKWebView *)webView didStartProvisionalNavigation:(WKNavigation *)navigation {
	wkWebViewDidStartProvisionalNavigation(webView, navigation);
}

- (void)webView:(WKWebView *)webView didReceiveServerRedirectForProvisionalNavigation:(WKNavigation *)navigation {
	wkWebViewDidReceiveServerRedirectForProvisionNavigation(webView, navigation);
}

- (void)webView:(WKWebView *)webView didReceiveAuthenticationChallenge:(NSURLAuthenticationChallenge *)challenge completionHandler:(void (^)(NSURLSessionAuthChallengeDisposition disposition, NSURLCredential *credential))completionHandler {
	NSURLSessionAuthChallengeDisposition disposition;
	NSURLCredentialPtr credential = wkWebViewDidReceiveAuthenticationChallenge(webView, challenge, &disposition);
	completionHandler(disposition, (NSURLCredential *)credential);
}

- (void)webView:(WKWebView *)webView didFailNavigation:(WKNavigation *)navigation withError:(NSError *)error {
	wkWebViewDidFailNavigationWithError(webView, navigation, (CFStringRef)[error localizedDescription]);
}

- (void)webView:(WKWebView *)webView didFailProvisionalNavigation:(WKNavigation *)navigation withError:(NSError *)error {
	wkWebViewDidFailProvisionalNavigationWithError(webView, navigation, (CFStringRef)[error localizedDescription]);
}

- (void)webView:(WKWebView *)webView didFinishNavigation:(WKNavigation *)navigation {
	wkWebViewDidFinishNavigation(webView, navigation);
}

- (void)webViewWebContentProcessDidTerminate:(WKWebView *)webView {
	wkWebViewWebContentProcessDidTerminate(webView);
}

- (void)webView:(WKWebView *)webView decidePolicyForNavigationAction:(WKNavigationAction *)navigationAction decisionHandler:(void (^)(WKNavigationActionPolicy))decisionHandler {
	decisionHandler(wkWebViewDecidePolicyForNavigationAction(webView, navigationAction));
}

- (void)webView:(WKWebView *)webView decidePolicyForNavigationResponse:(WKNavigationResponse *)navigationResponse decisionHandler:(void (^)(WKNavigationResponsePolicy))decisionHandler {
	decisionHandler(wkWebViewDecidePolicyForNavigationResponse(webView, navigationResponse));
}
@end

static WebViewNavDelegate *webViewNavDelegate = nil;

WKWebViewPtr wkWebViewInitWithFrameConfiguration(CGFloat x, CGFloat y, CGFloat width, CGFloat height, WKWebViewConfigurationPtr config) {
	WKWebView *view = [[WKWebView alloc] initWithFrame:NSMakeRect(x, y, width, height) configuration:(WKWebViewConfiguration *)config];
	if (!webViewNavDelegate) {
		webViewNavDelegate = [WebViewNavDelegate new];
	}
	[view setNavigationDelegate:webViewNavDelegate];
	return (WKWebViewPtr)view;
}

void wkWebViewLoadRequest(WKWebViewPtr view, NSURLRequestPtr request) {
	[(WKWebView *)view loadRequest:(NSURLRequest *)request];
}

void wkWebViewRelease(WKWebViewPtr view) {
	[(WKWebView *)view release];
}
*/
import "C"

const (
	NSURLSessionAuthChallengeUseCredential NSURLSessionAuthChallengeDisposition = iota
	NSURLSessionAuthChallengePerformDefaultHandling
	NSURLSessionAuthChallengeCancelAuthenticationChallenge
	NSURLSessionAuthChallengeRejectProtectionSpace
)

const (
	WKNavigationActionPolicyCancel WKNavigationActionPolicy = iota
	WKNavigationActionPolicyAllow
)

const (
	WKNavigationResponsePolicyCancel WKNavigationResponsePolicy = iota
	WKNavigationResponsePolicyAllow
)

var wkWebViewDelegateMap = make(map[C.WKWebViewPtr]*wkWebViewDelegateInfo)

type (
	NSURLSessionAuthChallengeDisposition int
	WKNavigationActionPolicy             int
	WKNavigationResponsePolicy           int
	WKNavigationPtr                      = C.WKNavigationPtr
)

type wkWebViewDelegateInfo struct {
	webView  *WKWebView
	delegate WKNavigationDelegate
}

type WKWebView struct {
	NSView
}

type WKNavigation struct {
	native C.WKNavigationPtr
}

func (a *WKNavigation) Native() WKNavigationPtr {
	return a.native
}

type WKNavigationAction struct {
	native C.WKNavigationActionPtr
}

type WKNavigationResponse struct {
	native C.WKNavigationResponsePtr
}

type WKNavigationDelegate interface {
	WebViewDidCommitNavigation(webView *WKWebView, nav *WKNavigation)
	WebViewDidStartProvisionalNavigation(webView *WKWebView, nav *WKNavigation)
	WebViewDidReceiveServerRedirectForProvisionNavigation(webView *WKWebView, nav *WKNavigation)
	WebViewDidReceiveAuthenticationChallenge(webView *WKWebView, challenge *NSURLAuthenticationChallenge) (disposition NSURLSessionAuthChallengeDisposition, credential *NSURLCredential)
	WebViewDidFailNavigationWithError(webView *WKWebView, nav *WKNavigation, errorMsg string)
	WebViewDidFailProvisionalNavigationWithError(webView *WKWebView, nav *WKNavigation, errorMsg string)
	WebViewDidFinishNavigation(webView *WKWebView, nav *WKNavigation)
	WebViewWebContentProcessDidTerminate(webView *WKWebView)
	WebViewDecidePolicyForNavigationAction(webView *WKWebView, action *WKNavigationAction) WKNavigationActionPolicy
	WebViewDecidePolicyForNavigationResponse(webView *WKWebView, response *WKNavigationResponse) WKNavigationResponsePolicy
}

func WKWebViewInitWithFrameConfiguration(x, y, width, height float64, config *WKWebViewConfiguration, delegate WKNavigationDelegate) *WKWebView {
	webView := &WKWebView{NSView: NSView{native: C.NSViewPtr(C.wkWebViewInitWithFrameConfiguration(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), config.native))}}
	wkWebViewDelegateMap[C.WKWebViewPtr(webView.native)] = &wkWebViewDelegateInfo{
		webView:  webView,
		delegate: delegate,
	}
	return webView
}

func (v *WKWebView) LoadRequest(request *NSURLRequest) {
	C.wkWebViewLoadRequest(C.WKWebViewPtr(v.native), request.native)
}

func (v *WKWebView) Release() {
	C.wkWebViewRelease(C.WKWebViewPtr(v.native))
	delete(wkWebViewDelegateMap, C.WKWebViewPtr(v.native))
}
