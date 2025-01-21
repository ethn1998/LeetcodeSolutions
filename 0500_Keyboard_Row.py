class Solution(object):
    def findWords(self, words):
        """
        :type words: List[str]
        :rtype: List[str]
        """
        fout = []
        r1 = set(list("qwertyuiop"))
        r2 = set(list("asdfghjkl"))
        r3 = set(list("zxcvbnm"))

        for word in words:
            lword = word.lower() #Make everything lowercase
            char0 = lword[0]
            if char0 in r1:
                i = 1
                test = True
                while i<len(word) and test:
                    if lword[i] not in r1:
                        test = False
                    i += 1
                if test:
                    fout.append(word)            

            if char0 in r2:
                i = 1
                test = True
                while i<len(word) and test:
                    if lword[i] not in r2:
                        test = False
                    i += 1
                if test:
                    fout.append(word)            

            if char0 in r3:
                i = 1
                test = True
                while i<len(word) and test:
                    if lword[i] not in r3:
                        test = False
                    i += 1
                if test:
                    fout.append(word)            

        return fout
