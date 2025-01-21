d = {
    "2": ["a","b","c"],
    "3": ["d","e","f"],
    "4": ["g","h","i"],
    "5": ["j","k","l"],
    "6": ["m","n","o"],
    "7": ["p","q","r","s"],
    "8": ["t","u","v"],
    "9": ["w","x","y","z"]
}

class Solution(object):
    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        if digits == "":
            return []
        fout = [""]
        for n in digits:
            dchars = d[n]
            perms = [] #Store earlier permutations here
            for char in dchars:
                for word in fout:
                    perms.append(word+char)
            fout = perms
        return fout
