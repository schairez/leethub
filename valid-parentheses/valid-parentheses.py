from typing import List 

"""
- use a stack
- add open brace/bracket opts to stack
- when encounter a closing brace/bracket pop() from stack
- verify

ex:
- s = '{[]}'
add '{' to stack
add '[' to stack 


"""


class Solution:
    def isValid(self, s: str) -> bool:
        if len(s) == 1: 
            return False 
        valid_keys = {
                      "}": "{", 
                      "]": "[",
                      ")": "("
                     }
        stack: List[str] = list() 
        for ch in s:
            if ch in valid_keys:
                if len(stack) > 0:
                    valid_open_brace = valid_keys[ch]
                    if not valid_open_brace == stack.pop():
                        return False 
                else:
                    return False 
            elif not ch in valid_keys: #append open brace types
                stack.append(ch)
            
        if len(stack) > 0: 
            return False 
        return True  