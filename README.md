# Go Essentials: Backend Engineering and Concurrency

This repository contains a comprehensive collection of implementations developed during the Go Essentials specialization. The project demonstrates the application of Go's core philosophy: simplicity, efficiency, and native support for highly concurrent systems.

## Core Technical Implementations

### 1. Concurrency and Synchronization
* **Goroutines and Channels:** Implementation of asynchronous workflows and the Fan-out/Fan-in pattern for parallel processing.
* **Synchronization Primitives:** Practical use of `sync.WaitGroup` for task orchestration and `sync.Map` for thread-safe in-memory data structures.
* **Context Management:** Utilizing the `context` package to handle timeouts, deadlines, and graceful cancellations in distributed operations.

### 2. Network Programming and API Integration
* **RESTful Services:** Development of HTTP servers from scratch, including custom routing, handlers, and status code management.
* **JSON Processing:** High-performance data serialization using `encoding/json` with a focus on `Encoders`, `Decoders`, and `io.Reader` streaming to minimize memory overhead.
* **External Integration:** Consuming third-party APIs (GitHub REST API) with robust error handling and URL path escaping.

### 3. Performance and Reliability
* **Unit Testing:** Comprehensive test suites using the native `testing` package and the `testify` assertion library.
* **Benchmarking:** Performance profiling of critical functions to measure execution time and throughput.
* **Observability:** CPU Profiling using `pprof` to identify and resolve execution bottlenecks.

### 4. Software Architecture and Idioms
* **Resource Management:** Effective use of `defer` for closing resources and `io.LimitReader` for security against oversized payloads.
* **Data Modeling:** Implementation of struct tags for JSON mapping, pointers for efficient state management, and empty interfaces for generic data handling.
* **Error Handling:** Idiomatic Go error propagation and the use of the `errors` package for wrapping context.

## Project Structure

* **/goroutines:** Implementations of concurrent downloads, channel-based communication, and the `select` statement.
* **/network:** HTTP servers, JSON-to-struct mapping, and an in-memory key-value database.
* **/functions:** Advanced function mechanics, including closures, deferred calls, and error returns.
* **/errors:** Patterns for logging, panic recovery, and custom error types.
* **/basics:** Fundamental algorithms including word counting, maps, slices, and string manipulation.

---
*Developed as part of the Go Essentials training series.*
