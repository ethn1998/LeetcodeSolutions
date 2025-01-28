class Solution(object):
    def numIslands(self, grid):
        """
        :type grid: List[List[str]]
        :rtype: int
        """
        """
        NOTE: Try to adapt DFS idea from Leetcode problem 695 Max Area of Island.
        NOTE: Cell entries are strings rather than ints or floats to troll the programmer. 
        """
        seen = set() #Track seen lands
        fout = 0
        for r0, row in enumerate(grid):
            for c0, val in enumerate(row):
                if val != "0" and (r0,c0) not in seen:
                    fout += 1
                    queue = [(r0,c0)] #Queue nearby lands that form island
                    seen.add((r0,c0)) 
                    while len(queue) > 0:
                        r,c = queue.pop() #Remove and add to island
                        for r1, c1 in ((r-1,c), (r+1,c),(r,c-1),(r,c+1)):
                            if (r1 >= 0 and r1 < len(grid) 
                            and c1 >= 0 and c1 < len(grid[0])
                            and grid[r1][c1] != "0" and (r1,c1) not in seen): #Check grid for adjecent unseen land cells
                                queue.append((r1,c1))
                                seen.add((r1,c1))
        return fout
