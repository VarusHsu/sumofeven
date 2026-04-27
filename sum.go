package sumofeven

// SumOfEven returns the sum of all even integers in nums.
func SumOfEven(nums []int) int {
	total := 0
	for _, n := range nums {
		if n%2 == 0 {
			total += n
		}
	}
	return total
}

// SumOfEvenBit returns the sum of all even integers in nums using bitwise operations.
func SumOfEvenBit(nums []int) int {
	total := 0
	for _, n := range nums {
		if n&1 == 0 {
			total += n
		}
	}
	return total
}

// SumOfEvenIfPrev returns the sum of all even integers in nums using a single addition and bitwise operations.
func SumOfEvenIfPrev(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n * (1 - n&1)
	}
	return total
}

/*
go test ./... && go test -bench . -benchmem ./...
ok      github.com/varushsu/sumOfEven   0.795s [no tests to run]
goos: darwin
goarch: arm64
pkg: github.com/varushsu/sumOfEven
cpu: Apple M2 Pro
BenchmarkSumOfEven/size=16-12   				164853006                6.845 ns/op           0 B/op          0 allocs/op
BenchmarkSumOfEven/size=1024-12                  1982019               595.8 ns/op             0 B/op          0 allocs/op
BenchmarkSumOfEven/size=65536-12                   30764             38869 ns/op               0 B/op          0 allocs/op
BenchmarkSumOfEvenBit/size=16-12                175173807                6.787 ns/op           0 B/op          0 allocs/op
BenchmarkSumOfEvenBit/size=1024-12               1871148               598.2 ns/op             0 B/op          0 allocs/op
BenchmarkSumOfEvenBit/size=65536-12                30978             39157 ns/op               0 B/op          0 allocs/op
BenchmarkSumOfEvenIfPrev/size=16-12             176719315                6.828 ns/op           0 B/op          0 allocs/op
BenchmarkSumOfEvenIfPrev/size=1024-12            2636635               459.5 ns/op             0 B/op          0 allocs/op
BenchmarkSumOfEvenIfPrev/size=65536-12             41421             28804 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/varushsu/sumOfEven   15.828s
*/
