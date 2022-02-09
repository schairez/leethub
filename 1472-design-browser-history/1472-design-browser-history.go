
type BrowserHistory struct {
    currIdx int
    pages []string
}


//D.S. holds O(n) space; traversal time is O(steps) forward and back
func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{ currIdx: 0, 
                          pages: []string{homepage},
                         }
}


func (this *BrowserHistory) Visit(url string)  {
    this.pages = this.pages[:this.currIdx+1]
    this.pages = append(this.pages, url)
    this.currIdx = len(this.pages) - 1
}



func (this *BrowserHistory) Back(steps int) string {
    if steps > this.currIdx {
        steps = this.currIdx 
    }
    this.currIdx -= steps
    return this.pages[this.currIdx]
}


func (this *BrowserHistory) Forward(steps int) string {
    //num steps to last idx
    remaining := len(this.pages)-1 - this.currIdx
    if steps > remaining {
        steps = remaining
    }
    this.currIdx += steps
    return this.pages[this.currIdx]
}












//////////////////////////////////////////////////


//doubly linked list version
//NOTE: nodes are stored on the memory heap, so below is 
// less efficient than an array based approach
type DoublyListNode struct {
    Website  string
    Next *DoublyListNode
    Prev *DoublyListNode
}


//D.S. holds O(n) space; traversal time is O(steps) forward and back
type BrowserHistoryNodeVersion struct {
    CurrPage *DoublyListNode
}


func ConstructorNodeVersion(homepage string) BrowserHistoryNodeVersion {
    node := &DoublyListNode{Website: homepage,
                            Next: nil, Prev: nil}
    
    return BrowserHistoryNodeVersion{CurrPage: node}
}


func (this *BrowserHistoryNodeVersion) Visit(url string)  {
    updatedCurrPage := &DoublyListNode{Website: url}
    prevPage := this.CurrPage
    //help out our GC
    if prevPage.Next != nil {
        tmp := prevPage.Next
        tmp.Prev = nil
    }
    prevPage.Next = nil
    prevPage.Next = updatedCurrPage
    updatedCurrPage.Prev = prevPage
    this.CurrPage = updatedCurrPage
}


func (this *BrowserHistoryNodeVersion) Back(steps int) string {
    i := 0
    for i != steps {
        if this.CurrPage.Prev == nil { 
            break 
        }
        this.CurrPage = this.CurrPage.Prev
        i++
    }
    return this.CurrPage.Website
}


func (this *BrowserHistoryNodeVersion) Forward(steps int) string {
    i := 0
    for i != steps {
        if this.CurrPage.Next == nil {
            break
        }
        this.CurrPage = this.CurrPage.Next
        i++
    }
    return this.CurrPage.Website
}


/**
 * Your BrowserHistoryNode  object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */