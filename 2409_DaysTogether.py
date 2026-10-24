def dateToInt(datestr):
    mdays = [0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    mtoi = 0
    dtoi = 0
    fout = 0
    for i in [0,1]:
        mtoi = 10 * mtoi
        mtoi += int(datestr[i])
    for i in [3,4]:
        dtoi = 10 * dtoi
        dtoi += int(datestr[i])        
    for i in list(range(mtoi)):
        fout += mdays[i]
    #print(mtoi,dtoi,fout)
    return fout + dtoi

class Solution(object):
    def countDaysTogether(self, arriveAlice, leaveAlice, arriveBob, leaveBob):
        """
        :type arriveAlice: str
        :type leaveAlice: str
        :type arriveBob: str
        :type leaveBob: str
        :rtype: int
        """
        a0 = dateToInt(arriveAlice)
        a1 = dateToInt(leaveAlice)
        b0 = dateToInt(arriveBob)
        b1 = dateToInt(leaveBob)

        #print(a0,a1,b0,b1)
        if a1 < b0 or b1 < a0:
            return 0

        return min(a1,b1)-max(a0,b0)+1
        
