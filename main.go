package main

import (
	"fmt"
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
func main() {
	// --- Testing MakeCounter ---
	counter1 := MakeCounter(0)
	fmt.Println(counter1()) // 1
	fmt.Println(counter1()) // 2

	counter2 := MakeCounter(10)
	fmt.Println(counter2()) // 11
	fmt.Println(counter1()) // 3 (Independent from counter2)

	// --- Testing MakeMultiplier ---
	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)
	fmt.Println(double(5)) // 10
	fmt.Println(triple(5)) // 15

	// --- Testing MakeAccumulator ---
	add, sub, get := MakeAccumulator(100)
	add(50)
	fmt.Println(get()) // 150
	sub(20)
	fmt.Println(get()) // 130
}
