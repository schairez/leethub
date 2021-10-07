/*
time: O(n)
space: O(n)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return head
    }
    //part1Dummy is vals < x; part 2 is vals >= x
    part1Dummy, part2Dummy := &ListNode{}, &ListNode{}
    ptr1, ptr2 := part1Dummy, part2Dummy
    for head != nil {
        if head.Val < x {
            ptr1.Next = head
            ptr1 = ptr1.Next
        } else {
            ptr2.Next = head
            ptr2 = ptr2.Next
        }
        head = head.Next
    }
    ptr2.Next = nil
    ptr1.Next = part2Dummy.Next
    return part1Dummy.Next
    
}