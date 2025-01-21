class Solution(object):
    def findErrorNums(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        Omega = set(range(1,len(nums)+1))
        for n in nums:
            if n in Omega:
                Omega.remove(n)
            else: #If seen already
                x0 = n #This is the repeated element
        x1 = Omega.pop()

        return([x0,x1])
        
