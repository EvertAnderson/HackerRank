package main

/*
 * Complete the 'divisibleSumPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY ar
 */

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	count := 0
	for i := int32(0); i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return int32(count)
}
