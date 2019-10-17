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
