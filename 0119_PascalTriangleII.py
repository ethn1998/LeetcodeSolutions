class Solution(object):
    def getRow(self, rowIndex):
        """
        :type rowIndex: int
        :rtype: List[int]
        """
        i = 0
        prevrow = [1]
        while i < rowIndex:
            newrow = [1]
            for j in range(len(prevrow)-1):
                newrow.append(prevrow[j] + prevrow[j+1])
            newrow.append(1)
            prevrow = newrow
            i += 1

        return prevrow
