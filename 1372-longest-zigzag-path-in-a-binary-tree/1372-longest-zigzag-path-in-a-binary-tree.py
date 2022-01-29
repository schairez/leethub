# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right



#time: O(n); numNodes
#space: O(treeHeight); Omega(logn) if height balanced; O(n) worst case if skewed

class Solution:
    def longestZigZag(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        
        max_path = 0
        def dfs(node: TreeNode, is_left: bool) -> int:
            """
            postorder dfs approach LRD
            """
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
   


"""
the dir param is used to designate whether we're on the left or the right child node in a binary tree. 
We're calculating the local_max_path and potentially updating the abs_max_path calculations prior to the return of our
dfs() fn call.
Keep in mind the zigzal length = numNodes in path minus 1. This amounts to the number of edges along this particular path.
if we're on a left child, we return 1 + the number of path returned from the right descendant of this left child. it's important
to look at postorder here, b/c we're processing our root last, and returning the call of our dfs() for root last in our processing order. 
but most importantly, before we return, we're potentially updating max_path prior to our return. Observe the following example.

note node labeled 'A' is our root. I ignored most nil child nodes for clarity.
                        
                  max(l,r)                    
        A   total=max(0,4) =4; max_path=max(3,4) =4; return 1+numRChild = 5 but max path is no longer updated once this call is popped off our stack
          \       ^
           C  total=max(3,0); max_path=max(2,3)=3; return 1 + numLChild -> 4 up the chain
          /        ^
        E      total=max(0,2)=2; max_path=max(1,2)=2; return 1 + numRChild -> 3 
         \          ^
          Y    total= max(1,0)=1; max_path=max(0,1); return 1 + numLChild -> 2
          /         ^
         T      total =0; max=max(0, 0) =0 ; return 1 + numRChild -> 1
           \        ^
            nil   return 0


our res, abs_max_path, is equal to 4


"""