# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

"""

note: if there are two middle nodes return 2nd

time: O(N)
space: O(1)

"""


class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow = fast = head 
        while fast is not None and fast.next is not None:
            fast = fast.next.next
            slow = slow.next 
        
        return slow 
        
        