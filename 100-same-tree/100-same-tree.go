
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// parallel version
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func dfs(root *TreeNode, ch chan int) {
    var walk func(*TreeNode, bool)
    walk = func(t *TreeNode,isLeft bool) {
        if t == nil {
            if isLeft {
                ch <- -1
            } else {
                ch <- 1
            }
            return
        }
        walk(t.Left, isLeft)
        ch <- t.Val
        walk(t.Right, !isLeft)
    }
    
	walk(root,true)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go dfs(p, ch1)
	go dfs(q, ch2)
    for {
		select {
		case n1, ok1 := <-ch1:
            n2, ok2 := <- ch2 
            if ok1 == ok2 {
                if ok1 == false && ok2 == false {
                    return true
                }
                if n1 != n2 {
                    return false
                }
            } else {
                return false
            }
        }
    }
}



//time: O(N)
// space: O(N)
/*
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil { 
        return true 
    } else if p == nil && q != nil || q == nil && p != nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
*/


