"""
time is len(s) + len(t)
time: O(N + M)
space: O(1); since size 26 arr stays constant 

"""
class Solution:
    def minSteps(self, s: str, t: str) -> int:
        res = 0 
        char_cnt = [0 for _ in range(26)]
        for ch in s:
            pt = ord(ch) - ord('a') 
            char_cnt[pt] +=1 
        for ch in t:
            pt = ord(ch) - ord('a') 
            char_cnt[pt] -=1 
            if char_cnt[pt] < 0:
                res +=1 
                
        return res  