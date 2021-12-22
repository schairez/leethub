/*
Input: heights = [10,6,8,5,11,9]
Output: [3,1,2,1,1,0]

res = [0,0,0,0,0,0]
end of i=0; st = [0]
i=1; heights[0] > heights[1] so skip inner while
 - res[0]++; res = [1,0,0,0,0,0]; st = [0, 1]

i = 2; heights[1] <= heights[2] True; 
 - res[st.pop()]++; res[1]++
 - st = [0]; res[st.top()]++;
 - res = [2, 1, 0, 0, 0, 0]   
 - st.append(st, i)
 - st = [0, 2]

i = 3; heights[2] > heights[3]
  - res[st.top()]++
  - res = [2, 1, 1, 0, 0, 0]
  - st.append(st, i)
  - st = [0, 2, 3]
  
i = 4; heights[3] <= heights[4] True 
  - res[st.pop()]++; res[3]++
  - st = [0, 2]
  - heights[2] <= heights[4] True
    - res[st.pop()]++; res[2]++
    - st = [0]
    - heights[0] <= heights[4] True
      - res[st.pop()]++; res[0]++
      - st = []; break
  - res = [3, 1, 2, 1, 0, 0]
  - st = [4]
  
i = 5; heights[4] > heights[5]
  - res[st.top()]++; res[4]++
  - res = [3, 1, 2, 1, 1, 0]

*/

//time: O(n)
//space: O(n)
func canSeePersonsCount(heights []int) []int {
    n := len(heights)
    res := make([]int, n)
    st := make(stack, 0)
    for i := 0; i < n; i++ {
        for !st.isEmpty() && heights[st.peek()] <= heights[i] {
            res[st.pop()]++
        } 
        if !st.isEmpty() {
            res[st.peek()]++
        }
        st.push(i)
    }
    return res
}


type stack []int

func (st *stack) isEmpty() bool {
    return len((*st)) == 0
}

func (st *stack) push(val int) {
    (*st) = append((*st), val)
}

func (st *stack) peek() int {
    return (*st)[len((*st))-1]
}

func (st *stack) pop() int {
    val := st.peek()
    (*st) = (*st)[:len((*st))-1]
    return val
}
