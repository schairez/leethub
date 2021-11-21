type MinStack struct {
    dataStack []int
    minStack []int
}


func Constructor() MinStack {
    return MinStack{make([]int, 0), make([]int, 0)}
}


func min(a, b int) int {if a <= b { return a }; return b}

//O(1)
func (this *MinStack) Push(val int)  {
    this.dataStack = append(this.dataStack, val)
    minV := val
    if len(this.minStack) > 0 {
        minV = min(this.minStack[len(this.minStack)-1], val)
    }
    this.minStack = append(this.minStack, minV )
}


//O(1)
func (this *MinStack) Pop()  {
    n := len(this.dataStack)
    //dataItem := this.dataStack[n-1]
    //minItem := this.minStack[n-1]
    this.dataStack[n-1] = -1 //avoid memory leaks
    this.dataStack[n-1] = -1
    this.dataStack = this.dataStack[:n-1]
    this.minStack = this.minStack[:n-1]
}

//Methods pop, top and getMin operations will always be called on non-empty stacks.

//O(1)
func (this *MinStack) Top() int {
    n := len(this.dataStack)
    return this.dataStack[n-1] 
}

//O(1)
func (this *MinStack) GetMin() int {
    n := len(this.minStack)
    return this.minStack[n-1] 
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */