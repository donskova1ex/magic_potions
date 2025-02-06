package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSqareNegative(t *testing.T) {
	n := -5
	expected := 25

	actual := square(n)

	require.Equal(t, expected, actual)
}
