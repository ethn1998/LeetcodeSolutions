class Solution(object):
    def longestPalindrome(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        fout = 0 #Initialize output
        seen = {} #Dictionary for storing seen words
        for w in words:
            if w[::-1] in seen: #If its reverse is already present
                seen[w[::-1]] -= 1 #Cancel out "antiword"
                fout += 4 #Increment output generated using the word and its reverse
                if seen[w[::-1]] == 0:
                    seen.pop(w[::-1]) #Delete zeroes from seen
            elif w in seen:
                seen[w] += 1
            else:
                seen[w] = 1

        for s in seen: #Check if any remaining words are palindromes themselves
            if s == s[::-1]: #If s itself is a palindrome
                fout += 2 #Put random middle palindrome in centre
                break
        return fout
        
