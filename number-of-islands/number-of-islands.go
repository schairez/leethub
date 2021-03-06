//m*n = total cells in grid; possible dirs = 4
//path compressed UF GetRoot() fn time: O(logM*N)
//time: O(m*n*4log(m*n) + n ) ~ O(M*N*log(M*N))
//space: O(N)

func numIslands(grid [][]byte) int {
    fmt.Println(grid)
    numR, numC := len(grid), len(grid[0])
    isValid := func(nextR, nextC int) bool {
        if nextR < 0 || nextC < 0 || nextR >= numR || nextC >= numC {
            return false
        }
        return true
    }
    get1DRep  := func(nextR, nextC int) int {
        idx := nextR * numC + nextC  //y * width + x
        return idx
    }
    //total cells
    totalItems := numR * numC
    uf := &UF{Parent: make([]int, totalItems), Sze: make([]int, totalItems)} //ValidVisitedItemMap: make(map[int]bool)}
    for i := range uf.Parent {
        uf.Parent[i] = -1
    }
    for i := range uf.Sze {
        uf.Sze[i] = 1
    }    
    
    dirs := [4][2]int{{0, 1}, //right
                      {0, -1}, //left
                      {1, 0}, //up
                      {-1, 0}, //down
                     }
    for y:=0; y < numR; y++ {
        for x:=0; x < numC; x++ {
            if grid[y][x] == '0' { continue }
            idxA := get1DRep(y, x)
            uf.ValidateNode(idxA)
            for _, dir := range dirs {
                dy, dx := dir[0], dir[1]
                nextR, nextC := dy + y, dx + x
                if isValid(nextR, nextC) && grid[nextR][nextC] == '1' {
                    idxB := get1DRep(nextR, nextC)
                    //uf.ValidateNode(idxA)
                    uf.Union(idxA, idxB)
                }
            }
        }
    }
   
    return uf.CntConnectedComponents()
}

type UF struct {
    Parent []int
    Sze []int
}
func (uf *UF) CntConnectedComponents() int {
    set := make(map[int]bool)
    for _, v := range uf.Parent {
        if v == -1 { continue } 
        root := uf.Root(v)
        set[root] = true
    }
    return len(set)
}
//make node valid
func (uf *UF) ValidateNode(p int) {
    if sentinel := uf.Parent[p]; sentinel == -1 {
        uf.Parent[p] = p
    }
}

//gets topmost Parent / Root
//path compression approach; O(logN)
func (uf *UF) Root(p int) int {
    if uf.Parent[p] == -1 {
        uf.Parent[p] = p
    }
    for p != uf.Parent[p] {
        uf.Parent[p] = uf.Parent[uf.Parent[p]]
        p = uf.Parent[p] 
    }
    return p
}

func (uf *UF) Union(p, q int) {
    uf.ValidateNode(p)
    uf.ValidateNode(q)
    pRoot := uf.Root(p)
    qRoot := uf.Root(q)
    if pRoot == qRoot { return }
    if uf.Sze[pRoot] < uf.Sze[qRoot] {
        uf.Parent[pRoot] = qRoot
        uf.Sze[qRoot] += uf.Sze[pRoot]
    } else {
        uf.Parent[qRoot] = pRoot
        uf.Sze[pRoot] += uf.Sze[qRoot]
        
    }
}


