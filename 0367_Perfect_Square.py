class Solution(object):
    def isPerfectSquare(self, num):
        """
        :type num: int
        :rtype: bool
        """
        s = 0
        while s*s <= num:
            if s*s == num:
                return True
            else:
                s += 1
        return False
        
