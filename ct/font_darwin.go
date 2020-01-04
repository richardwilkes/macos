// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ct

import (
	"unsafe"

	"github.com/richardwilkes/macos/cf"
	"github.com/richardwilkes/macos/cg"
)

// #import <CoreText/CoreText.h>
import "C"

const (
	FontUIFontUser FontUIFontType = iota
	FontUIFontUserFixedPitch
	FontUIFontSystem
	FontUIFontEmphasizedSystem
	FontUIFontSmallSystem
	FontUIFontSmallEmphasizedSystem
	FontUIFontMiniSystem
	FontUIFontMiniEmphasizedSystem
	FontUIFontViews
	FontUIFontApplication
	FontUIFontLabel
	FontUIFontMenuTitle
	FontUIFontMenuItem
	FontUIFontMenuItemMark
	FontUIFontMenuItemCmdKey
	FontUIFontWindowTitle
	FontUIFontPushButton
	FontUIFontUtilityWindowTitle
	FontUIFontAlertHeader
	FontUIFontSystemDetail
	FontUIFontEmphasizedSystemDetail
	FontUIFontToolbar
	FontUIFontSmallToolbar
	FontUIFontMessage
	FontUIFontPalette
	FontUIFontToolTip
	FontUIFontControlContent
	FontUIFontNone = ^FontUIFontType(0)
)

const (
	FontClassMaskShift                      = 28
	FontTraitItalic      FontSymbolicTraits = 1 << 0
	FontTraitBold        FontSymbolicTraits = 1 << 1
	FontTraitExpanded    FontSymbolicTraits = 1 << 5
	FontTraitCondensed   FontSymbolicTraits = 1 << 6
	FontTraitMonoSpace   FontSymbolicTraits = 1 << 10
	FontTraitVertical    FontSymbolicTraits = 1 << 11
	FontTraitUIOptimized FontSymbolicTraits = 1 << 12
	FontTraitColorGlyphs FontSymbolicTraits = 1 << 13
	FontTraitComposite   FontSymbolicTraits = 1 << 14
	FontTraitClassMask   FontSymbolicTraits = 15 << FontClassMaskShift
	FontItalicTrait                         = FontTraitItalic
	FontBoldTrait                           = FontTraitBold
	FontExpandedTrait                       = FontTraitExpanded
	FontCondensedTrait                      = FontTraitCondensed
	FontMonoSpaceTrait                      = FontTraitMonoSpace
	FontVerticalTrait                       = FontTraitVertical
	FontUIOptimizedTrait                    = FontTraitUIOptimized
	FontColorGlyphsTrait                    = FontTraitColorGlyphs
	FontCompositeTrait                      = FontTraitComposite
	FontClassMaskTrait                      = FontTraitClassMask
)

var FontFamilyNameAttribute = cf.String(C.kCTFontFamilyNameAttribute)

type (
	Font               = C.CTFontRef
	FontUIFontType     = uint32
	FontSymbolicTraits = uint32
)

func FontCreateWithName(name string, size float64, matrix *cg.AffineTransform) Font {
	nameStr := cf.StringCreateWithString(name)
	defer nameStr.Release()
	return C.CTFontCreateWithName(C.CFStringRef(nameStr), C.CGFloat(size), (*C.CGAffineTransform)(unsafe.Pointer(matrix)))
}

func FontCreateUIFontForLanguage(uiType FontUIFontType, size float64, language string) Font {
	var lang cf.String
	if language != "" {
		lang = cf.StringCreateWithString(language)
		defer lang.Release()
	}
	return C.CTFontCreateUIFontForLanguage(C.CTFontUIFontType(uiType), C.CGFloat(size), C.CFStringRef(lang))
}

func FontCreateWithFontDescriptor(descriptor FontDescriptor, size float64, matrix *cg.AffineTransform) Font {
	return C.CTFontCreateWithFontDescriptor(C.CTFontDescriptorRef(descriptor), C.CGFloat(size), (*C.CGAffineTransform)(unsafe.Pointer(matrix))) //nolint:unconvert
}

func (f Font) CreateCopyWithSymbolicTraits(size float64, matrix *cg.AffineTransform, value, mask FontSymbolicTraits) Font {
	return C.CTFontCreateCopyWithSymbolicTraits(f, C.CGFloat(size), (*C.CGAffineTransform)(unsafe.Pointer(matrix)), C.CTFontSymbolicTraits(value), C.CTFontSymbolicTraits(mask))
}

func (f Font) FamilyName() string {
	str := cf.String(C.CTFontCopyFamilyName(f))
	defer str.Release()
	return str.String()
}

func (f Font) GetAscent() float64 {
	return float64(C.CTFontGetAscent(f))
}

func (f Font) GetDescent() float64 {
	return float64(C.CTFontGetDescent(f))
}

func (f Font) GetLeading() float64 {
	return float64(C.CTFontGetLeading(f))
}

func (f Font) GetSize() float64 {
	return float64(C.CTFontGetSize(f))
}

func (f Font) GetSymbolicTraits() FontSymbolicTraits {
	return FontSymbolicTraits(C.CTFontGetSymbolicTraits(f))
}

func (f Font) Retain() {
	C.CFRetain(C.CFTypeRef(f))
}

func (f Font) Release() {
	C.CFRelease(C.CFTypeRef(f))
}
