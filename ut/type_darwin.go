package ut

import "github.com/richardwilkes/macos/cf"

// #import <Cocoa/Cocoa.h>
import "C"

func TypeCreatePreferredIdentifierForTagClassMimeType(mimeType string) string {
	tag := cf.StringCreateWithString(mimeType)
	defer tag.Release()
	uti := cf.String(C.UTTypeCreatePreferredIdentifierForTag(C.kUTTagClassMIMEType, C.CFStringRef(tag), 0))
	if uti == 0 {
		return ""
	}
	defer uti.Release()
	return uti.String()
}

func TypeCopyPreferredTagWithClassMimeType(uti string) string {
	tag := cf.StringCreateWithString(uti)
	defer tag.Release()
	mimeType := cf.String(C.UTTypeCopyPreferredTagWithClass(C.CFStringRef(tag), C.kUTTagClassMIMEType))
	if mimeType == 0 {
		return ""
	}
	defer mimeType.Release()
	return mimeType.String()
}
