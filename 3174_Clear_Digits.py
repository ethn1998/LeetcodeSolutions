class Solution(object):
    def clearDigits(self, s):
        """
        :type s: str
        :rtype: str
        NOTE: 
        Mark indices to be deleted, and then loop over string appending unmarked indices at end 
        Time complexity = O(2n) because we loop over the string twice.
        """
        marks = set() #Store marked indices in a set
        i = 0
        j = 0
        while i < len(s):
            char = s[i]
            if char.isnumeric():
                marks.add(i)
                search = True
                j = i-1
                while j >= 0:
                    if not s[j].isnumeric() and j not in marks:
                        marks.add(j)
                        break
                    j -= 1
            i += 1
        fout = ""
        k = 0
        while k < len(s):
            if k not in marks:
                fout += s[k]
            k += 1
        return fout        

