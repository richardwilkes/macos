package macos

/*
#import <Foundation/Foundation.h>
#import <Security/Security.h>

typedef void *NSURLProtectionSpacePtr;

SecTrustRef nsURLProtectionSpaceServerTrust(NSURLProtectionSpacePtr s) {
	return [(NSURLProtectionSpace *)s serverTrust];
}
*/
import "C"

type NSURLProtectionSpace struct {
	native C.NSURLProtectionSpacePtr
}

func (s *NSURLProtectionSpace) ServerTrust() SecTrust {
	return C.nsURLProtectionSpaceServerTrust(s.native)
}
