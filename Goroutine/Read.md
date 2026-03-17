
concurrency vs Parallelism

concurrency - Multiple tasks are in progress at the same time, but not necessarily running at the exact same instant.
Cooking, Watching YouTube 👉 You switch between both → this is concurrent

Parallelism - is a subset of concurrency. its when two or more threads execute at the same real time on two or more CPUs
ex- 3 threads executing on three diff cpu's simultaneously.
ex- You cook, Your friend watches YouTube 👉 Both happen at the same time → this is parallel

Goroutine in Golang - 
Goroutines are a key feature of the Go programming language that allow u to run functions concurrently or in parallel, with other parts of your program.
imagine u have a task that can be broken down into smaller sub-tasks that can be executed independently, instead of waiting for each sub-task to finish before starting the nextone, u can use goroutines to execute them concurrently. This can greatly improve the permormance of your program,especially when dealing with tasks that involve I/O operations or other types of blocking operations.