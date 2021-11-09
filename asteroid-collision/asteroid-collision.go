// time: O(n)

// idx0; append to res
// idx1; if both positive then do nothing res[len(res)-1]  
// idx2; -15; -15 with res[len(res)-1]; abs(i) abs(j); if j > i; append(j)
//  while abs(j) > abs(res[len(res)-1])
// monotonic stack

//[5,10,-15]
// [5]
// [5,10]
//  [5] ; 
// [-15]
    
// [8,-8]
// [8]

// [-2,-1,1,2]
//  |   |
//// [2, 1, -1, -2]
//   [2,1] ; [2]; -2 

//[5,10,-5]
//[5,10]
// 

type sign uint8

const (
    pos sign = iota
    neg
)

func asteroidCollision(asteroids []int) []int {
    getSign := func(v int) sign {
        if v < 0 { return neg }
        return pos
    }
    res := []int{}    
    addToRes := true 
    for _, val := range asteroids {
        addToRes = true
        for len(res) > 0 && val < 0 { //if neg
            lastV := res[len(res)-1]
            lastValSize := abs(lastV)
            newValSize := abs(val) 
            signNewVal, signLastVal := getSign(lastV), getSign(val) 
            if signNewVal == signLastVal {
                break
            } else if newValSize < lastValSize {
                addToRes = false
                break
            } else if newValSize > lastValSize {
                res = res[:len(res)- 1]
            } else { //else if newValSize == lastValSize {
                res = res[:len(res) -1]
                addToRes = false
                break
            }
        }
        if addToRes { 
            res = append(res, val) 
        }
    }
    return res
}


func abs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}