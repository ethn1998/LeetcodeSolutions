class Solution(object):
    def reverseStr(self, s, k):
        i = 0 #Starting point
        while i < len(s):
            l = i
            r = min(i+k,len(s))
            head = s[:l]
            body = s[l:r]
            tail = s[r:]
            body = body[::-1]
            s = head+body+tail
            i += 2*k
        return s


        
