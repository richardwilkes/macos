// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ns

import "github.com/richardwilkes/macos/cf"

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

type OpenPanel struct {
	native C.NSOpenPanelPtr
}

func NewOpenPanel() *OpenPanel {
	return &OpenPanel{native: C.nsOpenPanel()}
}

func (p *OpenPanel) AccessoryView() *View {
	if view := C.nsOpenPanelAccessoryView(p.native); view != nil {
		return &View{native: view}
	}
	return nil
}

func (p *OpenPanel) SetAccessoryView(view *View) {
	var v C.NSViewPtr
	if view != nil {
		v = view.native
	}
	C.nsOpenPanelSetAccessoryView(p.native, v)
}

func (p *OpenPanel) DirectoryURL() string {
	url := &URL{native: C.nsOpenPanelDirectoryURL(p.native)}
	return url.ResolveFilePath()
}

func (p *OpenPanel) SetDirectoryURL(url string) {
	nsurl := URLWithString(url)
	C.nsOpenPanelSetDirectoryURL(p.native, nsurl.native)
}

func (p *OpenPanel) AllowedFileTypes() []string {
	if types := cf.Array(C.nsOpenPanelAllowedFileTypes(p.native)); types != 0 {
		if count := types.GetCount(); count > 0 {
			result := make([]string, count)
			for i := 0; i < count; i++ {
				result[i] = cf.String(types.GetValueAtIndex(i)).String()
			}
			return result
		}
	}
	return nil
}

func (p *OpenPanel) SetAllowedFileTypes(types []string) {
	var a *cf.MutableArray
	if len(types) != 0 {
		a = cf.MutableArrayCreate(len(types))
		for _, t := range types {
			a.AppendStringValue(t)
		}
		C.nsOpenPanelSetAllowedFileTypes(p.native, C.CFArrayRef(a.AsCFArray()))
	} else {
		C.nsOpenPanelSetAllowedFileTypes(p.native, 0)
	}
}

func (p *OpenPanel) CanChooseFiles() bool {
	return bool(C.nsOpenPanelCanChooseFiles(p.native))
}

func (p *OpenPanel) SetCanChooseFiles(canChoose bool) {
	C.nsOpenPanelSetCanChooseFiles(p.native, C.bool(canChoose))
}

func (p *OpenPanel) CanChooseDirectories() bool {
	return bool(C.nsOpenPanelCanChooseDirectories(p.native))
}

func (p *OpenPanel) SetCanChooseDirectories(canChoose bool) {
	C.nsOpenPanelSetCanChooseDirectories(p.native, C.bool(canChoose))
}

func (p *OpenPanel) ResolvesAliases() bool {
	return bool(C.nsOpenPanelResolvesAliases(p.native))
}

func (p *OpenPanel) SetResolvesAliases(resolves bool) {
	C.nsOpenPanelSetResolvesAliases(p.native, C.bool(resolves))
}

func (p *OpenPanel) AllowsMultipleSelection() bool {
	return bool(C.nsOpenPanelAllowsMultipleSelection(p.native))
}

func (p *OpenPanel) SetAllowsMultipleSelection(allow bool) {
	C.nsOpenPanelSetAllowsMultipleSelection(p.native, C.bool(allow))
}

func (p *OpenPanel) CanDownloadUbiquitousContents() bool {
	return bool(C.nsOpenPanelCanDownloadUbiquitousContents(p.native))
}

func (p *OpenPanel) SetCanDownloadUbiquitousContents(canDownload bool) {
	C.nsOpenPanelSetCanDownloadUbiquitousContents(p.native, C.bool(canDownload))
}

func (p *OpenPanel) CanResolveUbiquitousConflicts() bool {
	return bool(C.nsOpenPanelCanResolveUbiquitousConflicts(p.native))
}

func (p *OpenPanel) SetCanResolveUbiquitousConflicts(canResolve bool) {
	C.nsOpenPanelSetCanResolveUbiquitousConflicts(p.native, C.bool(canResolve))
}

func (p *OpenPanel) AccessoryViewDisclosed() bool {
	return bool(C.nsOpenPanelAccessoryViewDisclosed(p.native))
}

func (p *OpenPanel) SetAccessoryViewDisclosed(disclosed bool) {
	C.nsOpenPanelSetAccessoryViewDisclosed(p.native, C.bool(disclosed))
}

func (p *OpenPanel) URLs() []string {
	urlRefs := cf.Array(C.nsOpenPanelURLs(p.native))
	count := urlRefs.GetCount()
	result := make([]string, count)
	for i := 0; i < count; i++ {
		url := &URL{native: C.NSURLPtr(urlRefs.GetValueAtIndex(i))}
		result[i] = url.ResolveFilePath()
	}
	return result
}

func (p *OpenPanel) RunModal() bool {
	return bool(C.nsOpenPanelRunModal(p.native))
}
