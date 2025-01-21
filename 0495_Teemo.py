class Solution(object):
    def findPoisonedDuration(self, timeSeries, duration):
        """
        :type timeSeries: List[int]
        :type duration: int
        :rtype: int
        """
        #if duration == 0:
        #    return 0 #Never poisoned
        #if duration == 1:
        #    return len(timeSeries) #Poisoned as long as hit
        fout = 0
        T1 = 0 #Time when poison wears off
        for t in timeSeries:
            if t < T1:
                fout += duration-(T1-t)
            else:
                fout += duration
            T1 = t+duration

        return fout
        
