package macos

/*
#import <Cocoa/Cocoa.h>
#import <CoreFoundation/CoreFoundation.h>

typedef void *NSPasteboardPtr;
typedef void *NSPasteboardItemPtr;

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

CFArrayRef nsPasteboardItems(NSPasteboardPtr pb) {
	return (CFArrayRef)[(NSPasteboard *)pb pasteboardItems];
}

bool nsPasteboardWriteObjects(NSPasteboardPtr pb, CFArrayRef objects) {
	return [(NSPasteboard *)pb writeObjects:(NSArray<id<NSPasteboardWriting>> *)objects];
}

NSPasteboardItemPtr nsPasteboardItemNew() {
	return (NSPasteboardItemPtr)[[NSPasteboardItem alloc] init];
}

CFArrayRef nsPasteboardItemTypes(NSPasteboardItemPtr pbi) {
	return (CFArrayRef)[(NSPasteboardItem *)pbi types];
}

CFDataRef nsPasteboardItemDataForType(NSPasteboardItemPtr pbi, CFStringRef dataType) {
	return (CFDataRef)[(NSPasteboardItem *)pbi dataForType:(NSString *)dataType];
}

void nsPasteboardItemSetDataForType(NSPasteboardItemPtr pbi, CFDataRef data, CFStringRef dataType) {
	[(NSPasteboardItem *)pbi setData:(NSData *)data forType:(NSString *)dataType];
}
*/
import "C"
import "unsafe"

type NSPasteboard struct {
	native C.NSPasteboardPtr
}

type NSPasteboardItem struct {
	native C.NSPasteboardItemPtr
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

func (pb *NSPasteboard) Items() []*NSPasteboardItem {
	a := C.nsPasteboardItems(pb.native)
	count := a.GetCount()
	items := make([]*NSPasteboardItem, count)
	for i := 0; i < count; i++ {
		items[i] = &NSPasteboardItem{native: C.NSPasteboardItemPtr(a.GetValueAtIndex(i))}
	}
	return items
}

func (pb *NSPasteboard) SetItems(items []*NSPasteboardItem) bool {
	a := CFMutableArrayCreate(len(items))
	for _, item := range items {
		a.AppendValue(unsafe.Pointer(item.native))
	}
	return bool(C.nsPasteboardWriteObjects(pb.native, a.AsCFArray()))
}

func NewPasteboardItem() *NSPasteboardItem {
	return &NSPasteboardItem{native: C.nsPasteboardItemNew()}
}

func (pbi *NSPasteboardItem) Types() CFArray {
	return C.nsPasteboardItemTypes(pbi.native)
}

func (pbi *NSPasteboardItem) DataForType(dataType string) CFData {
	str := CFStringCreateWithString(dataType)
	defer str.Release()
	return C.nsPasteboardItemDataForType(pbi.native, str)
}

func (pbi *NSPasteboardItem) SetDataForType(data CFData, dataType string) {
	str := CFStringCreateWithString(dataType)
	defer str.Release()
	C.nsPasteboardItemSetDataForType(pbi.native, data, str)
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
