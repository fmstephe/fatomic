package fatomic

import (
	"unsafe"
)

func CacheProtectedBytes(size int) []byte {
	return BufferedAlignedBytes(size, 64, 8)
}

func BufferedAlignedBytes(size, buffer, align int) []byte {
	if align <= 8 {
		b := make([]byte, buffer + size + buffer)
		return b[buffer:size+buffer]
	}
	total := size + (2*align) + (2*buffer)
	b := make([]byte, total)
	p := int(uintptr(unsafe.Pointer(&b[buffer])))
	rem := p % align
	offset := buffer
	if rem != 0 {
		offset += align - rem
	}
	bs := b[offset:size+offset]
	// These should be unit tests
	if int(uintptr(unsafe.Pointer(&bs[0])))% align != 0 {
		panic("Bad alignment")
	}
	if len(bs) != size {
		panic("Bad size")
	}
	return bs
}

func CacheProtectedPointers(size int) []unsafe.Pointer {
	b := make([]unsafe.Pointer, 64 + size + 64)
	return b[64:size+64]
}

func powerOfTwo(val int64) bool {
	pow := int64(2)
	for i := 0; i < 64; i++ {
		if pow == val {
			return true
		}
	}
	return false
}
