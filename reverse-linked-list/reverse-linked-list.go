//recursive implementation
// time: O(n) 
// space: O(n) stack trace
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil 
    }
    //curr = head; prev = nil
    return rev(head, nil)
}
 
func rev( currTail, prev *ListNode ) *ListNode {
    nextIter := currTail.Next
    currTail.Next = prev
    //recurse until we find tail nil ptr 
    if nextIter == nil {
        return currTail
    }
    head := rev(nextIter, currTail)
    return head
}




/*
 _rev(1, nil)   nil
     |          |
--rev(2,1)      1
     |          |   ^
   rev(3,2)     2
     |          |   ^
   rev(4,3)     3
     |          |
   rev(5,4)     |   ^
      |         4
   rev(nil, 5)  |   ^
                5
                
*/
