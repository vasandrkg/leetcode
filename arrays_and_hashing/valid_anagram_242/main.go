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
	isAnagram(s0, t0)
	isAnagram1(s0, t0)
	// true
	s1 := "listen"
	t1 := "silent"
	isAnagram(s1, t1)
	isAnagram1(s1, t1)

	// true
	s2 := "newyorktimes"
	t2 := "monkeyswrite"
	isAnagram(s2, t2)
	isAnagram1(s2, t2)
	// false
	s3 := "listen to your hurt"
	t3 := "silent hill"
	isAnagram(s3, t3)
	isAnagram1(s3, t3)

	//// more compact view of test cases
	//ss := []string{"cat", "listen", "newyorktimes", "listen toy your heart"}
	//tt := []string{"car", "silent", "monkeyswrite",  "silent hill"}
	//

}

// my own solution
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		fmt.Printf("false: Not an Anagram. The length of strings [%#s] and [%#s] is not the same.\n\n", s, t)
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
			fmt.Println("false: Not an Anagram.\n")
			return false
		}
	}
	fmt.Printf("true: \nms is: %#v\nmt is: %#v\n\n", ms, mt)
	return true
}

// my own solution about a year ago (November 2023)
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		fmt.Printf("false: Not an Anagram. The length of strings [%#s] and [%#s] is not the same.\n\n", s, t)
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
			fmt.Println("false: Not an Anagram.\n")
			return false
		}
	}
	fmt.Printf("true: \ncountS is: %#v\ncountT is: %#v\n\n", countS, countT)
	return true

}
