class Solution(object):
    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        fout = "" #Initialize output
        # Do place by place
        #Thousands
        th = num//1000
        for _ in list(range(th)):
            fout = fout + "M" #Append M
        #Hundreds
        h = (num%1000)//100
        if h >= 5:
            if h == 9:
                fout = fout + "CM"
            else:
                fout = fout + "D"
                for _ in list(range(h - 5)):
                    fout = fout + "C"
        else:
            if h == 4:
                fout = fout + "CD"
            else:
                for _ in list(range(h)):
                    fout = fout + "C"
        #Tens
        t = (num%100)//10
        if t >= 5:
            if t == 9:
                fout = fout + "XC"
            else:
                fout = fout + "L"
                for _ in list(range(t - 5)):
                    fout = fout + "X"
        else:
            if t == 4:
                fout = fout + "XL"
            else:
                for _ in list(range(t)):
                    fout = fout + "X"
        #Units
        u = num%10
        if u >= 5:
            if u == 9:
                fout = fout + "IX"
            else:
                fout = fout + "V"
                for _ in list(range(u - 5)):
                    fout = fout + "I"
        else:
            if u == 4:
                fout = fout + "IV"
            else:
                for _ in list(range(u)):
                    fout = fout + "I"
        return fout




        
