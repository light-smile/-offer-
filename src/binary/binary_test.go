/*
 * @Author: jlz
 * @Date: 2022-09-14 16:55:09
 * @LastEditors: jlz
 * @Description:
 */
package binary

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	words := []string{
		"abcw", "foo", "bar", "fxyz", "abcdef",
	}
	result := HashMax(words)
	fmt.Println(result)
}

func HashMax(words []string) (product int) {
	hashMap := make([]map[uint8]bool, len(words))
	var cur int
	for i := 0; i < len(words); i++ {
		hashMap[i] = make(map[uint8]bool, len(words[i]))
		for j := 0; j < len(words[i]); j++ {

			hashMap[i][words[i][j]] = false
		}
	}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			for _, b := range words[i] {
				if _, ok := hashMap[j][uint8(b)]; !ok {
					cur = len(words[i]) * len(hashMap[j])
					product = max(cur, product)
					fmt.Println(words[i], hashMap[j], cur, product)
				}

			}
		}
	}

	//for _, v := range words {
	//	for _, b := range v {
	//		for k, m := range hashMap {
	//			if v != k {
	//				if _, ok := m[b]; !ok {
	//					cur = len(v) * len(k)
	//					product = max(cur, product)
	//					fmt.Println(v, k, cur, product)
	//				}
	//			}
	//		}
	//	}
	//
	//}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
