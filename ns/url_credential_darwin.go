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
