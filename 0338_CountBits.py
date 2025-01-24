class Solution(object):
    def countBits(self, n):
        """
        :type n: int
        :rtype: List[int]
        """
        fout = []
        for i in range(n+1):
            b = 0
            k = i
            while k > 0:
                if k % 2 == 1:
                    b += 1
                k //= 2
            fout.append(b)
        return fout
        
