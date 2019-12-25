package ns

import "github.com/richardwilkes/macos/cf"

/*
#import <Cocoa/Cocoa.h>

typedef void *NSMenuPtr;
typedef void *NSMenuItemPtr;
typedef void *NSViewPtr;

void updateMenuCallback(NSMenuPtr menu);

@interface MenuDelegate : NSObject<NSMenuDelegate>
@end

@implementation MenuDelegate
- (void)menuNeedsUpdate:(NSMenu *)menu {
	updateMenuCallback((NSMenuPtr)(menu));
}
@end

static MenuDelegate *menuDelegate = nil;

NSMenuPtr nsMenuInitWithTitle(CFStringRef title) {
	NSMenu *menu = [[[NSMenu alloc] initWithTitle:(NSString *)title] retain];
	if (!menuDelegate) {
		menuDelegate = [MenuDelegate new];
	}
	[menu setDelegate:menuDelegate];
	return menu;
}

CFStringRef nsMenuTitle(NSMenuPtr m) {
	return (CFStringRef)[(NSMenu *)m title];
}

int nsMenuNumberOfItems(NSMenuPtr m) {
	return [(NSMenu *)m numberOfItems];
}

NSMenuItemPtr nsMenuItemAtIndex(NSMenuPtr m, int index) {
	return (NSMenuItemPtr)[(NSMenu *)m itemAtIndex:index];
}

void nsMenuInsertItemAtIndex(NSMenuPtr m, NSMenuItemPtr mi, int index) {
	[(NSMenu *)m insertItem:(NSMenuItem *)mi atIndex:index];
}

void nsMenuRemoveItem(NSMenuPtr m, int index) {
	[(NSMenu *)m removeItemAtIndex:index];
}

void nsMenuPopupMenuPositioningItemAtLocationInView(NSMenuPtr m, NSMenuItemPtr mi, CGFloat x, CGFloat y, CGFloat width, CGFloat height, NSViewPtr view) {
	// I don't use popupMenuPositioningItem:atLocation:inView: here because it
	// fails to work when a modal dialog is being used.
	NSPopUpButtonCell *popUpButtonCell = [[NSPopUpButtonCell alloc] initTextCell:@"" pullsDown:NO];
	[popUpButtonCell setAutoenablesItems:NO];
	[popUpButtonCell setAltersStateOfSelectedItem:NO];
	[popUpButtonCell setMenu:(NSMenu *)m];
	[popUpButtonCell selectItem:(NSMenuItem *)mi];
	[popUpButtonCell performClickWithFrame:NSMakeRect(x, y, width, height) inView:(NSView *)view];
	[popUpButtonCell release];
}

void nsMenuRelease(NSMenuPtr m) {
	[(NSMenu *)m release];
}
*/
import "C"

type MenuNative = C.NSMenuPtr

var menuUpdaters = make(map[C.NSMenuPtr]func(*Menu))

type Menu struct {
	native C.NSMenuPtr
}

func MenuInitWithTitle(title string, updater func(menu *Menu)) *Menu {
	str := cf.StringCreateWithString(title)
	defer str.Release()
	m := &Menu{native: C.nsMenuInitWithTitle(C.CFStringRef(str))}
	if updater != nil {
		menuUpdaters[m.native] = updater
	}
	return m
}

func (m *Menu) Native() MenuNative {
	return m.native
}

func (m *Menu) Title() string {
	return cf.String(C.nsMenuTitle(m.native)).String()
}

func (m *Menu) NumberOfItems() int {
	return int(C.nsMenuNumberOfItems(m.native))
}

func (m *Menu) ItemAtIndex(index int) *MenuItem {
	return &MenuItem{native: C.nsMenuItemAtIndex(m.native, C.int(index))}
}

func (m *Menu) InsertItemAtIndex(item *MenuItem, index int) {
	C.nsMenuInsertItemAtIndex(m.native, item.native, C.int(index))
}

func (m *Menu) RemoveItem(index int) {
	C.nsMenuRemoveItem(m.native, C.int(index))
}

func (m *Menu) PopupMenuPositioningItemAtLocationInView(item *MenuItem, x, y, width, height float64, view *View) {
	C.nsMenuPopupMenuPositioningItemAtLocationInView(m.native, item.native, C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height), view.native)
}

func (m *Menu) Release() {
	C.nsMenuRelease(m.native)
}
