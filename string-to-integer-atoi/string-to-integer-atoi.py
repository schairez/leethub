"""
time: O(n)
space: O(1)
"""
from typing import Tuple 


class Solution:
    def myAtoi(self, s: str) -> int:
        
        if len(s) == 0:
            return 0 
        # checks and clamps to closest i32 val
        def check_clamp_i32(num: int) -> int:
            MIN_I32 = -(1<<31)      # -2 **31
            MAX_I32 = (1<<31) - 1   # (2 **31) - 1
            print("input n", num)
            if num > MAX_I32:
                return MAX_I32
            elif num < MIN_I32:
                return MIN_I32
            else:
                return num 
            
        def check_code_pt_numeric(ch: str) -> Tuple[int, bool]:
            code = ord(ch) - ord('0')
            return (code, True) if 0 <= code <= 9 else (code, False)  
        
        res = 0 
        start = 0  
        # leading whitespace 
        while s[start] == " " and start != len(s)-1:
            start += 1 
            
        is_negative = False 
        if s[start] == "-" and start != len(s)-1:
            _, ok = check_code_pt_numeric(s[start+1])
            if not ok:
                return res 
            is_negative = True 
            start +=1 
        elif s[start] == "+":
            start +=1
        print('is_n', is_negative) 
        
        for i in range(start, len(s)):
            code, ok = check_code_pt_numeric(s[i])
            if not ok:
                if res != 0:
                    res = res*-1 if is_negative else res 
                    return check_clamp_i32(res)
                return 0 
            
            res = res*10 + code 
            
        res = res*-1 if is_negative else res 
        
        return check_clamp_i32(res)      
       
        