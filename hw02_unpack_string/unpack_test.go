package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abccd",
			expected: "abccd",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "aaa0b",
			expected: "aab",
		},
		{
			input:    "aaab0",
			expected: "aaa",
		},
		{
			input:    "_W5_",
			expected: "_WWWWW_",
		},
		{
			input:    "_1*2+3=4~0",
			expected: "_**+++====",
		},
		{
			input:    "d\n5abc",
			expected: "d\n\n\n\n\nabc",
		},
		{
			input:    "👍3👌2👎0",
			expected: "👍👍👍👌👌",
		},
		{
			input:    "こ1ん2に3ち4は5!",
			expected: "こんんにににちちちちははははは!",
		},
		{
			input:    `/2\2m/2\2`,
			expected: `//\\m//\\`,
		},
		{
			input:    `\ 3/`,
			expected: `\   /`,
		},
		{
			input:    "\n\r3",
			expected: "\n\r\r\r",
		},
		{
			input:    " 0",
			expected: "",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "aaab00", " 09"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
