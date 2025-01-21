idict = { #Dictionary converting string to integer value
    "0":0,
    "1":1,
    "2":2,
    "3":3,
    "4":4,
    "5":5,
    "6":6,
    "7":7,
    "8":8,
    "9":9
    }

class Solution(object):
    def sortSentence(self, s):
        """
        :type s: str
        :rtype: str
        """
        d = {} #Sigh, hashmap for index: word pair
        fout = ""
        tokens = s.split(" ")
        for t in tokens:
            d.update({idict[t[-1]]: t[:-1]})

        for j in list(range(len(tokens))):
            fout += "{} ".format(d[j+1])
        fout = fout.rstrip()
        return fout
        
