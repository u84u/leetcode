// https://leetcode.com/problems/move-zeroes

// this is a two pointer approach for this problem
// we start both the left and right pointers at the start of the array
// and for each non zero element on the right pointer we keep swapping the left and right values
// we do that for the entire array and move all zeroes to the end in one pass
// time complexity is O(n)
func moveZeroes(nums []int)  {
    l := 0
    for r := 0; r < len(nums); r++ {
        if nums[r] != 0 {
            nums[l], nums[r] = nums[r], nums[l]
            l += 1
        }
    }
    return
}
