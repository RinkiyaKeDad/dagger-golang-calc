package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	assert.Equal(t, 5, result, "Add(2, 3) should equal 5")
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 2)
	assert.Equal(t, 3, result, "Subtract(5, 2) should equal 3")
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)
	assert.Equal(t, 12, result, "Multiply(3, 4) should equal 12")
}

func TestDivide(t *testing.T) {
	// Test valid division
	result, err := Divide(10, 2)
	assert.NoError(t, err, "Divide(10, 2) should not return an error")
	assert.Equal(t, 5, result, "Divide(10, 2) should equal 5")

	// Test division by zero
	_, err = Divide(10, 0)
	assert.Error(t, err, "Divide(10, 0) should return an error for division by zero")
}
