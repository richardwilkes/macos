// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

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
