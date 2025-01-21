def mystrtoint(s): #Convert string to integer
    fout = 0
    for char in s:
        if char.isdigit():
            fout = 10*fout + ord(char) -48
        else:
            return -1 #Reference value for not an integer
    return fout

class Solution(object):
    def areNumbersAscending(self, s):
        """
        :type s: str
        :rtype: bool
        """
        
        #s = s.strip() #Clean input
        strlist = s.split() #Get list of tokens
        m = 0 #biggest number seen

        for word in strlist:
            v = mystrtoint(word)
            if v > 0: #If valid integer
                if v <= m:
                    return False
                else:
                    m = v #Update biggest number
        return True
