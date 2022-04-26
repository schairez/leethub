
type AutocompleteSystem struct {
    rootNode, currNode *TrieNode
    sentenceToIdx map[string]int
    sentenceData []pair
    sb strings.Builder
}


type pair struct {
    sentence string
    times int
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
    root := &TrieNode{}
    // condition: lenSentences == lenTimes 
    n := len(sentences)
    sentenceToIdx := make(map[string]int, n)
    sentenceData := make([]pair, n)
    for i := 0; i < n; i++ {
        sentenceToIdx[sentences[i]] = i
        sentenceData[i] = pair{sentences[i], times[i]}
        root.InsertSentence(i, sentences[i], times[i])
    }
    var sb strings.Builder
    return AutocompleteSystem{root, root, sentenceToIdx, sentenceData, sb }
}


func (this *AutocompleteSystem) Input(c byte) []string {
    if c == '#' {
        sentence := this.sb.String()
        if _, ok := this.sentenceToIdx[sentence]; !ok {
            this.sentenceData = append(this.sentenceData, pair{sentence, 0})
            this.sentenceToIdx[sentence] = len(this.sentenceData) - 1
        }
        wIdx := this.sentenceToIdx[sentence]
        this.sentenceData[wIdx].times++
        this.rootNode.InsertSentence(wIdx, sentence, 1)
        this.sb.Reset()
        this.currNode = this.rootNode
        return []string{}
    }
    //fmt.Println(string(c))
    //fmt.Println("data")
    //fmt.Println(this.sentenceData)
    this.sb.WriteByte(c)
    idx := 26
    if c != ' ' {
        idx = int(c - 'a')
    }
    // TODO
    if this.currNode.Children[idx] == nil {
        this.currNode.Children[idx] = &TrieNode{}
    }
    this.currNode = this.currNode.Children[idx]
    prefix := this.sb.String()
    // return top 3 historical hot sentences
    return this.topKPrefix(3, prefix)
    
}

func (this *AutocompleteSystem) topKPrefix(k int, prefix string) []string {
    // use the indexes for a tmp vector of sentences
    n := len(this.currNode.Indexes)
    if n < k {
        k = n
    }
    res := make([]string, 0, k)
    tmpArr := make([]pair, 0, n)
    for kIdx := range this.currNode.Indexes {
        tmpArr = append(tmpArr, this.sentenceData[kIdx])
    }
    sort.Slice(tmpArr, func(i, j int) bool {
        if tmpArr[i].times == tmpArr[j].times {
            return tmpArr[i].sentence < tmpArr[j].sentence
        }
        return tmpArr[i].times > tmpArr[j].times
    }) 
    //fmt.Println(tmpArr)
    
    for i := 0; i < k; i++ {
        res = append(res, tmpArr[i].sentence)
    }
    return res
}


type TrieNode struct {
    Indexes map[int]struct{}
    NumTimes int
    Children [27]*TrieNode
}

func (trie *TrieNode) InsertSentence(dataIdx int, sentence string, times int) {
    n := len(sentence)
    for i := 0; i < n; i++ {
        idx := 26
        if sentence[i] != ' ' {
            idx = int(sentence[i] - 'a')
        }
        if trie.Children[idx] == nil {
            trie.Children[idx] = &TrieNode{}
        }
        trie = trie.Children[idx]
        if trie.Indexes == nil {
            trie.Indexes = make(map[int]struct{})
        }
        trie.Indexes[dataIdx] = struct{}{}
    } 
    trie.NumTimes += times  
}


/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, Times);
 * param_1 := obj.Input(c);
 */