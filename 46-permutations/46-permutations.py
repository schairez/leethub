# time: O(n!)
# space: O(n!)

class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        res = []
        def helper(idx: int):
            nonlocal res
            nonlocal n
            if idx == n:
                res.append(nums[:])
                return
            for i in range(idx, n):
                nums[idx], nums[i] = nums[i], nums[idx]
                helper(idx+1)
                nums[idx], nums[i] = nums[i], nums[idx]
            
        helper(0)
        return res
