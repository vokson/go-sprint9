package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	s := make([]int, size)
	for i := range size {
		s[i] = rand.Int()
	}
	return s
}

func maximum(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("slice is empty")
	}
	res := data[0]
	if len(data) == 1 {
		return res, nil
	}
	for i := 1; i < len(data); i++ {
		if data[i] > res {
			res = data[i]
		}
	}
	return res, nil
}

func maxChunks(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("slice is empty")
	}

	if len(data) <= CHUNKS {
		return maximum(data)
	}

	var wg sync.WaitGroup
	sliceMax := make([]int, CHUNKS)
	size := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			start := i * size
			end := start + size
			if i == CHUNKS-1 {
				end = len(data)
			}

			var err error
			sliceMax[i], err = maximum(data[start:end])
			if err != nil {
				fmt.Println(err)
				return
			}
		}(i)
	}

	wg.Wait()
	return maximum(sliceMax)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	s := generateRandomElements(SIZE)

	var start, finish time.Time
	var duration time.Duration

	start = time.Now()
	fmt.Println("Ищем максимальное значение в один поток")
	maxSingle, err := maximum(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	finish = time.Now()

	duration = finish.Sub(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxSingle, int(duration.Microseconds()))

	start = time.Now()
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	maxMulti, err := maxChunks(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	finish = time.Now()

	duration = finish.Sub(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxMulti, int(duration.Microseconds()))
}
