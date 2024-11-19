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
	nums3 := []int{1}                   //false

	fmt.Println("\n\n==========================================")
	fmt.Println("\t First Solution")
	fmt.Println("==========================================\n\n")

	containsDuplicate1(nums0)
	containsDuplicate1(nums1)
	containsDuplicate1(nums2)
	containsDuplicate1(nums3)

	fmt.Println("\n\n=========================================")
	fmt.Println("\t Second Solution")
	fmt.Println("=========================================\n\n")

	containsDuplicate2(nums0)
	containsDuplicate2(nums1)
	containsDuplicate2(nums2)
	containsDuplicate2(nums3)

	fmt.Println("\n\n=========================================")
	fmt.Println("\t Third Solution")
	fmt.Println("=========================================\n\n")

	//containsDuplicate3(nums0)
	//containsDuplicate3(nums1)
	//containsDuplicate3(nums2)
	//containsDuplicate3(nums3)

}

// first solution
func containsDuplicate1(nums []int) bool {
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

// second solution (hash set length)
func containsDuplicate2(nums []int) bool {
	seen := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		seen[nums[i]] = true
	}
	b2 := (len(seen) != len(nums))
	fmt.Println(b2, "from function containsDuplicate2")
	return b2
}

// third solution

//func containsDuplicate3(nums []int) bool {
//
//}
