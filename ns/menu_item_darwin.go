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

typedef void *NSMenuPtr;
typedef void *NSMenuItemPtr;

bool menuItemValidateCallback(NSMenuItemPtr item);
void menuItemHandleCallback(NSMenuItemPtr item);

@interface MenuItemDelegate : NSObject<NSMenuItemValidation>
@end

@implementation MenuItemDelegate
- (BOOL)validateMenuItem:(NSMenuItem *)menuItem {
	return menuItemValidateCallback((NSMenuItemPtr)menuItem) ? YES : NO;
}

- (void)handleMenuItem:(id)sender {
	menuItemHandleCallback((NSMenuItemPtr)sender);
}
@end

static MenuItemDelegate *menuItemDelegate = nil;

NSMenuItemPtr nsMenuItemInitWithTitleActionKeyEquivalent(int tag, CFStringRef title, CFStringRef keyEquiv, int modifiers) {
	NSMenuItem *item = [[[NSMenuItem alloc] initWithTitle:(NSString *)title action:NSSelectorFromString(@"handleMenuItem:") keyEquivalent:(NSString *)keyEquiv] retain];
	[item setTag:tag];
	[item setKeyEquivalentModifierMask:modifiers];
	if (!menuItemDelegate) {
		menuItemDelegate = [MenuItemDelegate new];
	}
	[item setTarget:menuItemDelegate];
	return (NSMenuItemPtr)item;
}

NSMenuItemPtr nsMenuSeparatorItem() {
	return (NSMenuItemPtr)[[NSMenuItem separatorItem] retain];
}

NSMenuPtr nsMenuItemMenu(NSMenuItemPtr mi) {
	return (NSMenuPtr)[(NSMenuItem *)mi menu];
}

bool nsMenuItemIsSeparator(NSMenuItemPtr mi) {
	return [(NSMenuItem *)mi isSeparatorItem];
}

int nsMenuItemTag(NSMenuItemPtr mi) {
	return [(NSMenuItem *)mi tag];
}

NSControlStateValue nsMenuItemState(NSMenuItemPtr mi) {
	return [(NSMenuItem *)mi state];
}

void nsMenuItemSetState(NSMenuItemPtr mi, NSControlStateValue state) {
	[(NSMenuItem *)mi setState:(NSControlStateValue)state];
}

CFStringRef nsMenuItemTitle(NSMenuItemPtr mi) {
	return (CFStringRef)[(NSMenuItem *)mi title];
}

void nsMenuItemSetTitle(NSMenuItemPtr mi, CFStringRef title) {
	[(NSMenuItem *)mi setTitle:(NSString *)title];
}

bool nsMenuItemHasSubMenu(NSMenuItemPtr mi) {
	return [(NSMenuItem *)mi hasSubmenu] != 0;
}

NSMenuPtr nsMenuItemSubMenu(NSMenuItemPtr mi) {
	return [(NSMenuItem *)mi submenu];
}

void nsMenuItemSetSubMenu(NSMenuItemPtr mi, NSMenuPtr menu) {
	[(NSMenuItem *)mi setSubmenu:(NSMenu *)menu];
}
*/
import "C"

type MenuItemState int

const (
	MenuItemStateMixed MenuItemState = iota - 1
	MenuItemStateOff
	MenuItemStateOn
)

type MenuValidator func(item *MenuItem) bool
type MenuHandler func(item *MenuItem)

var (
	menuItemValidators = make(map[C.NSMenuItemPtr]MenuValidator)
	menuItemHandlers   = make(map[C.NSMenuItemPtr]MenuHandler)
)

type MenuItemNative = C.NSMenuItemPtr

type MenuItem struct {
	native C.NSMenuItemPtr
}

func MenuItemInitWithTitleActionKeyEquivalent(tag int, title, keyEquiv string, modifiers int, validator MenuValidator, handler MenuHandler) *MenuItem {
	tstr := cf.StringCreateWithString(title)
	defer tstr.Release()
	kstr := cf.StringCreateWithString(keyEquiv)
	defer kstr.Release()
	item := &MenuItem{native: C.nsMenuItemInitWithTitleActionKeyEquivalent(C.int(tag), C.CFStringRef(tstr), C.CFStringRef(kstr), C.int(modifiers))}
	if validator != nil {
		menuItemValidators[item.native] = validator
	} else {
		delete(menuItemValidators, item.native)
	}
	if handler != nil {
		menuItemHandlers[item.native] = handler
	} else {
		delete(menuItemHandlers, item.native)
	}
	return item
}

func (mi *MenuItem) Native() MenuItemNative {
	return mi.native
}

func MenuSeparatorItem() *MenuItem {
	return &MenuItem{native: C.nsMenuSeparatorItem()}
}

func (mi *MenuItem) Menu() *Menu {
	if m := C.nsMenuItemMenu(mi.native); m != nil {
		return &Menu{native: m}
	}
	return nil
}

func (mi *MenuItem) IsSeparator() bool {
	return bool(C.nsMenuItemIsSeparator(mi.native))
}

func (mi *MenuItem) Tag() int {
	return int(C.nsMenuItemTag(mi.native))
}

func (mi *MenuItem) State() MenuItemState {
	return MenuItemState(C.nsMenuItemState(mi.native))
}

func (mi *MenuItem) SetState(state MenuItemState) {
	C.nsMenuItemSetState(mi.native, C.NSControlStateValue(state))
}

func (mi *MenuItem) Title() string {
	return cf.String(C.nsMenuItemTitle(mi.native)).String()
}

func (mi *MenuItem) SetTitle(title string) {
	str := cf.StringCreateWithString(title)
	C.nsMenuItemSetTitle(mi.native, C.CFStringRef(str))
	str.Release()
}

func (mi *MenuItem) HasSubMenu() bool {
	return bool(C.nsMenuItemHasSubMenu(mi.native))
}

func (mi *MenuItem) SubMenu() *Menu {
	if m := C.nsMenuItemSubMenu(mi.native); m != nil {
		return &Menu{native: m}
	}
	return nil
}

func (mi *MenuItem) SetSubMenu(menu *Menu) {
	C.nsMenuItemSetSubMenu(mi.native, menu.native)
}
