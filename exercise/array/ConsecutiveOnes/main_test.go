package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	t.Run("expect the correct number", func(t *testing.T) {
		var nums []int = []int{0, 1, 0, 1, 1, 1, 0}
		assert.Equal(t, 3, findMaxConsecutiveOnes(nums))
	})
}
