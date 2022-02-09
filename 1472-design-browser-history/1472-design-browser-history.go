/*
int k -> linkedListNode -> <- 

forward history
backward history
--------------------> hackernews 
|
curr
goog -> LC -> FB
   ^          |
   ------------
*/

type DoublyListNode struct {
    Website  string
    Next *DoublyListNode
    Prev *DoublyListNode
}


type BrowserHistory struct {
    CurrPage *DoublyListNode
}


func Constructor(homepage string) BrowserHistory {
    node := &DoublyListNode{Website: homepage,
                            Next: nil, Prev: nil}
    
    return BrowserHistory{CurrPage: node}
}


func (this *BrowserHistory) Visit(url string)  {
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


func (this *BrowserHistory) Back(steps int) string {
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


func (this *BrowserHistory) Forward(steps int) string {
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
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */