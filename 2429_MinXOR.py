class Solution(object):
    def minimizeXor(self, num1, num2):
        """
        :type num1: int
        :type num2: int
        :rtype: int
        """
        n1,n2 = num1,num2
        b1 = list()
        setbits1 = 0
        while num1 > 0:
            if num1 % 2 == 1:
                setbits1 += 1
                b1.append(True)
            else:
                b1.append(False)
            num1 //= 2
        b0 = b1[::-1]
            
        #b2 = list()
        setbits2 = 0 #Get number of set bits in num2
        while num2 > 0:
            if num2 % 2 == 1:
                setbits2 += 1
            num2 //= 2
        #b2 = b2[::-1]
        
        #if setbits1 == setbits2:
        #    return n1
        if setbits1 < setbits2:
            diffbits = setbits2 - setbits1
            diff = 0
            i = 0
            omega = 1
            while diffbits > 0:
                if i < len(b1):
                    if not b1[i]:
                        diff += 2**i
                        diffbits -= 1
                else:
                    diff += omega
                    diffbits -= 1
                i += 1
                omega *= 2
            return n1 + diff

        else:
            i = 0
            fout = 0
            while setbits2 > 0:
                if b0[i]:
                    fout += 2**(len(b0)-1-i)
                    setbits2 -= 1
                i += 1
            return fout
