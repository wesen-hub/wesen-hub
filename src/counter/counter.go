package counter

import "sync/atomic"

var count uint32 = 0

func IncrementAndGetOld() uint32 {
	return atomic.AddUint32(&count, 1)
}

func Get() uint32 {
	return atomic.LoadUint32(&count)
}

func DecrementAndGetNew() uint32 {
	return atomic.AddUint32(&count, -1)
}

func Reset() {
	atomic.SwapUint32(&count, 0)
}
