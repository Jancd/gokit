package threading

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJobsGroup_Start(t *testing.T) {
	var counter int32
	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())
	group := NewJobsGroup(func() {
		goroutineId := GetGoRoutineId()
		if goroutineId == 0 {
			return
		}
		atomic.AddInt32(&counter, 1)
		wg.Done()
	}, runtime.NumCPU())
	go group.Start()
	wg.Wait()

	assert.Equal(t, int(counter), runtime.NumCPU())
}
