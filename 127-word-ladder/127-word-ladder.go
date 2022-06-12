
func ladderLength(beginWord string, endWord string, wordList []string) int {   
    wordSet := make(map[string]struct{}, len(wordList))
    for i := range wordList {
        wordSet[wordList[i]] = struct{}{}
    }
    lvl := 0
    var node []byte
    queue := make([][]byte, 0, len(wordList) + 1)
    queue = append(queue, []byte(beginWord))
    for len(queue) != 0 {
        lvl++
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            if string(node) == endWord {
                return lvl
            }
            cands := getCandidates([]byte(node), wordSet)
            for _, cand := range cands {
                queue = append(queue, cand)
                delete(wordSet, string(cand))
            }
            /*
            for wordKey := range wordSet {
                if canTransform(node, wordKey) {
                    queue = append(queue, wordKey)
                    delete(wordSet, wordKey)
                } 
            }
            */
        }
    }
    return 0
}

// get next 1 degree away candidates
func getCandidates(word []byte, wordSet map[string]struct{}) [][]byte {
    n := len(word)
    var cands [][]byte
    for i := 0; i < 26; i++ {
        for j := 0; j < len(word); j++ {
            origChar := word[j]
            nextChar := byte(int('a')+i)
            if origChar != nextChar {
                word[j] = nextChar
                candWord := string(word)
                if _, exists := wordSet[candWord]; exists {
                    tmp := make([]byte, n)
                    copy(tmp, word)
                    cands = append(cands, tmp)
                }
            }
            word[j] = origChar
        }
    }
    return cands
}


/*
func canTransform(w1, w2 string) bool {
    differs := 0
    for i := 0; i < len(w1); i++ {
        if w1[i] != w2[i] {
            differs++
            if differs > 1 {
                return false
            }
        }
    }
    return differs == 1
}
*/