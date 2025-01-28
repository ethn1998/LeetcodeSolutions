class Solution(object):
    def maxAreaOfIsland(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        seen = set() #Track seen lands
        fout = 0
        for r0, row in enumerate(grid):
            for c0, val in enumerate(row):
                if val != 0 and (r0,c0) not in seen:
                    area = 0
                    queue = [(r0,c0)] #Queue nearby lands that form island
                    seen.add((r0,c0)) 
                    while len(queue) > 0:
                        r,c = queue.pop() #Remove and add to area
                        area += 1
                        for r1, c1 in ((r-1,c), (r+1,c),(r,c-1),(r,c+1)):
                            if (r1 >= 0 and r1 < len(grid) 
                            and c1 >= 0 and c1 < len(grid[0])
                            and grid[r1][c1] != 0 and (r1,c1) not in seen): #Check grid for adjecent unseen land squares
                                queue.append((r1,c1))
                                seen.add((r1,c1))
                    if area > fout:
                        fout = area
        return fout
