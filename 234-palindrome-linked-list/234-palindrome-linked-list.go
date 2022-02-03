/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//time: O(n)
//space: O(1)

func isPalindrome(head *ListNode) bool {
    if head.Next == nil {
        return true
    }
    midNode := getMidNode(head)
    //reverse 2nd half of list
    reversedHalf := reverseList(midNode.Next)
    for reversedHalf != nil {
        if reversedHalf.Val != head.Val {
            return false
        }
        reversedHalf = reversedHalf.Next
        head = head.Next
    }
    return true
}
func getMidNode(head *ListNode) *ListNode {
    //if even len, returns first mid node
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
} 

func reverseList(node *ListNode) *ListNode {
    if node == nil {
        return nil
    }
    var currHead, prev, tmp *ListNode
    currHead = node
    for currHead != nil {
        tmp = currHead.Next
        currHead.Next = prev
        prev = currHead
        currHead = tmp
    } 
    return prev
}

/*
     slow fst       fst       fst
     |    |         |         |
1 -> 2 -> 3 -> 4 -> 3 -> 2 -> 1 -> nil
|         |
slow      fst
fast      slow

reverse from slow.next onward
===  1 -> 2 -> 3 -> 4 -> nil
head is
====  1 -> 2 -> 3 -> ...
loop till slow != nil

fast,
slow
|         
1 -> 2 -> 2 -> 1 -> nil
     |    |
    slow fst
- reverse from slow.next onward 
head 1 -> 2
- slow = rev(slow.next)
- iter through head and slow ptrs and check vals

find the midPt via slow/fast ptr approach

*/