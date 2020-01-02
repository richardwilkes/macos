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

typedef void *NSURLRequestPtr;
typedef void *NSURLPtr;

NSURLRequestPtr nsURLRequestWithURL(NSURLPtr url) {
	return (NSURLRequestPtr)[NSURLRequest requestWithURL:(NSURL *)url];
}
*/
import "C"

type URLRequestNative = C.NSURLRequestPtr

type URLRequest struct {
	native URLRequestNative
}

func URLRequestWithURL(url *URL) *URLRequest {
	return &URLRequest{native: C.nsURLRequestWithURL(url.native)}
}

func (ur *URLRequest) Native() URLRequestNative {
	return ur.native
}
