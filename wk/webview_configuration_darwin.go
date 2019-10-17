package wk

/*
#import <WebKit/WebKit.h>

typedef void *WKWebViewConfigurationPtr;

WKWebViewConfigurationPtr wkWebViewConfigurationNew() {
	return (WKWebViewConfigurationPtr)[WKWebViewConfiguration new];
}
*/
import "C"

type WebViewConfiguration struct {
	native C.WKWebViewConfigurationPtr
}

func NewWebViewConfiguration() *WebViewConfiguration {
	return &WebViewConfiguration{native: C.wkWebViewConfigurationNew()}
}
