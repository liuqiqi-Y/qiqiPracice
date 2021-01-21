// 一个数组和一个给定值，求出数组中两数之和等于给定值的对应数组下标
package twoSum

type index struct {
	index1 int
	index2 int
}

func twoSum1(arr []int, target int) []index {
	if len(arr) < 2 {
		return nil
	}
	result := []index{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == (target - arr[i]) {
				result = append(result, index{i, j})
			}
		}
	}
	return result
}
func twoSum2(arr []int, target int) []index {
	if len(arr) < 2 {
		return nil
	}
	m := make(map[int]int)
	result := []index{}
	for i := 0; i < len(arr); i++ {
		if in, ok := m[target-arr[i]]; ok {
			result = append(result, index{i, in})
		} else {
			m[arr[i]] = i
		}
	}
	return result
}
