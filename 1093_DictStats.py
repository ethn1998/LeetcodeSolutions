class Solution(object):
    def sampleStats(self, count):
        """
        :type count: List[int]
        :rtype: List[float]
        """
        #count is actually a dictionary
        minnotfound = True #Check if minimum has been found
        s = 0.0 #Sum. Used in computing mean.
        n = sum(count) #Sample size, used in calculating mean and median
        evencount = (n % 2 == 0) 
        if evencount:
            m0 = n/2
            m1 = n/2 + 1   
        else:
            m = n//2 + 1
        median = -1
        findm1 = False #Boolean triggering special phase of median calculation
        k = 0 #Position in array
        mode = 0
        biggest = -1
        for i in range(len(count)):
            if count[i] > 0:
                if findm1:
                    median = (i + m0val)/2.0
                    findm1 = False
                if minnotfound:
                    minimum = float(i)
                    minnotfound = False
                maximum = i
                s += i * count[i] #Used to compute mean
                k += count[i]
                if median < 0:
                    if evencount:
                        if k == m0:
                            m0val = i
                            findm1 = True
                        elif k >= m1: #If overshot
                            median = i
                    else:
                        if k >= m:
                            median = i
                if count[i] > biggest: #Used to compute mode
                    biggest = count[i]
                    mode = i
        mean = s/n
        return [minimum,maximum,mean,median,mode]
