package services_test

import (
	"challenge/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPath(t *testing.T) {

	t.Run("test example", func(t *testing.T) {
		input := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
		expected := 237
		actual := services.MaxPathSum(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("test max path", func(t *testing.T) {
		input := services.MaxPath("../files/hard.json")
		assert.Equal(t, 7273, input)
	})
}
