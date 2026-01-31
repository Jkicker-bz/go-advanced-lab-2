package main

import (
	"fmt"
	"os"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}

	result := 1
	for k := 1; k <= n; k++ {
		result *= k
	}
	return result, nil
}

func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("prime check requires number >= 2")
	}
	if n == 2 {
		return true, nil
	}
	if n%2 == 0 {
		return false, nil
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, fmt.Errorf("negative exponents not supported")
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	total := initial
	add = func(x int) {
		total += x
	}
	subtract = func(x int) {
		total -= x
	}
	get = func() int {
		return total
	}
	return add, subtract, get
}

func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = operation(num)
	}
	return result
}

func Filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range nums {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
	accumulator := initial
	for _, num := range nums {
		accumulator = operation(accumulator, num)
	}
	return accumulator
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func ExploreProcess() {
	fmt.Printf("Current rocess ID: %d\n", os.Getpid())
	fmt.Printf("Parent Process ID: %d\n", os.Getppid())

	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("Memory address of slice header: %p\n", &data)
	fmt.Printf("Memory address of first element: %p\n", &data[0])

	fmt.Printf("\nNote: Other processes cannot access these memory addresses due to process isolation.\n")
}

// This function will not modify the original variable, because Go uses "pass by value".
func DoubleValue(x int) {
	x = x * 2
}

// This function will modify the original variable, because we are passing a pointer to it,
// allowing us to change the value at that memory address.
func DoublePointer(x *int) {
	*x = *x * 2
}

// This variable stays on the stack
func CreateOnStack() int {
	x := 42
	return x
}

// This variable escapes to the heap
func CreateOnHeap() *int {
	x := 42
	return &x
}

func SwapValues(a, b int) (int, int) {
	return b, a
}

func SwapPointers(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

/*
  - Which variables escaped? The variable 'x' in CreateOnHeap() escaped this is because the functions returns a pointer
    to a local variable. Basically, the address of 'x' is returned, so 'x' must live longer than the function call.
    If 'x' stayed on the stack, it would be deleted when the function finishes, leaving the pointer dangling.

  - What does "escapes to heap" mean?
    It means the variable is moved from the stock to the heap so it can live lponger than the function that created it.
*/
func AnalyzeEscape() {
	_ = CreateOnStack()
	_ = CreateOnHeap()
}

func main() {
	ExploreProcess()
}
