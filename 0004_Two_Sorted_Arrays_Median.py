class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """

        i = 0
        j = 0
        m = len(nums1)
        n = len(nums2)
        nums1out = (m == 0)
        nums2out = (n == 0)
        
        if (m+n)%2 == 1: #Total length is odd
            k = (m+n)//2
            while i+j <= k: #Count until we reach index of median or one list runs out
                if not nums1out and not nums2out: #Compare
                    if nums1[i]<nums2[j]:
                        v = nums1[i]
                        i += 1
                        if i == m:
                            nums1out = True
                    else:
                        v = nums2[j]
                        j += 1
                        if j == n:
                            nums2out = True
                elif nums1out:
                    v = nums2[j]
                    j += 1
                else:
                    v = nums1[i]
                    i += 1
            return v
        else:
            k = (m+n)//2 - 1
            while i+j <= k: #Count until we reach index of median or one list runs out
                if not nums1out and not nums2out: #Compare
                    if nums1[i] < nums2[j]:
                        v = nums1[i]
                        i += 1
                        if i == m:
                            nums1out = True
                    else:
                        v = nums2[j]
                        j += 1
                        if j == n:
                            nums2out = True
                elif nums1out:
                    v = nums2[j]
                    j += 1
                else:
                    v = nums1[i]
                    i += 1
            if nums1out:
                return (v + nums2[j]) / 2.0
            elif nums2out:
                return (v + nums1[i]) / 2.0
            else:
                if nums2[j] < nums1[i]:
                    return (v + nums2[j]) / 2.0
                else:
                    return (v + nums1[i]) / 2.0


