package macos

/*
#import <Foundation/Foundation.h>

typedef void *NSURLPtr;

NSURLPtr nsURLWithString(CFStringRef url) {
	return (NSURLPtr)[NSURL URLWithString:(NSString *)url];
}
*/
import "C"

type NSURL struct {
	native C.NSURLPtr
}

func NSURLWithString(url string) *NSURL {
	str := CFStringCreateWithString(url)
	defer str.Release()
	return &NSURL{native: C.nsURLWithString(str)}
}
