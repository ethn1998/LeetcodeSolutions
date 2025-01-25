class Solution(object):
    def lemonadeChange(self, bills):
        """
        :type bills: List[int]
        :rtype: bool
        """
        x = 0 #Five bills
        y = 0 #Ten bills
        for i in range(len(bills)):
            if bills[i] == 5:
                x += 1
            elif bills[i] == 10:
                x -= 1
                y += 1
            else:
                if y > 0:
                    x -= 1
                    y -= 1
                else:
                    x -= 3
            if x < 0 or y < 0:
                return False
        return True
