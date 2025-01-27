class Solution(object):
    def groupAnagrams(self, strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """
        mydict = {}
        for word in strs:
            sortedword = ''.join(sorted(word))
            if sortedword in mydict:
                mydict[sortedword].append(word)
            else:
                mydict[sortedword] = [word]
        fout = []
        for key in mydict:
            fout.append(mydict[key])
        return fout
