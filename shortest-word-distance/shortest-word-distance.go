func abs(a int) int {
    if a < 0 { return -a }
    return a
}

func dist(idx1, idx2 int) int {
    return abs(idx1 - idx2)
}
func min(a, b int) int {
    if a <= b { return a}
    return b
}

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    var word1Idx, word2Idx int
    globalDist := len(wordsDict) -1 
    found1, found2 := false, false
    for i := range wordsDict {
        switch {
        case wordsDict[i] == word1:
            if !found1 { found1 = true }
            word1Idx = i
        case wordsDict[i] == word2:
            if !found2 { found2 = true }
            word2Idx = i
        }
        if found1 && found2 {
            //calc dist
            globalDist = min(dist(word2Idx,word1Idx), globalDist)
        }
    }
    return globalDist
}