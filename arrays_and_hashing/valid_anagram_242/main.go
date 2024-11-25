// LeetCode 242.
// Valid Anagram.

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

	// true
	s1 := "listen"
	t1 := "silent"
	isAnagram(s1, t1)

	// true
	s2 := "newyorktimes"
	t2 := "monkeyswrite"
	isAnagram(s2, t2)

	//false
	s3 := "listen to your hurt"
	t3 := "silent hill"
	isAnagram(s3, t3)

}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		fmt.Println(len(s), len(t))
		fmt.Printf("false - Not an Anagram. The length of strings [%#s] and [%#s] is not the same.\n\n", s, t)
		return false
	}
	ms := make(map[string]int)
	for _, vs := range s {
		ms[string(vs)] += 1
	}
	fmt.Printf("ms is: %#v\n", ms)
	mt := make(map[string]int)
	for _, vt := range t {
		mt[string(vt)] += 1
	}
	fmt.Printf("mt is: %#v\n", mt)

	//for keyt, _ := range mt {
	//	for keys, _ := range ms {
	//		if keyt != keys {
	//			fmt.Printf("\nfalse - Not an Anagram. There is no [%s] and [%s] is not the same", keyt, keys)
	//			return false
	//		}
	//
	//	}
	//}

	return true
}
