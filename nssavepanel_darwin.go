package macos

/*
#import <Cocoa/Cocoa.h>
#import <CoreFoundation/CoreFoundation.h>

typedef void *NSSavePanelPtr;
typedef void *NSViewPtr;
typedef void *NSURLPtr;

NSSavePanelPtr nsSavePanel() {
	return (NSSavePanelPtr)[NSSavePanel savePanel];
}

NSViewPtr nsSavePanelAccessoryView(NSSavePanelPtr savePanel) {
	return [(NSSavePanel *)savePanel accessoryView];
}

void nsSavePanelSetAccessoryView(NSSavePanelPtr savePanel, NSViewPtr view) {
	[(NSSavePanel *)savePanel setAccessoryView:view];
}

NSURLPtr nsSavePanelDirectoryURL(NSSavePanelPtr savePanel) {
	return [(NSSavePanel *)savePanel directoryURL];
}

void nsSavePanelSetDirectoryURL(NSSavePanelPtr savePanel, NSURLPtr url) {
	[(NSSavePanel *)savePanel setDirectoryURL:url];
}

CFArrayRef nsSavePanelAllowedFileTypes(NSSavePanelPtr savePanel) {
	return (CFArrayRef)([(NSSavePanel *)savePanel allowedFileTypes]);
}

void nsSavePanelSetAllowedFileTypes(NSSavePanelPtr savePanel, CFArrayRef types) {
	[(NSSavePanel *)savePanel setAllowedFileTypes:(NSArray<NSString *>*)(types)];
}

NSURLPtr nsSavePanelURL(NSSavePanelPtr savePanel) {
	return [(NSSavePanel *)savePanel URL];
}

bool nsSavePanelRunModal(NSSavePanelPtr savePanel) {
	return [(NSSavePanel *)savePanel runModal] == NSModalResponseOK;
}
*/
import "C"

type NSSavePanel struct {
	native C.NSSavePanelPtr
}

func NewSavePanel() *NSSavePanel {
	return &NSSavePanel{native: C.nsSavePanel()}
}

func (p *NSSavePanel) AccessoryView() *NSView {
	if view := C.nsSavePanelAccessoryView(p.native); view != nil {
		return &NSView{native: view}
	}
	return nil
}

func (p *NSSavePanel) SetAccessoryView(view *NSView) {
	var v C.NSViewPtr
	if view != nil {
		v = view.native
	}
	C.nsSavePanelSetAccessoryView(p.native, v)
}

func (p *NSSavePanel) DirectoryURL() string {
	url := &NSURL{native: C.nsSavePanelDirectoryURL(p.native)}
	return url.ResolveFilePath()
}

func (p *NSSavePanel) SetDirectoryURL(url string) {
	nsurl := NSURLWithString(url)
	C.nsSavePanelSetDirectoryURL(p.native, nsurl.native)
}

func (p *NSSavePanel) AllowedFileTypes() []string {
	if types := CFArray(C.nsSavePanelAllowedFileTypes(p.native)); types != 0 {
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

func (p *NSSavePanel) SetAllowedFileTypes(types []string) {
	var a *CFMutableArray
	if len(types) != 0 {
		a = CFMutableArrayCreate(len(types))
		for _, t := range types {
			a.AppendStringValue(t)
		}
		C.nsSavePanelSetAllowedFileTypes(p.native, a.AsCFArray())
	} else {
		C.nsSavePanelSetAllowedFileTypes(p.native, 0)
	}
}

func (p *NSSavePanel) URL() string {
	url := &NSURL{native: C.nsSavePanelURL(p.native)}
	return url.ResolveFilePath()
}

func (p *NSSavePanel) RunModal() bool {
	return bool(C.nsSavePanelRunModal(p.native))
}
