/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {  
    // len == 1
    if head == nil || head.Next == nil {
        return head
    }
    headPtr := head.Next //for return
    n1 := head 
    //currTail := n1
    n2 := head.Next 
    n3 := n2.Next 
    var currTail *ListNode 
    for {
        n2.Next = n1
        n1.Next = n3
        currTail = n1 
        if n3 == nil || n3.Next == nil {
            break
        }
        n1 = n3 
        n2 = n3.Next
        n3 = n2.Next
        currTail.Next = n2
        
    }
    return headPtr
    
}

