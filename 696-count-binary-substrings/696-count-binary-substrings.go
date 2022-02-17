
/*
       --  --
s = "00110011" = min(cnt(0),cnt(1)) + min(cnt(1), cnt(0)) + min(cnt(0),cnt(1)) 
     --  --    =           2        +       2             +   2
               =  6

*/


//time: O(n)
//space: O(1)

func min(a, b int) int { if a <= b { return a}; return b}

func countBinarySubstrings(s string) int {
    n := len(s)
    if n == 1 {
        return 0
    }
    //get cnt of searchChar and return next idx
    getCntAndIdx := func(idx int, iscurrZero bool) (int, int) {
        var searchChar byte
        if iscurrZero {
            searchChar = '0'
        } else {
            searchChar = '1' 
        } 
        cnt := 0
        for idx < n && s[idx] == searchChar {
            cnt++
            idx++
        } 
        return cnt, idx
    }
    var res int
    iter1, iter2 := 0, 0
    for iter1 < n && iter2 < n {
        char1Cnt, char2Cnt := 0, 0
        char1Cnt, iter1 = getCntAndIdx(iter1, s[iter1] == '0')
        iter2 = iter1
        if iter2 == n {
            break
        }
        char2Cnt, iter2 = getCntAndIdx(iter2, s[iter2] == '0') 
        
        res += min(char1Cnt, char2Cnt)
    }
    return res
}










































/*
r   reset cnt0
l r V
  --- ---
0 1 0 1 0 1
--- --- ---
  ^
  "01" res++
 (l,r)   
  r
  l   i
      V 
 "00110011"
s[0]=2 s[1]=2
   l=r; l =3
   s[0] = 0
   i =r = 4 s[0]++ -> 1
   
  s[1] = 2
  s[0] = 0 
     V
     l=3
     r=4  s[0]++ -> 1; s[1]=2
 "0011011"    "110" || "001" +1
     --    
  ----
   +2
  
res := 0
arr = [2]int
l,r = 0,0
while r < len(s):
    arr[s[r]]++
    if r != 0 && s[r-1] != s[r] {
        if arr[0] == arr[1] { //"0,1"
            res = res + arr[0]
            arr[s[l]] = 0
            l = r
        }
    }
    r++



    if cntL != cntR && cntL != 0 || cntR != 0 {
        arr[s[r]]++
    } else {
        
    }







l,r = 0; move r until a breaking cond
   i3             cnt1, cnt0
   V   V      as we move r; we check cnt1 == cnt0         
        r
 l    r     
    l  r  3 0's
1111000111 
   --
  ----
 ------
4 1's 3 0's



cntSoFar of 0 -> 2; and of 1 is also 2
+1 to res; +(n-1) to res
n =

- the whole string 000111
- 01
- 0011

00001111  +4 to our res
  ^^
 
  
*/