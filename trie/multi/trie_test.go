package multi

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"testing"
)

func decode(s string) []string {
	return strings.Split(s, " ")
}

func TestNode(t *testing.T) {
	root := New[string, int]()
	root.Add(decode("hello"), 1)
	root.Add(decode("nice to meet you"), 1)
	root.Add(decode("nice boat"), 1, 2, 3, 4)
	root.Add(decode("hello world"), 2)
	root.Add(decode("hello how are you"), 1)
	root.Add(decode("hello world"), 4)
	root.Add(decode("bravo"), 5)
	root.Add(decode("bravo"), 5)

	t.Run("some", func(t *testing.T) {
		got := root.Get(decode("hello world"))
		want := []int{2, 4}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("all", func(t *testing.T) {
		got := make(map[string][]int)
		for k, v := range root.All() {
			fmt.Printf("%v -> %v\n", k, v)
			got[strings.Join(k, " ")] = v
		}
		want := map[string][]int{
			"hello":             {1},
			"hello how are you": {1},
			"hello world":       {2, 4},
			"nice boat":         {1, 2, 3, 4},
			"nice to meet you":  {1},
			"bravo":             {5, 5},
		}
		if !maps.EqualFunc(got, want, slices.Equal) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
