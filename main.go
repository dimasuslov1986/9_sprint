package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		fmt.Println("ошибочное значение размера ожидаемого слайса случайных чисел")
		return nil
	}
	src := rand.NewSource(time.Now().Unix())
	var sl []int
	for i := 0; i < size; i++ {
		// получаем случайное число
		randomNumber := src.Int63()
		sl = append(sl, int(randomNumber))
	}
	return sl
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	// в задании сказано: Обязательно обработайте крайние случаи, например, если был передан размер слайса,
	// равный 0. но это же тоже не является проблемой. код отработает и вернет 0
	if len(data) == 0 {
		fmt.Println("ошибочный размер слайса")
		return 0
	}

	var max int
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		fmt.Println("ошибочный размер слайса")
		return 0
	}

	var wg sync.WaitGroup

	var arMaxCh [CHUNKS]int
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		go func() {
			slCh := data[i*len(data)/CHUNKS : (i+1)*len(data)/CHUNKS]

			maxCh := maximum(slCh)

			arMaxCh[i] = maxCh

			wg.Done()
		}()
	}
	wg.Wait()

	slMaxCh := arMaxCh[:CHUNKS]

	max := maximum(slMaxCh)

	return max
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	sl := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(sl)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(sl)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
