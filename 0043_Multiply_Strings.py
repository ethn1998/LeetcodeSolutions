class Solution(object):
    def multiply(self, num1, num2):
        """
        :type num1: str
        :type num2: str
        :rtype: str
        """
        #Implement archaic multiplication algorithm
        fout = 0
        #How "directly" is directly? Can we not convert string into integer at all or what?
        inum1 = 0
        for i in num1:
            inum1 = 10*inum1 + ord(i)-48 #Convert first number into an integer object.

        for j in num2:
            p = 0
            for m in range(0,ord(j)-48):
                p += inum1
            fout = 10*fout + p
        
        return str(fout)

