package threading

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSema(t *testing.T) {
	s := NewSema(3)

	sharedChan := make(chan struct{}, 10)

	worker := func() {
		defer s.Done()
		s.Add(1)

		sharedChan <- struct{}{}
		time.Sleep(time.Second)
	}

	jobCount := 10
	for i := 0; i < jobCount; i++ {
		go worker()
	}

	counter := 0
	for {
		select {
		case <-sharedChan:
			counter += 1
			continue
		case <-time.After(time.Millisecond * 500):
			goto b
		}
	b:
		break
	}

	assert.Equal(t, 3, counter)

	s.Wait()
}
