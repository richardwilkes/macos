package macos

/*
#import <Foundation/Foundation.h>

typedef void *NSURLPtr;

NSURLPtr nsURLWithString(CFStringRef url) {
	return (NSURLPtr)[NSURL URLWithString:(NSString *)url];
}

CFStringRef nsURLFilePathURL(NSURLPtr u) {
	NSURL *url = [(NSURL *)u filePathURL];
	if (url) {
		return (CFStringRef)[url absoluteString];
	}
	return (CFStringRef)@"";
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

func (u *NSURL) ResolveFilePath() string {
	return C.nsURLFilePathURL(u.native).String()
}
