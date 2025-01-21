class Solution(object):
    def pancakeSort(self, arr):
        """
        :type arr: List[int]
        :rtype: List[int]
        """
        """
        NOTE:
        Take advantage of fact that arr is a permutation of 1,...,arr.length.
        """
        fout = []
        target = len(arr)
        while target > 1:
            if arr[target-1] == target: #If target already in correct place
                target -= 1 #Go to next target
            elif arr[0] == target:
                #Perform pancake flip target-1 to bring target to correct index 
                l = 0
                r = target - 1
                while l < r:
                    arr[l], arr[r] = arr[r], arr[l]
                    l += 1
                    r -= 1
                fout.append(target)
            else:
                l = 0
                # Look for target and perform pancake flip r to bring target to index 0.
                r = arr.index(target)
                fout.append(r+1)
                while l < r:
                    arr[l], arr[r] = arr[r], arr[l]
                    l += 1
                    r -= 1
        return fout
