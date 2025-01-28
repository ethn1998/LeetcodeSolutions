class Solution(object):
    def rotate(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: None Do not return anything, modify matrix in-place instead.
        """
        """
        NOTE: Problem constraints: Square matrices only
        NOTE: Rotate 90 degrees = Transpose and reverse.
        """
        #Step 1: Transpose
        n = len(matrix) #Square matrices
        i = 0
        while i < n:
            j = i
            while j < n:
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j] 
                j += 1
            i += 1

        #Step 2: Reverse
        i = 0
        while i < n:
            j = 0
            k = n-1
            while j < k:
                matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
                j += 1
                k -= 1
            i += 1
