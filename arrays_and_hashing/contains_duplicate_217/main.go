// LeetCode 217.
// Contains Duplicate.
// https://leetcode.com/problems/contains-duplicate/description/

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

// Example 1:

// Input: nums = [1,2,3,1]
// Output: true
// Example 2:

// Input: nums = [1,2,3,4]
// Output: false
// Example 3:

// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true

// Constraints:

// 1 <= nums.length <= 105
// -109 <= nums[i] <= 109

package main

import "fmt"

func main() {

	// test data
	nums0 := []int{1, 2, 5, 3, 4, 5}    // true
	nums1 := []int{1, 2, 3, 4, 5}       // false
	nums2 := []int{1, 2, 5, 3, 4, 4, 5} // true
	nums3 := []int{1}                   // false

	fmt.Println("\n\n=======================================================")
	fmt.Println("\t First Solution. My Own Solution. My First One")
	fmt.Println("===========================================================\n\n")

	containsDuplicate1(nums0)
	containsDuplicate1(nums1)
	containsDuplicate1(nums2)
	containsDuplicate1(nums3)

	fmt.Println("\n\n=====================================================")
	fmt.Println("\t Second Solution. Hash Set Length")
	fmt.Println("=========================================================\n\n")

	containsDuplicate2(nums0)
	containsDuplicate2(nums1)
	containsDuplicate2(nums2)
	containsDuplicate2(nums3)

	fmt.Println("\n\n=========================================")
	fmt.Println("\t Third Solution. Brute Force")
	fmt.Println("=========================================\n\n")

	containsDuplicate3(nums0)
	containsDuplicate3(nums1)
	containsDuplicate3(nums2)
	containsDuplicate3(nums3)

	fmt.Println("\n\n=========================================")
	fmt.Println("\t Fourth Solution. Hash set")
	fmt.Println("=========================================\n\n")

	containsDuplicate4(nums0)
	containsDuplicate4(nums1)
	containsDuplicate4(nums2)
	containsDuplicate4(nums3)

}

func containsDuplicate1(nums []int) bool {
	x := make(map[int]int)

	if len(nums) == 1 {
		fmt.Println(false, " - from function containsDuplicate1, My Own Solution, First Solution")
		return false
	}

	for i := 0; i < len(nums); i++ {
		x[nums[i]] += 1
	}

	for _, val := range x {
		if val > 1 {
			fmt.Println(true, " - from function containsDuplicate1, My Own Solution, First Solution")
			return true
		}
	}
	fmt.Println(false, " - from function containsDuplicate1, My Own Solution, First Solution")
	return false
}

func containsDuplicate2(nums []int) bool {
	seen := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		seen[nums[i]] = true
	}
	fmt.Println(len(seen) != len(nums), " - from function containsDuplicate2, Hash Set Length, Second Solution")
	return len(seen) != len(nums)
}

func containsDuplicate3(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				fmt.Println(true, " - from function containsDuplicate3, Brute Force, Third Solution")
				return true
			}
		}
	}
	fmt.Println(false, " - from function containsDuplicate3, Brute Force, Third Solution")
	return false
}

func containsDuplicate4(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			fmt.Println(true, " - from function containsDuplicate4, Hash Set, Fourth Solution")
			return true
		}
		seen[num] = true
	}
	fmt.Println(false, " - from function containsDuplicate4, Hash Set, Fourth Solution")
	return false
}
