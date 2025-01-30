class Solution(object):
    def detectCapitalUse(self, word):
        """
        :type word: str
        :rtype: bool
        """
        if len(word) < 2:
            return True
        #Use ASCII encoding to check if letter is capital
        if ord(word[0]) <= 90 : #If first is capital
            if ord(word[1]) <= 90: #If second is capital
                i = 2
                while i < len(word):
                    if ord(word[i]) > 90:
                        return False
                    i += 1
            else: #If second is not capital
                i = 2
                while i < len(word):
                    if ord(word[i]) <= 90:
                        return False
                    i += 1
        else:
            i = 1
            while i < len(word):
                if ord(word[i]) <= 90:
                    return False
                i += 1
        return True

        
