// 238. Product of Array Except Self
// time: O(n)
// space: O(n)

class Solution {
    public int[] productExceptSelf(int[] nums) {
        int n  = nums.length;
        int[] res = new int[n];
        int curr = 1;
        int prev = 1;
        for (int idx = 0; idx < n; idx++) {
            prev = curr;
            res[idx] = prev;
            curr = curr * nums[idx];
        }
        curr = nums[n-1]; // 2 <= nums.length <= 10^5
        for (int idx = n-2; idx >= 0; idx--) {
            res[idx] = res[idx] * curr;
            curr = curr * nums[idx];
        }
        return res;
    }
    
}