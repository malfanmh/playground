package pg01

import (
	"testing"
)

func TestSum(t *testing.T) {
	tt := []struct {
		name    string
		numbers []int
		result  int
	}{
		{"one and two", []int{1, 2}, 3},
		{"no numbers", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := Sum(tc.numbers)
			if res != tc.result {
				t.Fatalf("sum of %v should be %v; got %v", tc.name, tc.result, res)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	//test table
	tt := []struct {
		name    string
		numbers []int
		result  int
	}{
		{"one and two", []int{1, 2}, 2},
		{"no numbers", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}
	//loop test case
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := Multiply(tc.numbers)
			if res != tc.result {
				t.Fatalf("multiply of %v should be %v; got %v", tc.name, tc.result, res)
			}
		})
	}
}

func TestFiboSeries(t *testing.T) {
	//test table
	tt := []struct {
		name   string
		n      int
		series []int
	}{
		{"fibo series ", 4, []int{0, 1, 1, 2}},
	}
	//loop test case
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := FiboSeries(tc.n)
			if !eqSlice(res, tc.series) {
				t.Fatalf("the first %v of fibonnaci series should be %v; got %v", tc.n, tc.series, res)
			}
		})
	}
}

func TestPrimeSeries(t *testing.T) {
	//test table
	tt := []struct {
		name   string
		n      int
		series []int
	}{
		{"prime series ", 4, []int{2, 3, 5, 7}},
	}
	//loop test case
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := PrimeSeries(tc.n)
			if !eqSlice(res, tc.series) {
				t.Fatalf("the first %v of Prime series should be %v; got %v", tc.name, tc.series, res)
			}
		})
	}
}

// compare between two slice of int
func eqSlice(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
