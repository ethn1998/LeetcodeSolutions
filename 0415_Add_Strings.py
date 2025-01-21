class Solution(object):
    def addStrings(self, num1, num2):
        """
        :type num1: str
        :type num2: str
        :rtype: str
        """
        #Add seperately
        n1 = 0
        n2 = 0
        for i in range(len(num1)):
            n1 = 10*n1 + ord(num1[i]) - 48
        for i in range(len(num2)):
            n2 = 10*n2 + ord(num2[i]) - 48
        return str(n1+n2)
        
