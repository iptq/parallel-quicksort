package main

import (
	"sync"
)

func ParallelQuicksort(arr *[]int, left, right int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go parallelQuicksort(arr, left, right, &wg)
	wg.Wait()
}

func parallelQuicksort(arr *[]int, left, right int, pwg *sync.WaitGroup) {
	defer pwg.Done()

	var wg sync.WaitGroup
	piv := partition(arr, left, right)
	if left < piv-1 {
		wg.Add(1)
		go parallelQuicksort(arr, left, piv-1, &wg)
	}
	if piv < right {
		wg.Add(1)
		go parallelQuicksort(arr, piv, right, &wg)
	}
	wg.Wait()
}

func partition(arr *[]int, left, right int) int {
	piv := (*arr)[(left+right)/2]

	for left <= right {
		for (*arr)[left] < piv {
			left++
		}
		for (*arr)[right] > piv {
			right--
		}
		if left <= right {
			// swap
			t := (*arr)[left]
			(*arr)[left] = (*arr)[right]
			(*arr)[right] = t

			left++
			right--
		}
	}

	return left
}
