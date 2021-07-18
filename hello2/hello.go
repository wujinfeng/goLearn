package main

import (
	"fmt"

	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func fibonacci(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}
	nums := make([]int, n)
	nums[0], nums[1] = 1, 1
	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums
}

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	var num int
	fmt.Println("输入数字：")
	fmt.Scanln(&num)
	fmt.Println("输出：", fibonacci(num))

}
