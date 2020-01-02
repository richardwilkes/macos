// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package sec

import "github.com/richardwilkes/macos/cf"

// #import <Security/Security.h>
import "C"

type Trust = C.SecTrustRef

func (t Trust) CopyExceptions() cf.Data {
	return cf.Data(C.SecTrustCopyExceptions(t))
}

func (t Trust) SetExceptions(exceptions cf.Data) bool {
	return bool(C.SecTrustSetExceptions(t, C.CFDataRef(exceptions)))
}
