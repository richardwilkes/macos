// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

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
