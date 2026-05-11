# sumofeven

A small Go project for comparing different implementations of "sum of even numbers", with benchmarks.

## Features

The project provides 3 implementations. Each takes `[]int` and returns the sum of all even numbers:

- `SumOfEven(nums []int) int`: uses `%2 == 0`
- `SumOfEvenBit(nums []int) int`: uses bitwise check `n&1 == 0`
- `SumOfEvenIfPrev(nums []int) int`: uses a branchless expression `total += n * (1 - n&1)`

## Project structure

- `sum.go`: core implementations
- `sum_benchmark_test.go`: benchmarks
- `go.mod`: module definition

## Run benchmarks

From the repository root, run:

```bash
go test -run '^$' -bench . -benchmem
```

Common variants:

```bash
# Run only benchmarks with names matching Sum
go test -run '^$' -bench 'Sum' -benchmem

# Run multiple times for more stable numbers
go test -run '^$' -bench . -benchmem -count=5
```

## Notes

- Benchmarks compare the three implementations across multiple input sizes (for example: 16 / 1024 / 65536).
- The goal is to compare performance differences between coding styles, not to provide a general-purpose math library.

## Environment

- Go 1.22+

