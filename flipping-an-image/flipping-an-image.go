//time: O(numR*numC)
//space:O(1)


func flipAndInvertImage(image [][]int) [][]int {
    invertNum := func(num int) int {
        if num == 1 { return 0 }
        return 1
    }
    //conditions
    // numR == numC
    //1 <= n <= 20
    numCols := len(image[0])
    if numCols == 1 { 
        image[0][0] = invertNum(image[0][0])
        return image
    }
    var isEven bool
    if numCols % 2 == 0 {
        isEven = true
    }
    for row := range image {
        colStart, colEnd := 0, len(image[row])-1
        for colStart < colEnd {
            image[row][colStart], image[row][colEnd] = image[row][colEnd], image[row][colStart]
            image[row][colStart] = invertNum(image[row][colStart])
            image[row][colEnd] = invertNum(image[row][colEnd])
            if !isEven && colStart+1 == colEnd-1 {
                image[row][colStart+1] = invertNum(image[row][colStart+1])
            }
            colStart++
            colEnd--
        }
    }
    return image
}
