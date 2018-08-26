package main

import "fmt"

/*

893. 特殊等价字符串组
用户通过次数 47
用户尝试次数 60
通过次数 48
提交次数 83
题目难度 Easy
你将得到一个字符串数组 A。

如果经过任意次数的移动，S == T，那么两个字符串 S 和 T 是特殊等价的。



一次移动包括选择两个索引 i 和 j，且 i％2 == j％2，并且交换 S[j] 和 S [i]。

现在规定，A 中的特殊等效字符串组是 A 的非空子集，这样任何不在 A 中的字符串都不与 A 中的任何字符串特殊等效。



返回 A 中特殊等效字符串组的数量。



示例 1：

输入：["a","b","c","a","c","c"]
输出：3
解释：3 组 ["a","a"]，["b"]，["c","c","c"]
示例 2：

输入：["aa","bb","ab","ba"]
输出：4
解释：4 组 ["aa"]，["bb"]，["ab"]，["ba"]
示例 3：

输入：["abc","acb","bac","bca","cab","cba"]
输出：3
解释：3 组 ["abc","cba"]，["acb","bca"]，["bac","cab"]
示例 4：

输入：["abcd","cdab","adcb","cbad"]
输出：1
解释：1 组 ["abcd","cdab","adcb","cbad"]


提示：

1 <= A.length <= 1000
1 <= A[i].length <= 20
所有 A[i] 都具有相同的长度。
所有 A[i] 都只由小写字母组成。
*/
func isValidSpecialStrings(A, B string) bool {
	if len(A) != len(B) || len(A) == 0 {
		return false
	}

	if A == B {
		return true
	}

	jiA := make(map[int]int)
	ouA := make(map[int]int)

	jiB := make(map[int]int)
	ouB := make(map[int]int)

	for i := 0; i < len(A); i += 2 {
		jiA[int(A[i]-'a')] += 1
		jiB[int(B[i]-'a')] += 1
	}

	if len(A) > 1 {
		for i := 1; i < len(A); i += 2 {
			ouA[int(A[i]-'a')] += 1
			ouB[int(B[i]-'a')] += 1
		}
	}

	for k, v := range jiA {
		if jiB[k] != v {
			return false
		}
	}

	for k, v := range ouA {
		if ouB[k] != v {
			return false
		}
	}

	return true
}

func sum(A string) int {
	su := 0
	for _, v := range A {
		su += int(v - 'a')
	}
	return su
}

func numSpecialEquivGroups(A []string) int {
	ret := 0
	mapa := make(map[string]bool)
	for i := 0; i < len(A); i++ {
		first := true
		flag := false
		for j := i; j < len(A); j++ {
			if isValidSpecialStrings(A[i], A[j]) {
				_, ok := mapa[A[j]]
				if first != true {
					A[j] = ""
				}
				first = false
				if ok {
					continue
				}
				if flag == false {
					ret++
					flag = true
					mapa[A[i]] = true
				}
			}
		}
	}
	return ret
}

func main() {
	A := []string{"a", "b", "c", "a", "c", "c"}
	fmt.Println(numSpecialEquivGroups(A))
}
