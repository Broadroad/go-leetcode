/*
1. Two Sum
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].


Solution:
One Pass way hashmap

It turns out we can do it in one-pass. While we iterate and inserting elements into the table, we also look back to check if current element's complement already exists in the table. If it exists, we have found a solution and return immediately.

Complexity Analysis:

Time complexity : O(n)O(n). We traverse the list containing nn elements only once. Each look up in the table costs only O(1)O(1) time.

Space complexity : O(n)O(n). The extra space required depends on the number of items stored in the hash table, which stores at most nn elements.
*/

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for k, num := range nums {
		t := target - nums[k]

		v, ok := hashMap[t]
		if ok {

			return []int{k, v}

		}
		hashMap[num] = k

	}
	return nil
}
