package main

import (
	"flag"
	"fmt"
	"math"
	"math/bits"
	"strconv"
)

func main() {
	flag.Parse()
	nums := argInt(flag.Args())
	fmt.Println(nums)
	fmt.Println(calc(nums))
}

func calc(nums []int) (result []int) {
	s := sum(nums)
	target := s / 2
	max := math.MinInt
	for i := 1; i <= len(nums); i++ {
		for _, c := range combinations(nums, i) {
			ss := sum(c)
			if ss > max && ss <= target {
				max = ss
				result = c
			}
		}
	}
	return
}

func argInt(numsStr []string) (nums []int) {
	for _, s := range numsStr {
		n, _ := strconv.Atoi(s)
		if n != 0 {
			nums = append(nums, n)
		}
	}
	return nums
}

func sum(v []int) (sum int) {
	for _, vv := range v {
		sum += vv
	}
	return
}

// https://github.com/mxschmitt/golang-combinations
func combinations[T any](array []T, n int) (combis [][]T) {
	length := uint(len(array))
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}
		var c []T
		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				c = append(c, array[object])
			}
		}
		combis = append(combis, c)
	}
	return combis
}
