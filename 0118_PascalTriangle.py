class Solution(object):
    def generate(self, numRows):
        """
        :type numRows: int
        :rtype: List[List[int]]
        """
        fout = [[1]]
        i = 1
        #NOTE: ith row has i elements
        while i < numRows:
            newrow = [1]
            for j in range(len(fout[-1])-1):
                newrow.append(fout[-1][j] + fout[-1][j+1])
            newrow.append(1)
            fout.append(newrow)
            row = newrow
            i += 1

        return fout
        
