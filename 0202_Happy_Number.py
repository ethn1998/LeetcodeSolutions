class Solution(object):
    def isHappy(self, n):
        """
        :type n: int
        :rtype: bool
        """
        if n == 0: #Degenerate cases
            return False
        if n == 1:
            return True
        seen = {n}
        while True:
            p = math.log10(n)
            s = 0
            i = 0
            while(i<=p):
                d = (n//10**i) % 10
                s += d*d
                i += 1
            if s in seen: #Assuming that it will always loop endlessly
                return False
            elif s == 1:
                return True
            seen.add(s)
            n = s #Repeat with new value
