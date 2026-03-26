package worker

import (
	"sync"
	"testing"

	"task-queue/internal/job"
)


func TestWorkerPoolBasicProcessing(t *testing.T) {
	numJobs := 20
	numWorkers := 2

	jobs := make(chan job.Job, numJobs)

	var wg sync.WaitGroup
	var successCount int64
	var failedCount int64

	mockProcess := func(payload string) error {
	return nil
	}

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
			Retries: 0,
		}
	}

	close(jobs)

	wg.Wait()

	total := successCount + failedCount

	if total != int64(numJobs) {
		t.Fatalf("expected %d processed jobs, got %d", numJobs, total)
	}
}
