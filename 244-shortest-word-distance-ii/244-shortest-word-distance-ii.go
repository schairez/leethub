type pair struct {
    w1, w2 string
}

type WordDistance struct {
    wordsLocMap map[string][]int
    //cacheMinDist map[pair]int
}

//time: O(n)
//space: O(n) where n = size of constructor input 

func Constructor(wordsDict []string) WordDistance {
    wordsLocMap := make(map[string][]int)
    for idx, word := range wordsDict {
        if _, ok := wordsLocMap[word]; !ok {
            wordsLocMap[word] = []int{idx}
        } else {
            wordsLocMap[word] = append(wordsLocMap[word], idx)
        }
    }
    return WordDistance{wordsLocMap: wordsLocMap,
                      //  cacheMinDist: make(map[pair]int),
                       }
    
}


/*
note: At most 5000 calls will be made to shortest.
- it's good to cache the distances and not repeat the same calculation again


time: O(lenW1locs + lenW2locs) where m = lenW1 idx locs; n = lenW2locs
space: O(1)
*/

func (this *WordDistance) Shortest(word1 string, word2 string) int {
    //Q: should we cache the min distance? question does not make clear whether
    //we'll have repeated queries or not
    /*
    if _, ok := this.cacheMinDist[pair{word1, word2}]; ok {
        return this.cacheMinDist[pair{word1, word2}]
    }
    if _, ok := this.cacheMinDist[pair{word2, word1}]; ok {
        return this.cacheMinDist[pair{word2, word1}]
    }
    */
    
    shortestDist := (1 << 31)-1
    //word1 and word2 are in wordsDict.
    word1Locs, word2Locs := this.wordsLocMap[word1], this.wordsLocMap[word2]
    w1Idx, w2Idx := 0, 0
    for w1Idx != len(word1Locs) && w2Idx != len(word2Locs) {
        distW1 := word1Locs[w1Idx]
        distW2 := word2Locs[w2Idx]
        
        shortestDist = min(shortestDist, abs(distW1 - distW2))
        if shortestDist == 1 { break }
        
        if distW1 < distW2 {
            w1Idx++
        } else {
            w2Idx++
        }
    }
   // this.cacheMinDist[pair{word1, word2}] = shortestDist
    
    return shortestDist
    
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func min(a, b int) int { if a <= b { return a}; return b }



/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */