package macos

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

var (
	nsMenuItemValidators = make(map[int]func() bool)
	nsMenuItemHandlers   = make(map[int]func())
)

type NSMenuItem struct {
	native C.NSMenuItemPtr
}

func NSMenuItemInitWithTitleActionKeyEquivalent(tag int, title, keyEquiv string, modifiers int, validator func() bool, handler func()) *NSMenuItem {
	tstr := CFStringCreateWithString(title)
	defer tstr.Release()
	kstr := CFStringCreateWithString(keyEquiv)
	defer kstr.Release()
	item := &NSMenuItem{native: C.nsMenuItemInitWithTitleActionKeyEquivalent(C.int(tag), tstr, kstr, C.int(modifiers))}
	if validator != nil {
		nsMenuItemValidators[tag] = validator
	} else {
		delete(nsMenuItemValidators, tag)
	}
	if handler != nil {
		nsMenuItemHandlers[tag] = handler
	} else {
		delete(nsMenuItemHandlers, tag)
	}
	return item
}

func NSMenuSeparatorItem() *NSMenuItem {
	return &NSMenuItem{native: C.nsMenuSeparatorItem()}
}

func (mi *NSMenuItem) Tag() int {
	return int(C.nsMenuItemTag(mi.native))
}

func (mi *NSMenuItem) Title() string {
	return C.nsMenuItemTitle(mi.native).String()
}

func (mi *NSMenuItem) SetTitle(title string) {
	str := CFStringCreateWithString(title)
	C.nsMenuItemSetTitle(mi.native, str)
	str.Release()
}

func (mi *NSMenuItem) HasSubmenu() bool {
	return bool(C.nsMenuItemHasSubmenu(mi.native))
}

func (mi *NSMenuItem) Submenu() *NSMenu {
	if m := C.nsMenuItemSubmenu(mi.native); m != nil {
		return &NSMenu{native: m}
	}
	return nil
}

func (mi *NSMenuItem) SetSubmenu(menu *NSMenu) {
	C.nsMenuItemSetSubmenu(mi.native, menu.native)
}
