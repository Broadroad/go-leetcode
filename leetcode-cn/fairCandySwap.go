/*
888. 公平的糖果交换
用户通过次数 0
用户尝试次数 2
通过次数 0
提交次数 3
题目难度 Easy
爱丽丝和鲍勃有不同大小的糖果棒：A[i] 是爱丽丝拥有的第 i 块糖的大小，B[j] 是鲍勃拥有的第 j 块糖的大小。

因为他们是朋友，所以他们想交换一个糖果棒，这样交换后，他们都有相同的糖果总量。（一个人拥有的糖果总量是他们拥有的糖果棒大小的总和。）

返回一个整数数组 ans，其中 ans[0] 是爱丽丝必须交换的糖果棒的大小，ans[1] 是 Bob 必须交换的糖果棒的大小。

如果有多个答案，你可以返回其中任何一个。保证答案存在。


示例 1：

输入：A = [1,1], B = [2,2]
输出：[1,2]
示例 2：

输入：A = [1,2], B = [2,3]
输出：[1,2]
示例 3：

输入：A = [2], B = [1,3]
输出：[2,3]
示例 4：

输入：A = [1,2,5], B = [2,4]
输出：[5,4]
*/

package main

import "fmt"

func calSum(A []int) int {
	sum := 0
	for _, v := range A {
		sum += v
	}
	return sum
}

func fairCandySwap(A []int, B []int) []int {

	a := make([]int, 2)
	sumA := calSum(A)
	sumB := calSum(B)
	fmt.Println(sumA, sumB)
	var distance int
	if sumA > sumB {
		distance = (sumA - sumB) / 2
	} else {
		distance = (sumB - sumA) / 2
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {

			if sumA < sumB && A[i] == B[j]-distance {
				a[0] = A[i]
				a[1] = B[j]
			}

			if sumA > sumB && A[i] == B[j]+distance {
				a[0] = A[i]
				a[1] = B[j]
			}
		}
	}
	return a
}

func main() {
	A := []int{1, 2}
	B := []int{2, 3}
	fmt.Println(fairCandySwap(A, B))
}
