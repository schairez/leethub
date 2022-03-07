// 127. Word Ladder
// BFS+hashmap approach
// N = numWords; W = largestWordLen in wordList
// time: 
// O(N*W) to build wordSet
// O(W^2) to iterate and build a word transformations out of 26 dict chars
// and worst case we do this for every word in our wordList
// thus, time: O(W^2*N) + O(N*W) ≈ O(W^2*N)
// space: O(N*W)


func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordSet := make(map[string]struct{}, len(wordList))
    for _, word := range wordList {
        wordSet[word] = struct{}{}
    }
    if _, exists := wordSet[endWord]; !exists {
        return 0
    }
    queue := [][]byte{[]byte(beginWord)}
    var numWords int //numWords in path
    for len(queue) != 0 {
        numWords++
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[i]
            if string(node) == endWord {
                return numWords
            }
            cands := getCandidates(node, wordSet)
            for _, cand := range cands {
                queue = append(queue, cand)
                delete(wordSet, string(cand))
            }
        }
        queue = queue[size:]
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
