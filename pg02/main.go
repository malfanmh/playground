package main

import (
	"fmt"
)

// Series interface obj
type Series interface {
	CalcMedian() float64
}

// Median :
type Median struct {
	Nums1 []float64
	Nums2 []float64
}

// CalcMedian : calculate median by combined arrays
func (i Median) CalcMedian() float64 {
	//concate nums1 + nums2
	nums := append(i.Nums1, i.Nums2...)
	// length of nums
	ln := len(nums)
	if ln == 0 {
		// "blank input"
		return 0
	} else if ln%2 != 0 {
		//odd case
		return nums[ln/2]
	}
	return (nums[((ln)/2)-1] + nums[ln/2]) / 2
}
func main() {
	series := &Median{
		Nums1: []float64{1, 2, 3, 4},
		Nums2: []float64{5, 6, 7, 8},
	}
	fmt.Println(series.CalcMedian())
}
