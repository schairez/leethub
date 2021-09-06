"""
space: O(1) 
time: O(N)
"""


class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        prev_last = 0 
        len_last = 0 
        for ch in s:
            if ch == " ":
                if len_last != 0:
                    prev_last = len_last 
                len_last = 0 
            else:
                len_last += 1  
                
            
        if s[len(s)-1] == " ":
            return prev_last 
        return len_last 