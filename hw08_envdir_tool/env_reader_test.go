package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	env, err := ReadDir("./testdata/env")

	require.Nil(t, err)

	require.NotNil(t, env)

	require.Equal(t, 5, len(env))

	require.False(t, env["HELLO"].NeedRemove)
	require.Equal(t, "\"hello\"", env["HELLO"].Value)

	require.True(t, env["EMPTY"].NeedRemove)
	require.Empty(t, env["EMPTY"].Value)

	require.True(t, env["UNSET"].NeedRemove)
	require.Empty(t, env["UNSET"].Value)

	require.Contains(t, env["FOO"].Value, "\n")
}

func TestReadDirInvalidPath(t *testing.T) {
	invalidPaths := []string{"/100500", "\\/\\/\\/", "main.go", "&"}

	for _, tc := range invalidPaths {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			env, err := ReadDir(tc)
			require.NotNil(t, err)
			require.Nil(t, env)
		})
	}
}
