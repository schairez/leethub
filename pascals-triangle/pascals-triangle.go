//time:O(n^2)
//space:O(n^2)


func generate(numRows int) [][]int {
    res := make([][]int, numRows)
    res[0] = append(res[0], 1)
    if numRows == 1 {
        return res
    }
    res[1] = append(res[1], []int{1,1}...)
    
    for i:=2; i < len(res); i++ {
        res[i] = make([]int, i+1)
        res[i][0], res[i][len(res[i])-1] = 1, 1    
        for j:=1; j < len(res[i])-1; j++ {
            res[i][j] = res[i-1][j-1] + res[i-1][j]    
        }
    }
     return res
}


/*

res[0] = [] -> [1]
res[0] = append(res[0], 1)
res[1] = append(res[1], []int{1,1}...) 

//len =5 
//2; 3; 4; 5 breaks

//   start;     end;   step
      1     --> res[0][0] = 1 
     1 1    --> res[1][0..1] = res[0][]
    1 2 1   --> res[2][0..2] 
  1  3  3  1 --> res[3][0..3]
 1  4  6  4  1  
 
  
 
  0   1   2   3   4 
 [[], [], [], [], []]
 
 []int
 [][]int --> res[0][0]
  
*/



