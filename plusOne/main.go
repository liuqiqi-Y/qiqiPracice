// 数组加1。
/*
	给定一个由非负整数组成的数组，其中每个元素只存储单个
	整数，然后对这个数组进行加1运算。
	例：
	输入：[1,2,3]
	输出：[1,2,4]
*/
package main

import "fmt"

func plusOne(arr []int) []int {
	if arr == nil {
		return nil
	}
	for j := len(arr) - 1; j >= 0; j-- {
		if arr[j] < 9 {
			arr[j]++
			return arr
		}
		arr[j] = 0
	}
	return append([]int{1}, arr...)
}
func main() {
	arr := []int{1, 2, 3}
	result := plusOne(arr)
	fmt.Printf("%v\n", result)
}
