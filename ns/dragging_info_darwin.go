package ns

/*
#import <Cocoa/Cocoa.h>

typedef void *NSWindowPtr;
typedef void *NSDraggingInfoPtr;
typedef void *NSPasteboardPtr;

NSPasteboardPtr nsDraggingInfoPasteboard(NSDraggingInfoPtr di) {
	return (NSPasteboardPtr)[(id<NSDraggingInfo>)di draggingPasteboard];
}

NSInteger nsDraggingInfoSequenceNumber(NSDraggingInfoPtr di) {
	return [(id<NSDraggingInfo>)di draggingSequenceNumber];
}

NSDragOperation nsDraggingInfoSourceOperationMask(NSDraggingInfoPtr di) {
	return [(id<NSDraggingInfo>)di draggingSourceOperationMask];
}

NSPoint nsDraggingInfoLocation(NSDraggingInfoPtr di) {
	return [(id<NSDraggingInfo>)di draggingLocation];
}

NSWindowPtr nsDraggingInfoDestinationWindow(NSDraggingInfoPtr di) {
	return (NSWindowPtr)[(id<NSDraggingInfo>)di draggingDestinationWindow];
}

NSInteger nsDraggingInfoGetNumberOfValidItemsForDrop(NSDraggingInfoPtr di) {
	return [(id<NSDraggingInfo>)di numberOfValidItemsForDrop];
}

void nsDraggingInfoSetNumberOfValidItemsForDrop(NSDraggingInfoPtr di, NSInteger count) {
	[(id<NSDraggingInfo>)di setNumberOfValidItemsForDrop:count];
}

NSPoint nsDraggingInfoImageLocation(NSDraggingInfoPtr di) {
	return [(id<NSDraggingInfo>)di draggedImageLocation];
}
*/
import "C"

type DraggingInfo struct {
	native C.NSDraggingInfoPtr
}

func (di *DraggingInfo) Pasteboard() *Pasteboard {
	return &Pasteboard{native: C.nsDraggingInfoPasteboard(di.native)}
}

func (di *DraggingInfo) SequenceNumber() int {
	return int(C.nsDraggingInfoSequenceNumber(di.native))
}

func (di *DraggingInfo) SourceOperationMask() DragOperation {
	return DragOperation(C.nsDraggingInfoSourceOperationMask(di.native))
}

func (di *DraggingInfo) Location() (x, y float64) {
	p := C.nsDraggingInfoLocation(di.native)
	return float64(p.x), float64(p.y)
}

func (di *DraggingInfo) DestinationWindow() *Window {
	return &Window{native: C.nsDraggingInfoDestinationWindow(di.native)}
}

func (di *DraggingInfo) NumberOfValidItemsForDrop() int {
	return int(C.nsDraggingInfoGetNumberOfValidItemsForDrop(di.native))
}

func (di *DraggingInfo) SetNumberOfValidItemsForDrop(count int) {
	C.nsDraggingInfoSetNumberOfValidItemsForDrop(di.native, C.NSInteger(count))
}

func (di *DraggingInfo) ImageLocation() (x, y float64) {
	p := C.nsDraggingInfoImageLocation(di.native)
	return float64(p.x), float64(p.y)
}
