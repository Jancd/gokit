package threading

type (
	JobsGroup struct {
		Jobs  func()
		Count int
	}
)

func NewJobsGroup(job func(), count int) JobsGroup {
	return JobsGroup{
		Jobs:  job,
		Count: count,
	}
}

func (jg JobsGroup) Start() {
	group := NewGoroutineGroup()
	for i := 0; i < jg.Count; i++ {
		group.RunSafe(jg.Jobs)
	}
	group.Wait()
}
