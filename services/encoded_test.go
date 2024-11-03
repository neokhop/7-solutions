package services_test

import (
	"fmt"
	"testing"

	"challenge/services"

	"github.com/stretchr/testify/assert"
)

func TestEncoded(t *testing.T) {
	testCases := []struct {
		pattern string
		result  string
	}{
		{"LLRR=", "210122"},
		{"==RLL", "000210"},
		{"=LLRR", "221012"},
		{"RRL=R", "012001"},
	}

	for _, testCase := range testCases {
		label := fmt.Sprintf("test case: %s", testCase.pattern)
		t.Run(label, func(t *testing.T) {
			result := services.Encoded(testCase.pattern)
			assert.Equal(t, testCase.result, result)
		})
	}
}
