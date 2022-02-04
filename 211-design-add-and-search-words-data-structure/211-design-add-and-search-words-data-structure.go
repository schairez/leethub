
type WordDictionary struct {
    Trie *TrieNode
    MaxWordLen int
}

func Constructor() WordDictionary {
    return WordDictionary{&TrieNode{}, 0}
    
}


func (this *WordDictionary) AddWord(word string)  {
    if len(word) > this.MaxWordLen {
        this.MaxWordLen = len(word)
    }
    this.Trie.AddWord(word)
}


func (this *WordDictionary) Search(word string) bool {
    if len(word) > this.MaxWordLen {
        return false
    }
    return this.Trie.Search(word)
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */


type TrieNode struct {
    Children [26]*TrieNode
    IsWord bool
}

func (trie *TrieNode) AddWord(word string) {
    wordLen := len(word)
    for i := 0; i < wordLen; i++ {
        idx := word[i] - 'a'
        if trie.Children[idx] == nil {
            trie.Children[idx] = &TrieNode{}
        }
        trie = trie.Children[idx]
    }
    trie.IsWord = true
}

//time: O(lenWord) if alpha bytes input; 
// O(26^N) where N = lenWord if '.' in bytesequence
//space: O(N) N = lenWord  

func (trie *TrieNode) Search(word string) bool {
    wordLen := len(word)
    
    var dfs func(wordIdx int, node *TrieNode) bool
    dfs = func(wordIdx int, node *TrieNode) bool {
        if wordIdx == wordLen {
            return node.IsWord
        }
        if word[wordIdx] == '.' {
            for i :=0; i < 26; i++ {
                if node.Children[i] == nil {
                    continue
                }
                if dfs(wordIdx+1, node.Children[i] ) {
                    return true
                }
            }
        } else if node.Children[word[wordIdx] - 'a'] != nil {
            node = node.Children[word[wordIdx] - 'a']
            return dfs(wordIdx+1, node)
        }
        return false
    }
    return dfs(0, trie)
}
