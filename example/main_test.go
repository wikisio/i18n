package main

import (
	"testing"

	"github.com/wikisio/i18n/i18n"
)

func TestName(t *testing.T) {
	i18n.MustInit("locale", &fs)

	testCases := []struct {
		keyId    string
		lng      []string
		expected string
	}{
		{
			keyId:    "helloId",
			lng:      []string{"fr", "en-US"},
			expected: "hello",
		},
		{
			keyId:    "helloId",
			lng:      []string{"en-US", "fr,zh-CN;q=0.7,en;q=0.8"},
			expected: "hello",
		},
		{
			keyId:    "helloId",
			lng:      []string{"fr", "fr,zh-CN;q=0.7,en;q=0.8"},
			expected: "hello",
		},
		{
			keyId:    "helloId",
			lng:      []string{"fr", "fr,zh-CN;q=0.9,en;q=0.8"},
			expected: "你好",
		},
		{
			keyId:    "helloId-X",
			lng:      []string{"fr", "fr,zh-CN;q=0.9,en;q=0.8"},
			expected: "helloId-X",
		},
	}

	for i, tc := range testCases {
		if got := i18n.Get(tc.keyId, tc.lng...); got != tc.expected {
			t.Errorf("Testcase %d, Expected %s but got %s", i, tc.expected, got)
		}
	}
}

func BenchmarkName(b *testing.B) {
	i18n.MustInit("locale", &fs)
	for i := 0; i < b.N; i++ {

		testCases := []struct {
			keyId    string
			lng      []string
			expected string
		}{
			{
				keyId:    "helloId",
				lng:      []string{"fr", "en-US"},
				expected: "hello",
			},
			{
				keyId:    "helloId",
				lng:      []string{"en-US", "fr,zh-CN;q=0.7,en;q=0.8"},
				expected: "hello",
			},
			{
				keyId:    "helloId",
				lng:      []string{"fr", "fr,zh-CN;q=0.7,en;q=0.8"},
				expected: "hello",
			},
			{
				keyId:    "helloId",
				lng:      []string{"fr", "fr,zh-CN;q=0.9,en;q=0.8"},
				expected: "你好",
			},
			{
				keyId:    "helloId-X",
				lng:      []string{"fr", "fr,zh-CN;q=0.9,en;q=0.8"},
				expected: "helloId-X",
			},
		}

		for i, tc := range testCases {
			if got := i18n.Get(tc.keyId, tc.lng...); got != tc.expected {
				b.Errorf("Testcase %d, Expected %s but got %s", i, tc.expected, got)
			}
		}
	}
}
