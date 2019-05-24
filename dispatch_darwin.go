package macos

import "sync"

/*
#import <dispatch/dispatch.h>

void dispatchTaskCallback(uint64_t id);

static void dispatchAsyncFOnMainQueue(uint64_t id) {
	dispatch_async_f(dispatch_get_main_queue(), (void *)id, (dispatch_function_t)dispatchTaskCallback);
}
*/
import "C"

var (
	dispatchLock            sync.Mutex
	dispatchID              uint64 = 1
	dispatchMap                    = make(map[uint64]func())
	dispatchRecoverCallback func(interface{})
)

func DispatchAsyncFOnMainQueue(callback func()) {
	dispatchLock.Lock()
	id := dispatchID
	dispatchID++
	dispatchMap[id] = callback
	dispatchLock.Unlock()
	C.dispatchAsyncFOnMainQueue(C.uint64_t(id))
}

func SetDispatchRecoverCallback(f func(interface{})) {
	dispatchLock.Lock()
	dispatchRecoverCallback = f
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
		if recoverCallback != nil {
			defer func() {
				if err := recover(); err != nil {
					recoverCallback(err)
				}
			}()
		}
		callback()
	}
}
