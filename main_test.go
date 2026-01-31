package main

import (
	"slices"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{name: "factorial of 0", input: 0, want: 1, wantErr: false},
		{name: "factorial of 1", input: 1, want: 1, wantErr: false},
		{name: "factorial of 3", input: 3, want: 6, wantErr: false},
		{name: "factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "factorial of 10", input: 10, want: 3628800, wantErr: false},
		{name: "factorial of negative number", input: -4, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{name: "Prime number: 2", input: 2, want: true, wantErr: false},
		{name: "Prime number: 3", input: 3, want: true, wantErr: false},
		{name: "Composite number: 6", input: 6, want: false, wantErr: false},
		{name: "Number less than 2", input: 1, want: false, wantErr: true},
		{name: "Negative number", input: -5, want: false, wantErr: true},
		{name: "Large prime number: 29", input: 29, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool
	}{
		{name: "Positive exponent", base: 2, exponent: 3, want: 8, wantErr: false},
		{name: "Exponent zero", base: 5, exponent: 0, want: 1, wantErr: false},
		{name: "Base zero", base: 0, exponent: 2, want: 0, wantErr: false},
		{name: "Exponent equals 1", base: 25, exponent: 1, want: 25, wantErr: false},
		{name: "Negative exponent", base: 2, exponent: -3, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Power() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakecounter(t *testing.T) {
	tests := []struct {
		name      string
		start     int
		calls     int
		wantFirst int
		wantLast  int
	}{
		{name: "Start at 0, 3 calls", start: 0, calls: 3, wantFirst: 1, wantLast: 3},
		{name: "Start at 5, 5 calls", start: 5, calls: 5, wantFirst: 6, wantLast: 10},
		{name: "Start at -2, 4 calls", start: -2, calls: 4, wantFirst: -1, wantLast: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := MakeCounter(tt.start)
			var got int
			for i := 0; i < tt.calls; i++ {
				got = counter()
				if i == 0 &&
					got != tt.wantFirst {
					t.Errorf("MakeCounter() first call = %v, want %v", got, tt.wantFirst)
				}
			}
			if got != tt.wantLast {
				t.Errorf("MakeCounter() last call = %v, want %v", got, tt.wantLast)
			}
		})
	}
}

func TestCounterIndependence(t *testing.T) {
	c1 := MakeCounter(0)
	c2 := MakeCounter(10)

	c1() // 1
	c1() // 2

	val2 := c2() // Should be 11, not influenced by c1
	if val2 != 11 {
		t.Errorf("Independence check failed: c2 got %d, want 11", val2)
	}
	val1 := c1() // Should be 3
	if val1 != 3 {
		t.Errorf("Independence check failed: c1 got %d, want 3", val1)
	}
}

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		input  int
		want   int
	}{
		{name: "Double 4", factor: 2, input: 4, want: 8},
		{name: "Triple 5", factor: 3, input: 5, want: 15},
		{name: "Quadruple 6", factor: 4, input: 6, want: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiplier := MakeMultiplier(tt.factor)
			got := multiplier(tt.input)
			if got != tt.want {
				t.Errorf("MakeMultiplier(%d) with input %d = %v, want %v", tt.factor, tt.input, got, tt.want)
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	add, subtract, get := MakeAccumulator(100)

	tests := []struct {
		name      string
		operation string
		value     int
		want      int
	}{
		{name: "Add 50", operation: "add", value: 50, want: 150},
		{name: "Subtract 20", operation: "subtract", value: 20, want: 130},
		{name: "Add 30", operation: "add", value: 30, want: 160},
		{name: "Subtract 60", operation: "subtract", value: 60, want: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.operation == "add" {
				add(tt.value)
			} else if tt.operation == "subtract" {
				subtract(tt.value)
			} else {
				t.Fatalf("Unknown operation: %s", tt.operation)
			}
			got := get()
			if got != tt.want {
				t.Errorf("After %s %d, get() = %v, want %v", tt.operation, tt.value, got, tt.want)
			}
		})
	}
}

func TestApply(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		operation func(int) int
		want      []int
	}{
		{"Square", []int{1, 2, 3}, func(x int) int { return x * x }, []int{1, 4, 9}},
		{"Double", []int{10, 20}, func(x int) int { return x * 2 }, []int{20, 40}},
		{"Negate", []int{1, -1}, func(x int) int { return -x }, []int{-1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.nums, tt.operation)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		want      []int
	}{
		{"Evens", []int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 }, []int{2, 4}},
		{"Positive", []int{-1, 0, 1}, func(x int) bool { return x > 0 }, []int{1}},
		{"Greater than 10", []int{5, 15, 20}, func(x int) bool { return x > 10 }, []int{15, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.nums, tt.predicate)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReduce(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		reducer func(int, int) int
		initial int
		want    int
	}{
		{"Sum", []int{1, 2, 3, 4}, func(acc, x int) int { return acc + x }, 0, 10},
		{"Product", []int{1, 2, 3, 4}, func(acc, x int) int { return acc * x }, 1, 24},
		{"Max", []int{1, 5, 3, 4}, func(acc, x int) int {
			if x > acc {
				return x
			}
			return acc
		}, 0, 5},
		{"Min", []int{4, 2, 8, 1}, func(acc, x int) int {
			if x < acc {
				return x
			}
			return acc
		}, 100, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.nums, tt.initial, tt.reducer)
			if got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	double := func(x int) int { return x * 2 }
	square := func(x int) int { return x * x }

	tests := []struct {
		name string
		f    func(int) int
		g    func(int) int
		in   int
		want int
	}{
		{"Square then Double: 2(x^2)", double, square, 3, 18}, // 2 * (3*3)
		{"Double then Square: (2x)^2", square, double, 3, 36}, // (2 * 3)^2
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			combined := Compose(tt.f, tt.g)
			if got := combined(tt.in); got != tt.want {
				t.Errorf("%s failed: got %d, want %d", tt.name, got, tt.want)
			}
		})
	}
}

func TestSwapValues(t *testing.T) {
	tests := []struct {
		name         string
		a, b         int
		wantA, wantB int
	}{
		{"Swap positives", 5, 10, 10, 5},
		{"Swap negatives", -1, -5, -5, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := SwapValues(tt.a, tt.b)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf("SwapValues(%d, %d) = %d, %d; want %d, %d",
					tt.a, tt.b, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}

func TestSwapPointers(t *testing.T) {
	tests := []struct {
		name         string
		a, b         int
		wantA, wantB int
	}{
		{"Swap positives", 5, 10, 10, 5},
		{"Swap negatives", -1, -5, -5, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ValA := tt.a
			ValB := tt.b
			SwapPointers(&ValA, &ValB)
			if ValA != tt.wantA || ValB != tt.wantB {
				t.Errorf("SwapPointers(%d, %d) = %d, %d; want %d, %d",
					tt.a, tt.b, ValA, ValB, tt.wantA, tt.wantB)
			}
		})
	}
}
