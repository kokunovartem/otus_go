package main

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestAtoi(t *testing.T) {
	const str, want = "42", 42
	got, err := strconv.Atoi(str)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
