import "sort"

//time: O(n + nlogn + k ) ~ O(nlogn)
//space: O(2n) ~ O(n)

type wordData struct {
    word string
    freq int
}

func topKFrequent(words []string, k int) []string {
    //map a word to idx on wordFreq Arr
    wordIdxLookup := make(map[string]int, 0)
    wordFreq := make([]wordData, 0)
    for _, word := range words {
        if idx, exists := wordIdxLookup[word]; exists {
            wordFreq[idx].freq++
            continue
        }
        wordFreq = append(wordFreq, 
                          wordData{word: word, freq: 1})
        idx := len(wordFreq)-1
        wordIdxLookup[word] = idx 
    }
    sort.Slice(wordFreq, func(i, j int) bool {
        if wordFreq[i].freq == wordFreq[j].freq {
            return wordFreq[i].word < wordFreq[j].word
        }
        return wordFreq[i].freq > wordFreq[j].freq
    })
    res := make([]string, k)
    for i := 0; i < k; i++ {
        res[i] = wordFreq[i].word
    }
    return res
    
}