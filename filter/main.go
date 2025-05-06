package main

import (
	"fmt"
	"github.com/hyisen/how-to-go/trie"
	"iter"
	"maps"
	"slices"
	"strings"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	minValue := 2

	// If you use FilterV3 only as WrapperV3, the abstraction is redundant.
	_ = slices.Collect(FilterV3(slices.Values(nums), GreaterEqual(minValue)))

	fmt.Println(nums)
	// First, using the iter directly in the destination could reduce memory footage.
	for num := range FilterV3(slices.Values(nums), GreaterEqual(minValue)) {
		fmt.Println(num)
	}

	var nameToScore = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	fmt.Println(nameToScore)
	// Secondly, it works on any kind of container that could provide iterators.
	for score := range FilterV3(maps.Values(nameToScore), GreaterEqual(minValue)) {
		fmt.Println(score)
	}

	data := map[string]int{
		"hello":             1,
		"hello how are you": 2,
		"hello world":       3,
		"nice boat":         4,
		"nice to meet you":  5,
		"bravo":             6,
	}
	root := trie.New[string, int]()
	for str, num := range data {
		root.Set(num, strings.Split(str, " ")...)
	}
	for k, v := range root.All() {
		fmt.Printf("%v -> %v\n", k, *v)
	}
	// Any kind of container.
	for num := range FilterV3(root.Values(), GreaterEqual(minValue)) {
		fmt.Println(num)
	}
}

func FilterV0(nums []int, minValue int) []int {
	var ret []int
	for _, num := range nums {
		if num >= minValue {
			ret = append(ret, num)
		}
	}
	return ret
}

func FilterV1(nums []int, validator func(num int) bool) []int {
	var ret []int
	for _, num := range nums {
		if validator(num) {
			ret = append(ret, num)
		}
	}
	return ret
}

func FilterV2[T any](nums []T, validator func(num T) bool) []T {
	var ret []T
	for _, num := range nums {
		if validator(num) {
			ret = append(ret, num)
		}
	}
	return ret
}

func FilterV3[T any](seq iter.Seq[T], validator func(num T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if validator(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func GreaterEqual(minValue int) func(num int) bool {
	return func(num int) bool {
		return num >= minValue
	}
}
