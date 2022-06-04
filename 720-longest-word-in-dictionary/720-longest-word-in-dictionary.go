// 720. Longest Word in Dictionary
// time: O(nlogn) to sort words
// to insert words into trie n words K = longest length of word
// to insert into trie : O(nk)
// to traverse dfs: we're bound to longest word with the most character growing prefix sequences O(w)
// thus inserting into trie is our largest time bound
// space to sort O(logn) but traversing the dfs and building the trie is the largest recursive bound
// thus space: O(w)

type TrieNode struct {
    Children [26]*TrieNode
    IsWord bool
}

func (trie *TrieNode) InsertWord(word string) {
    n := len(word)
    for i := 0; i < n; i++ {
        if trie.Children[word[i]-'a'] == nil {
            trie.Children[word[i]-'a'] = &TrieNode{}
        }
        trie = trie.Children[word[i]-'a']
    } 
    trie.IsWord = true
}


func longestWord(words []string) string {
    sort.Strings(words)
    root := &TrieNode{}
    for _, word := range words {
        root.InsertWord(word)
    }
    var (
        res string
        dfs func(*TrieNode, []byte)
    )
    dfs = func(node *TrieNode, currWord []byte) {
        if len(currWord) > len(res) || 
        len(currWord) == len(res) && string(currWord) < res {
            res = string(currWord)
        } 
        for i := 0; i < 26; i++ {
            if childNode := node.Children[i]; childNode != nil && childNode.IsWord {
                currWord = append(currWord, byte(i+'a'))
                dfs(node.Children[i], currWord)
                currWord = currWord[:len(currWord)-1]
            }
        }
    }
    dfs(root, []byte{})
    return res
}