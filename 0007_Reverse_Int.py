import math

class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """

        if x == 0: #Degenerate case
            return 0

        isNegative = False
        m = 0 #Magnitude of input
        if x < 0:
            isNegative = True
            m = -x #Remove negative sign for process
        else:
            m = x
        p = 0 #Do place by place
        mout = 0 #Magnitude of output
        M = math.log10(m)

        while p <= M:
            d = (m//10**p)%10
            mout = mout*10 + d
            p += 1
        if isNegative:
            if mout > 2**31:
                return 0
            else:
                return -mout
        else:
            if mout >= 2**31:
                return 0
            else:
                return mout
