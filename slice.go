package fatomic

const cacheLen = 64

func CacheProtectedBytes(size int) []byte {
	b := make([]byte, cacheLen + size + cacheLen)
	return b[cacheLen:size+cacheLen]
}
