"""
get() and put() operations time: O(1)

space: O(n)
"""

from typing import Optional 
        
class Node(object):
    def __init__(self, key:int, val:int):
        self.key = key
        self.val = val
        self.next: Optional[Node] = None 
        self.prev: Optional[Node] = None



class LRUCache:

    def __init__(self, capacity: int):
        self.capacity = capacity
        self._head = Node(-1,-1) #dummyhead
        self._tail = Node(-1,-1) #dummytail
        self._head.next = self._tail 
        self._tail.prev = self._head
        self._cache: dict[int, Node] = dict()
        self._size = 0 
    
    def get(self, key: int) -> int:
        if key not in self._cache:
            return -1 
        # cache hit
        node: Node = self._cache[key]
        self._remove_node(node) 
        # key is most-recently-used; move to head 
        self._add_new_head_node(node)
        return node.val
  

    def put(self, key: int, value: int) -> None:
        if not isinstance(key, int) or not isinstance(value, int):
            return None
        
        if key in self._cache: #cache hit
            self._remove_node(self._cache[key])
        else: # cache miss
            if self.capacity == self._size:
                # evict LRU 
                self._remove_lru_node()
        newNode = Node(key=key, val=value)
        self._add_new_head_node(newNode)
        self._cache[newNode.key] = newNode
        
    def _move_to_head(self, node: Node) -> None:
        node.next.prev = node.prev
        node.prev.next = node.next
        node.next = self._head.next 
        self._head.next = node
        node.prev = self._head
        
    def _remove_node(self, node: Node) -> None:
        k = node.key
        node.prev.next = node.next
        node.next.prev = node.prev 
        self._size -=1 
        del self._cache[k]

    def _remove_lru_node(self) -> None:
        node: Node = self._tail.prev
        prev_node = node.prev 
        prev_node.next = self._tail 
        self._tail.prev = prev_node 
        del self._cache[node.key]
        self._size -=1 

    
    def _add_new_head_node(self, node: Node) -> None:
        node.next = self._head.next 
        self._head.next = node 
        node.prev = self._head 
        node.next.prev = node 
        self._cache[node.key] = self._head.next
        self._size += 1

    def printNodes(self) -> None:
        n = self._head
        print("from printnodes")
        while n is not None:
            print(n.val, "-->",  end= " ")
            n = n.next
        print()

