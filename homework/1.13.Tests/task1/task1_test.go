package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	result := Divide(45.45, 12.12)

	assert.NotZero(t, result, "result should not be zero")
	assert.Equalf(t, 45.45/12.12, result, "%v divided by %v should be %v", 45.45, 12.12, 45.45/12.12)
}

func TestDivideByZero(t *testing.T) {
	assert.Panics(t, func() {
		_ = Divide(45.45, 0)
	})
}
