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
	if len(data) < 2 {
		fmt.Println("ошибочный размер слайса")
		return 0
	}

	var max int
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	if max == 0 {
		fmt.Println("ошибка: пустой слайс")
		return 0
	}
	return max
}

var mu sync.Mutex

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) < 2 {
		fmt.Println("ошибочный размер слайса")
		return 0
	}

	// решил убрать проверку на "< 16" и "%8" потому что даже при невыполнении этих условий код работает корректно
	// и находит наибольшее значение. будть то слайс {0, 1} или любой другой. ниже добавил проверку если все значения слайса == 0

	//if len(data) < 16 {
	//	fmt.Println("ошибочный размер слайса, количество элементов слайса недостаточно для корректного сравнения в 8 горутинах")
	//	return 0
	//}
	//if len(data)%8 != 0 {
	//	fmt.Println("ошибочный размер слайса, количество элементов слайса не кратно 8")
	//	return 0
	//}
	var wg sync.WaitGroup
	var max int
	var slMaxCh []int
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		go func() {
			slCh := data[i*len(data)/CHUNKS : (i+1)*len(data)/CHUNKS]
			var maxCh int
			for j := 0; j < len(slCh); j++ { // или range лучше?
				if slCh[j] > maxCh {
					maxCh = slCh[j]
				}
			}
			mu.Lock()
			slMaxCh = append(slMaxCh, maxCh)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	for h := 0; h < len(slMaxCh); h++ {
		if slMaxCh[h] > max {
			max = slMaxCh[h]
		}
	}
	if max == 0 {
		fmt.Println("ошибка: пустой слайс")
		return 0
	}

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
	elapsed := time.Since(start) / time.Millisecond

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(sl)
	elapsed = time.Since(start) / time.Millisecond
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
