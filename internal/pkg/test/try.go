package test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Try[T any](t *testing.T, fn func() (T, error)) T {
	val, err := fn() // Выполняем переданную функцию
	require.Nil(t, err)
	return val
}
