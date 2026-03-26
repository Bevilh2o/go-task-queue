package queue

import "task-queue/internal/job"

type Queue struct {
	Jobs []job.Job
}

func (q *Queue) Push(j job.Job) {
	q.Jobs = append(q.Jobs, j)
}

func (q *Queue) Pop() (job.Job, bool) {
	if len(q.Jobs) == 0 {
		return job.Job{}, false
	}

	j := q.Jobs[0]
	q.Jobs = q.Jobs[1:]
	return j, true
}
