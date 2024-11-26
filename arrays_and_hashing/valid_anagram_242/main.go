// LeetCode 242.
// Valid Anagram.
// https://leetcode.com/problems/valid-anagram/description/
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
//
// Example 1:
//
// Input: s = "anagram", t = "nagaram"
//
// Output: true
//
// Example 2:
//
// Input: s = "rat", t = "car"
//
// Output: false
//
//
//
// Constraints:
//
// 1 <= s.length, t.length <= 5 * 104
// s and t consist of lowercase English letters.
//
// Follow up:
// What if the inputs contain Unicode characters?
// How would you adapt your solution to such a case?

package main

import "fmt"

func main() {

	// test cases

	// false
	s0 := "cat"
	t0 := "car"

	// true
	s1 := "listen"
	t1 := "silent"

	// true
	s2 := "newyorktimes"
	t2 := "monkeyswrite"

	// false
	s3 := "listen to your hurt"
	t3 := "silent hill"

	fmt.Println("\n\n==========================================================")
	fmt.Println("\t First Solution. My Own Solution. My First One")
	fmt.Println("===========================================================\n\n")
	isAnagram0(s0, t0)
	isAnagram0(s1, t1)
	isAnagram0(s2, t2)
	isAnagram0(s3, t3)

	fmt.Println("\n\n==========================================================")
	fmt.Println("\t Second Solution. My Second Solution. My Second One")
	fmt.Println("===========================================================\n\n")
	isAnagram1(s0, t0)
	isAnagram1(s1, t1)
	isAnagram1(s2, t2)
	isAnagram1(s3, t3)

}

// my own solution
func isAnagram0(s, t string) bool {
	if len(s) != len(t) {
		fmt.Printf("false: Solution0. Not an Anagram. The length of strings [%#s] and [%#s] is not the same.\n\n", s, t)
		return false
	}

	ms := make(map[string]int)
	mt := make(map[string]int)

	for _, char := range s {
		ms[string(char)] += 1
	}

	for _, char := range t {
		mt[string(char)] += 1
	}

	for key, _ := range ms {
		if ms[key] != mt[key] {
			fmt.Println("false: Solution0. Not an Anagram.\n")
			return false
		}
	}
	fmt.Printf("true: Solution0. \nms is: %#v\nmt is: %#v\n\n", ms, mt)
	return true
}

// my own solution about a year ago (November 2023)
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		fmt.Printf("false: Solution1. Not an Anagram. The length of strings [%#s] and [%#s] is not the same.\n\n", s, t)
		return false
	}

	countS := map[rune]int{}
	countT := map[rune]int{}

	for _, c := range s {
		_, ok := countS[c]
		if ok == false {
			countS[c] = 0
		}
		countS[c] += 1

	}

	for _, c := range t {
		_, ok := countT[c]
		if ok == false {
			countT[c] = 0
		}
		countT[c] += 1

	}

	for key, _ := range countS {
		_, ok := countT[key]

		if ok == false || countT[key] != countS[key] {
			fmt.Println("false: Solution1. Not an Anagram.\n")
			return false
		}
	}
	fmt.Printf("true: Solution1. \ncountS is: %#v\ncountT is: %#v\n\n", countS, countT)
	return true

}
