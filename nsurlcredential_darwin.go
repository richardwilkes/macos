package macos

/*
#import <WebKit/WebKit.h>
#import "authentication_challenge_response.h"

NSURLCredentialPtr nsURLCredentialForTrust(SecTrustRef trust) {
	return (NSURLCredentialPtr)[NSURLCredential credentialForTrust:trust];
}
*/
import "C"

type NSURLCredential struct {
	native C.NSURLCredentialPtr
}

func NSURLCredentialForTrust(trust SecTrust) *NSURLCredential {
	return &NSURLCredential{native: C.nsURLCredentialForTrust(trust)}
}
