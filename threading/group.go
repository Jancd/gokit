package threading

import "sync"

const (
	one = 1
)

type GoroutineGroup struct {
	waitGroup sync.WaitGroup
}

func NewGoroutineGroup() *GoroutineGroup {
	return new(GoroutineGroup)
}

func (g *GoroutineGroup) Run(fn func()) {
	g.waitGroup.Add(one)

	go func() {
		defer g.waitGroup.Done()
		fn()
	}()
}

func (g *GoroutineGroup) RunSafe(fn func()) {
	g.waitGroup.Add(one)

	RunSafe(func() {
		defer g.waitGroup.Done()
		fn()
	})
}

func (g *GoroutineGroup) Wait() {
	g.waitGroup.Wait()
}
