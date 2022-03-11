
// 232. Implement Queue using Stacks
// time: avg/amortized O(1) worst O(n) for every empty pop/peek 
// space: O(n) store 

type MyQueue struct {
    tailStack []int
    headStack []int
}


func Constructor() MyQueue {
    return MyQueue{make([]int,0), make([]int, 0)}
}


func (this *MyQueue) Push(x int)  {
    this.tailStack = append(this.tailStack, x)
}



func (this *MyQueue) Pop() int {
    item := this.Peek()
    n := len(this.headStack)
    this.headStack = this.headStack[:n-1]
    return item
}


func (this *MyQueue) Peek() int {
    if len(this.headStack) == 0 {
        n := len(this.tailStack)
        for n != 0 {
            v := this.tailStack[n-1]
            this.tailStack = this.tailStack[:n-1] 
            this.headStack = append(this.headStack, v)
            n--
        }
    }
    idx := len(this.headStack)-1
    return this.headStack[idx]
}


func (this *MyQueue) Empty() bool {
    return len(this.tailStack) == 0 && len(this.headStack) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */