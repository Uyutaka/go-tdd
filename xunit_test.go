package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	t.Run("should return true if testMethod is called", func(t *testing.T) {
		test := NewWasRun("TestMethod")
		assert.False(t, test.WasRun)
		test.Run()
		assert.True(t, test.WasRun)
	})
}
