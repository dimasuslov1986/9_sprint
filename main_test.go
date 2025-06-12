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

	data := []struct {
		count []int
		want  int
	}{
		{make([]int, 0), 0},
		{[]int{0, 1, 2, 3, 4, 5}, 5},
	}

	for _, v := range data {
		x := maximum(v.count)
		assert.Equal(t, v.want, x)
	}
}

func TestMaxChunks(t *testing.T) {
	data := []struct {
		count []int
		want  int
	}{
		{make([]int, 0), 0},
		{[]int{0, 1, 2, 3, 4, 5}, 5},
	}

	for _, v := range data {
		x := maxChunks(v.count)
		assert.Equal(t, v.want, x)
	}
}
