package counter

import (
	"sync/atomic"
)

//Counter structure
type Counter uint64

//Count func
func (c *Counter) Count() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}

//Empty func
func (c *Counter) Empty() {
	atomic.StoreUint64((*uint64)(c), 0)
}

//Put func
func (c *Counter) Put() {
	atomic.AddUint64((*uint64)(c), 1)
}

//Set func
func (c *Counter) Set(delta uint64) {
	atomic.AddUint64((*uint64)(c), delta)
}
