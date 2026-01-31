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
	fmt.Println("\n=== Math Operations ===")

	// Factorials
	for _, n := range []int{0, 5, -2} {
		val, _ := Factorial(n)
		fmt.Printf("Factorial(%d) = %d\n", n, val)
	}
	// Prime Numbers
	for _, n := range []int{17, 20, 25} {
		res, _ := IsPrime(n)
		fmt.Printf("IsPrime(%d) = %v\n", n, res)
	}
	// Powers
	p1, _ := Power(2, 8)
	p2, _ := Power(5, -3)
	fmt.Printf("Power(2, 8) = %d\n", p1)
	fmt.Printf("Power(5, -3) = %d\n", p2)

	fmt.Println("\n=== Closure Demonstration ===")
	count0 := MakeCounter(0)
	count100 := MakeCounter(100)
	fmt.Printf("Counter1 (Start 0): %d, %d\n", count0(), count0())
	fmt.Printf("Counter2 (Start 100): %d\n", count100())

	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)

	fmt.Printf("Multiplier Demo: %d doubled is %d, tripled is %d\n", 5, double(5), triple(5))

	fmt.Println("\n=== Higher-Order Functions ===")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original: %v\n", nums)

	fmt.Printf("Squared: %v\n", Apply(nums, func(x int) int { return x * x }))
	fmt.Printf("Evens only: %v\n", Filter(nums, func(x int) bool { return x%2 == 0 }))
	fmt.Printf("Sum of all: %d\n", Reduce(nums, 0, func(acc, x int) int { return acc + x }))

	add10 := func(x int) int { return x + 10 }
	doubleThenAdd10 := Compose(add10, double)
	fmt.Printf("Compose (Double then Add 10) on 5: %d\n", doubleThenAdd10(5))

	fmt.Println("\n=== Pointer Demonstration ===")
	x, y := 5, 10
	fmt.Printf("Before SwapValues: a=%d, b=%d\n", x, y)
	SwapValues(x, y)
	fmt.Printf("After SwapValues: a=%d, b=%d (originals unchanged)\n", x, y)

	SwapPointers(&x, &y)
	fmt.Printf("After SwapPointers: a=%d, b=%d (originals swapped)\n", x, y)
}
