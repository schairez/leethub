
//consists of only english letters so [26]*Trie 

type Trie struct {
    children [26]*Trie
    isWord bool
    
}


func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string)  {
    curr := this
    for i:=0; i < len(word); i++ {
        idx := word[i] - 'a'
        if v := curr.children[idx]; v == nil {
            curr.children[idx] = &Trie{}
        }
        curr = curr.children[idx]
    }
    curr.isWord = true
    return
}


func (this *Trie) Search(word string) bool {
    curr := this
    for i:=0; i < len(word); i++ {
        idx := word[i] - 'a'
        if v := curr.children[idx]; v == nil {
            return false
        }
        curr = curr.children[idx]
    }
    return curr.isWord
}


func (this *Trie) StartsWith(prefix string) bool {
    curr := this
    for i:=0; i < len(prefix); i++ {
        idx := prefix[i] - 'a'
        if v := curr.children[idx]; v == nil {
            return false
        }
        curr = curr.children[idx]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */