package macos

/*
#import <Cocoa/Cocoa.h>
#import <CoreFoundation/CoreFoundation.h>

typedef void *NSOpenPanelPtr;
typedef void *NSViewPtr;
typedef void *NSURLPtr;

NSOpenPanelPtr nsOpenPanel() {
	return (NSOpenPanelPtr)[NSOpenPanel openPanel];
}

NSViewPtr nsOpenPanelAccessoryView(NSOpenPanelPtr openPanel) {
	return [(NSOpenPanel *)openPanel accessoryView];
}

void nsOpenPanelSetAccessoryView(NSOpenPanelPtr openPanel, NSViewPtr view) {
	[(NSOpenPanel *)openPanel setAccessoryView:view];
}

NSURLPtr nsOpenPanelDirectoryURL(NSOpenPanelPtr openPanel) {
	return [(NSOpenPanel *)openPanel directoryURL];
}

void nsOpenPanelSetDirectoryURL(NSOpenPanelPtr openPanel, NSURLPtr url) {
	[(NSOpenPanel *)openPanel setDirectoryURL:url];
}

CFArrayRef nsOpenPanelAllowedFileTypes(NSOpenPanelPtr openPanel) {
	return (CFArrayRef)([(NSOpenPanel *)openPanel allowedFileTypes]);
}

void nsOpenPanelSetAllowedFileTypes(NSOpenPanelPtr openPanel, CFArrayRef types) {
	[(NSOpenPanel *)openPanel setAllowedFileTypes:(NSArray<NSString *>*)(types)];
}

bool nsOpenPanelCanChooseFiles(NSOpenPanelPtr openPanel) {
	return [(NSOpenPanel *)openPanel canChooseFiles];
}

void nsOpenPanelSetCanChooseFiles(NSOpenPanelPtr openPanel, bool set) {
	[(NSOpenPanel *)openPanel setCanChooseFiles:set];
}

bool nsOpenPanelCanChooseDirectories(NSOpenPanelPtr openPanel) {
	return [(NSOpenPanel *)openPanel canChooseDirectories];
}

void nsOpenPanelSetCanChooseDirectories(NSOpenPanelPtr openPanel, bool set) {
	[(NSOpenPanel *)openPanel setCanChooseDirectories:set];
}

bool nsOpenPanelResolvesAliases(NSOpenPanelPtr openPanel) {
	return [(NSOpenPanel *)openPanel resolvesAliases];
}

void nsOpenPanelSetResolvesAliases(NSOpenPanelPtr openPanel, bool set) {
	[(NSOpenPanel *)openPanel setResolvesAliases:set];
}

bool nsOpenPanelAllowsMultipleSelection(NSOpenPanelPtr openPanel) {
	return [(NSOpenPanel *)openPanel allowsMultipleSelection];
}

void nsOpenPanelSetAllowsMultipleSelection(NSOpenPanelPtr openPanel, bool set) {
	[(NSOpenPanel *)openPanel setAllowsMultipleSelection:set];
}

bool nsOpenPanelCanDownloadUbiquitousContents(NSOpenPanelPtr openPanel) {
	return [(NSOpenPanel *)openPanel canDownloadUbiquitousContents];
}

void nsOpenPanelSetCanDownloadUbiquitousContents(NSOpenPanelPtr openPanel, bool set) {
	[(NSOpenPanel *)openPanel setCanDownloadUbiquitousContents:set];
}

bool nsOpenPanelCanResolveUbiquitousConflicts(NSOpenPanelPtr openPanel) {
	return [(NSOpenPanel *)openPanel canResolveUbiquitousConflicts];
}

void nsOpenPanelSetCanResolveUbiquitousConflicts(NSOpenPanelPtr openPanel, bool set) {
	[(NSOpenPanel *)openPanel setCanResolveUbiquitousConflicts:set];
}

bool nsOpenPanelAccessoryViewDisclosed(NSOpenPanelPtr openPanel) {
	return [(NSOpenPanel *)openPanel isAccessoryViewDisclosed];
}

void nsOpenPanelSetAccessoryViewDisclosed(NSOpenPanelPtr openPanel, bool set) {
	[(NSOpenPanel *)openPanel setAccessoryViewDisclosed:set];
}

CFArrayRef nsOpenPanelURLs(NSOpenPanelPtr openPanel) {
	return (CFArrayRef)[(NSOpenPanel *)openPanel URLs];
}

bool nsOpenPanelRunModal(NSOpenPanelPtr openPanel) {
	return [(NSOpenPanel *)openPanel runModal] == NSModalResponseOK;
}
*/
import "C"

type NSOpenPanel struct {
	native C.NSOpenPanelPtr
}

func NewOpenPanel() *NSOpenPanel {
	return &NSOpenPanel{native: C.nsOpenPanel()}
}

func (p *NSOpenPanel) AccessoryView() *NSView {
	if view := C.nsOpenPanelAccessoryView(p.native); view != nil {
		return &NSView{native: view}
	}
	return nil
}

func (p *NSOpenPanel) SetAccessoryView(view *NSView) {
	var v C.NSViewPtr
	if view != nil {
		v = view.native
	}
	C.nsOpenPanelSetAccessoryView(p.native, v)
}

func (p *NSOpenPanel) DirectoryURL() string {
	url := &NSURL{native: C.nsOpenPanelDirectoryURL(p.native)}
	return url.ResolveFilePath()
}

func (p *NSOpenPanel) SetDirectoryURL(url string) {
	nsurl := NSURLWithString(url)
	C.nsOpenPanelSetDirectoryURL(p.native, nsurl.native)
}

func (p *NSOpenPanel) AllowedFileTypes() []string {
	if types := CFArray(C.nsOpenPanelAllowedFileTypes(p.native)); types != 0 {
		if count := types.GetCount(); count > 0 {
			result := make([]string, count)
			for i := 0; i < count; i++ {
				result[i] = CFString(types.GetValueAtIndex(i)).String()
			}
			return result
		}
	}
	return nil
}

func (p *NSOpenPanel) SetAllowedFileTypes(types []string) {
	var a *CFMutableArray
	if len(types) != 0 {
		a = CFMutableArrayCreate(len(types))
		for _, t := range types {
			a.AppendStringValue(t)
		}
		C.nsOpenPanelSetAllowedFileTypes(p.native, a.AsCFArray())
	} else {
		C.nsOpenPanelSetAllowedFileTypes(p.native, 0)
	}
}

func (p *NSOpenPanel) CanChooseFiles() bool {
	return bool(C.nsOpenPanelCanChooseFiles(p.native))
}

func (p *NSOpenPanel) SetCanChooseFiles(canChoose bool) {
	C.nsOpenPanelSetCanChooseFiles(p.native, C.bool(canChoose))
}

func (p *NSOpenPanel) CanChooseDirectories() bool {
	return bool(C.nsOpenPanelCanChooseDirectories(p.native))
}

func (p *NSOpenPanel) SetCanChooseDirectories(canChoose bool) {
	C.nsOpenPanelSetCanChooseDirectories(p.native, C.bool(canChoose))
}

func (p *NSOpenPanel) ResolvesAliases() bool {
	return bool(C.nsOpenPanelResolvesAliases(p.native))
}

func (p *NSOpenPanel) SetResolvesAliases(resolves bool) {
	C.nsOpenPanelSetResolvesAliases(p.native, C.bool(resolves))
}

func (p *NSOpenPanel) AllowsMultipleSelection() bool {
	return bool(C.nsOpenPanelAllowsMultipleSelection(p.native))
}

func (p *NSOpenPanel) SetAllowsMultipleSelection(allow bool) {
	C.nsOpenPanelSetAllowsMultipleSelection(p.native, C.bool(allow))
}

func (p *NSOpenPanel) CanDownloadUbiquitousContents() bool {
	return bool(C.nsOpenPanelCanDownloadUbiquitousContents(p.native))
}

func (p *NSOpenPanel) SetCanDownloadUbiquitousContents(canDownload bool) {
	C.nsOpenPanelSetCanDownloadUbiquitousContents(p.native, C.bool(canDownload))
}

func (p *NSOpenPanel) CanResolveUbiquitousConflicts() bool {
	return bool(C.nsOpenPanelCanResolveUbiquitousConflicts(p.native))
}

func (p *NSOpenPanel) SetCanResolveUbiquitousConflicts(canResolve bool) {
	C.nsOpenPanelSetCanResolveUbiquitousConflicts(p.native, C.bool(canResolve))
}

func (p *NSOpenPanel) AccessoryViewDisclosed() bool {
	return bool(C.nsOpenPanelAccessoryViewDisclosed(p.native))
}

func (p *NSOpenPanel) SetAccessoryViewDisclosed(disclosed bool) {
	C.nsOpenPanelSetAccessoryViewDisclosed(p.native, C.bool(disclosed))
}

func (p *NSOpenPanel) URLs() []string {
	urlRefs := CFArray(C.nsOpenPanelURLs(p.native))
	count := urlRefs.GetCount()
	result := make([]string, count)
	for i := 0; i < count; i++ {
		url := &NSURL{native: C.NSURLPtr(urlRefs.GetValueAtIndex(i))}
		result[i] = url.ResolveFilePath()
	}
	return result
}

func (p *NSOpenPanel) RunModal() bool {
	return bool(C.nsOpenPanelRunModal(p.native))
}
