// Code generated from "tmpl/nscursor.go.tmpl" - DO NOT EDIT.
package macos

/*
#import <Cocoa/Cocoa.h>

typedef void *NSCursorPtr;

NSCursorPtr getarrowCursor() {
	return [NSCursor arrowCursor];
}

NSCursorPtr getIBeamCursor() {
	return [NSCursor IBeamCursor];
}

NSCursorPtr getIBeamCursorForVerticalLayout() {
	return [NSCursor IBeamCursorForVerticalLayout];
}

NSCursorPtr getcrosshairCursor() {
	return [NSCursor crosshairCursor];
}

NSCursorPtr getclosedHandCursor() {
	return [NSCursor closedHandCursor];
}

NSCursorPtr getopenHandCursor() {
	return [NSCursor openHandCursor];
}

NSCursorPtr getpointingHandCursor() {
	return [NSCursor pointingHandCursor];
}

NSCursorPtr getresizeLeftCursor() {
	return [NSCursor resizeLeftCursor];
}

NSCursorPtr getresizeRightCursor() {
	return [NSCursor resizeRightCursor];
}

NSCursorPtr getresizeLeftRightCursor() {
	return [NSCursor resizeLeftRightCursor];
}

NSCursorPtr getresizeUpCursor() {
	return [NSCursor resizeUpCursor];
}

NSCursorPtr getresizeDownCursor() {
	return [NSCursor resizeDownCursor];
}

NSCursorPtr getresizeUpDownCursor() {
	return [NSCursor resizeUpDownCursor];
}

NSCursorPtr getdisappearingItemCursor() {
	return [NSCursor disappearingItemCursor];
}

NSCursorPtr getoperationNotAllowedCursor() {
	return [NSCursor operationNotAllowedCursor];
}

NSCursorPtr getdragLinkCursor() {
	return [NSCursor dragLinkCursor];
}

NSCursorPtr getdragCopyCursor() {
	return [NSCursor dragCopyCursor];
}

NSCursorPtr getcontextualMenuCursor() {
	return [NSCursor contextualMenuCursor];
}
*/
import "C"

func NSCursorArrowCursor() *NSCursor {
	return &NSCursor{native: C.getarrowCursor()}
}

func NSCursorIBeamCursor() *NSCursor {
	return &NSCursor{native: C.getIBeamCursor()}
}

func NSCursorIBeamCursorForVerticalLayout() *NSCursor {
	return &NSCursor{native: C.getIBeamCursorForVerticalLayout()}
}

func NSCursorCrosshairCursor() *NSCursor {
	return &NSCursor{native: C.getcrosshairCursor()}
}

func NSCursorClosedHandCursor() *NSCursor {
	return &NSCursor{native: C.getclosedHandCursor()}
}

func NSCursorOpenHandCursor() *NSCursor {
	return &NSCursor{native: C.getopenHandCursor()}
}

func NSCursorPointingHandCursor() *NSCursor {
	return &NSCursor{native: C.getpointingHandCursor()}
}

func NSCursorResizeLeftCursor() *NSCursor {
	return &NSCursor{native: C.getresizeLeftCursor()}
}

func NSCursorResizeRightCursor() *NSCursor {
	return &NSCursor{native: C.getresizeRightCursor()}
}

func NSCursorResizeLeftRightCursor() *NSCursor {
	return &NSCursor{native: C.getresizeLeftRightCursor()}
}

func NSCursorResizeUpCursor() *NSCursor {
	return &NSCursor{native: C.getresizeUpCursor()}
}

func NSCursorResizeDownCursor() *NSCursor {
	return &NSCursor{native: C.getresizeDownCursor()}
}

func NSCursorResizeUpDownCursor() *NSCursor {
	return &NSCursor{native: C.getresizeUpDownCursor()}
}

func NSCursorDisappearingItemCursor() *NSCursor {
	return &NSCursor{native: C.getdisappearingItemCursor()}
}

func NSCursorOperationNotAllowedCursor() *NSCursor {
	return &NSCursor{native: C.getoperationNotAllowedCursor()}
}

func NSCursorDragLinkCursor() *NSCursor {
	return &NSCursor{native: C.getdragLinkCursor()}
}

func NSCursorDragCopyCursor() *NSCursor {
	return &NSCursor{native: C.getdragCopyCursor()}
}

func NSCursorContextualMenuCursor() *NSCursor {
	return &NSCursor{native: C.getcontextualMenuCursor()}
}
