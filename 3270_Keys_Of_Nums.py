class Solution(object):
    def generateKey(self, num1, num2, num3):
        """
        :type num1: int
        :type num2: int
        :type num3: int
        :rtype: int
        """
        s1 = str(num1).zfill(4)
        s2 = str(num2).zfill(4)
        s3 = str(num3).zfill(4)
        i = 0
        fout = 0
        while i < 4:
            x = 9
            if x > int(s1[i]):
                x = int(s1[i])
            if x > int(s2[i]):
                x = int(s2[i])
            if x > int(s3[i]):
                x = int(s3[i])
            fout += 9*fout + x 
            i += 1
        return fout
