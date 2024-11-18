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
	nums0 := []int{1, 2, 5, 3, 4, 5}
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{1, 2, 5, 3, 4, 4, 5}
	nums3 := []int{1}

	fmt.Println("\n\n======================================")
	fmt.Println("\t First Solution")
	fmt.Println("==========================================\n\n")

	containsDuplicate(nums0)	  
	containsDuplicate(nums1)
  	containsDuplicate(nums2)
  	containsDuplicate(nums3)

	fmt.Println("\n\n=====================================")
	fmt.Println("\t Second Solution")	
	fmt.Println("=========================================\n\n")
	
	containsDuplicate2(nums0)
	containsDuplicate2(nums1)
	containsDuplicate2(nums2)
	containsDuplicate2(nums3)

}

// first solution
func containsDuplicate(nums []int) bool {
	x := make(map[int]int)
	var b bool

	if len(nums) == 1 {
		b = false
		fmt.Println(b, "from function containsDuplicate")
		return b
	}

	for i := 0; i < len(nums); i++ {
		x[nums[i]] += 1
	}
	for _, val := range x {
		if val > 1 {
			b = true
			fmt.Println(b, "from function containsDuplicate")
			return b
		}
	}
	b = false
	fmt.Println(b, "from function containsDuplicate")
	return b
}


//second solution
func containsDuplicate2(nums []int) bool {
	s := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		s[nums[i]] = true
	}
	b2 := (len(s) != len(nums))
	fmt.Println(b2, "from function containsDuplicate2")
	return b2
}
