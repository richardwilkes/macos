package ns

import (
	"unsafe"

	"github.com/richardwilkes/macos/cf"
)

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
*/
import "C"

type Pasteboard struct {
	native C.NSPasteboardPtr
}

func PasteboardGeneral() *Pasteboard {
	return &Pasteboard{native: C.nsPasteboardGeneral()}
}

func (pb *Pasteboard) Types() cf.Array {
	return cf.Array(C.nsPasteboardTypes(pb.native))
}

func (pb *Pasteboard) ChangeCount() int {
	return int(C.nsPasteboardChangeCount(pb.native))
}

func (pb *Pasteboard) ClearContents() {
	C.nsPasteboardClearContents(pb.native)
}

func (pb *Pasteboard) DataForType(dataType string) cf.Data {
	str := cf.StringCreateWithString(dataType)
	defer str.Release()
	return cf.Data(C.nsPasteboardDataForType(pb.native, C.CFStringRef(str)))
}

func (pb *Pasteboard) SetDataForType(data cf.Data, dataType string) {
	str := cf.StringCreateWithString(dataType)
	defer str.Release()
	C.nsPasteboardSetDataForType(pb.native, C.CFDataRef(data), C.CFStringRef(str))
}

func (pb *Pasteboard) Items() []*PasteboardItem {
	a := cf.Array(C.nsPasteboardItems(pb.native))
	count := a.GetCount()
	items := make([]*PasteboardItem, count)
	for i := 0; i < count; i++ {
		items[i] = &PasteboardItem{native: C.NSPasteboardItemPtr(a.GetValueAtIndex(i))}
	}
	return items
}

func (pb *Pasteboard) SetItems(items []*PasteboardItem) bool {
	a := cf.MutableArrayCreate(len(items))
	for _, item := range items {
		a.AppendValue(unsafe.Pointer(item.native))
	}
	return bool(C.nsPasteboardWriteObjects(pb.native, C.CFArrayRef(a.AsCFArray())))
}
