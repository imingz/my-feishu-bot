package slogor

import "sync"

const (
	maxBufferSize     = 16 << 10 // 16384
	initialBufferSize = 1024
)

// bufPool is a sync.Pool that holds a pool of byte slices to be used for logging.
var bufPool = sync.Pool{
	// New initializes a new buffer of size 1024 in the pool.
	New: func() any {
		b := make([]byte, 0, initialBufferSize)
		return &b
	},
}

// allocBuf retrieves a buffer from the pool.
// It type-asserts the result of the pool's Get method to a byte slice pointer.
func allocBuf() *[]byte {
	return bufPool.Get().(*[]byte)
}

// freeBuf puts the buffer back into the pool after truncating it.
// This ensures that allocBuf always returns a zero-length slice.
// Additionally, it prevents large buffers from being returned to the pool.
// This is important to avoid potential memory wastage and long-term storage of large unused buffers.
func freeBuf(b *[]byte) {
	// To prevent excessive memory usage, only smaller buffers are returned to the pool.
	if cap(*b) <= maxBufferSize {
		// Truncate the buffer to length 0 before returning it to the pool.
		*b = (*b)[:0]
		// Put the buffer back into the pool.
		bufPool.Put(b)
	}
}
