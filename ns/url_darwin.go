package ns

import "github.com/richardwilkes/macos/cf"

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

const (
	URLSessionAuthChallengeUseCredential URLSessionAuthChallengeDisposition = iota
	URLSessionAuthChallengePerformDefaultHandling
	URLSessionAuthChallengeCancelAuthenticationChallenge
	URLSessionAuthChallengeRejectProtectionSpace
)

type URLSessionAuthChallengeDisposition int

type URL struct {
	native C.NSURLPtr
}

func URLWithString(url string) *URL {
	str := cf.StringCreateWithString(url)
	defer str.Release()
	return &URL{native: C.nsURLWithString(C.CFStringRef(str))}
}

func (u *URL) ResolveFilePath() string {
	return cf.String(C.nsURLFilePathURL(u.native)).String()
}
