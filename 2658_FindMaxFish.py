class Solution(object):
    def findMaxFish(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        #NOTE: This is a generalization of Leetcode problem 695 Max Area of Island.
        seen = set() #Track seen water cells
        fout = 0
        for r0, row in enumerate(grid):
            for c0, val in enumerate(row):
                if val > 0 and (r0,c0) not in seen:
                    fish = 0 #Start fishing here
                    queue = [(r0,c0)] #Queue nearby lands that form island
                    seen.add((r0,c0)) 
                    while len(queue) > 0:
                        r,c = queue.pop() #Remove and add to area
                        fish += grid[r][c] #Catch all fish in cell
                        for r1, c1 in ((r-1,c), (r+1,c),(r,c-1),(r,c+1)):
                            if (r1 >= 0 and r1 < len(grid) 
                            and c1 >= 0 and c1 < len(grid[0])
                            and grid[r1][c1] > 0 and (r1,c1) not in seen): #Check for adjecent unvisited water cells
                                queue.append((r1,c1))
                                seen.add((r1,c1))
                    if fish > fout:
                        fout = fish
        return fout        
