#import math

class Solution(object):
    def selfDividingNumbers(self, left, right):
        """
        :type left: int
        :type right: int
        :rtype: List[int]
        """
        fout = []
        for num in range(left,right+1):
            isSelfDividing = True
            tenp = 1
            while tenp < num and isSelfDividing: #Innocent until proven guilty
                d = (num//tenp)%10
                if d != 0: #Sanity check
                    if num % d != 0:
                        isSelfDividing = False
                else:
                    isSelfDividing = False
                tenp *= 10
            if isSelfDividing:
                fout.append(num)
        return fout


        
