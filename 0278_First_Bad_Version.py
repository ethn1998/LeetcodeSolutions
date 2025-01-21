# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):

class Solution(object):
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """

        if n == 1: #Degenerate case
            return 1

        d = {} #Store seen values in a dictionary
        i = 1 #Lower bound
        j = n #Upper bound

        #notfound = True
        while True:
            k = (i+j)//2
            if d.get(k) is None:
                kbad = isBadVersion(k)
            else:
                kbad = d[k]
        
            if kbad:
                if d.get(k-1) is None:
                    k0bad = isBadVersion(k-1)
                else:
                    k0bad = d[k-1]
                if k0bad:
                    d.update({k-1:True, k:True}) #Update seen values
                    j = k #Check smaller values
                else:
                    return k
            else:
                if d.get(k+1) is None:
                    k0bad = isBadVersion(k+1)
                else:
                    k0bad = d[k+1]
                if k0bad:
                    return k+1
                else:
                    d.update({k:False, k+1:False})
                    i = k #Check larger values
