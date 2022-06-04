/*
Given a string s, sort it in decreasing order based on the frequency of the
characters. The frequency of a character is the number of times it appears in the
string.

Return the sorted string. If there are multiple answers, return any of them.

Example 1:

Input: s = "book"
Output: "oobk" or "ookb"

Example 2:

Input: s = "zwzwzw"
Output: "wwwzzz" or "zzzwww"


Example 3:
Input: s = "mississippi"
Output: "iiiissssppm" or "ssssiiiippm"
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func strSlice(st string) string {
	m := map[string]int{}

	for _, s := range st {
		m[string(s)]++
	}

	// To use the count as the key
	n := map[int][]string{}
	// To have a slice that counts each occurences
	var a []int

	// range through m and create data for map n where the count is the key
	for k, v := range m {
		n[v] = append(n[v], k)
	}

	// Now loop through the n map to store the key which are integers in slice a
	for k := range n {
		a = append(a, k)
	}

	b := sort.Reverse(sort.IntSlice(a))
	// Sort accepts a parameter that implements the sort interface
	// and IntSlice and Reverse implements this same interface
	sort.Sort(b)

	var sortedSliceString []string
	var stringDescendingCharacterFrequency string

	for _, k := range a {

		for _, s := range n[k] {
			fmt.Printf("%s, %d\n", s, k)
			// mini for loop to loop thru k for each string/character instance
			for i := 0; i < k; i++ {
				sortedSliceString = append(sortedSliceString, s)
			}
		}
	}

	stringDescendingCharacterFrequency = strings.Join(sortedSliceString, "")

	return stringDescendingCharacterFrequency

}

func main() {

	fmt.Println(strSlice("book"))
	fmt.Println(strSlice("zwzwzw"))
	fmt.Println(strSlice("mississippi"))
}
