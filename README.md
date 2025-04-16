# 🚀 Go Concurrency Patterns

A collection of practical **concurrency patterns** and implementations in Go (Golang). This repository serves as both a learning resource and a reference for common (and advanced) concurrency paradigms using goroutines, channels, and sync primitives.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go) 
[![License](https://img.shields.io/badge/license-MIT-blue)](LICENSE)

## 📌 Why This Repo?

Concurrency is one of Go's strongest features, but writing **correct, efficient, and maintainable** concurrent code requires understanding key patterns. This repo documents my journey exploring these patterns with practical examples.

## 🧩 Patterns Included

### Basic Patterns
- [Generator Pattern](/generator/) – Lazy data generation using channels
- [Fan-Out/Fan-In](/fanout-fanin/) – Parallel processing with multiple workers
- [Worker Pool](/workerpool/) – Fixed pool of goroutines for task processing
- [Pipeline](/pipeline/) – Chained processing stages with channels

### Advanced Patterns
- [Pub/Sub](/pubsub/) – Publish-Subscribe model in Go
- [Semaphore Pattern](/semaphore/) – Controlling concurrency with buffered channels
- [Error Handling](/error-handling/) – Robust error propagation in concurrent workflows
- [Context Cancellation](/context-cancellation/) – Graceful shutdown using `context.Context`

### Sync Primitives
- [Mutex Patterns](/sync-mutex/) – Safe shared state access
- [WaitGroup](/sync-waitgroup/) – Coordinating goroutine completion
- [Atomic Counters](/sync-atomic/) – Lock-free counters with `sync/atomic`

## 🛠️ Usage

Each pattern includes:
- A standalone **example implementation**
