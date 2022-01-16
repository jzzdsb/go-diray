/*
 * @Author: vic
 * @Date: 2022-01-14 17:40:46
 * @LastEditors: vic
 * @LastEditTime: 2022-01-16 10:17:52
 * @Description:
 */
package main

import "fmt"

/**
 * @author : vic
 * @description :最大前缀
 * 1 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] 仅由小写英文字母组成
 * @createDate : 2022/1/14 4:19 下午
 * @params : [strs]
 **/
func longestCommonPrefix(strs []string) string {

	if len(strs) == 1 {
		return strs[0]
	}

	temp := strs[0]
	index := 0
	key := true
w:
	for true {
		for i := 1; i < len(strs); i++ {

			if index >= len(strs[i]) {
				index--
				break w
			}

			if temp[index] != strs[i][index] {
				index--
				break w
			}
		}
		key = false
		index++
	}

	if key {
		return ""
	} else {
		return strs[0][0 : index+1]
	}

}

func isValid(s string) bool {

	if len(s)%2 == 1 {
		return false
	}

	m := map[byte]byte{'(': ')', '[': ']', '{': '}'}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		_, k := m[s[i]]
		if k {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			} else {
				b := stack[len(stack)-1]
				if m[b] == s[i] {
					stack = stack[1:]
				} else {
					return false
				}
			}
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Printf("%s",isValid("()"))
}
