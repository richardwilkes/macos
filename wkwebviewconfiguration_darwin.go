package macos

/*
#import <WebKit/WebKit.h>

typedef void *WKWebViewConfigurationPtr;

WKWebViewConfigurationPtr wkWebViewConfigurationNew() {
	return (WKWebViewConfigurationPtr)[WKWebViewConfiguration new];
}
*/
import "C"

type WKWebViewConfiguration struct {
	native C.WKWebViewConfigurationPtr
}

func NewWKWebViewConfiguration() *WKWebViewConfiguration {
	return &WKWebViewConfiguration{native: C.wkWebViewConfigurationNew()}
}
