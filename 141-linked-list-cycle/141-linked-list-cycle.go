/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time: O(n) 
// space: O(1)

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true 
        }
    }
    return false
}