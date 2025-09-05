package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNumbers(t *testing.T) {
	t.Run("Check how many number of even number of digits in array", func(t *testing.T) {
		nums := []int{12, 345, 2, 6, 7896}
		assert.Equal(t, 2, findNumbers(nums))
	})
}
