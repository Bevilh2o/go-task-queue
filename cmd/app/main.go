package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
	"task-queue/internal/job"
	"task-queue/internal/worker"
)

func main() {
	// CLI flags
	numJobs := flag.Int("jobs", 100, "number of jobs to process")
	numWorkers := flag.Int("workers", 4, "number of workers")

	flag.Parse()

	jobs := make(chan job.Job, *numJobs)
	var wg sync.WaitGroup

	var successCount int64
	var failedCount int64

	start := time.Now()

	worker.StartWorkerPool(*numWorkers, jobs, &wg, &successCount, &failedCount)

	// Produce jobs
	for i := 0; i < *numJobs; i++ {
		wg.Add(1)

		jobs <- job.Job{
			ID:      i,
			Payload: fmt.Sprintf("job-%d", i),
		}
	}

	// graceful shutdown
	go func() {
		wg.Wait()
		close(jobs)
	}()

	wg.Wait()

	elapsed := time.Since(start)

	fmt.Println("-----")
	fmt.Println("Workers:", *numWorkers)
	fmt.Println("Jobs:", *numJobs)
	fmt.Println("Success:", successCount)
	fmt.Println("Failed:", failedCount)
	fmt.Println("Total time:", elapsed)
	fmt.Println("Jobs/sec:", float64(*numJobs)/elapsed.Seconds())
}
