package macos

// #import <CoreFoundation/CoreFoundation.h>
import "C"

type (
	CFAttributedString        = C.CFAttributedStringRef
	CFMutableAttributedString = C.CFMutableAttributedStringRef
)

func CFAttributedStringCreateMutable(maxLength int) CFMutableAttributedString {
	return C.CFAttributedStringCreateMutable(0, C.CFIndex(maxLength))
}

func (s CFMutableAttributedString) BeginEditing() {
	C.CFAttributedStringBeginEditing(s)
}

func (s CFMutableAttributedString) ReplaceString(location, length int, replacement CFString) {
	C.CFAttributedStringReplaceString(s, C.CFRangeMake(C.CFIndex(location), C.CFIndex(length)), replacement)
}

func (s CFMutableAttributedString) SetAttribute(location, length int, attrName CFString, value CFType) {
	C.CFAttributedStringSetAttribute(s, C.CFRangeMake(C.CFIndex(location), C.CFIndex(length)), attrName, value)
}

func (s CFMutableAttributedString) EndEditing() {
	C.CFAttributedStringEndEditing(s)
}

func (s CFMutableAttributedString) Release() {
	C.CFRelease(CFType(s))
}
