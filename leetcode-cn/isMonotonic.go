/*896. 单调数列
用户通过次数 0
用户尝试次数 1
通过次数 0
提交次数 1
题目难度 Easy
如果数组是单调递增或单调递减的，那么它是单调的。

如果对于所有 i <= j，A[i] <= A[j]，那么数组 A 是单调递增的。 如果对于所有 i <= j，A[i]> = A[j]，那么数组 A 是单调递减的。

当给定的数组 A 是单调数组时返回 true，否则返回 false。



示例 1：

输入：[1,2,2,3]
输出：true
示例 2：

输入：[6,5,4,4]
输出：true
示例 3：

输入：[1,3,2]
输出：false
示例 4：

输入：[1,2,4,5]
输出：true
示例 5：

输入：[1,1,1]
输出：true
*/

package main

import "fmt"

func isMonotonic(A []int) bool {
	length := len(A)
	if length <= 1 {
		return true
	}

	prevv := A[0]
	iszeng := false
	isfirst := true

	flag := true
	for i := 1; i < length; i++ {
		if isfirst {
			if prevv < A[i] {
				iszeng = true
				prevv = A[i]
				isfirst = false
				continue
			} else if prevv > A[i] {
				prevv = A[i]
				isfirst = false
				continue
			} else {
				prevv = A[i]
				continue
			}
		}

		if iszeng {
			if prevv > A[i] {
				return false
			}
			prevv = A[i]
		} else {
			if prevv < A[i] {
				return false
			}
			prevv = A[i]
		}
	}
	return flag
}

func main() {
	A := []int{1, 2, 5, 3, 3}
	fmt.Println(isMonotonic(A))
}
