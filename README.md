# 🚀 Concurrent Task Queue in Go

![CI](https://github.com/Bevilh2o/go-task-queue/actions/workflows/go.yml/badge.svg)

A high-performance concurrent task queue built in Go, featuring worker pools, retry logic, and graceful shutdown.

---

## 📌 Overview

This project demonstrates how to design and implement a scalable job processing system using Go’s concurrency primitives.

It simulates a distributed workload where multiple workers process jobs with retry logic and failure handling.

---

## ⚙️ Features

- 🧵 Worker pool with configurable concurrency
- 🔁 Retry mechanism with max retry limit
- ❌ Permanent failure handling
- 📦 Channel-based job queue
- 🧠 Deterministic behavior under concurrency
- 🛑 Graceful shutdown (no job loss, no goroutine leaks)
- 📊 Performance benchmarking (jobs/sec)
- 🔌 Decoupled processing logic via dependency injection

---

## 🏗️ Architecture

```
                +-------------------+
                |     Job Queue     |
                |     (channel)     |
                +--------+----------+
                         |
        +----------------+----------------+
        |        |        |        |       |
     Worker   Worker   Worker   Worker   ...
        |        |        |        |
        +--------+--------+--------+
                         |
                 Result Aggregation
                (success / failure)
```

---

## 🔁 Retry Logic

Each job:

- is retried up to **3 times**
- logs retry attempts
- is marked as **FAILED permanently** after max retries

---

## 🛑 Graceful Shutdown

The system ensures:

- all jobs are processed before exit  
- no goroutines are leaked  
- clean termination of workers  

---

## 📊 Benchmark

Run:

```bash
go test -bench=. -benchmem ./internal/worker
```

Example result:

```
BenchmarkWorkerPool-8   	     148	   8269285 ns/op	   39598 B/op	     765 allocs/op
```

### Throughput

- ~120,000 jobs/sec (8 workers, lightweight processing)

---

## 📈 Observations

- Throughput increases with concurrency  
- Scaling is **not linear** due to:
  - scheduling overhead  
  - shared resource contention  
- Results are **consistent across runs**, proving deterministic behavior  

---

## 🧪 How to Run

```bash
go run cmd/app/main.go -workers=4 -jobs=100
```

### Parameters

- `-workers` → number of concurrent workers  
- `-jobs` → total number of jobs  

---

## 🧠 What This Demonstrates

- Go concurrency (goroutines + channels)  
- Worker pool pattern  
- Dependency injection (decoupled processing logic)  
- Fault tolerance via retries  
- Synchronization and coordination  
- Performance measurement  
- Clean system shutdown  

---

## 🚀 Future Improvements

- Exponential backoff for retries  
- Context-based cancellation  
- Metrics export (Prometheus)  
- Persistent queue (Redis / Kafka)  
- Rate limiting  

---

## 👤 Author

Michel Bevilacqua
