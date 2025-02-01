# Go Race Condition in Concurrent Append to Slice

This repository demonstrates a common concurrency bug in Go: a race condition when multiple goroutines concurrently append to a slice without proper synchronization.

## Description
The `bug.go` file contains a program that launches multiple goroutines, each appending a number to a shared slice.  Without proper synchronization, this leads to a race condition, resulting in an incorrect final slice length.

## Solution
The `bugSolution.go` file demonstrates how to solve this using a `sync.Mutex` to protect the shared slice during append operations. This ensures that only one goroutine can modify the slice at a time, preventing data corruption.

## How to Run
1. Clone this repository.
2. Navigate to the repository's directory.
3. Run `go run bug.go` to see the race condition in action.
4. Run `go run bugSolution.go` to observe the corrected behavior with the mutex in place.
