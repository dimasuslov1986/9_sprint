package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	sl := generateRandomElements(5)
	assert.Len(t, sl, 5)
	assert.NotEqual(t, len(sl), 0)
	sl1 := generateRandomElements(-5)
	require.Equal(t, []int(nil), sl1)
}

func TestMaximum(t *testing.T) {

	x := maximum(make([]int, 0))
	assert.Equal(t, 0, x)

	x1 := maximum(make([]int, 1))
	assert.Equal(t, 0, x1)

	x2 := maximum(make([]int, 3))
	assert.Equal(t, 0, x2)

	sl3 := []int{0, 1, 2, 3, 4, 5}
	x3 := maximum(sl3)
	assert.Equal(t, 5, x3)
}

func TestMaxChunks(t *testing.T) {
	x := maximum(make([]int, 0))
	assert.Equal(t, 0, x)

	x1 := maximum(make([]int, 1))
	assert.Equal(t, 0, x1)

	x2 := maximum(make([]int, 3))
	assert.Equal(t, 0, x2)

	sl4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}
	x4 := maximum(sl4)
	assert.Equal(t, 23, x4)
}
