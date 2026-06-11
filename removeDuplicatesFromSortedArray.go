// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

// this is a two pointer approach for this problem
// we start at the 2nd element (index 1) since we know the array is sorted
// and we need to compare the previous element with the current
// if the adjacent elements arent same we know that the element on right pointer is unique
// so we set the left pointer element to the right pointer element and increment left pointer
// since left pointer keeps the index of the last unique element we return it as the answer
func removeDuplicates(nums []int) int {
    l := 1
    for r := 1; r < len(nums); r++ {
        if nums[r] != nums[r - 1] {
            nums[l] = nums[r]
            l += 1
        }
    }
    return l
}
