// https://leetcode.com/problems/valid-palindrome/

// this is the bruteforce way to solve this problem
// we first lower the string
// then we build a new string containing only alphanumeric characters [A-Z,a-z,0-9]
// since go doesnt have a built-in for reversing like python's [::-1] we build the reversed string manually
// time complexity is O(n)
import "strings"
// https://medium.com/@saharat.paynok/how-to-check-if-the-character-is-alphanumeric-in-go-6783b92ec412
func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
func isPalindrome(s string) bool {
	lowered := strings.ToLower(s)
	newS := ""
	for _, c := range lowered {
		if isAlphaNumeric(byte(c)) {
			newS += string(c)
		}
	}
	reversed := ""
	for i := len(newS) - 1; i >= 0; i-- {
		reversed += string(newS[i])
	}
	return newS == reversed

}


// this is the optimized and the fast way for this problem
// instead of reversing the string or anything
// we simply have two pointers working inwards, left to right, right to left
// we skip any non alphanumeric characters [A-Z,a-z,0-9]
// we also implement a ToLower function because strings.ToLower works on the entire string at once, added memory and time complexity
// time complexity is still O(n) but memory complexity is O(1)

// https://medium.com/@saharat.paynok/how-to-check-if-the-character-is-alphanumeric-in-go-6783b92ec412
func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
func ToLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
func isPalindrome(s string) bool {
    l := 0
    r := len(s) - 1
    for l < r {
        for l < r && !isAlphaNumeric(s[l]){
            l += 1
        }
        for r > l && !isAlphaNumeric(s[r]){
            r -= 1
        }
        if ToLower(s[l]) != ToLower(s[r]) { return false }
        l += 1
        r -= 1
    }    
    return true
}
