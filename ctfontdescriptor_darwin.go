package macos

// #import <CoreText/CoreText.h>
import "C"

type CTFontDescriptor = C.CTFontDescriptorRef

func (fd CTFontDescriptor) CopyAttribute(attribute CFString) CFType {
	return C.CTFontDescriptorCopyAttribute(fd, attribute)
}
