package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	M := 800000
	if len(os.Args) > 2 {
		M, _ = strconv.Atoi(os.Args[2])
	}

	log.Println("generating random numbers")
	buffer := make([]int, N)
	for i, _ := range buffer {
		buffer[i] = rand.Intn(M)
	}

	if N < 100 {
		for _, n := range buffer {
			fmt.Println(n)
		}
	}

	log.Println("sorting")
	begin := time.Now()
	ParallelQuicksort(&buffer, 0, N-1)
	end := time.Now()

	log.Println("finished in", end.Sub(begin))

	if N < 100 {
		for _, n := range buffer {
			fmt.Println(n)
		}
	}

	for i, _ := range buffer {
		if i == 0 {
			continue
		}
		if buffer[i-1] > buffer[i] {
			log.Fatal("didn't sort correctly.")
		}
	}
}
