package ns

import (
	"github.com/richardwilkes/macos/cf"
)

/*
#import <Cocoa/Cocoa.h>
#import <CoreFoundation/CoreFoundation.h>

typedef void *NSPasteboardItemPtr;

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

type PasteboardItem struct {
	native C.NSPasteboardItemPtr
}

func NewPasteboardItem() *PasteboardItem {
	return &PasteboardItem{native: C.nsPasteboardItemNew()}
}

func (pbi *PasteboardItem) Types() cf.Array {
	return cf.Array(C.nsPasteboardItemTypes(pbi.native))
}

func (pbi *PasteboardItem) DataForType(dataType string) cf.Data {
	str := cf.StringCreateWithString(dataType)
	defer str.Release()
	return cf.Data(C.nsPasteboardItemDataForType(pbi.native, C.CFStringRef(str)))
}

func (pbi *PasteboardItem) SetDataForType(data cf.Data, dataType string) {
	str := cf.StringCreateWithString(dataType)
	defer str.Release()
	C.nsPasteboardItemSetDataForType(pbi.native, C.CFDataRef(data), C.CFStringRef(str))
}
