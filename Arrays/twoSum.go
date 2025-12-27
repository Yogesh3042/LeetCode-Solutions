// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	nums := []int{2, 4, 4}
	target := 6
	arr := twoSum(nums, target)
	fmt.Println(arr)
}
func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, v := range nums {
		diff := target - v

		if idx, ok := m[diff]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return []int{}
}
