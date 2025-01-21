class Solution(object):
    def fizzBuzz(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        answer = [] #Initialize output
        i = 1
        while i <= n:
            m = ""
            if i % 3 == 0:
                m += "Fizz"
            if i % 5 == 0:
                m += "Buzz"
            if m == "":
                m = str(i)
            answer.append(m)
            i += 1          
        return answer  

        
