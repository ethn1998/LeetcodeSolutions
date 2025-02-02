class Solution(object):
    """
    STRATEGY 1:
    detect 'b', and then detect 'a'
    More robust, but could be slower in some cases
    """
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
    """
    STRATEGY 2:
    two pointers: 
    i: count 'a's from start and terminate at 'b'
    j: count 'b's from end and terminate at 'a'
    return i > j (first b appears after final a)
    PROBLEM: Breaks if all entries are 'a' or 'b'
    """

    def checkStringij(self, s):
        """
        :type s: str
        :rtype: bool
        """
        b0 = True #Searching for first b
        a1 = True #Searching for last a
        i = 0
        j = len(s)-1
        while (b0 or a1) and (i<len(s)) and (j >= 0) :
            if b0:
                if s[i] == 'b':
                    b0 = False
                else:     
                    i += 1    
            if a1:
                if s[j] == 'a':
                    a1 = False
                else:
                    j -= 1
        return ((b0 or a1) or (j < i))

        
