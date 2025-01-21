# The rand7() API is already defined for you.
# def rand7():
# @return a random integer in the range 1 to 7

def coin():
    n = rand7()
    if n == 7 :
        return coin()
    else:
        return n%2 

def rand5():
    n = rand7()
    if n >5:
        return rand5()
    else:
        return n

class Solution(object):
    def rand10(self):
        """
        :rtype: int
        """
        u = coin()
        v = rand5()
        return 5*u+v

        
