//N = len(edges)
//time: O(N*logN)
//space: O(N)

func countComponents(n int, edges [][]int) int {
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }
    size := make([]int, n)
    for i := range size {
        size[i] = 1
    }
    cnt := n
    
    //path compression; time: O(logN)
    findRoot := func(p int) int {
        for p != parent[p] {
            parent[p] = parent[parent[p]]
            p = parent[p]
        } 
        return p
    }
    union := func(p, q int) {
        pRoot := findRoot(p)
        qRoot := findRoot(q)
        if pRoot == qRoot { return }
        if size[pRoot] < size[qRoot] {
            parent[pRoot] = qRoot
            size[qRoot] += size[pRoot]
        } else {
            parent[qRoot] = pRoot
            size[pRoot] += size[qRoot]
        }
        cnt--
    }
    
    for i := range edges {
        edge := edges[i]
        union(edge[0], edge[1])
    }
    
    return cnt
    
}