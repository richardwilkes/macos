package cf

// #import <CoreFoundation/CoreFoundation.h>
import "C"

type (
	AttributedString        = C.CFAttributedStringRef
	MutableAttributedString = C.CFMutableAttributedStringRef
)

func AttributedStringCreateMutable(maxLength int) MutableAttributedString {
	return C.CFAttributedStringCreateMutable(0, C.CFIndex(maxLength))
}

func (s MutableAttributedString) BeginEditing() {
	C.CFAttributedStringBeginEditing(s)
}

func (s MutableAttributedString) ReplaceString(location, length int, replacement String) {
	C.CFAttributedStringReplaceString(s, C.CFRangeMake(C.CFIndex(location), C.CFIndex(length)), replacement)
}

func (s MutableAttributedString) SetAttribute(location, length int, attrName String, value Type) {
	C.CFAttributedStringSetAttribute(s, C.CFRangeMake(C.CFIndex(location), C.CFIndex(length)), attrName, value)
}

func (s MutableAttributedString) EndEditing() {
	C.CFAttributedStringEndEditing(s)
}

func (s MutableAttributedString) Release() {
	C.CFRelease(Type(s))
}
