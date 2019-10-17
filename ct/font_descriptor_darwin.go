package ct

import "github.com/richardwilkes/macos/cf"

// #import <CoreText/CoreText.h>
import "C"

type FontDescriptor = C.CTFontDescriptorRef

func (fd FontDescriptor) CopyAttribute(attribute cf.String) cf.Type {
	return cf.Type(C.CTFontDescriptorCopyAttribute(fd, C.CFStringRef(attribute)))
}
