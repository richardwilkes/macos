package macos

/*
#import <Cocoa/Cocoa.h>

typedef void *NSApplicationPtr;
typedef void *NSRunningApplicationPtr;
typedef void *NSNotificationPtr;
typedef void *NSWindowPtr;
typedef void *NSMenuPtr;

// Prototypes for app_callbacks_darwin.go
void applicationWillFinishLaunchingCallback(NSNotificationPtr aNotification);
void applicationDidFinishLaunchingCallback(NSNotificationPtr aNotification);
NSApplicationTerminateReply applicationShouldTerminateCallback(NSApplicationPtr sender);
BOOL applicationShouldTerminateAfterLastWindowClosedCallback(NSApplicationPtr theApplication);
void applicationWillTerminateCallback(NSNotificationPtr aNotification);
void applicationWillBecomeActiveCallback(NSNotificationPtr aNotification);
void applicationDidBecomeActiveCallback(NSNotificationPtr aNotification);
void applicationWillResignActiveCallback(NSNotificationPtr aNotification);
void applicationDidResignActiveCallback(NSNotificationPtr aNotification);
void themeChangedCallback(NSNotificationPtr aNotification);

@interface AppDelegate : NSObject<NSApplicationDelegate>
@end

@implementation AppDelegate
- (void)applicationWillFinishLaunching:(NSNotification *)aNotification {
	applicationWillFinishLaunchingCallback((NSNotificationPtr)aNotification);
}

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
	applicationDidFinishLaunchingCallback((NSNotificationPtr)aNotification);
}

- (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication *)sender {
	return applicationShouldTerminateCallback((NSApplicationPtr)sender);
}

- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication *)theApplication {
	return applicationShouldTerminateAfterLastWindowClosedCallback((NSApplicationPtr)theApplication);
}

- (void)applicationWillTerminate:(NSNotification *)aNotification {
	applicationWillTerminateCallback((NSNotificationPtr)aNotification);
}

- (void)applicationWillBecomeActive:(NSNotification *)aNotification {
	applicationWillBecomeActiveCallback((NSNotificationPtr)aNotification);
}

- (void)applicationDidBecomeActive:(NSNotification *)aNotification {
	applicationDidBecomeActiveCallback((NSNotificationPtr)aNotification);
}

- (void)applicationWillResignActive:(NSNotification *)aNotification {
	applicationWillResignActiveCallback((NSNotificationPtr)aNotification);
}

- (void)applicationDidResignActive:(NSNotification *)aNotification {
	applicationDidResignActiveCallback((NSNotificationPtr)aNotification);
}

- (void)themeChanged:(NSNotification *)aNotification {
	themeChangedCallback((NSNotificationPtr)aNotification);
}
@end

NSApplicationPtr nsSharedApplication() {
	return (NSApplicationPtr)[NSApplication sharedApplication];
}

void nsApplicationSetDelegate(NSApplicationPtr app) {
	AppDelegate *delegate = [AppDelegate new];
	[(NSApplication *)app setDelegate:delegate];
	[NSDistributedNotificationCenter.defaultCenter addObserver:delegate selector:@selector(themeChanged:) name:@"AppleInterfaceThemeChangedNotification" object: nil];
	[NSDistributedNotificationCenter.defaultCenter addObserver:delegate selector:@selector(themeChanged:) name:@"AppleColorPreferencesChangedNotification" object: nil];
}

void nsApplicationSetActivationPolicy(NSApplicationPtr app, NSApplicationActivationPolicy policy) {
	[(NSApplication *)app setActivationPolicy:policy];
}

NSWindowPtr nsApplicationKeyWindow(NSApplicationPtr app) {
	return (NSWindowPtr)[(NSApplication *)app keyWindow];
}

NSWindowPtr nsApplicationMainWindow(NSApplicationPtr app) {
	return (NSWindowPtr)[(NSApplication *)app mainWindow];
}

void nsApplicationHideOtherApplications(NSApplicationPtr app, NSApplicationPtr sender) {
	[(NSApplication *)app hideOtherApplications:(NSApplication *)sender];
}

void nsApplicationUnhideAllApplications(NSApplicationPtr app, NSApplicationPtr sender) {
	[(NSApplication *)app unhideAllApplications:(NSApplication *)sender];
}

void nsApplicationRun(NSApplicationPtr app) {
	[(NSApplication *)app run];
}

void nsApplicationReplyToApplicationShouldTerminate(NSApplicationPtr app, bool shouldTerminate) {
	[(NSApplication *)app replyToApplicationShouldTerminate:shouldTerminate ? YES : NO];
}

void nsApplicationTerminate(NSApplicationPtr app, NSApplicationPtr sender) {
	[(NSApplication *)app terminate:(NSApplication *)sender];
}

void nsApplicationSetMainMenu(NSApplicationPtr app, NSMenuPtr menu) {
	[(NSApplication *)app setMainMenu:(NSMenu *)menu];
}

void nsApplicationSetServicesMenu(NSApplicationPtr app, NSMenuPtr menu) {
	[(NSApplication *)app setServicesMenu:(NSMenu *)menu];
}

void nsApplicationSetWindowsMenu(NSApplicationPtr app, NSMenuPtr menu) {
	[(NSApplication *)app setWindowsMenu:(NSMenu *)menu];
}

void nsApplicationSetHelpMenu(NSApplicationPtr app, NSMenuPtr menu) {
	[(NSApplication *)app setHelpMenu:(NSMenu *)menu];
}

NSRunningApplicationPtr nsRunningApplicationCurrent() {
	return (NSRunningApplicationPtr)[NSRunningApplication currentApplication];
}

void nsRunningApplicationActivateWithOptions(NSRunningApplicationPtr app, NSApplicationActivationOptions options) {
	[(NSRunningApplication *)app activateWithOptions:options];
}

bool nsRunningApplicationHide(NSRunningApplicationPtr app) {
	return [(NSRunningApplication *)app hide] != 0;
}
*/
import "C"

const (
	NSTerminateCancel NSApplicationTerminateReply = iota
	NSTerminateNow
	NSTerminateLater
)

const (
	NSApplicationActivationPolicyRegular NSApplicationActivationPolicy = iota
	NSApplicationActivationPolicyAccessory
	NSApplicationActivationPolicyProhibited
)

const (
	NSApplicationActivateAllWindows NSApplicationActivationOptions = 1 << iota
	NSApplicationActivateIgnoringOtherApps
)

var currentAppDelegate NSApplicationDelegate

type (
	NSApplicationTerminateReply    int
	NSApplicationActivationPolicy  int
	NSApplicationActivationOptions int
)

type NSApplicationDelegate interface {
	ApplicationWillFinishLaunching(notification *NSNotification)
	ApplicationDidFinishLaunching(notification *NSNotification)
	ApplicationShouldTerminate(sender *NSApplication) NSApplicationTerminateReply
	ApplicationShouldTerminateAfterLastWindowClosed(app *NSApplication) bool
	ApplicationWillTerminate(notification *NSNotification)
	ApplicationWillBecomeActive(notification *NSNotification)
	ApplicationDidBecomeActive(notification *NSNotification)
	ApplicationWillResignActive(notification *NSNotification)
	ApplicationDidResignActive(notification *NSNotification)
	ThemeChanged(notification *NSNotification)
}

type NSApplication struct {
	native C.NSApplicationPtr
}

func NSSharedApplication() *NSApplication {
	return &NSApplication{native: C.nsSharedApplication()}
}

func (app *NSApplication) SetDelegate(delegate NSApplicationDelegate) {
	currentAppDelegate = delegate
	C.nsApplicationSetDelegate(app.native)
}

func (app *NSApplication) SetActivationPolicy(policy NSApplicationActivationPolicy) {
	C.nsApplicationSetActivationPolicy(app.native, C.NSApplicationActivationPolicy(policy))
}

func (app *NSApplication) KeyWindow() *NSWindow {
	if w := C.nsApplicationKeyWindow(app.native); w != nil {
		return &NSWindow{native: w}
	}
	return nil
}

func (app *NSApplication) MainWindow() *NSWindow {
	if w := C.nsApplicationMainWindow(app.native); w != nil {
		return &NSWindow{native: w}
	}
	return nil
}

func (app *NSApplication) Run() {
	C.nsApplicationRun(app.native)
}

func (app *NSApplication) UnhideAllApplications(sender *NSApplication) {
	var s C.NSApplicationPtr
	if sender != nil {
		s = sender.native
	}
	C.nsApplicationUnhideAllApplications(app.native, s)
}

func (app *NSApplication) HideOtherApplications(sender *NSApplication) {
	var s C.NSApplicationPtr
	if sender != nil {
		s = sender.native
	}
	C.nsApplicationHideOtherApplications(app.native, s)
}

func (app *NSApplication) ReplyToApplicationShouldTerminate(shouldTerminate bool) {
	C.nsApplicationReplyToApplicationShouldTerminate(app.native, C.bool(shouldTerminate))
}

func (app *NSApplication) Terminate(sender *NSApplication) {
	var s C.NSApplicationPtr
	if sender != nil {
		s = sender.native
	}
	C.nsApplicationTerminate(app.native, s)
}

func (app *NSApplication) SetMainMenu(menu *NSMenu) {
	C.nsApplicationSetMainMenu(app.native, menu.native)
}

func (app *NSApplication) SetServicesMenu(menu *NSMenu) {
	C.nsApplicationSetServicesMenu(app.native, menu.native)
}

func (app *NSApplication) SetWindowsMenu(menu *NSMenu) {
	C.nsApplicationSetWindowsMenu(app.native, menu.native)
}

func (app *NSApplication) SetHelpMenu(menu *NSMenu) {
	C.nsApplicationSetHelpMenu(app.native, menu.native)
}

type NSRunningApplication struct {
	native C.NSRunningApplicationPtr
}

func NSRunningApplicationCurrent() *NSRunningApplication {
	return &NSRunningApplication{native: C.nsRunningApplicationCurrent()}
}

func (app *NSRunningApplication) ActivateWithOptions(options NSApplicationActivationOptions) {
	C.nsRunningApplicationActivateWithOptions(app.native, C.NSApplicationActivationOptions(options))
}

func (app *NSRunningApplication) Hide() bool {
	return bool(C.nsRunningApplicationHide(app.native))
}
