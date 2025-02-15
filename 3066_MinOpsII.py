import heapq

class Solution(object):
    def minOperations(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        fout = 0
        heapq.heapify(nums)
        while nums[0] < k:
            x = heapq.heappop(nums)
            y = heapq.heappop(nums)
            z = 2*min(x,y) + max(x,y)
            heapq.heappush(nums,z)
            # Update number of operations
            fout += 1
        return fout
        
