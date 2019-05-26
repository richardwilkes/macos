package macos

/*
#import <Foundation/Foundation.h>

typedef void *NSURLRequestPtr;
typedef void *NSURLPtr;

NSURLRequestPtr nsURLRequestWithURL(NSURLPtr url) {
	return (NSURLRequestPtr)[NSURLRequest requestWithURL:(NSURL *)url];
}
*/
import "C"

type NSURLRequest struct {
	native C.NSURLRequestPtr
}

func NSURLRequestWithURL(url *NSURL) *NSURLRequest {
	return &NSURLRequest{native: C.nsURLRequestWithURL(url.native)}
}
