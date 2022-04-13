# 47. Permutations II
# time: O(n*n!)
# space: O(n*n!)

from collections import defaultdict

class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        res = []
        n = len(nums)
        freqMap = defaultdict(int)
        for num in nums:
            freqMap[num] += 1
        def helper(perm: List[int]):
            nonlocal res
            nonlocal freqMap
            if len(perm) == n:
                res.append(perm[:])
                return
            for num_key, cnt_val in freqMap.items():
                if cnt_val == 0:
                    continue
                perm.append(num_key)
                freqMap[num_key] -= 1
                helper(perm)
                freqMap[num_key] += 1
                perm.pop()
        
        helper([])
        return res