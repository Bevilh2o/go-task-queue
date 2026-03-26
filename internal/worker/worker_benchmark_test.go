package worker

import (
	"sync"
	"testing"

	"task-queue/internal/job"
)

func BenchmarkWorkerPool(b *testing.B) {
	numWorkers := 8
	numJobs := 1000

	mockProcess := func(payload string) error {
		return nil
	}

	for n := 0; n < b.N; n++ {
		jobs := make(chan job.Job, numJobs)

		var wg sync.WaitGroup
		var successCount int64
		var failedCount int64

		wg.Add(numJobs)

		StartWorkerPool(
			numWorkers,
			jobs,
			&wg,
			&successCount,
			&failedCount,
			mockProcess,
		)

		for i := 0; i < numJobs; i++ {
			jobs <- job.Job{
				ID:      i,
				Payload: "test",
			}
		}

		close(jobs)
		wg.Wait()
	}
}
