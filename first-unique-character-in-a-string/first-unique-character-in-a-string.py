"""
time: O(n)
space: O(n)
"""

class Node:
    def __init__(self, key:str, idx_loc:int, prev=None, next=None):
        self.key = key
        self.idx_loc = idx_loc
        self.prev = prev 
        self.next = next 
        

class Llist:
    def __init__(self):
        self.cur_len = 0 
        #self.char =  char
        self._head = Node(-1,-1)
        self._tail = Node(-1,-1)
        self._head.next = self._tail
        self._tail.prev = self._head
        
    def __len__(self) ->int:
        return self.cur_len
    
    @property
    def head_node(self) -> Node:
        return self._head.next
        
    def add_to_tail(self, node: Node):
        node.prev = self._tail.prev 
        self._tail.prev.next = node 
        node.next = self._tail 
        self._tail.prev = node
        self.cur_len +=1 

class Solution:
    def firstUniqChar(self, s: str) -> int:
        if len(s) == 1:
            return 0 #idx 0
        tbl: list[Llist] = [Llist() for i in range(26)]
            
        for idx, ch in enumerate(s):
            code = ord(ch) - ord('a')
            n = Node(key=ch, idx_loc=idx)
            tbl[code].add_to_tail(n)
        
        res_opts = []
        for llist in tbl:
            if len(llist) == 0 or len(llist) > 1:
                continue 
            #only one node in llist
            res_opts.append(llist.head_node)
                                 
        if len(res_opts) == 0:
            return -1                           
        print(res_opts) 
        res = res_opts[0]
        # chr(i + ord('a')
        for opt in res_opts[1:]:
            if opt.idx_loc < res.idx_loc:
                res = opt 
                
        return res.idx_loc 