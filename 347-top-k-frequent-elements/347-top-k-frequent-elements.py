class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        n = len(nums)
        if n == 1:
            return nums
        
        freq_tbl = {}
        for num in nums:
            freq_tbl[num] = freq_tbl.get(num, 0) + 1
        
        inv_freq = [[] for _ in range(n+1)]
        max_key = float("-inf")
        for num_key, freq_val in freq_tbl.items():
            inv_freq[freq_val].append(num_key)
            max_key = max(max_key, freq_val)
       
        print(freq_tbl)
        print(inv_freq)
        res = []
        cap = k
        curr_key = max_key
        while cap != 0:
            n_key = len(inv_freq[curr_key]) 
            if n_key > cap:
                res.extend(inv_freq[curr_key][:cap])
                break
            elif n_key != 0:
                res.extend(inv_freq[curr_key])
                cap -= n_key
            curr_key -= 1 
            
        return res