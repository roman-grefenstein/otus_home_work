package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	env := make(map[string]EnvValue, 2)
	env["FOO"] = EnvValue{Value: "FOO_ENV_VALUE"}
	env["BAR"] = EnvValue{Value: "BAR_ENV_VALUE", NeedRemove: true}

	var tests = []struct {
		name   string
		reqCmd []string
		reqEnv Environment
		result int
	}{
		{"successful", []string{"printenv", "FOO"}, env, 0},
		{"error 1", []string{"printenv", "BAR"}, env, 1},
		{"error 2", []string{"printenv", "BAR"}, make(map[string]EnvValue), 1},
		{"misuse of shell builtins", []string{"ls", "d"}, env, 2},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := RunCmd(tc.reqCmd, tc.reqEnv)
			require.Equal(t, tc.result, result)
		})
	}
}
