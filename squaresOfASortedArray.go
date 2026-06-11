// https://leetcode.com/problems/squares-of-a-sorted-array 

// the simple solution
// we create a new array and keep appending the squares of the array while looping
// after that we sort the `res` array
// since sorting has added time complexity
// time complexity is O(n logn)
import "sort"
func sortedSquares(nums []int) []int {
    res := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        res = append(res, nums[i] * nums[i])
    }
    sort.Ints(res)
    return res
}


// this is the optimal solution
// we create the resulting array first with the correct length from the input
// then we have two pointers left and right
// then while left and right pointers do not cross each other
// we check if the left square is larger than right square
// if so we add the left square from the back to the `res` array
// else we add the right square
// we are essentially building the return array from the back
// time complexity is O(n)
func sortedSquares(nums []int) []int {
    res := make([]int, len(nums))
    l := 0
    r := len(nums) - 1
    pos := len(nums) - 1
    for l <= r {
        lsqur := nums[l] * nums[l]
        rsqur := nums[r] * nums[r]
        if lsqur > rsqur { 
            res[pos] = lsqur
            l += 1
        } else {
            res[pos] = rsqur
            r -= 1
        }
        pos -= 1
    }
    return res
}
