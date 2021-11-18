//time: O(N)
//space:O(N)

/*

Return an edge that can be removed so that the resulting graph
is a tree of n nodes. If there are multiple answers, return the answer
that occurs last in the input.

*/

func findRedundantConnection(edges [][]int) []int {
	candidateEdges := [][]int{}
	uf := NewUnionFind(len(edges) + 1)
	for i := 0; i < len(edges); i++ {
		src, dst := edges[i][0], edges[i][1]
		//if uf.Union fails, we have an extraneous component
		isNewConnection := uf.Union(src, dst)
		if !isNewConnection {
			candidateEdges = append(candidateEdges, edges[i])
		}
	}
	return candidateEdges[len(candidateEdges)-1]

}

type UnionFind struct {
	Parent []int
	Sze    []int
}

func NewUnionFind(numNodes int) *UnionFind {
	uf := &UnionFind{Parent: make([]int, numNodes), Sze: make([]int, numNodes)}
	for i := range uf.Parent {
		uf.Parent[i] = i
		uf.Sze[i] = 1
	}
	return uf
}

//path compression approach
func (uf *UnionFind) Root(u int) int {
	for u != uf.Parent[u] {
		uf.Parent[u] = uf.Parent[uf.Parent[u]]
		u = uf.Parent[u]
	}
	return u
}

func (uf *UnionFind) Union(u, v int) bool {
	parU := uf.Root(u)
	parV := uf.Root(v)
	if parU == parV {
		return false
	}
	szeParU := uf.Sze[parU]
	szeParV := uf.Sze[parV]
	switch {
	case szeParU <= szeParV:
		uf.Sze[parU] += uf.Sze[parV]
		uf.Parent[parU] = parV

	case szeParV > szeParU:
		uf.Sze[parU] += uf.Sze[parV]
		uf.Parent[parU] = parV
	}
	return true
}
