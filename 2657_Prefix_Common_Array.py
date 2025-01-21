class Solution(object):
    def findThePrefixCommonArray(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: List[int]
        """
        seenA = set()
        seenB = set()
        fout = []
        for i in range(len(A)):
            seenA.add(A[i])
            seenB.add(B[i])
            fout.append(len(seenA & seenB))
        return fout
