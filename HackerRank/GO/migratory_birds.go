package main

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	var pMap = make(map[int32]int32)
	for i := 1; i <= 5; i++ {
		pMap[int32(i)] = 0
	}

	for _, val := range arr {
		pMap[int32(val)]++
	}

	maxRepeat := int32(0)
	var minVal int32 = 5
	for key, val := range pMap {
		if val > maxRepeat {
			maxRepeat = val
			minVal = key
		}
	}

	return minVal
}
