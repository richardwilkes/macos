package ns

/*
#import <Cocoa/Cocoa.h>

typedef void *NSApplicationPtr;
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

bool nsApplicationActive(NSApplicationPtr app) {
	return [(NSApplication *)app isActive];
}

void nsApplicationSetActivationPolicy(NSApplicationPtr app, NSApplicationActivationPolicy policy) {
	[(NSApplication *)app setActivationPolicy:policy];
}

void nsApplicationActivateIgnoringOtherApps(NSApplicationPtr app, bool force) {
	[(NSApplication *)app activateIgnoringOtherApps:force];
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

NSModalResponse nsApplicationRunModalForWindow(NSApplicationPtr app, NSWindowPtr window) {
	return [(NSApplication *)app runModalForWindow:(NSWindow *)window];
}

void nsApplicationStopModalWithCode(NSApplicationPtr app, NSModalResponse code) {
	return [(NSApplication *)app stopModalWithCode:code];
}
*/
import "C"

const (
	TerminateCancel ApplicationTerminateReply = iota
	TerminateNow
	TerminateLater
)

const (
	ApplicationActivationPolicyRegular ApplicationActivationPolicy = iota
	ApplicationActivationPolicyAccessory
	ApplicationActivationPolicyProhibited
)

const (
	ApplicationActivateAllWindows ApplicationActivationOptions = 1 << iota
	ApplicationActivateIgnoringOtherApps
)

var currentAppDelegate ApplicationDelegate

type (
	ApplicationTerminateReply    int
	ApplicationActivationPolicy  int
	ApplicationActivationOptions int
)

type ApplicationDelegate interface {
	ApplicationWillFinishLaunching(notification *Notification)
	ApplicationDidFinishLaunching(notification *Notification)
	ApplicationShouldTerminate(sender *Application) ApplicationTerminateReply
	ApplicationShouldTerminateAfterLastWindowClosed(app *Application) bool
	ApplicationWillTerminate(notification *Notification)
	ApplicationWillBecomeActive(notification *Notification)
	ApplicationDidBecomeActive(notification *Notification)
	ApplicationWillResignActive(notification *Notification)
	ApplicationDidResignActive(notification *Notification)
	ThemeChanged(notification *Notification)
}

type Application struct {
	native C.NSApplicationPtr
}

func SharedApplication() *Application {
	return &Application{native: C.nsSharedApplication()}
}

func (app *Application) SetDelegate(delegate ApplicationDelegate) {
	currentAppDelegate = delegate
	C.nsApplicationSetDelegate(app.native)
}

func (app *Application) Active() bool {
	return bool(C.nsApplicationActive(app.native))
}

func (app *Application) SetActivationPolicy(policy ApplicationActivationPolicy) {
	C.nsApplicationSetActivationPolicy(app.native, C.NSApplicationActivationPolicy(policy))
}

func (app *Application) ActivateIgnoringOtherApps(force bool) {
	C.nsApplicationActivateIgnoringOtherApps(app.native, C.bool(force))
}

func (app *Application) KeyWindow() *Window {
	if w := C.nsApplicationKeyWindow(app.native); w != nil {
		return &Window{native: w}
	}
	return nil
}

func (app *Application) MainWindow() *Window {
	if w := C.nsApplicationMainWindow(app.native); w != nil {
		return &Window{native: w}
	}
	return nil
}

func (app *Application) Run() {
	C.nsApplicationRun(app.native)
}

func (app *Application) RunModalForWindow(wnd *Window) int {
	return int(C.nsApplicationRunModalForWindow(app.native, wnd.native))
}

func (app *Application) StopModalWithCode(code int) {
	C.nsApplicationStopModalWithCode(app.native, C.NSModalResponse(code))
}

func (app *Application) UnhideAllApplications(sender *Application) {
	var s C.NSApplicationPtr
	if sender != nil {
		s = sender.native
	}
	C.nsApplicationUnhideAllApplications(app.native, s)
}

func (app *Application) HideOtherApplications(sender *Application) {
	var s C.NSApplicationPtr
	if sender != nil {
		s = sender.native
	}
	C.nsApplicationHideOtherApplications(app.native, s)
}

func (app *Application) ReplyToApplicationShouldTerminate(shouldTerminate bool) {
	C.nsApplicationReplyToApplicationShouldTerminate(app.native, C.bool(shouldTerminate))
}

func (app *Application) Terminate(sender *Application) {
	var s C.NSApplicationPtr
	if sender != nil {
		s = sender.native
	}
	C.nsApplicationTerminate(app.native, s)
}

func (app *Application) SetMainMenu(menu *Menu) {
	C.nsApplicationSetMainMenu(app.native, menu.native)
}

func (app *Application) SetServicesMenu(menu *Menu) {
	C.nsApplicationSetServicesMenu(app.native, menu.native)
}

func (app *Application) SetWindowsMenu(menu *Menu) {
	C.nsApplicationSetWindowsMenu(app.native, menu.native)
}

func (app *Application) SetHelpMenu(menu *Menu) {
	C.nsApplicationSetHelpMenu(app.native, menu.native)
}
