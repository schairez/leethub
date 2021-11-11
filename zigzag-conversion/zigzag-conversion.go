
import "strings"

type dir uint8

const (
    north dir =  iota
    south
)

func convert(s string, numRows int) string {
    if numRows == 1 { return s }
    
    
    nextIdxAndDir := func(idx int, curDir dir) (int, dir) {
        switch curDir {
            case north:
                if nextV := idx -1; nextV < 0 { return idx+1, south }
            case south:
                if nextV := idx +1; nextV > numRows-1 { return idx-1, north }
        }
        if curDir == north { return idx-1, curDir }
        return idx+1, curDir
    }
    arrSBs := make([]strings.Builder, numRows)
    for i :=0; i < numRows; i++ {
        sb := strings.Builder{}
        arrSBs = append(arrSBs, sb )
    }
    idx := 0
    curDir := south
    for i := range []byte(s) {
        arrSBs[idx].WriteByte(s[i])
        idx, curDir = nextIdxAndDir(idx, curDir)
    }
    var sb strings.Builder
    sb.Grow(len(s))
    for len(arrSBs) > 0 {
        nextRowStr := arrSBs[0].String()
        sb.WriteString(nextRowStr)
        arrSBs = arrSBs[1:]
    }
    
    return sb.String()
    
}



/*

"PAYPALISHIRING"  numRows = 3
Output: "PAHNAPLSIIGYIR"



P   A   H   N
A P L S I I G
Y   I   R


*/