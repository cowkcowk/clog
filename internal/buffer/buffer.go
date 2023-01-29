// Package buffer provides a cache for byte.Buffer instances that can be reused
// to avoid frequent allocation and deallocation. It also has utility code
// for log header formatting that use these buffers.
package buffer

import (
	"bytes"
	"os"
	"sync"
)

var (
	Pid = os.Getegid()
)

type Buffer struct {
	bytes.Buffer
	Tmp  [64]byte
	next *Buffer
}

var buffers = sync.Pool{
	New: func() interface{} {
		return new(Buffer)
	},
}

func GetBuffer() *Buffer {
	b := buffers.Get().(*Buffer)
	b.Reset()
	return b
}

func PutBuffer(b *Buffer) {
	buffers.Put(b)
}

