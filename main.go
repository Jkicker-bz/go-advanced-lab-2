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

func main() {

}
