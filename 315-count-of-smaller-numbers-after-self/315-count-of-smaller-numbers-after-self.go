// 315. Count of Smaller Numbers After Self
// 1 <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4

const (
    size = 10_000 * 2 + 1 // max arr input size
)

func offset(num int) int {
    return num + 10_000
}

// dynamic segTree approach ?

type SegTree struct {
    Node [4*size]int
}

func countSmaller(nums []int) []int {
    n := len(nums)
    if n == 1 {
        return []int{0}
    }
    res := make([]int, n)
    segTree := &SegTree{}
    for i := n-1; i >= 0; i-- {
        index := offset(nums[i])
        cntLessRight := segTree.QueryRange(0, 0, size-1, 0, index-1)        
        res[i] = cntLessRight
        segTree.Update(0, 0, size-1, index)
    }
    return res
}


func (segTree *SegTree) QueryRange(treeIndex, lo, hi, ql, qr int) int {
    if qr < lo || ql > hi {
        return 0
    }
    
    if ql == lo && qr == hi {
        return segTree.Node[treeIndex]
    }
    
    mid := lo + (hi-lo)>>1
    
    if qr <= mid {
        return segTree.QueryRange(2*treeIndex+1, lo, mid, ql, qr)
    } else if ql > mid {
        return segTree.QueryRange(2*treeIndex+2, mid+1, hi, ql, qr)
    }
    leftQuery := segTree.QueryRange(2*treeIndex+1, lo, mid, ql, mid)
    rightQuery := segTree.QueryRange(2*treeIndex+2, mid+1, hi, mid+1, qr)
    
    return leftQuery + rightQuery
}



func (segTree *SegTree) Update(treeIndex, lo, hi, arrIndex int) {
    if lo == hi {
        segTree.Node[treeIndex] += 1
        return
    }
    mid := lo + (hi-lo) >> 1
    if mid >= arrIndex {
        segTree.Update(2*treeIndex+1, lo, mid, arrIndex)
    } else if arrIndex > mid {
        segTree.Update(2*treeIndex+2, mid+1, hi, arrIndex)
    }
    
    segTree.Node[treeIndex] = segTree.Node[2*treeIndex+1] + segTree.Node[2*treeIndex+2]
}




// augmented merge sort approach
/*
func countSmaller(nums []int) []int {
	res, index := make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		index[i] = i
	}

	mergeSort(&res, &index, 0, len(nums)-1, nums)
	return res
}

func mergeSort(res, index *[]int, l, r int, nums []int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2
	mergeSort(res, index, l, mid, nums)
	mergeSort(res, index, mid+1, r, nums)
	merge(res, index, l, mid, mid+1, r, nums)
}

func merge(res, index *[]int, l1, r1, l2, r2 int, nums []int) {
	tmp := make([]int, r2-l1+1)
	start, count, p := l1, 0, 0
	for l1 <= r1 || l2 <= r2 {
		if l1 > r1 {
			tmp[p] = (*index)[l2]
			p += 1
			l2 += 1
		} else if l2 > r2 {
			(*res)[(*index)[l1]] += count
			tmp[p] = (*index)[l1]
			p += 1
			l1 += 1
		} else if nums[(*index)[l1]] > nums[(*index)[l2]] {
			tmp[p] = (*index)[l2]
			p += 1
			l2 += 1
			count++
		} else {
			(*res)[(*index)[l1]] += count
			tmp[p] = (*index)[l1]
			p += 1
			l1 += 1
		}
	}
	for i := 0; i < len(tmp); i++ {
		(*index)[start+i] = tmp[i]
	}
}
*/
/*
Input: nums = [5,2,6,1]
Output: [2,1,1,0]
*/

/*
// naiive BST approach; TLE's with the T.C.'s with a large size input'

func countSmaller(nums []int) []int {
    n := len(nums)
    if n == 1 {
        return []int{0}
    }
    res := make([]int, n)
    root := &Node{Val: nums[n-1], EqCnt: 1}
    for i := n-2; i >= 0; i-- {
        res[i] = insertNode(root, nums[i])
    }
    return res
}


type Node struct {
    Val, EqCnt, LessCnt int
    Left, Right *Node
}

// inserts Node data and returns cnt smaller numbers
func insertNode(root *Node, val int) int {
    res := 0
    for root != nil {
        if val > root.Val {
            res += root.EqCnt + root.LessCnt
            if root.Right == nil {
                root.Right = &Node{Val: val, EqCnt: 1}
                return res
            } else {
                root = root.Right
            }
        } else if val == root.Val {
            root.EqCnt++
            return res + root.LessCnt
        } else {
            root.LessCnt++
            if root.Left == nil {
                root.Left = &Node{Val: val, EqCnt: 1}
                return res
            } else {
                root = root.Left
            }
        }
    }
    return 0
}

*/
