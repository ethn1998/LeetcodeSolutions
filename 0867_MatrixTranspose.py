class Solution(object):
    def transpose(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[List[int]]
        """
        nrows = len(matrix)
        ncols = len(matrix[0])
        fout = []

        for i in range(ncols): #Columns become rows
            row = []
            for j in range(nrows):
                row.append(matrix[j][i])
            fout.append(row)
        
        return fout
        
