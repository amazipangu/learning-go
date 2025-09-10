package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	t.Run("Return an array of the squares of each number", func(t *testing.T) {
		nums := []int{12, 345, 2, 6, 7896}
		expect := []int{4, 36, 144, 119025, 62346816}
		assert.Equal(t, expect, sortedSquares(nums))
	})
}
