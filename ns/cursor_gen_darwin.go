// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ns

/*
#import <Cocoa/Cocoa.h>

typedef void *NSCursorPtr;

NSCursorPtr getArrowCursor() {
	return [NSCursor arrowCursor];
}

NSCursorPtr getIBeamCursor() {
	return [NSCursor IBeamCursor];
}

NSCursorPtr getIBeamCursorForVerticalLayout() {
	return [NSCursor IBeamCursorForVerticalLayout];
}

NSCursorPtr getCrosshairCursor() {
	return [NSCursor crosshairCursor];
}

NSCursorPtr getClosedHandCursor() {
	return [NSCursor closedHandCursor];
}

NSCursorPtr getOpenHandCursor() {
	return [NSCursor openHandCursor];
}

NSCursorPtr getPointingHandCursor() {
	return [NSCursor pointingHandCursor];
}

NSCursorPtr getResizeLeftCursor() {
	return [NSCursor resizeLeftCursor];
}

NSCursorPtr getResizeRightCursor() {
	return [NSCursor resizeRightCursor];
}

NSCursorPtr getResizeLeftRightCursor() {
	return [NSCursor resizeLeftRightCursor];
}

NSCursorPtr getResizeUpCursor() {
	return [NSCursor resizeUpCursor];
}

NSCursorPtr getResizeDownCursor() {
	return [NSCursor resizeDownCursor];
}

NSCursorPtr getResizeUpDownCursor() {
	return [NSCursor resizeUpDownCursor];
}

NSCursorPtr getDisappearingItemCursor() {
	return [NSCursor disappearingItemCursor];
}

NSCursorPtr getOperationNotAllowedCursor() {
	return [NSCursor operationNotAllowedCursor];
}

NSCursorPtr getDragLinkCursor() {
	return [NSCursor dragLinkCursor];
}

NSCursorPtr getDragCopyCursor() {
	return [NSCursor dragCopyCursor];
}

NSCursorPtr getContextualMenuCursor() {
	return [NSCursor contextualMenuCursor];
}
*/
import "C"

var (
	arrowCursor                  *Cursor
	iBeamCursor                  *Cursor
	iBeamCursorForVerticalLayout *Cursor
	crosshairCursor              *Cursor
	closedHandCursor             *Cursor
	openHandCursor               *Cursor
	pointingHandCursor           *Cursor
	resizeLeftCursor             *Cursor
	resizeRightCursor            *Cursor
	resizeLeftRightCursor        *Cursor
	resizeUpCursor               *Cursor
	resizeDownCursor             *Cursor
	resizeUpDownCursor           *Cursor
	disappearingItemCursor       *Cursor
	operationNotAllowedCursor    *Cursor
	dragLinkCursor               *Cursor
	dragCopyCursor               *Cursor
	contextualMenuCursor         *Cursor
)

func ArrowCursor() *Cursor {
	if arrowCursor == nil {
		arrowCursor = cursorInit(C.getArrowCursor())
	}
	return arrowCursor
}

func IBeamCursor() *Cursor {
	if iBeamCursor == nil {
		iBeamCursor = cursorInit(C.getIBeamCursor())
	}
	return iBeamCursor
}

func IBeamCursorForVerticalLayout() *Cursor {
	if iBeamCursorForVerticalLayout == nil {
		iBeamCursorForVerticalLayout = cursorInit(C.getIBeamCursorForVerticalLayout())
	}
	return iBeamCursorForVerticalLayout
}

func CrosshairCursor() *Cursor {
	if crosshairCursor == nil {
		crosshairCursor = cursorInit(C.getCrosshairCursor())
	}
	return crosshairCursor
}

func ClosedHandCursor() *Cursor {
	if closedHandCursor == nil {
		closedHandCursor = cursorInit(C.getClosedHandCursor())
	}
	return closedHandCursor
}

func OpenHandCursor() *Cursor {
	if openHandCursor == nil {
		openHandCursor = cursorInit(C.getOpenHandCursor())
	}
	return openHandCursor
}

func PointingHandCursor() *Cursor {
	if pointingHandCursor == nil {
		pointingHandCursor = cursorInit(C.getPointingHandCursor())
	}
	return pointingHandCursor
}

func ResizeLeftCursor() *Cursor {
	if resizeLeftCursor == nil {
		resizeLeftCursor = cursorInit(C.getResizeLeftCursor())
	}
	return resizeLeftCursor
}

func ResizeRightCursor() *Cursor {
	if resizeRightCursor == nil {
		resizeRightCursor = cursorInit(C.getResizeRightCursor())
	}
	return resizeRightCursor
}

func ResizeLeftRightCursor() *Cursor {
	if resizeLeftRightCursor == nil {
		resizeLeftRightCursor = cursorInit(C.getResizeLeftRightCursor())
	}
	return resizeLeftRightCursor
}

func ResizeUpCursor() *Cursor {
	if resizeUpCursor == nil {
		resizeUpCursor = cursorInit(C.getResizeUpCursor())
	}
	return resizeUpCursor
}

func ResizeDownCursor() *Cursor {
	if resizeDownCursor == nil {
		resizeDownCursor = cursorInit(C.getResizeDownCursor())
	}
	return resizeDownCursor
}

func ResizeUpDownCursor() *Cursor {
	if resizeUpDownCursor == nil {
		resizeUpDownCursor = cursorInit(C.getResizeUpDownCursor())
	}
	return resizeUpDownCursor
}

func DisappearingItemCursor() *Cursor {
	if disappearingItemCursor == nil {
		disappearingItemCursor = cursorInit(C.getDisappearingItemCursor())
	}
	return disappearingItemCursor
}

func OperationNotAllowedCursor() *Cursor {
	if operationNotAllowedCursor == nil {
		operationNotAllowedCursor = cursorInit(C.getOperationNotAllowedCursor())
	}
	return operationNotAllowedCursor
}

func DragLinkCursor() *Cursor {
	if dragLinkCursor == nil {
		dragLinkCursor = cursorInit(C.getDragLinkCursor())
	}
	return dragLinkCursor
}

func DragCopyCursor() *Cursor {
	if dragCopyCursor == nil {
		dragCopyCursor = cursorInit(C.getDragCopyCursor())
	}
	return dragCopyCursor
}

func ContextualMenuCursor() *Cursor {
	if contextualMenuCursor == nil {
		contextualMenuCursor = cursorInit(C.getContextualMenuCursor())
	}
	return contextualMenuCursor
}
