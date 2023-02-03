package threading

import "sync"

// Sema 小demo，避免在大量 goroutine 情景使用
type Sema struct {
	ch chan struct{}
	wg *sync.WaitGroup
}

func NewSema(routineCount int) *Sema {
	return &Sema{
		ch: make(chan struct{}, routineCount),
		wg: new(sync.WaitGroup),
	}
}

func (s *Sema) Add(delta int) {
	s.wg.Add(delta)
	for i := 0; i < delta; i++ {
		s.ch <- struct{}{}
	}
}

func (s *Sema) Done() {
	_ = <-s.ch
	s.wg.Done()
}

func (s *Sema) Wait() {
	s.wg.Wait()
}
