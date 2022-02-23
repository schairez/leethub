/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */


// 364. Nested List Weight Sum II
// bfs approach
// time: O(n)
// space: O(n)

type item struct {
    val int
    depth int
}

func max(a, b int) int { if a >= b { return a}; return b }

func depthSumInverse(nestedList []*NestedInteger) int {
    var items []item
    var queue [][]*NestedInteger
    queue = append(queue, nestedList)
    depth := 1
    var nodeList []*NestedInteger
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            nodeList, queue = queue[0], queue[1:]
            for _, node := range nodeList {
                if !node.IsInteger() {
                    queue = append(queue, node.GetList())
                } else {
                    items = append(items, item{node.GetInteger(), depth})
                }
            }
        }
        if len(queue) != 0 {
            depth++
        }
    }
    maxDepth := depth
    res := 0
    for _, item := range items {
        weight := maxDepth - item.depth + 1
        res += weight * item.val
    }
    return res
}

/*
// 364. Nested List Weight Sum II
// dfs approach
// time: O(n)
// space: O(n)

func dfs(nestedList []*NestedInteger, depthSoFar int, items *[]item) int {
    maxDepth := depthSoFar
    for _, elem := range nestedList {
        if !elem.IsInteger() {
            maxDepth = max(dfs(elem.GetList(), depthSoFar+1, items), maxDepth)
        } else {
            elemVal := elem.GetInteger()
            *items = append(*items, item{elemVal, depthSoFar})
        }
    }
    return maxDepth
}

func depthSumInverse(nestedList []*NestedInteger) int {
    res := 0
    var items []item
    maxDepth := dfs(nestedList, 1, &items)
    for _, item := range items {
        weight := maxDepth - item.depth + 1
        res += item.val * weight
    }
    return res
}
*/