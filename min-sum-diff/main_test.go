package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalc(t *testing.T) {
	values := [][2][]int{
		{{1, 1}, {1}},
		{{1, 2, 2}, {2}},
		{{1, 1, 2, 3, 5, 6, 8}, {5, 8}},
	}
	for _, v := range values {
		res := calc(v[0])
		if !reflect.DeepEqual(res, v[1]) {
			t.Error("error:", v, res)
		}
	}
}

func TestCombinations(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		v := []int{1, 2, 3}
		n := 2
		res := [][]int{
			{1, 2}, {1, 3}, {2, 3},
		}
		combi := combinations(v, n)
		if !reflect.DeepEqual(combi, res) {
			t.Error("unexpected combinations:", v, res)
		}
	})
	t.Run("2", func(t *testing.T) {
		v := []int{1, 2, 3, 4}
		n := 3
		res := [][]int{
			{1, 2, 3}, {1, 2, 4}, {1, 3, 4},
			{2, 3, 4},
		}
		combi := combinations(v, n)
		if !reflect.DeepEqual(combi, res) {
			t.Error("unexpected combinations:", v, res)
		}
	})
	t.Run("3", func(t *testing.T) {
		v := []int{1, 2, 3, 4}
		n := 2
		res := [][]int{
			{1, 2}, {1, 3}, {2, 3}, {1, 4}, {2, 4}, {3, 4},
		}
		combi := combinations(v, n)
		if !reflect.DeepEqual(combi, res) {
			t.Error("unexpected combinations:", v, "res:", combi, "expected:", res)
		}
	})
}

func BenchmarkCalc(b *testing.B) {
	v := []int{1, 1}
	for n := 2; n < 8; n++ {
		v = append(v, n)
		b.Run(fmt.Sprintf("num: %d", len(v)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				calc(v)
			}
		})
	}
}
