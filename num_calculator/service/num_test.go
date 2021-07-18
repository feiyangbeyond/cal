package service

import (
	"fmt"
	"testing"
)

func TestCalAll(t *testing.T) {
	var nums = make([]int, 0)
	nums = append(nums, 1, 123, 241, 120, 487, 11, 90)
	all := Cal(nums)
	fmt.Println(all)
}

func TestT(t *testing.T) {
	maxCol:=8
	for i := 0; i < maxCol; i++ {
		a := ranks[len(ranks)-i-1]
		b := string(byte(66 + i*2))
		cell := b+"1"

		fmt.Print(cell)
		fmt.Println(a)

	}
}

func TestCalOne1(t *testing.T) {

	println(CalOne1(1, "个"))
}
