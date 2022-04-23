
# 317. Shortest Distance from All Buildings
# time: O(m^2*n^2)
# space: O(m*n)

from collections import deque

class Solution:
    def shortestDistance(self, grid: List[List[int]]) -> int:
        def bfs(grid: List[List[int]], distance_grid: List[List[int]],
                reach_grid: List[List[int]], house: List[int], num_reach: int):

            num_x = len(grid)
            num_y = len(grid[0])
            dirs = [1, 0, -1, 0, 1]
            queue = deque([])
            queue.append(house)

            dist = 0
            while len(queue) != 0:
                curr_len = len(queue)
                for _ in range(curr_len):
                    node = queue.popleft()
                    x, y = node
                    distance_grid[x][y] += dist
                    for i in range(1, 5):
                        new_x = x + dirs[i-1]
                        new_y = y + dirs[i]
                        if new_x < 0 or new_x >= num_x or new_y < 0 or new_y >= num_y:
                            continue
                        if grid[new_x][new_y] == 1 or grid[new_x][new_y] == 2:
                            continue

                        if reach_grid[new_x][new_y] == num_reach - 1:
                            queue.append([new_x, new_y])
                            reach_grid[new_x][new_y] = num_reach
                dist += 1

        num_x, num_y = len(grid), len(grid[0])
        house_q = deque([])
        for x in range(num_x):
            for y in range(num_y):
                if grid[x][y] == 1:
                    house_q.append([x, y])
                    
        total = len(house_q) 
        distance_grid = [[0 for _ in range(num_y)] for _ in range(num_x)]
        reach_grid = [[0 for _ in range(num_y)] for _ in range(num_x)]
        
        min_dist  = float('inf')
        
        num_reach = 1
        while len(house_q) != 0:
            house = house_q.popleft()
            bfs(grid, distance_grid, reach_grid, house, num_reach)
            num_reach += 1
            
        print(distance_grid) 
        print(reach_grid)
        for x in range(num_x):
            for y in range(num_y):
                if reach_grid[x][y] == total and distance_grid[x][y] < min_dist:
                    min_dist = distance_grid[x][y]
        
        if min_dist == float('inf'):
            return -1
        
        return min_dist
        
            