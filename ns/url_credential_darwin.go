package ns

import "github.com/richardwilkes/macos/sec"

/*
#import <WebKit/WebKit.h>

typedef void *NSURLCredentialPtr;

NSURLCredentialPtr nsURLCredentialForTrust(SecTrustRef trust) {
	return (NSURLCredentialPtr)[NSURLCredential credentialForTrust:trust];
}
*/
import "C"

type URLCredentialNative = C.NSURLCredentialPtr

type URLCredential struct {
	native URLCredentialNative
}

func URLCredentialFromNative(native URLCredentialNative) *URLCredential {
	return &URLCredential{native: native}
}

func URLCredentialForTrust(trust sec.Trust) *URLCredential {
	return &URLCredential{native: C.nsURLCredentialForTrust(C.SecTrustRef(trust))}
}

func (uc *URLCredential) Native() URLCredentialNative {
	return uc.native
}
