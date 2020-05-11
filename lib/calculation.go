package lib

import (
	"math"
	"sync/atomic"
	"unsafe"
)

// AddFloat64 is an adaptation of the "atomic" add function
func AddFloat64(val *float64, delta float64) (new float64) {
    for {
        old := *val
        new = old + delta
        if atomic.CompareAndSwapUint64(
			(*uint64)(unsafe.Pointer(val)), 
			math.Float64bits(old),
			math.Float64bits(new),
        ) {
            break
        }
    }
    return
}