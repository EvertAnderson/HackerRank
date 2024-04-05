package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	var result int32
	result = 0

	var maxV, minV float64

	if len(a) > 0 && len(b) > 0 {
		maxV = float64(a[0])
		minV = float64(b[0])

		for _, val := range a {
			maxV = math.Max(float64(val), float64(maxV))
		}

		for _, val := range b {
			minV = math.Min(float64(val), float64(minV))
		}

		fmt.Printf("1) Max: %.1f - Min: %.1f", maxV, minV)
	}

	for i := int32(maxV); i < int32(minV+1); i++ {
		isFactorMultiple := true

		for _, val := range a {
			if i%val != 0 {
				isFactorMultiple = false
				break
			}
		}

		for _, val := range b {
			if val%i != 0 {
				isFactorMultiple = false
				break
			}
		}

		if isFactorMultiple == true {
			result++
		}
	}

	return int32(result)
}
