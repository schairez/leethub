
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func walk(t *TreeNode, ch chan int, isLeft bool) {
    if t == nil {
        if isLeft {
            ch <- -1
        } else {
            ch <- 1
        }
        return
    }
    walk(t.Left, ch, isLeft)
    ch <- t.Val
    walk(t.Right, ch, !isLeft)
}
func Walk(t *TreeNode, ch chan int) {
	walk(t, ch, true)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go Walk(p, ch1)
	go Walk(q, ch2)
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


/*
func isSameTree(p *TreeNode, q *TreeNode) bool {
   	ch1, ch2 := make(chan int), make(chan int)
	go Walk(p, ch1)
	go Walk(q, ch2)
	v1, v2, ok1, ok2 := 0, 0, true, true
	for ok1 && ok2 {
		v1, ok1 = <-ch1
		v2, ok2 = <-ch2
		if ok1 != ok2 || v1 != v2 {
			return false
		}

	}
	return true
}

func Walk(t *TreeNode, ch chan int) {
	defer close(ch)
	var walk func(*TreeNode)
	walk = func(t *TreeNode) {
		if t == nil {
            return
		}
        if t.Left == nil {
           ch <- -1
        }
        if t.Right == nil {
            ch <- 1
        }
        walk(t.Left)
		ch <- t.Val
		walk(t.Right)
	}
	walk(t)
}
*/

/*

func dfs(root *TreeNode, ch chan int) {
    defer close(ch)
    var inorder func(*TreeNode, chan int)
    inorder = func(node *TreeNode, ch chan int) {
        if node == nil {
            ch <- -1
            return
        }
        inorder(node.Left, ch)
        ch <- node.Val
        inorder(node, ch)
    }
    inorder(root, ch)
}


// parallel version

func isSameTree(p *TreeNode, q *TreeNode) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go dfs(p, ch1)
    go dfs(q, ch2)
    
    for {
        val1, ok1 := <-ch1
        val2, ok2 := <-ch1
        
        if val1 != val2 || ok1 != ok2 {
            return false
        }
    }
    return true
}
*/



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


