package verbosity

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func New() *VState {

}

type VState struct {
	mu sync.Mutex

	vmodule mo
}

// Level must be an int32 to support atomic read/writes.
type Level int32

type levelSpec struct {
	vs *VState
	l  Level
}

// get returns the value of the level.
func (l *levelSpec) get() Level {
	return Level(atomic.LoadInt32((*int32)(&l.l)))
}

type moduleSpec struct {
	vs *VState
	filter []
}

type modulePat struct {
	pattern string
	literal bool
	level 
}