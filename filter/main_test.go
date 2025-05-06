package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"slices"
	"testing"
)

func WrapperV0(nums []int, minValue int) []int {
	return FilterV0(nums, minValue)
}

func WrapperV1(nums []int, minValue int) []int {
	return FilterV1(nums, GreaterEqual(minValue))
}

func WrapperV2(nums []int, minValue int) []int {
	return FilterV2(nums, GreaterEqual(minValue))
}

func WrapperV3(nums []int, minValue int) []int {
	return slices.Collect(FilterV3(slices.Values(nums), GreaterEqual(minValue)))
}

type args struct {
	nums     []int
	minValue int
}

func generateArgs(size int) []args {
	random := rand.New(rand.NewPCG(1, 2))
	var ret []args
	for range size {
		nums := make([]int, random.IntN(1000))
		for i := range nums {
			nums[i] = random.IntN(1_000_000)
		}
		var minValue int
		if len(nums) > 0 {
			minValue = nums[random.IntN(len(nums))]
		}
		ret = append(ret, args{
			nums:     nums,
			minValue: minValue,
		})
	}
	return ret
}

var wrappers = []func([]int, int) []int{
	WrapperV0,
	WrapperV1,
	WrapperV2,
	WrapperV3,
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"happy path", args{[]int{1, 2, 3, 4, 5}, 4}, []int{4, 5}},
		{"full", args{[]int{1, 3, 4, 5}, 0}, []int{1, 3, 4, 5}},
		{"none", args{[]int{1, 2, 3, 5}, 6}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, wrapper := range wrappers {
				t.Run(fmt.Sprintf("v=%d", i), func(t *testing.T) {
					if got := wrapper(tt.args.nums, tt.args.minValue); !reflect.DeepEqual(got, tt.want) {
						t.Errorf("got %v, want %v", got, tt.want)
					}
				})
			}
		})
	}

	inputs := generateArgs(100)
	t.Run("random", func(t *testing.T) {
		for _, input := range inputs {
			v0 := WrapperV0(input.nums, input.minValue)
			v1 := WrapperV1(input.nums, input.minValue)
			v2 := WrapperV2(input.nums, input.minValue)
			v3 := WrapperV3(input.nums, input.minValue)
			if !reflect.DeepEqual(v0, v1) {
				t.Errorf("v0 %v != v1 %v", v0, v1)
			}
			if !reflect.DeepEqual(v0, v2) {
				t.Errorf("v0 %v != v1 %v", v0, v2)
			}
			if !reflect.DeepEqual(v0, v3) {
				t.Errorf("v0 %v != v1 %v", v0, v3)
			}
		}
	})
}

func BenchmarkFilter(b *testing.B) {
	inputs := generateArgs(1000)
	for i, wrapper := range wrappers {
		b.Run(fmt.Sprintf("v=%d", i), func(b *testing.B) {
			for b.Loop() {
				for _, input := range inputs {
					wrapper(input.nums, input.minValue)
				}
			}
		})
	}
}
