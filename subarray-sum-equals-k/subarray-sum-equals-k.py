"""
time: O(n)
space: O(n) for auxiliary dict
"""
from typing import Dict

class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        # using a map to store the sum as keys and the number of occurrences of that sum as v
        d: Dict[int,int] = {}
        d[0] = 1 # there is a value that sums to 0; e.g. the absence of a val
        cum_sum = 0
        total = 0
        for v in nums:
            cum_sum += v 
            complement = cum_sum - k
            if complement in d:
                total += d[complement]
            d[cum_sum] = d.get(cum_sum, 0) + 1
            
        return total
        