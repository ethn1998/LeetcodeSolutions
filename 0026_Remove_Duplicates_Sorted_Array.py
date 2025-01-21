class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        k = 0
        m = {} #Must I really use a dict for this?
        M = -101 #Biggest number seen so far
        for v in nums:
            if v > M: #If we see a new value
                m.update({k:v}) #Store new value and corresponding index
                M = v #Update biggest number seen
                k += 1
        for i in m.keys():
            nums[i] = m[i] #Update nums
        return k
        
