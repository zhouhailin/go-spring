package lic

import "testing"

func TestTo(t *testing.T) {
	ToXml(map[string]any{
		"hello": "world",
		"je":    "x",
		"hello1": map[string]any{
			"tom": "cat",
		},
	}, nil)
}
