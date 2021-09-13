"""
we turn clockwise when out of bounds
time: O(m*n); m = num_rows; n = num_cols
space: O(n) result array; but O(1) extra space if we don't include the res_arr
"""

class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        m = len(matrix)
        n = len(matrix[0])
        top, bottom = 0, m-1
        left, right = 0, n-1
        res : list[int] = list()
        while True:
            if left > right:
                break
            for i in range(left, right+1):
                res.append(matrix[top][i])
            # update top row
            top +=1
            if top > bottom:
                break
            for i in range(top, bottom+1):
                res.append(matrix[i][right])
            right -= 1
            if left > right:
                break
            # bottom row right to left 
            for i in range(right, left-1, -1):
                res.append(matrix[bottom][i])
            bottom -= 1
            if top > bottom:
                break
            for i in range(bottom, top-1, -1):
                res.append(matrix[i][left])
            left +=1 
        return res
            
            
            
        