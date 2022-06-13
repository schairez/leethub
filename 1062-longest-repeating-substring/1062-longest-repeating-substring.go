
func longestRepeatingSubstring(s string) int {
    sa := build(s)
    lcp := kasai(s, sa)
    res := 0
    for i := range lcp {
        if lcp[i] > res {
            res = lcp[i]
        }
    }
    return res
}

func kasai(txt string, sa []int) []int {
    n := len(sa)
    lcp := make([]int, n)
    saInv := make([]int, n) // inverse suffix arr
    for i := 0; i < n; i++ {
        saInv[sa[i]] = i
    }
    k := 0
    for i := 0; i < n; i++ {
        if saInv[i] == n-1 {
            k = 0
            continue
        }
        j := sa[saInv[i]+1]
        for i + k < n && j + k < n && txt[i+k] == txt[j+k] {
            k++
        }
        lcp[saInv[i]] = k
        if k > 0 {
            k--
        }
    }
    
    return lcp
}


type SuffixNode struct {
    Idx int
    Rank [2]int // store Rank of char at Idx and next Idx
}

type BySuffix []*SuffixNode

func (s BySuffix) Len() int {
    return len(s)
} 
func (s BySuffix) Less(i, j int) bool {
    if s[i].Rank[0] == s[j].Rank[0] {
        return s[i].Rank[1] < s[j].Rank[1]
    }
    return s[i].Rank[0] < s[j].Rank[0]
}

func (s BySuffix) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}


// build suffix arr
func build(txt string) []int {
    n := len(txt)
    //suffixData := &SuffixData{Data: make([]*SuffixNode, 0, n)}
    suffixData := make([]*SuffixNode, 0, n)
    for i := range txt {
        r0 := int(txt[i]-'a')
        r1 := -1
        if i+1 < n {
            r1 = int(txt[i+1]-'a')
        }
        node := &SuffixNode{Idx: i, Rank: [2]int{r0, r1}}
        suffixData = append(suffixData, node)
    }
    sort.Sort(BySuffix(suffixData))
    indices := make([]int, n)
    for j := 4; j < 2*n; j *= 2 {
        currRank := 0
        prevRank := suffixData[0].Rank[0]
        suffixData[0].Rank[0] = currRank 
        indices[suffixData[0].Idx] = 0
        for i := 1; i < n; i++ {
            if suffixData[i].Rank[0] == prevRank &&
            suffixData[i].Rank[1] == suffixData[i-1].Rank[1] {
                prevRank = suffixData[i].Rank[0]
                suffixData[i].Rank[0] = currRank
            } else {
                prevRank = suffixData[i].Rank[0]
                currRank++
                suffixData[i].Rank[0] = currRank
            }
            indices[suffixData[i].Idx] = i
        }
        for i := 0; i < n; i++ {
            nextIdx := suffixData[i].Idx + (j / 2) 
            r1 := -1
            if nextIdx < n {
                r1 = suffixData[indices[nextIdx]].Rank[0]
            }
            suffixData[i].Rank[1] = r1
        }
        sort.Sort(BySuffix(suffixData))
    }
    sa := make([]int, 0, n)
    for _, node := range suffixData {
        sa = append(sa, node.Idx)
    }
    return sa
}



