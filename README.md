# 🚀 Go Concurrency Patterns

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-blue)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

A collection of production-ready concurrency patterns in Go and real-world use cases.

## 📖 Table of Contents

- [Features](#-features)
- [Patterns](#-patterns-included)
  - [Basic](#basic-patterns)
  - [Advanced](#advanced-patterns)
  - [Sync Primitives](#sync-primitives)
- [Usage](#-usage)
- [Contributing](#-contributing)
- [Resources](#-learning-resources)
- [License](#-license)

## 🌟 Features

✔️ Production-grade implementations  
✔️ Example use cases  
✔️ Common pitfalls explained  

## 🧩 Patterns Included

### Basic Patterns
| Pattern | Description | Example Use |
|---------|------------|-------------|
| [Generator](/basic_patterns/generator/) | Lazy data generation | Streaming large datasets |
| [Fan-Out/Fan-In](/basic_patterns/fanout-fanin/) | Parallel processing | CPU-bound workloads |
| [Worker Pool](/basic_patterns/workerpool/) | Controlled goroutine pool | API rate limiting |
| [Pipeline](/basic_patterns/pipeline/) | Chained processing | ETL workflows |

### Advanced Patterns
| Pattern | Description | Concurrency Control |
|---------|------------|---------------------|
| [Pub/Sub](/pubsub/) | Event broadcasting | `select` with channels |
| [Semaphore](/semaphore/) | Resource limiting | Buffered channels |
| [Error Group](/errgroup/) | Error propagation | `sync/errgroup` |
| [Circuit Breaker](/circuit-breaker/) | Fault tolerance | State machine |

### Sync Primitives
- [Mutex](/sync-mutex/) - Shared state protection
- [RWMutex](/sync-mutex/) - Optimized read-heavy cases
- [WaitGroup](/sync-waitgroup/) - Goroutine synchronization
- [Atomic](/sync-atomic/) - Lock-free operations

## 🛠️ Usage

you can try different implementation by changing the implementation's function call in main.go
```bash
# Run any pattern
go run ./basic_patterns/main.go
```

```bash
# Run benchmarks
cd ./basic_patterns
go test -bench=. -benchmem
