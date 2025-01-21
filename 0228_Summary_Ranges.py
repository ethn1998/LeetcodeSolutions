class Solution(object):
    def summaryRanges(self, nums):
        """
        :type nums: List[int]
        :rtype: List[str]
        """
        fout = []
        if nums == []:
            return fout
        i,j = 0,0
        while j < len(nums)-1:
            if nums[j+1] != nums[j] + 1:
                if i == j:
                    fout.append(str(nums[i]))
                else:
                    fout.append("{:d}->{:d}".format(nums[i],nums[j]))
                i = j+1
            j += 1
        #Final entry
        if i == j:
            fout.append(str(nums[-1]))
        else:
            fout.append("{:d}->{:d}".format(nums[i],nums[-1]))
        return fout

