package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestItShouldReturnAJsonFromError(t *testing.T) {
	msg := "Hello Json Error Message"
	expected := []byte(`{"message":"Hello Json Error Message"}`)

	result := jsonError(msg)

	require.Equal(t, expected, result)
}
