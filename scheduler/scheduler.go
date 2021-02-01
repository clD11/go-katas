package scheduler

type Scheduler struct {
	jobs []Job
}

type Job struct {
	f    func(...int) int
	args []int
}

func New() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Schedule(f func(i ...int) int, args ...int) {
	s.jobs = append(s.jobs, Job{f, args})
}

func (s *Scheduler) Invoke() []int {
	result := []int{}
	for _, j := range s.jobs {
		result = append(result, j.f(j.args...))
	}
	return result
}
