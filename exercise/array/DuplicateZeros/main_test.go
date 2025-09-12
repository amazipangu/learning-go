package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateZeros(t *testing.T) {
	t.Run("Duplicate each occurrence of zero, shifting the remaining elements to the right", func(t *testing.T) {
		arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
		expect := []int{1, 0, 0, 2, 3, 0, 0, 4}
		assert.Equal(t, expect, duplicatedZeros(arr))
	})

	t.Run("There is no shift to the right if 0 wouldn't appear in the arr", func(t *testing.T) {
		arr := []int{1, 2, 3}
		expect := []int{1, 2, 3}
		assert.Equal(t, expect, duplicatedZeros(arr))
	})
}
