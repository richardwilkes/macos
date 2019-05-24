package macos

/*
#import <Quartz/Quartz.h>

void caTransactionFlush() {
	[CATransaction flush];
}
*/
import "C"

func CATransactionFlush() {
	C.caTransactionFlush()
}
