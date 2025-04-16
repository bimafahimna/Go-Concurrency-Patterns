# 🚀 Go Concurrency Patterns

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-blue)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

A collection of production-ready concurrency patterns in Go, with benchmarks and real-world use cases.

## 📖 Table of Contents

- [Features](#-features)
- [Patterns](#-patterns-included)
  - [Basic](#basic-patterns)
  - [Advanced](#advanced-patterns)
  - [Sync Primitives](#sync-primitives)
- [Usage](#-usage)
- [Benchmarks](#-benchmarks)
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
| [Generator](/generator/) | Lazy data generation | Streaming large datasets |
| [Fan-Out/Fan-In](/fanout-fanin/) | Parallel processing | CPU-bound workloads |
| [Worker Pool](/workerpool/) | Controlled goroutine pool | API rate limiting |
| [Pipeline](/pipeline/) | Chained processing | ETL workflows |

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

```bash
# Run any pattern
go run ./patterns/<pattern-name>/main.go

# Run benchmarks
cd ./patterns/<pattern-name>
go test -bench=. -benchmem
