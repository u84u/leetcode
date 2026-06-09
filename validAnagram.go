// https://leetcode.com/problems/valid-anagram/


// this is the approach for this problem
// basically we first check if the string lenghts are equal or not
// if not we return early
// if yes then we make a hashmap for both strings such that
// hashmap[char] = occurences
// if the strings are anagrams the hashmap key from string S
// should match the value of hashmap for string T
// time complexity is O(S + T)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countS := make(map[byte]int)
	countT := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		cs := s[i]
		ct := t[i]
		countS[cs]++
		countT[ct]++
	}
	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
}
