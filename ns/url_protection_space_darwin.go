package ns

import "github.com/richardwilkes/macos/sec"

/*
#import <Foundation/Foundation.h>
#import <Security/Security.h>

typedef void *NSURLProtectionSpacePtr;

SecTrustRef nsURLProtectionSpaceServerTrust(NSURLProtectionSpacePtr s) {
	return [(NSURLProtectionSpace *)s serverTrust];
}
*/
import "C"

type URLProtectionSpace struct {
	native C.NSURLProtectionSpacePtr
}

func (s *URLProtectionSpace) ServerTrust() sec.Trust {
	return sec.Trust(C.nsURLProtectionSpaceServerTrust(s.native))
}
