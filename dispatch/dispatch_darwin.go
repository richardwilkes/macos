package dispatch

import (
	"sync"

	"github.com/richardwilkes/toolbox/errs"
)

/*
#import <dispatch/dispatch.h>

void dispatchTaskCallback(uint64_t id);

static void dispatchAsyncFunctionOnMainQueue(uint64_t id) {
	dispatch_async_f(dispatch_get_main_queue(), (void *)id, (dispatch_function_t)dispatchTaskCallback);
}
*/
import "C"

var (
	dispatchLock            sync.Mutex
	dispatchID              uint64 = 1
	dispatchMap                    = make(map[uint64]func())
	dispatchRecoverCallback errs.RecoveryHandler
)

func AsyncFunctionOnMainQueue(callback func()) {
	dispatchLock.Lock()
	id := dispatchID
	dispatchID++
	dispatchMap[id] = callback
	dispatchLock.Unlock()
	C.dispatchAsyncFunctionOnMainQueue(C.uint64_t(id))
}

func SetDispatchRecoverCallback(recoveryHandler errs.RecoveryHandler) {
	dispatchLock.Lock()
	dispatchRecoverCallback = recoveryHandler
	dispatchLock.Unlock()
}

//export dispatchTaskCallback
func dispatchTaskCallback(id uint64) {
	dispatchLock.Lock()
	callback, ok := dispatchMap[id]
	if ok {
		delete(dispatchMap, id)
	}
	recoverCallback := dispatchRecoverCallback
	dispatchLock.Unlock()
	if callback != nil {
		defer errs.Recovery(recoverCallback)
		callback()
	}
}
