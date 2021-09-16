/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ex: 1 -> 2 -> nil 
//     h 

func swapPairs(head *ListNode) *ListNode {  
    // len == 1
    if head == nil || head.Next == nil {
        return head
    }
    headPtr := head.Next //for return
    n1 := head 
    currTail := n1
    n2 := head.Next 
    n3 := n2.Next 
   // var currTail *ListNode 
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
    //var l1Head, l1Tail, l2Head, l2Tail *ListNode 
    //var prevTail *ListNode
    //var localHead *ListNode
    //var localTail *ListNode
    //localHead = n2 
    // len == 2 
    //if n3 == nil {
    /*
    for {
        localHead = n2 
        localHead.Next = n1 
        n1.Next = n3
        if n3 == nil {
            break
        }
        n2 = n3.Next 
        if n2 == nil {
            break
        }
        n1.Next = n2
        n2.Next = n3
        n1 = n3 
        if n1 == nil {
            break
        } 
        n2 = n1.Next 
        if n2 == nil {
            break
        }
        localHead = n2 
        n3 = n2.Next 
        //if n3 == nil {
         //   break 
        }
        */
    //return headPtr
    
    //len >= 3 
    //for head.Next != nil && head.Next.Next != nil {
        
          
    
    return headPtr
    
    // we swap every
    
    
}

    /*
    if head == nil {
        return nil 
    }
    var nextSwap *ListNode
    
    // swap first two nodes to mark head 
    if head.Next != nil {
        moveMe := head //moveMe
        currNext := head.Next //after
        afterNext := currNext.Next
        head = currNext 
        head.Next = moveMe
        moveMe.Next = afterNext 
        nextSwap = afterNext
    }
    for nextSwap != nil || nextSwap.Next != nil {
        moveMe := nextSwap //moveMe
        currNext := moveMe.Next //after
        afterNext := currNext.Next
        head = currNext 
        head.Next = moveMe
        moveMe.Next = afterNext 
        nextSwap = afterNext
    } 
    
    return head
    */

/*
    
    
    if head == nil {//|| head.next == nil {
        return nil
    }
    nextNode := head.Next
    if nextNode == nil {
        return head
    }
    after := nextNode.Next
    if nextNode != head {
        prev := head 
        //after := nextNode.Next
        //next := head.Next
        head = nextNode 
        head.Next = prev
        prev.Next = after 
        
    }
    nextNode = swapPairs(after.Next)
    return head
    
    
    //for node.Next != nil {
    //}
}
*/
/*

node.Next = f()


nil
|
2
|
1


*/

// ex: 1 -> 2 -> nil 
// @f(2) -> ret 2 
// 
//f(1)
//  |
//  --> f(2)
// h -> 1 
// next -> 2 
// prev -> 1             h; prev    next
// next = h.next               |    |
// h = next                    1 -> 2 -> nil 

//  h = next                        h    next
// next = h.next          |         |
// h = next                    1 -> 2 -> nil 
// 
/* 
func recurse(head *ListNode) *ListNode {
    if head == nil {//|| head.next == nil {
        return nil
    }
    if head.Next == nil {
        return head
    }
    nextNode := recurse(head.Next)
    if nextNode != head {
        prev := head 
        after := nextNode.Next
        //next := head.Next
        head = nextNode 
        head.Next = prev
        prev.Next = after 
        
    }
    return head
    
}
*/