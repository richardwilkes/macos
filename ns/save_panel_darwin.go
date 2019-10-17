package ns

import "github.com/richardwilkes/macos/cf"

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

type SavePanel struct {
	native C.NSSavePanelPtr
}

func NewSavePanel() *SavePanel {
	return &SavePanel{native: C.nsSavePanel()}
}

func (p *SavePanel) AccessoryView() *View {
	if view := C.nsSavePanelAccessoryView(p.native); view != nil {
		return &View{native: view}
	}
	return nil
}

func (p *SavePanel) SetAccessoryView(view *View) {
	var v C.NSViewPtr
	if view != nil {
		v = view.native
	}
	C.nsSavePanelSetAccessoryView(p.native, v)
}

func (p *SavePanel) DirectoryURL() string {
	url := &URL{native: C.nsSavePanelDirectoryURL(p.native)}
	return url.ResolveFilePath()
}

func (p *SavePanel) SetDirectoryURL(url string) {
	nsurl := URLWithString(url)
	C.nsSavePanelSetDirectoryURL(p.native, nsurl.native)
}

func (p *SavePanel) AllowedFileTypes() []string {
	if types := cf.Array(C.nsSavePanelAllowedFileTypes(p.native)); types != 0 {
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

func (p *SavePanel) SetAllowedFileTypes(types []string) {
	var a *cf.MutableArray
	if len(types) != 0 {
		a = cf.MutableArrayCreate(len(types))
		for _, t := range types {
			a.AppendStringValue(t)
		}
		C.nsSavePanelSetAllowedFileTypes(p.native, C.CFArrayRef(a.AsCFArray()))
	} else {
		C.nsSavePanelSetAllowedFileTypes(p.native, 0)
	}
}

func (p *SavePanel) URL() string {
	url := &URL{native: C.nsSavePanelURL(p.native)}
	return url.ResolveFilePath()
}

func (p *SavePanel) RunModal() bool {
	return bool(C.nsSavePanelRunModal(p.native))
}
