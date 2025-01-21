class Solution(object):
    def canPlaceFlowers(self, flowerbed, n):
        """
        :type flowerbed: List[int]
        :type n: int
        :rtype: bool
        """
        if len(flowerbed) == 0: #Degenerate philosophical case
            return False
        if len(flowerbed) == 1:
            if flowerbed[0] == 0:
                return True
            else:
                return n == 0

        if flowerbed[0] == 0:
            if flowerbed[1] == 0:
                n -= 1
                flowerbed[0] = 1
        i = 1
        while i < len(flowerbed)-1 and n>0:
            if flowerbed[i] == 0:
                if flowerbed[i-1] == 0:
                    if flowerbed[i+1] == 0:
                        n -= 1
                        flowerbed[i] = 1
            i += 1
        if flowerbed[-1] == 0:
            if flowerbed[-2] == 0:
                n -=1
                flowerbed[-1] = 1
        return n <= 0
