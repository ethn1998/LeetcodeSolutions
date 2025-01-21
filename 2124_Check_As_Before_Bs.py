class Solution(object):
    def checkString(self, s):
        """
        :type s: str
        :rtype: bool
        """
        seenB = False
        for char in s:
            if char == "b":
                seenB = True
            if seenB:
                if char == "a":
                    return False
        return True
        
