
"""

in = [1,2,3]
     [4,5,6]
     [7,8,9]

out = [12,21,16]
      [27,45,33]
      [24,39,28]
      
k = 1 

@ mat[0][0] 
r >= 0 and r <= 1; 
c >= 0 and c <= 1 

"""

"""
time: O(n^4)
space: O(n*m)

"""

class Solution:
    def matrixBlockSum(self, mat: List[List[int]], k: int) -> List[List[int]]:
        m = len(mat)
        n = len(mat[0])
        res = [[0 for c in range(n)] for r in range(m)]
        for i in range(0, m):
            for j in range(0, n):
                rmin, rmax  = i - k , i+k
                rmin = 0 if rmin < 0 else rmin
                rmax = m-1 if rmax > m-1 else rmax 
                cmin, cmax = j - k, j + k 
                cmin = 0 if cmin < 0 else cmin
                cmax = n-1 if cmax > n-1 else cmax 
                val = 0 
                for r in range(rmin, rmax+1):
                    for c in range(cmin, cmax+1):
                        val += mat[r][c]
                
                res[i][j] = val 
        return res 
        