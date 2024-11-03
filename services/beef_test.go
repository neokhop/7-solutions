package services_test

import (
	"challenge/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeef(t *testing.T) {
	rawWords := `Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.`
	counts := services.GetBeef(rawWords)

	expected := map[string]int{
		"t-bone":   4,
		"fatback":  1,
		"pastrami": 1,
		"pork":     1,
		"meatloaf": 1,
		"jowl":     1,
		"enim":     1,
		"bresaola": 1,
	}

	t.Run("test get beef", func(t *testing.T) {
		assert.Equal(t, expected, counts)
	})
}
