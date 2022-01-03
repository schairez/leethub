//time: O(n) for popMax(); O(1) for other methods 
//space: O(2n) ~ O(n)

//There will be at least one element in the stack when pop, top, peekMax, or popMax is called.

type MaxStack struct {
    Items Stack
    MaxItems Stack
}


func Constructor() MaxStack {
    itemStack := NewStack()
    maxStack := NewStack()
    return MaxStack{itemStack, maxStack}
    
}


func (this *MaxStack) Push(x int)  {
    var maxVal int
    if this.MaxItems.IsEmpty() {
        maxVal = x
    } else if this.MaxItems.Peek() > x {
        maxVal = this.MaxItems.Peek()
    } else {
        maxVal = x
    }
    this.MaxItems.Push(maxVal)
    this.Items.Push(x)
}


func (this *MaxStack) Pop() int {
    this.MaxItems.Pop()
    v := this.Items.Pop()
    return v
}


func (this *MaxStack) Top() int {
    return this.Items.Peek()
}


func (this *MaxStack) PeekMax() int {
    return this.MaxItems.Peek()
}


func (this *MaxStack) PopMax() int {
    buffer := NewStack()
    maxVal := this.MaxItems.Peek()
    for this.Top() != maxVal {
        buffer.Push(this.Pop())
    }
    this.Pop()
    for !buffer.IsEmpty() {
        this.Push(buffer.Pop())
    }
    return maxVal
}

type Stack []int

func NewStack() Stack {
    return make(Stack, 0)
}
func (st *Stack) IsEmpty() bool {
    return len(*st) == 0
}

func (st *Stack) Len() int {
    return len(*st)
}
func (st *Stack) Pop() int {
    n := st.Len()
    if !st.IsEmpty() {
        item := (*st)[n-1]
        (*st) = (*st)[:n-1]
        return item
    }
    return -1
}

func (st *Stack) Push(val int) {
    (*st) = append((*st), val)
}

func (st *Stack) Peek() int {
    n := st.Len()
    if !st.IsEmpty() {
        item := (*st)[n-1]
        return item
    }
    return -1
}





/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */