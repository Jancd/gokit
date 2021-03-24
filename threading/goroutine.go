package threading

import (
	"bytes"
	"runtime"
	"strconv"

	"gokit/chars"
	"gokit/logger"
)

const (
	goroutinePrefix = "goroutine "
	byteHolder      = ' '
)

type GoRoutineId uint64

func GoSafe(fn func()) {
	go RunSafe(fn)
}

func RunSafe(fn func()) {
	defer recoverWithLog()
	fn()
}

func GetGoRoutineId() GoRoutineId {
	tmp := make([]byte, 64)
	tmp = tmp[:runtime.Stack(tmp, false)]
	tmp = bytes.TrimPrefix(tmp, chars.StringToBytes(goroutinePrefix))
	tmp = tmp[:bytes.IndexByte(tmp, byteHolder)]

	num, err := strconv.ParseUint(chars.BytesToString(tmp), 10, 64)
	if err != nil {
		logger.Error(err)
		return GoRoutineId(0)
	}
	return GoRoutineId(num)
}

func recoverWithLog(sweepFns ...func()) {
	for _, f := range sweepFns {
		f()
	}

	if r := recover(); r != nil {
		logger.ErrorStack(r)
	}
}
