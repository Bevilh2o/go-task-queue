package worker

import (
	"fmt"
	"sync"
	"sync/atomic"
	"task-queue/internal/job"
	"task-queue/internal/processor"
)

func StartWorkerPool(
	numWorkers int,
	jobs chan job.Job,
	wg *sync.WaitGroup,
	successCount *int64,
	failedCount *int64,
) {
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for j := range jobs {

				err := processor.Process(j.Payload)

				if err != nil {
					if j.Retries < 3 {
						j.Retries++
						fmt.Printf("Worker %d - Job %d RETRY (%d)\n", workerID, j.ID, j.Retries)

						jobs <- j
					} else {
						fmt.Printf("Worker %d - Job %d FAILED permanently\n", workerID, j.ID)

						atomic.AddInt64(failedCount, 1)
						wg.Done()
					}
				} else {
					fmt.Printf("Worker %d - Job %d DONE\n", workerID, j.ID)

					atomic.AddInt64(successCount, 1)
					wg.Done()
				}
			}
		}(i)
	}
}
