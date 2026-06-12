// https://leetcode.com/problems/longest-substring-without-repeating-characters

// yet another sliding window approach
// we create a hashmap of our window, cuz golang dont got sets in it (mid language)
// then starting at the left pointer, we loop the entire array
// then for each element, we check if the element is in the hashmap
// if yes we remove the left ptr and increment its position else we break knowing that the character is unique
// then we add the right ptr character to the hashmap (set)
// and update `res` to the max of `res` and difference between left and right ptr indexes (1-based)
// note that duplicates will work as long as they are the same length
// as we only care about the actual length of the subarray not its contents
// time complexity is O(n)
func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]struct{})
	l := 0
	res := 0
	for r := 0; r < len(s); r++ {
		for {
			_, ok := set[s[r]]
			if ok {
				delete(set, s[l])
				l += 1
			} else {
				break
			}
		}
		set[s[r]] = struct{}{}
		res = max(res, r-l+1)
	}
	return res
}
