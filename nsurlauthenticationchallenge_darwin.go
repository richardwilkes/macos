package macos

/*
#import <Foundation/Foundation.h>

typedef void *NSURLAuthenticationChallengePtr;
typedef void *NSURLProtectionSpacePtr;

NSURLProtectionSpacePtr nsURLAuthenticationChallengeProtectionSpace(NSURLAuthenticationChallengePtr challenge) {
	return (NSURLProtectionSpacePtr)[(NSURLAuthenticationChallenge *)challenge protectionSpace];
}
*/
import "C"

type NSURLAuthenticationChallenge struct {
	native C.NSURLAuthenticationChallengePtr
}

func (c *NSURLAuthenticationChallenge) ProtectionSpace() *NSURLProtectionSpace {
	return &NSURLProtectionSpace{native: C.nsURLAuthenticationChallengeProtectionSpace(c.native)}
}
