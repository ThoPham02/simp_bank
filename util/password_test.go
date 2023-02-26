package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := RandomString(8)

	hashPassword, err := HashPassword(password)
	require.NoError(t, err)

	isCompare := ComparePassword(password, hashPassword)
	require.True(t, isCompare)

	wrongPassword := RandomString(8)

	isNotCompare := ComparePassword(wrongPassword, hashPassword)
	require.False(t, isNotCompare)
}
