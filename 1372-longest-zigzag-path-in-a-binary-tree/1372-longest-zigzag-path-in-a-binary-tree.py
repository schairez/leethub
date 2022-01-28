# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def longestZigZag(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        
        max_path = 0
        def dfs(node: TreeNode, is_left: bool) -> int:
            nonlocal max_path
            if node is None:
                return 0
            num_left = dfs(node.left, is_left=True)
            num_right = dfs(node.right, is_left=False)
            
            total = max(num_left, num_right)
            max_path = max(max_path, total)
            
            if is_left:
                return 1 + num_right
            return 1 + num_left
        
        dfs(root, is_left=True)
        return max_path