"""
time: O(m*n)
space: O(m*n)
"""


class Solution:
    def numBusesToDestination(self, routes: List[List[int]], source: int, target: int) -> int:
        graph = defaultdict(set)
        for i, route in enumerate(routes):
            for bus_stop in route:
                graph[bus_stop].add(i)
        num_routes = 0
        queue = collections.deque([source])
        visited_stops = set()
        visited_routes = set()
        visited_stops.add(source)
        while len(queue) != 0:
            for _ in range(len(queue)):
                bus_stop = queue.popleft()
                if bus_stop == target:
                    return num_routes
                for route_id in graph[bus_stop]:
                    if route_id not in visited_routes:
                        for next_stop in routes[route_id]:
                            if next_stop not in visited_stops:
                                queue.append(next_stop)
                                visited_stops.add(next_stop)
                        visited_routes.add(route_id)
            num_routes+=1
                
        return -1