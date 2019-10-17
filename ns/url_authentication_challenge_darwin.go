package ns

/*
#import <Foundation/Foundation.h>

typedef void *NSURLAuthenticationChallengePtr;
typedef void *NSURLProtectionSpacePtr;

NSURLProtectionSpacePtr nsURLAuthenticationChallengeProtectionSpace(NSURLAuthenticationChallengePtr challenge) {
	return (NSURLProtectionSpacePtr)[(NSURLAuthenticationChallenge *)challenge protectionSpace];
}
*/
import "C"

type URLAuthenticationChallengeNative = C.NSURLAuthenticationChallengePtr

type URLAuthenticationChallenge struct {
	native URLAuthenticationChallengeNative
}

func NewURLAuthenticationChallengeFromNative(native URLAuthenticationChallengeNative) *URLAuthenticationChallenge {
	return &URLAuthenticationChallenge{native: native}
}

func (c *URLAuthenticationChallenge) ProtectionSpace() *URLProtectionSpace {
	return &URLProtectionSpace{native: C.nsURLAuthenticationChallengeProtectionSpace(c.native)}
}
