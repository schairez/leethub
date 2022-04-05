/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 23. Merge k Sorted Lists
// time: O(kN); k == numLists; N = totalNodes 
// space: O(1)

func mergeKLists(lists []*ListNode) *ListNode {
    n := len(lists)
    if n == 0 {
        return nil
    }
    l1 := lists[0]
    if n == 1 {
        return l1
    }
    for i := 1; i < n; i++ {
        l1 = merge2Lists(l1, lists[i])
        lists[i] = nil
    } 
    return l1
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }
    dummyNode := &ListNode{}
    currNode := dummyNode
    for l1 != nil || l2 != nil {
        if l1 != nil && l2 != nil {
            if l1.Val <= l2.Val {
                currNode.Next = l1
                l1 = l1.Next
            } else {
                currNode.Next = l2
                l2 = l2.Next
            }
        } else if l1 != nil {
            currNode.Next = l1
            l1 = l1.Next
        } else if l2 != nil {
            currNode.Next = l2
            l2 = l2.Next
        }
        currNode = currNode.Next
    }
    
    return dummyNode.Next
}