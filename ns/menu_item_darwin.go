package ns

import (
	"github.com/richardwilkes/macos/cf"
)

/*
#import <Cocoa/Cocoa.h>

typedef void *NSMenuPtr;
typedef void *NSMenuItemPtr;

bool menuItemValidateCallback(int tag);
void menuItemHandleCallback(int tag);

@interface MenuItemDelegate : NSObject<NSMenuItemValidation>
@end

@implementation MenuItemDelegate
- (BOOL)validateMenuItem:(NSMenuItem *)menuItem {
	return menuItemValidateCallback([menuItem tag]) ? YES : NO;
}

- (void)handleMenuItem:(id)sender {
	menuItemHandleCallback([(NSMenuItem *)sender tag]);
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

bool nsMenuItemHasSubmenu(NSMenuItemPtr mi) {
	return [(NSMenuItem *)mi hasSubmenu] != 0;
}

NSMenuPtr nsMenuItemSubmenu(NSMenuItemPtr mi) {
	return [(NSMenuItem *)mi submenu];
}

void nsMenuItemSetSubmenu(NSMenuItemPtr mi, NSMenuPtr menu) {
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

var (
	menuItemValidators = make(map[int]func() bool)
	menuItemHandlers   = make(map[int]func())
)

type MenuItem struct {
	native C.NSMenuItemPtr
}

func MenuItemInitWithTitleActionKeyEquivalent(tag int, title, keyEquiv string, modifiers int, validator func() bool, handler func()) *MenuItem {
	if validator != nil {
		menuItemValidators[tag] = validator
	} else {
		delete(menuItemValidators, tag)
	}
	if handler != nil {
		menuItemHandlers[tag] = handler
	} else {
		delete(menuItemHandlers, tag)
	}
	tstr := cf.StringCreateWithString(title)
	defer tstr.Release()
	kstr := cf.StringCreateWithString(keyEquiv)
	defer kstr.Release()
	return &MenuItem{native: C.nsMenuItemInitWithTitleActionKeyEquivalent(C.int(tag), C.CFStringRef(tstr), C.CFStringRef(kstr), C.int(modifiers))}
}

func MenuSeparatorItem() *MenuItem {
	return &MenuItem{native: C.nsMenuSeparatorItem()}
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

func (mi *MenuItem) HasSubmenu() bool {
	return bool(C.nsMenuItemHasSubmenu(mi.native))
}

func (mi *MenuItem) Submenu() *Menu {
	if m := C.nsMenuItemSubmenu(mi.native); m != nil {
		return &Menu{native: m}
	}
	return nil
}

func (mi *MenuItem) SetSubmenu(menu *Menu) {
	C.nsMenuItemSetSubmenu(mi.native, menu.native)
}
