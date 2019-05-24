package main

//go:generate go run codegen_darwin.go

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

var (
	systemColorNames = []string{
		"alternateSelectedControlTextColor",
		"controlAccentColor",
		"controlBackgroundColor",
		"controlColor",
		"controlTextColor",
		"disabledControlTextColor",
		"findHighlightColor",
		"gridColor",
		"headerTextColor",
		"highlightColor",
		"keyboardFocusIndicatorColor",
		"labelColor",
		"linkColor",
		"placeholderTextColor",
		"quaternaryLabelColor",
		"secondaryLabelColor",
		"selectedContentBackgroundColor",
		"selectedControlColor",
		"selectedControlTextColor",
		"selectedMenuItemTextColor",
		"selectedTextBackgroundColor",
		"selectedTextColor",
		"separatorColor",
		"shadowColor",
		"systemBlueColor",
		"systemBrownColor",
		"systemGrayColor",
		"systemGreenColor",
		"systemOrangeColor",
		"systemPinkColor",
		"systemPurpleColor",
		"systemRedColor",
		"systemYellowColor",
		"tertiaryLabelColor",
		"textBackgroundColor",
		"textColor",
		"underPageBackgroundColor",
		"unemphasizedSelectedContentBackgroundColor",
		"unemphasizedSelectedTextBackgroundColor",
		"unemphasizedSelectedTextColor",
		"windowBackgroundColor",
		"windowFrameTextColor",
	}
	systemCursorNames = []string{
		"arrowCursor",
		"IBeamCursor",
		"IBeamCursorForVerticalLayout",
		"crosshairCursor",
		"closedHandCursor",
		"openHandCursor",
		"pointingHandCursor",
		"resizeLeftCursor",
		"resizeRightCursor",
		"resizeLeftRightCursor",
		"resizeUpCursor",
		"resizeDownCursor",
		"resizeUpDownCursor",
		"disappearingItemCursor",
		"operationNotAllowedCursor",
		"dragLinkCursor",
		"dragCopyCursor",
		"contextualMenuCursor",
	}
)

func main() {
	dir, err := os.Open("..")
	fatalIfErr(err)
	fis, err := dir.Readdir(-1)
	fatalIfErr(err)
	fatalIfErr(dir.Close())
	for _, fi := range fis {
		if !fi.IsDir() && strings.HasSuffix(strings.ToLower(fi.Name()), "_gen.go") {
			fatalIfErr(os.Remove(filepath.Join("..", fi.Name())))
		}
	}
	processTemplate("nscolor", systemColorNames)
	processTemplate("nscursor", systemCursorNames)
}

func processTemplate(name string, arg interface{}) {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "// Code generated from \"tmpl/%s.go.tmpl\" - DO NOT EDIT.\n", name)
	tmpl, err := template.New(name + ".go.tmpl").Funcs(template.FuncMap{
		"firstToUpper": firstToUpper,
	}).ParseFiles(name + ".go.tmpl")
	fatalIfErr(err)
	fatalIfErr(tmpl.Execute(&buffer, arg))
	data, err := format.Source(buffer.Bytes())
	fatalIfErr(err)
	fatalIfErr(ioutil.WriteFile(filepath.Join("..", name+"_gen_darwin.go"), data, 0644))
}

func firstToUpper(in string) string {
	return string(unicode.ToUpper(rune(in[0]))) + in[1:]
}

func fatalIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
