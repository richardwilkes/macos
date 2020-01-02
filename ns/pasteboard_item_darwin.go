// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

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
