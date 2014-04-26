package fatomic

import (
	"unsafe"
)

const cacheLen = 64

func CacheProtectedBytes(size int) []byte {
	b := make([]byte, cacheLen + size + cacheLen)
	return b[cacheLen:size+cacheLen]
}

func CacheProtectedPointers(size int) []unsafe.Pointer {
	b := make([]unsafe.Pointer, cacheLen + size + cacheLen)
	return b[cacheLen:size+cacheLen]
}
