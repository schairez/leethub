
func fourSum(nums []int, target int) [][]int {
    n := len(nums)
    if n < 4 {
        return [][]int{}
    }
    keysMap := make(map[[4]int]struct{})
    sort.Ints(nums)
    var res [][]int
    for fourthIdx := n-1; fourthIdx > 2; fourthIdx-- {
        /*
        if fourthIdx != n-1 && nums[fourthIdx] == nums[fourthIdx+1] {
            continue
        }
        */
        n4 := nums[fourthIdx]
        for thirdIdx := fourthIdx-1; thirdIdx > 1; thirdIdx-- {
            /*
            if thirdIdx != n-2 && nums[thirdIdx] == nums[thirdIdx+1] {
                continue
            }
            */
            n3 := nums[thirdIdx]
            start, end := 0, thirdIdx-1
            for start < end {
                n1, n2 := nums[start], nums[end]
                fourSumCand := n4 + n3 + n2 + n1
                if fourSumCand == target {
                    keyCand := [4]int{n1, n2, n3, n4}
                    if _, ok := keysMap[keyCand]; !ok {
                        keysMap[keyCand] = struct{}{}
                    }
                    //res = append(res, []int{n1, n2, n3, n4})
                    start++
                    //end--
                    for start < end && n1 == nums[start] {
                        start++
                    } 
                    /*
                    for end > start && n2 == nums[end] {
                        end--
                    }
                    */
                } else if fourSumCand < target {
                    start++
                } else {
                    end--
                }
            }
        }
    }
    fmt.Println(keysMap)
    res = make([][]int, 0, len(keysMap))
    for key := range keysMap {
        tmp := make([]int, 0, 4)
        for i := 0; i < 4; i++ {
            tmp = append(tmp, key[i])
        }
        res = append(res, tmp)
    }
    
    return res
}