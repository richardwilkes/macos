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

var (
	nsCursorArrowCursor                  *NSCursor
	nsCursorIBeamCursor                  *NSCursor
	nsCursorIBeamCursorForVerticalLayout *NSCursor
	nsCursorCrosshairCursor              *NSCursor
	nsCursorClosedHandCursor             *NSCursor
	nsCursorOpenHandCursor               *NSCursor
	nsCursorPointingHandCursor           *NSCursor
	nsCursorResizeLeftCursor             *NSCursor
	nsCursorResizeRightCursor            *NSCursor
	nsCursorResizeLeftRightCursor        *NSCursor
	nsCursorResizeUpCursor               *NSCursor
	nsCursorResizeDownCursor             *NSCursor
	nsCursorResizeUpDownCursor           *NSCursor
	nsCursorDisappearingItemCursor       *NSCursor
	nsCursorOperationNotAllowedCursor    *NSCursor
	nsCursorDragLinkCursor               *NSCursor
	nsCursorDragCopyCursor               *NSCursor
	nsCursorContextualMenuCursor         *NSCursor
)

func NSCursorArrowCursor() *NSCursor {
	if nsCursorArrowCursor == nil {
		nsCursorArrowCursor = nsCursorInit(C.getarrowCursor())
	}
	return nsCursorArrowCursor
}

func NSCursorIBeamCursor() *NSCursor {
	if nsCursorIBeamCursor == nil {
		nsCursorIBeamCursor = nsCursorInit(C.getIBeamCursor())
	}
	return nsCursorIBeamCursor
}

func NSCursorIBeamCursorForVerticalLayout() *NSCursor {
	if nsCursorIBeamCursorForVerticalLayout == nil {
		nsCursorIBeamCursorForVerticalLayout = nsCursorInit(C.getIBeamCursorForVerticalLayout())
	}
	return nsCursorIBeamCursorForVerticalLayout
}

func NSCursorCrosshairCursor() *NSCursor {
	if nsCursorCrosshairCursor == nil {
		nsCursorCrosshairCursor = nsCursorInit(C.getcrosshairCursor())
	}
	return nsCursorCrosshairCursor
}

func NSCursorClosedHandCursor() *NSCursor {
	if nsCursorClosedHandCursor == nil {
		nsCursorClosedHandCursor = nsCursorInit(C.getclosedHandCursor())
	}
	return nsCursorClosedHandCursor
}

func NSCursorOpenHandCursor() *NSCursor {
	if nsCursorOpenHandCursor == nil {
		nsCursorOpenHandCursor = nsCursorInit(C.getopenHandCursor())
	}
	return nsCursorOpenHandCursor
}

func NSCursorPointingHandCursor() *NSCursor {
	if nsCursorPointingHandCursor == nil {
		nsCursorPointingHandCursor = nsCursorInit(C.getpointingHandCursor())
	}
	return nsCursorPointingHandCursor
}

func NSCursorResizeLeftCursor() *NSCursor {
	if nsCursorResizeLeftCursor == nil {
		nsCursorResizeLeftCursor = nsCursorInit(C.getresizeLeftCursor())
	}
	return nsCursorResizeLeftCursor
}

func NSCursorResizeRightCursor() *NSCursor {
	if nsCursorResizeRightCursor == nil {
		nsCursorResizeRightCursor = nsCursorInit(C.getresizeRightCursor())
	}
	return nsCursorResizeRightCursor
}

func NSCursorResizeLeftRightCursor() *NSCursor {
	if nsCursorResizeLeftRightCursor == nil {
		nsCursorResizeLeftRightCursor = nsCursorInit(C.getresizeLeftRightCursor())
	}
	return nsCursorResizeLeftRightCursor
}

func NSCursorResizeUpCursor() *NSCursor {
	if nsCursorResizeUpCursor == nil {
		nsCursorResizeUpCursor = nsCursorInit(C.getresizeUpCursor())
	}
	return nsCursorResizeUpCursor
}

func NSCursorResizeDownCursor() *NSCursor {
	if nsCursorResizeDownCursor == nil {
		nsCursorResizeDownCursor = nsCursorInit(C.getresizeDownCursor())
	}
	return nsCursorResizeDownCursor
}

func NSCursorResizeUpDownCursor() *NSCursor {
	if nsCursorResizeUpDownCursor == nil {
		nsCursorResizeUpDownCursor = nsCursorInit(C.getresizeUpDownCursor())
	}
	return nsCursorResizeUpDownCursor
}

func NSCursorDisappearingItemCursor() *NSCursor {
	if nsCursorDisappearingItemCursor == nil {
		nsCursorDisappearingItemCursor = nsCursorInit(C.getdisappearingItemCursor())
	}
	return nsCursorDisappearingItemCursor
}

func NSCursorOperationNotAllowedCursor() *NSCursor {
	if nsCursorOperationNotAllowedCursor == nil {
		nsCursorOperationNotAllowedCursor = nsCursorInit(C.getoperationNotAllowedCursor())
	}
	return nsCursorOperationNotAllowedCursor
}

func NSCursorDragLinkCursor() *NSCursor {
	if nsCursorDragLinkCursor == nil {
		nsCursorDragLinkCursor = nsCursorInit(C.getdragLinkCursor())
	}
	return nsCursorDragLinkCursor
}

func NSCursorDragCopyCursor() *NSCursor {
	if nsCursorDragCopyCursor == nil {
		nsCursorDragCopyCursor = nsCursorInit(C.getdragCopyCursor())
	}
	return nsCursorDragCopyCursor
}

func NSCursorContextualMenuCursor() *NSCursor {
	if nsCursorContextualMenuCursor == nil {
		nsCursorContextualMenuCursor = nsCursorInit(C.getcontextualMenuCursor())
	}
	return nsCursorContextualMenuCursor
}
