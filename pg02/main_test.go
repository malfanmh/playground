package main

import "testing"

func TestCalcMedian(t *testing.T) {
	// test table
	tt := []struct {
		name string  // test case
		in   Series  // input process
		want float64 // expected result
	}{
		{"odd array of [1,2] and [3]", &Median{Nums1: []float64{1, 2}, Nums2: []float64{3}}, 2},
		{"even array of [1,2] and [3,4]", &Median{Nums1: []float64{1, 2}, Nums2: []float64{3, 4}}, 2.5},
		{"blank input", &Median{Nums1: []float64{}, Nums2: []float64{}}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.in.CalcMedian()
			if res != tc.want {
				t.Fatalf("Median of %v should be %v; got %v", tc.name, tc.want, res)
			}
		})
	}
}
