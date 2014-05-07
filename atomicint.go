package fatomic

import "sync/atomic"

type AtomicInt struct {
	before [7]int64
	Value  int64
	after  [8]int64
}

// Sequentially consistent load of Value
func (a *AtomicInt) ALoad() int64 {
	return atomic.LoadInt64(&a.Value)
}

// Sequentially consistent store of Value
func (a *AtomicInt) AStore(val int64) {
	atomic.StoreInt64(&a.Value, val)
}

// Relaxed store of Value, equivalent to java's AtomicInteger.LazySet(...)
func (a *AtomicInt) LazyStore(val int64) {
	LazyStore(&a.Value, val)
}

// Atomically adds to Value - also provides sequentially consistent memory ordering
func (a *AtomicInt) Add(add int64) {
	atomic.AddInt64(&a.Value, add)
}

// Relaxed add to value
func (a *AtomicInt) LazyAdd(add int64) {
	val := a.Value + add
	LazyStore(&a.Value, val)
}
