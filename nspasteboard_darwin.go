package macos

/*
#import <Cocoa/Cocoa.h>
#import <CoreFoundation/CoreFoundation.h>

typedef void *NSPasteboardPtr;

NSPasteboardPtr nsPasteboardGeneral() {
	return [NSPasteboard generalPasteboard];
}

CFArrayRef nsPasteboardTypes(NSPasteboardPtr pb) {
	return (CFArrayRef)[(NSPasteboard *)pb types];
}

int nsPasteboardChangeCount(NSPasteboardPtr pb) {
	return [(NSPasteboard *)pb changeCount];
}

void nsPasteboardClearContents(NSPasteboardPtr pb) {
	[(NSPasteboard *)pb clearContents];
}

CFDataRef nsPasteboardDataForType(NSPasteboardPtr pb, CFStringRef dataType) {
	return (CFDataRef)[(NSPasteboard *)pb dataForType:(NSString *)dataType];
}

void nsPasteboardSetDataForType(NSPasteboardPtr pb, CFDataRef data, CFStringRef dataType) {
	[(NSPasteboard *)pb setData:(NSData *)data forType:(NSString *)dataType];
}
*/
import "C"

type NSPasteboard struct {
	native C.NSPasteboardPtr
}

func NSPasteboardGeneral() *NSPasteboard {
	return &NSPasteboard{native: C.nsPasteboardGeneral()}
}

func (pb *NSPasteboard) Types() CFArray {
	return C.nsPasteboardTypes(pb.native)
}

func (pb *NSPasteboard) ChangeCount() int {
	return int(C.nsPasteboardChangeCount(pb.native))
}

func (pb *NSPasteboard) ClearContents() {
	C.nsPasteboardClearContents(pb.native)
}

func (pb *NSPasteboard) DataForType(dataType string) CFData {
	str := CFStringCreateWithString(dataType)
	defer str.Release()
	return C.nsPasteboardDataForType(pb.native, str)
}

func (pb *NSPasteboard) SetDataForType(data CFData, dataType string) {
	str := CFStringCreateWithString(dataType)
	defer str.Release()
	C.nsPasteboardSetDataForType(pb.native, data, str)
}

func UTTypeCreatePreferredIdentifierForTagClassMimeType(mimeType string) string {
	tag := CFStringCreateWithString(mimeType)
	defer tag.Release()
	uti := C.UTTypeCreatePreferredIdentifierForTag(C.kUTTagClassMIMEType, tag, 0)
	if uti == 0 {
		return ""
	}
	defer uti.Release()
	return uti.String()
}

func UTTypeCopyPreferredTagWithClassMimeType(uti string) string {
	tag := CFStringCreateWithString(uti)
	defer tag.Release()
	mimeType := C.UTTypeCopyPreferredTagWithClass(tag, C.kUTTagClassMIMEType)
	if mimeType == 0 {
		return ""
	}
	defer mimeType.Release()
	return mimeType.String()
}
