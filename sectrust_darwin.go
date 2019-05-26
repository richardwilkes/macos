package macos

// #import <Security/Security.h>
import "C"

type SecTrust = C.SecTrustRef

func (st SecTrust) CopyExceptions() CFData {
	return C.SecTrustCopyExceptions(st)
}

func (st SecTrust) SetExceptions(exceptions CFData) bool {
	return bool(C.SecTrustSetExceptions(st, exceptions))
}
