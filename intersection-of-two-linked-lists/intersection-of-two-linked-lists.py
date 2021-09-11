"""
time: O(m+n); m = len(headA); n = len(headB)
space: O(1)
"""

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


"""

          3 -> nil
          ^
     2 -> | 

"""
class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        if headA is None:
            return None 
        
        intersect = False 
        lenA, lenB = 0, 0 
        pA, pB = headA, headB 
        while pA is not None:
            lenA += 1 
            pA = pA.next
        while pB is not None:
            lenB += 1 
            pB = pB.next 
            
        pA, pB = headA, headB 
        if lenA > lenB:
            diff = lenA - lenB
            while diff != 0:
                pA = pA.next 
                diff -= 1 
        elif lenB > lenA:
            diff = lenB - lenA
            while diff != 0:
                pB = pB.next 
                diff -= 1
        while pA is not None and pB is not None:
            if pA is pB:
                intersect = True 
                break 
            pA = pA.next 
            pB = pB.next 
            
        if intersect is True:
            return pA 
        
        return None
