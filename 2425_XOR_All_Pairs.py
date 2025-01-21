class Solution(object):
    def xorAllNums(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: int
        """
        d = {} #Dictionary storing counts of powers of 2.
        if len(nums2)%2 != 0:
            for a in nums1:
                i = 0
                while a > 0:
                    if a%2 == 1: 
                        if i in d.keys():
                            d[i] += 1
                        else:
                            d[i] = 1
                    i += 1
                    a //= 2
        if len(nums1)%2 != 0:
            for b in nums2:
                i = 0
                while b > 0:
                    if b%2 == 1:
                        if i in d.keys():
                            d[i] += 1
                        else:
                            d[i] = 1
                    i += 1
                    b //= 2
        fout = 0
        twop = 1
        for key in d:
            if d[key] % 2 != 0:
                fout += twop
            twop *= 2
        return fout
