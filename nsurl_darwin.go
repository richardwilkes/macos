package macos

/*
#import <Foundation/Foundation.h>

typedef void *NSURLPtr;

NSURLPtr nsURLWithString(CFStringRef url) {
	return (NSURLPtr)[NSURL URLWithString:(NSString *)url];
}

CFStringRef nsURLFilePathURL(NSURLPtr u) {
	return (CFStringRef)[[(NSURL *)u filePathURL] absoluteString];
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

func (u *NSURL) FilePath() string {
	return C.nsURLFilePathURL(u.native).String()
}
