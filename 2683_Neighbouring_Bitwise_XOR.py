class Solution(object):
    def doesValidArrayExist(self, derived):
        """
        :type derived: List[int]
        :rtype: bool
        """
        bi = 0
        for i in range(len(derived)-1):
            bi ^= derived[i] #This generates b(n-1)
        return bi ^ 0 == derived[-1]
