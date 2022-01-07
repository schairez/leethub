/*
graph forms a valid tree if:
- acyclic
- n-1 edges to qualify fully connected w/o cycles
- all nodes are connected, so a path exists from u to v for any vertex
  in the graph
- no cycles, so no redundant connections or backedges

*/
//below is a unionfind approach
//time: O(E)
//space: O(N)

func validTree(n int, edges [][]int) bool {
    //n-1 edges for a graph to be a tree
    if n-1 != len(edges) {
        return false
    }
    uf := NewUF(n)
    for _,edge := range edges {
        if ok := uf.Union(edge[0], edge[1]); !ok {
            return false
        }
    }
    return true
}


type UnionFind struct {
    parent []int
    sze []int
    
}

func NewUF(numNodes int) UnionFind {
    parent := make([]int, numNodes)
    sze := make([]int, numNodes)
    for node := range parent {
        parent[node] = node
        sze[node] = 1
    }
    return UnionFind{parent, sze}
}

func (uf *UnionFind) Root(node int) int {
    for node != uf.parent[node] {
        uf.parent[node] = uf.parent[uf.parent[node]]
        node = uf.parent[node]
    }
    return node
}

func (uf *UnionFind) Union(v1, v2 int) bool {
    v1Root := uf.Root(v1)
    v2Root := uf.Root(v2)
    //already connected, edge is a cycle
    if v1Root == v2Root {
        return false
    }
    if uf.sze[v1Root] <= uf.sze[v2Root] {
        uf.sze[v2Root] += uf.sze[v1Root]
        uf.parent[v1Root] = v2Root
    } else {
        uf.sze[v1Root] += uf.sze[v2Root]
        uf.parent[v2Root] = v1Root
    }
    
    return true
} 