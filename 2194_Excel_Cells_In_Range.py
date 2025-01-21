class Solution(object):
    def cellsInRange(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        fout = []
        corners = s.split(":") #Split at : to get limits
        #Parse top-left cell
        topstr = ""
        leftstr = ""
        for char in corners[0]:
            if char.isdigit(): #Numbers go here
                topstr += char
            else: #Letters go here
                leftstr = char
        leftstr = leftstr[::-1] #Reverse
        left = 0
        v = 1
        for char in leftstr:
            left += (ord(char) - 64)*v
            v *= 26
        top = int(topstr)
        
        #Parse bottom-right cell
        botstr = ""
        rightstr = ""
        for char in corners[1]:
            if char.isdigit(): #Numbers go here
                botstr += char
            else: #Letters go here
                rightstr = char
        rightstr = rightstr[::-1] #Reverse
        right = 0
        v = 1
        for char in rightstr:
            right += (ord(char) - 64)*v
            v *= 26
        bot = int(botstr)
    
        for i in range(left,right+1): #Wrap characters around numbers
            word = ""
            n = i
            while n > 0:
                if n % 26 == 0:
                    word = "Z" + word
                    n = n//26 -1
                else:
                    char = chr(n % 26 + 64)
                    word = char + word
                    n //= 26

            for j in range(top, bot+1):
                fout.append(word + "{:d}".format(j))
        return fout
        
