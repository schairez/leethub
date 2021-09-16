"""
time: O(len(linked-list)) ~ O(n)
space: O(n)
"""

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

# ex input: 1 -> 2 -> 3 -> nil 
class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None:
            return head
        tmp = self.swapPairs(head.next.next) 
        newHead = head.next 
        newHead.next = head
        head.next = tmp
        return newHead
