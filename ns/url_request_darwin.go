package ns

/*
#import <Foundation/Foundation.h>

typedef void *NSURLRequestPtr;
typedef void *NSURLPtr;

NSURLRequestPtr nsURLRequestWithURL(NSURLPtr url) {
	return (NSURLRequestPtr)[NSURLRequest requestWithURL:(NSURL *)url];
}
*/
import "C"

type URLRequestNative = C.NSURLRequestPtr

type URLRequest struct {
	native URLRequestNative
}

func URLRequestWithURL(url *URL) *URLRequest {
	return &URLRequest{native: C.nsURLRequestWithURL(url.native)}
}

func (ur *URLRequest) Native() URLRequestNative {
	return ur.native
}
