package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadData(t *testing.T) {
	// Read input file from test resources
	input, err := os.ReadFile("data.json")
	if err != nil {
		t.Fatalf("failed to read input file: %v", err)
	}

	// Call function that reads from file
	output, err := readData("data.json")

	// Check that output matches expected output
	assert.Nil(t, err)
	assert.Equal(t, input, output)
}
