import math

# Declare dictionaries FIRST
p = { #Implement powers for easy future scaling.
    0: "", #We don't say units
    3: "Thousand",
    6: "Million",
    9: "Billion",
    12: "Trillion",
    15: "Quadrillion",
    18: "Quintillion"}

#Not that we need to go this far

d = {
    0: "", #We don't say zero
    1: "One",
    2: "Two",
    3: "Three",
    4: "Four",
    5: "Five",
    6: "Six",
    7: "Seven",
    8: "Eight",
    9: "Nine",
} #Generic dictionary
dt = {
    0: "", #We don't say zero
    2: "Twenty",
    3: "Thirty",
    4: "Forty",
    5: "Fifty",
    6: "Sixty",
    7: "Seventy",
    8: "Eighty",
    9: "Ninety",
} #Dictionary of ten-tys (default case)
d10 = {
    0: "Ten",
    1: "Eleven",
    2: "Twelve",
    3: "Thirteen",
    4: "Fourteen",
    5: "Fifteen",
    6: "Sixteen",
    7: "Seventeen",
    8: "Eighteen",
    9: "Nineteen"
} #Dictionary for 11-20

def hundredsWords(x): #Words between 1 and 999, called multiple times in actual function.
    fout = "" #Initialize empty string
    #Get digits
    h = x // 100 % 10
    t = x // 10 % 10
    u = x % 10

    if h != 0:
        fout += "{} Hundred ".format(d[h]) #Can't do sprintf formatting on Python

    if t == 0:
        fout += d[u]
    elif t == 1:
        fout += d10[u]
    else:
        fout += "{} {}".format(dt[t],d[u])
    fout = fout.strip() #Clean output

    return fout

class Solution(object):
    def numberToWords(self, num):
        """
        :type num: int
        :rtype: str
        """
        if num == 0: #Special case
            return "Zero"

        Omega = math.log10(num)
        fout = "" #Initialize output
        i = 0 #Track power
        while i <= Omega:
            h = (num // 10**i) % 1000
            if h != 0:
                fout = "{} {} ".format(hundredsWords(h),p[i]) + fout
            i += 3
        
        #Clean output
        fout = fout.strip()
        return fout
