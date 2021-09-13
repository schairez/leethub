"""
time: O(1)
space: O(n)
"""
class Node(object):
    def __init__(self, t_val:int, prev=None, next=None) :
        self.t_val = t_val
        self.prev = prev
        self.next = next
        
class Deque(object):
    def __init__(self):
        self._head = Node(-1) #dummyHead
        self._tail = Node(-1) #dummyTail
        self._head.next = self._tail
        self._tail.prev = self._head
        self.size = 0
        
    def __len__(self):
        return self.size
    
    def append_to_head(self, node:Node):
        """
        new nodes added to head; older nodes toward right (back)
        """
        node.next = self._head.next 
        node.prev = self._head 
        self._head.next = node 
        node.next.prev = node 
        self.size +=1 
        
    def remove_node(self, node:Node):
        print("removing node", node)
        node.prev.next = node.next 
        node.next.prev = node.prev 
        self.size -= 1 
        
    def remove_expired_nodes(self, time:int):
        valid_start = time - 3000
        valid_end = time 
        cur_node = self._tail.prev 
        #print("looking @", cur_node.t_val)
        while not valid_start <= cur_node.t_val <= valid_end:
            if cur_node.t_val == -1: #reaches dummyNode
                break 
            next_node = cur_node.prev
            self.remove_node(cur_node)
            cur_node = next_node
    def _print_nodes(self):
        start = self._head 
        print(start._t)


class RecentCounter:

    def __init__(self):
        self.deque = Deque()
        
    def ping(self, t: int) -> int:
        n = Node(t)
        self.deque.append_to_head(n)
        self.deque.remove_expired_nodes(t)
        return len(self.deque)
        
        # reqs in range[t -3000, t]


# Your RecentCounter object will be instantiated and called as such:
# obj = RecentCounter()
# param_1 = obj.ping(t)