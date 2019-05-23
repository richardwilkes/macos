package macos

// #import <CoreText/CoreText.h>
import "C"

const (
	CTFontUIFontUser CTFontUIFontType = iota
	CTFontUIFontUserFixedPitch
	CTFontUIFontSystem
	CTFontUIFontEmphasizedSystem
	CTFontUIFontSmallSystem
	CTFontUIFontSmallEmphasizedSystem
	CTFontUIFontMiniSystem
	CTFontUIFontMiniEmphasizedSystem
	CTFontUIFontViews
	CTFontUIFontApplication
	CTFontUIFontLabel
	CTFontUIFontMenuTitle
	CTFontUIFontMenuItem
	CTFontUIFontMenuItemMark
	CTFontUIFontMenuItemCmdKey
	CTFontUIFontWindowTitle
	CTFontUIFontPushButton
	CTFontUIFontUtilityWindowTitle
	CTFontUIFontAlertHeader
	CTFontUIFontSystemDetail
	CTFontUIFontEmphasizedSystemDetail
	CTFontUIFontToolbar
	CTFontUIFontSmallToolbar
	CTFontUIFontMessage
	CTFontUIFontPalette
	CTFontUIFontToolTip
	CTFontUIFontControlContent
	CTFontUIFontNone = ^CTFontUIFontType(0)
)

const (
	CTFontClassMaskShift                        = 28
	CTFontTraitItalic      CTFontSymbolicTraits = 1 << 0
	CTFontTraitBold        CTFontSymbolicTraits = 1 << 1
	CTFontTraitExpanded    CTFontSymbolicTraits = 1 << 5
	CTFontTraitCondensed   CTFontSymbolicTraits = 1 << 6
	CTFontTraitMonoSpace   CTFontSymbolicTraits = 1 << 10
	CTFontTraitVertical    CTFontSymbolicTraits = 1 << 11
	CTFontTraitUIOptimized CTFontSymbolicTraits = 1 << 12
	CTFontTraitColorGlyphs CTFontSymbolicTraits = 1 << 13
	CTFontTraitComposite   CTFontSymbolicTraits = 1 << 14
	CTFontTraitClassMask   CTFontSymbolicTraits = 15 << CTFontClassMaskShift
	CTFontItalicTrait                           = CTFontTraitItalic
	CTFontBoldTrait                             = CTFontTraitBold
	CTFontExpandedTrait                         = CTFontTraitExpanded
	CTFontCondensedTrait                        = CTFontTraitCondensed
	CTFontMonoSpaceTrait                        = CTFontTraitMonoSpace
	CTFontVerticalTrait                         = CTFontTraitVertical
	CTFontUIOptimizedTrait                      = CTFontTraitUIOptimized
	CTFontColorGlyphsTrait                      = CTFontTraitColorGlyphs
	CTFontCompositeTrait                        = CTFontTraitComposite
	CTFontClassMaskTrait                        = CTFontTraitClassMask
)

var CTFontFamilyNameAttribute = C.kCTFontFamilyNameAttribute

type (
	CTFont               = C.CTFontRef
	CTFontUIFontType     = uint32
	CTFontSymbolicTraits = uint32
)

func CTFontCreateWithName(name string, size float64, matrix *CGAffineTransform) CTFont {
	nameStr := CFStringCreateWithString(name)
	defer nameStr.Release()
	return C.CTFontCreateWithName(nameStr, C.CGFloat(size), matrix)
}

func CTFontCreateUIFontForLanguage(uiType CTFontUIFontType, size float64, language string) CTFont {
	var lang CFString
	if language != "" {
		lang = CFStringCreateWithString(language)
		defer lang.Release()
	}
	return C.CTFontCreateUIFontForLanguage(C.CTFontUIFontType(uiType), C.CGFloat(size), lang)
}

func (f CTFont) CreateCopyWithSymbolicTraits(size float64, matrix *CGAffineTransform, value, mask CTFontSymbolicTraits) CTFont {
	return C.CTFontCreateCopyWithSymbolicTraits(f, C.CGFloat(size), matrix, C.CTFontSymbolicTraits(value), C.CTFontSymbolicTraits(mask))
}

func (f CTFont) FamilyName() string {
	str := C.CTFontCopyFamilyName(f)
	defer str.Release()
	return str.String()
}

func (f CTFont) GetAscent() float64 {
	return float64(C.CTFontGetAscent(f))
}

func (f CTFont) GetDescent() float64 {
	return float64(C.CTFontGetDescent(f))
}

func (f CTFont) GetLeading() float64 {
	return float64(C.CTFontGetLeading(f))
}

func (f CTFont) GetSize() float64 {
	return float64(C.CTFontGetSize(f))
}

func (f CTFont) GetSymbolicTraits() CTFontSymbolicTraits {
	return CTFontSymbolicTraits(C.CTFontGetSymbolicTraits(f))
}

func (f CTFont) Retain() {
	C.CFRetain(CFType(f))
}

func (f CTFont) Release() {
	C.CFRelease(CFType(f))
}
