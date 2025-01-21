class Solution(object):
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        i = 0
        k = len(nums)

        while i<k:
            if nums[i] == val:
                nums.remove(val)
                k -= 1 #Update length of list
            else:
                i += 1
        return k
