package Array

import (
	"fmt"
	"testing"
)

/*
题目：输入一个递增排序的数组和一个值k，请问如何在数组中找出两个和为k的数字并返回它们的下标？假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。例如，输入数组[1，2，4，6，10]，k的值为8，数组中的数字2与6的和为8，它们的下标分别为1与3。
*/
func TestTwoSum(t *testing.T) {
	arr := []int{
		1, 2, 4, 6, 10,
	}
	//r1, r2 := TwoSum(arr, 8)
	//r1, r2 := HashSum(arr, 8)
	r1, r2 := dichotomy(arr, 8)
	fmt.Println(r1, r2)
}

// 双指针
func TwoSum(arr []int, k int) (index1, index2 int) {
	var a, b = 0, len(arr) - 1
	for a < b {
		cur := arr[a] + arr[b]
		if cur == k {
			return arr[a], arr[b]
		}
		if cur > k {
			b--
		}
		a++
	}
	return
}

// 哈希表
func HashSum(arr []int, k int) (index1, index2 int) {
	h := make(map[int]int, len(arr))
	for _, v := range arr {
		h[v] = k - v
	}
	for k, v := range h {
		if _, ok := h[v]; ok {
			return k, v
		}
	}
	return
}

// 二分法
func dichotomy(arr []int, k int) (index1, index2 int) {
	for i := 0; i < len(arr); i++ {
		left := i + 1
		right := len(arr) - 1
		middle := (left + right) / 2
		for left < right {
			if (arr[i] + arr[middle]) == k {
				return arr[i], arr[middle]
			} else if arr[i]+arr[middle] > k {
				right = middle
				middle = (left + right) / 2
			} else {
				left = middle
				middle = (left + right) / 2
			}
		}
	}
	return
}

/*
题目：输入一个数组，如何找出数组中所有和为0的3个数字的三元组？需要注意的是，返回值中不得包含重复的三元组。例如，在数组[-1，0，1，2，-1，-4]中有两个三元组的和为0，它们分别是[-1，0，1]和[-1，-1，2]。
*/
func TestThreeSum(t *testing.T) {
	arr := []int{
		-4, -1, -1, 0, 1, 2,
	}
	res := ThreeSum(arr, 0)
	fmt.Println(res)
}

func ThreeSum(arr []int, k int) (res [][]int) {
	l := len(arr) // 数组长度
	var need int  // 剩余2位的和
	var temp int  // 记录开头数字，如果再次出现直接跳过，避免重复
	for i := 0; i < l-2; i++ {

		need = k - arr[i]                        //
		surplus := arr[i+1:]                     // 剩余需要判断的数组
		twoSum := TwoSum2(surplus, need, arr[i]) // 得到满足条件的数组

		if len(twoSum) != 0 {
			res = append(res, twoSum...)

		}
		temp = arr[i]
		for temp == arr[i+1] {
			i++
		}
	}
	return
}
func TwoSum2(arr []int, k int, header int) (res [][]int) {
	var a, b = 0, len(arr) - 1
	var temp int
	for a < b {
		cur := arr[a] + arr[b]
		if cur == k { // 判断是否满足条件
			accord := []int{
				header, arr[a], arr[b],
			}
			res = append(res, accord) // 满足条件添加到结果中
			temp = arr[a]
			for temp == arr[a+1] {
				a++
			}
			// 跳过当前值，避免重复
		} else if cur > k {
			b--
		} else {
			a++
		}
	}
	return
}

/*
乘积小于k的子数组
题目：输入一个正整数组成的数组和一个正整数k，请问数组中和大于或等于k的连续子数组的最短长度是多少？如果不存在所有数字之和大于或等于k的子数组，则返回0。例如，输入数组[5，1，4，3]，k的值为7，和大于或等于7的最短连续子数组是[4，3]，因此输出它的长度2。
*/
func TestShortestSum(t *testing.T) {
	arr := []int{
		1, 2, 3, 4, 5,
	}
	k := 15
	res := ShortestSum(arr, k)
	fmt.Println(res)
}
func ShortestSum(nums []int, target int) int {
	length := len(nums)
	shortest := length
	l, r := 0, 0
	curLen := 0
	cur := nums[0]
	flag := false

	for r < length {
		if cur >= target {
			tempCur := cur
			tempL := l
			flag = true
			for tempCur >= target {
				curLen = r - tempL + 1
				if curLen <= shortest {
					shortest = curLen

					l = tempL
					cur = tempCur
				}
				tempCur -= nums[tempL]
				tempL++
			}
		}
		if !flag {
			r++
			if r < length {
				cur += nums[r]
			}
		} else {
			cur -= nums[l]
			l++
			r++
			if r < length {
				cur += nums[r]
			}
		}

	}
	if flag {
		return shortest
	}
	return 0
}
