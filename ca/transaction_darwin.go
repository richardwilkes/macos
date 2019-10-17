package ca

/*
#import <Quartz/Quartz.h>

void transactionFlush() {
	[CATransaction flush];
}
*/
import "C"

func TransactionFlush() {
	C.transactionFlush()
}
