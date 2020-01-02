// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package cg

import "github.com/richardwilkes/macos/cf"

// #import <ImageIO/ImageIO.h>
import "C"

type DataProvider = C.CGDataProviderRef

func DataProviderCreateWithCFData(data cf.Data) DataProvider {
	return C.CGDataProviderCreateWithCFData(C.CFDataRef(data))
}

func (dp DataProvider) Release() {
	C.CGDataProviderRelease(dp)
}
