package worker

import (
	"errors"
	"sync"
	"testing"

	"task-queue/internal/job"
)

func TestWorkerPoolFailsAfterMaxRetries(t *testing.T) {
	numWorkers := 1

	jobs := make(chan job.Job, 10)

	var wg sync.WaitGroup
	var successCount int64
	var failedCount int64

	attempts := 0

	mockProcess := func(payload string) error {
		attempts++
		return errors.New("always fail")
	}

	wg.Add(1)

	StartWorkerPool(
		numWorkers,
		jobs,
		&wg,
		&successCount,
		&failedCount,
		mockProcess,
	)

	jobs <- job.Job{
		ID:      1,
		Payload: "test",
		Retries: 0,
	}

	wg.Wait()

	if successCount != 0 {
		t.Fatalf("expected 0 success, got %d", successCount)
	}

	if failedCount != 1 {
		t.Fatalf("expected 1 failure, got %d", failedCount)
	}

	if attempts != 4 {
		t.Fatalf("expected 4 attempts, got %d", attempts)
	}
}
