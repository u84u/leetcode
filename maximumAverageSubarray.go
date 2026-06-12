// https://leetcode.com/problems/maximum-average-subarray-i/ 

// this is a fixed window algorithm
// we first compute the sum of the first `k` elements
// then we loop from starting on `k` to `n`
// we add the element at `i` and subtract `i-k` to remove the first element
// this ensures that at any given point we are only holding the sum
// of `k` elements, then we check if this is the max sum we could find
// if yes we update the value and continue if not we still continue
// time complexity is O(n)
func findMaxAverage(nums []int, k int) float64 {
    n := len(nums)
    sum := 0
    for i := 0; i < k; i++ { sum += nums[i] }
    maxSum := sum
    for i := k; i < n; i++ {
        sum += nums[i] - nums[i-k]
        maxSum = max(sum, maxSum)
    }
    return float64(maxSum) / float64(k)
}
