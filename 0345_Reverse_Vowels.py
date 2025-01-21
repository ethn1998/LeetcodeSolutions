vowels = {"a","e","i","o","u","A","E","I","O","U"}

class Solution(object):
    def reverseVowels(self, s):
        """
        :type s: str
        :rtype: str
        """
        fout = ""
        j = len(s)
        for char in s:
            if char in vowels:
                vowelNotFound = True
                while vowelNotFound:
                    j -= 1
                    if s[j] in vowels:
                        fout += s[j]
                        vowelNotFound = False
            else:
                fout += char

        return fout
