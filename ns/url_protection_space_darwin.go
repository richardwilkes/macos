// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

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
