class Solution(object):
    def constructRectangle(self, area):
        """
        :type area: int
        :rtype: List[int]
        """
        w = 1
        while w*w <= area:
            if area % w == 0:
                W = w
                L = area/w
            w += 1
        return [L,W]

        
