package macos

/*
#import <Cocoa/Cocoa.h>

typedef void *NSMenuPtr;
typedef void *NSMenuItemPtr;
typedef void *NSViewPtr;

NSMenuPtr nsMenuInitWithTitle(CFStringRef title) {
	return [[[NSMenu alloc] initWithTitle:(NSString *)title] retain];
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

void nsMenuPopupMenuPositioningItemAtLocationInView(NSMenuPtr m, NSMenuItemPtr mi, CGFloat x, CGFloat y, NSViewPtr view) {
	[(NSMenu *)m popUpMenuPositioningItem:(NSMenuItem *)mi atLocation:NSMakePoint(x, y) inView:(NSView *)view];
}

void nsMenuRelease(NSMenuPtr m) {
	[(NSMenu *)m release];
}
*/
import "C"

type NSMenuNative = C.NSMenuPtr

type NSMenu struct {
	native C.NSMenuPtr
}

func NSMenuInitWithTitle(title string) *NSMenu {
	str := CFStringCreateWithString(title)
	defer str.Release()
	return &NSMenu{native: C.nsMenuInitWithTitle(str)}
}

func (m *NSMenu) Native() NSMenuNative {
	return m.native
}

func (m *NSMenu) Title() string {
	return C.nsMenuTitle(m.native).String()
}

func (m *NSMenu) NumberOfItems() int {
	return int(C.nsMenuNumberOfItems(m.native))
}

func (m *NSMenu) ItemAtIndex(index int) *NSMenuItem {
	return &NSMenuItem{native: C.nsMenuItemAtIndex(m.native, C.int(index))}
}

func (m *NSMenu) InsertItemAtIndex(item *NSMenuItem, index int) {
	C.nsMenuInsertItemAtIndex(m.native, item.native, C.int(index))
}

func (m *NSMenu) RemoveItem(index int) {
	C.nsMenuRemoveItem(m.native, C.int(index))
}

func (m *NSMenu) PopupMenuPositioningItemAtLocationInView(item *NSMenuItem, x, y float64, view *NSView) {
	C.nsMenuPopupMenuPositioningItemAtLocationInView(m.native, item.native, C.CGFloat(x), C.CGFloat(y), view.native)
}

func (m *NSMenu) Release() {
	C.nsMenuRelease(m.native)
}
