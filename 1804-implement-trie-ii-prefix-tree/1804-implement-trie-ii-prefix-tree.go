
type Trie struct {
    Children  [26]*Trie
    PrefixCnt int
    WordCnt   int
}


func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string)  {
    for i := 0; i < len(word); i++ {
        if this.Children[word[i]-'a'] == nil {
            this.Children[word[i]-'a'] = &Trie{}
        }
        this = this.Children[word[i]-'a']
        this.PrefixCnt++
    }
    this.WordCnt++
    fmt.Println(this.WordCnt, this.PrefixCnt)
}

func (this *Trie) findNode(word string) *Trie {
    node := this
    for i := 0; i < len(word); i++ {
        if node == nil {
            return nil
        }
        node = node.Children[word[i]-'a']
    }
    return node
}

func (this *Trie) CountWordsEqualTo(word string) int {
    node := this.findNode(word)
    if node == nil {
        return 0
    }
    return node.WordCnt
}


func (this *Trie) CountWordsStartingWith(prefix string) int {
    node := this.findNode(prefix)
    if node == nil {
        return 0
    }
    return node.PrefixCnt
}


func (this *Trie) Erase(word string)  {
    node := this
    for i := 0; i < len(word); i++ {
        if node.Children[word[i]-'a'] == nil {
            return
        }
        node = node.Children[word[i]-'a']
        node.PrefixCnt--
    }
    node.WordCnt--
    if node.WordCnt == 0 {
        node = this 
        type pair struct {trieNode *Trie; key int }
        var key int
        var stack []pair
        for i := 0; i < len(word); i++ {
            key = int(word[i] - 'a')
            if node.Children[key] == nil {
                break
            }
            next := node.Children[key]
            if next.WordCnt == 0 && next.PrefixCnt == 0 {
                stack = append(stack, pair{node, key})
            }
        }
        for len(stack) != 0 {
            fmt.Println("wh")
            n := len(stack)
            p := stack[n-1].trieNode
            fmt.Println(p)
            c := stack[n-1].key
            fmt.Println(c)
            stack = stack[:n-1]
            p.Children[c] = nil
        }
    }
}

/*

func (this *Trie) hasChildren() bool {
    for i := 0; i < 26; i++ {
        if this.Children[i] != nil {
            return true
        }
    }
    return false
}

*/

// condition: guaranteed word is in trie
// postorder dfs traverse down the word path
// if cnt - 1 == 0 send signal up the chain to delete
// - this node, continue deleting until either backtrack 
// - up the root, or break when reaching a node with
// - isWord == true
// else just subtract cnt and dont delete nodes and break early

/*
func (this *Trie) Erase(word string)  {
    n := len(word)
    var dfs func(int, *Trie) *Trie 
    dfs = func(idx int, node *Trie) *Trie {
        if node == nil {
            return nil
        }
        node.PrefixCnt--
        if idx == n {
            //condition: guaranteed word is in trie
            node.WordCnt--
            if node.WordCnt == 0 {
                //fmt.Println("wha")
                //node.IsWord = false
                if !node.hasChildren() {
                    //erase up the chain
                    node = nil
                    return node
                }
            }
            return node
        }
        
        node.Children[word[0]-'a'] = dfs(idx+1, node.Children[word[idx]-'a'])
        //node.PrefixCnt--
        if !node.hasChildren() {
            return nil
        }
        return node
    }
    this = dfs(0, this)
    //dfs(0, this)
}
*/

/*
func (this *Trie) Erase(word string)  {
    n := len(word)
    var dfs func(int, *Trie) bool
    dfs = func(idx int, node *Trie) bool {
        if node == nil {
            return false
        }
        node.PrefixCnt--
        if idx == n {
            //condition: guaranteed word is in trie
            node.WordCnt--
            if node.WordCnt == 0 {
                //fmt.Println("wha")
                //node.IsWord = false
                if !node.hasChildren() {
                    //erase up the chain
                    //node = nil
                    return true
                }
            }
            return false
        }
        
        //node.PrefixCnt--
        eraseCond := dfs(idx+1, node.Children[word[idx]-'a'])
        if eraseCond {
            node.Children[word[0]-'a'] = nil
        }
        return false
    }
    
    dfs(0, this)
}
*/

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.CountWordsEqualTo(word);
 * param_3 := obj.CountWordsStartingWith(prefix);
 * obj.Erase(word);
 */