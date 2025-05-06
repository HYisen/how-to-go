package trie

import (
	"maps"
	"slices"
	"strings"
	"testing"
)

func decode(s string) []string {
	return strings.Split(s, " ")
}

func TestNode(t *testing.T) {
	data := map[string]int{
		"hello":             1,
		"hello how are you": 2,
		"hello world":       3,
		"nice boat":         4,
		"nice to meet you":  5,
		"bravo":             6,
	}
	root := New[string, int]()
	for str, num := range data {
		root.Set(num, decode(str)...)
	}

	t.Run("some", func(t *testing.T) {
		str := "hello world"
		got := root.Get(decode(str)...)
		want := data[str]
		if got == nil || *got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("all", func(t *testing.T) {
		got := make(map[string]int)
		for k, v := range root.All() {
			got[strings.Join(k, " ")] = *v
		}

		want := data
		if !maps.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Values", func(t *testing.T) {
		got := slices.Sorted(root.Values())
		want := slices.Sorted(maps.Values(data))
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
