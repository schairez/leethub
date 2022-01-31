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

//time: O(n)
//space: O(n)
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    var recurseRev func(prev, curr *ListNode) *ListNode
    recurseRev = func(prev, curr *ListNode) *ListNode {
        if curr == nil {
            return prev
        }
        tmp := curr.Next
        curr.Next = prev
        return recurseRev(curr, tmp)
    }
    
    return recurseRev(nil, head)
}
    
//time: O(n)
//space: O(1)
func reverseListIter(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var prev, curr *ListNode
    curr = head
    for curr != nil {
        curr.Next, prev, curr  = prev, curr, curr.Next  
    }
    return prev
}
    
    



/*
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

*/


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
