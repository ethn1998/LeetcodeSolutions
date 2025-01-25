class Solution(object):
    def maxDistance(self, colors):
        """
        :type colors: List[int]
        :rtype: int
        """
        j = len(colors)-1
        while j > 0:
            if colors[0] != colors[j]:
                leftMax = j
                break
            j -= 1
        i = 0
        while i < len(colors):
            if colors[i] != colors[-1]:
                rightMax = len(colors)-i-1
                break
            i += 1
        return max(leftMax,rightMax)
        
