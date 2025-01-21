class Solution(object):
    def filterRestaurants(self, restaurants, veganFriendly, maxPrice, maxDistance):
        """
        :type restaurants: List[List[int]]
        :type veganFriendly: int
        :type maxPrice: int
        :type maxDistance: int
        :rtype: List[int]
        """
        candidates = {}
        for restaurant in restaurants:
            if restaurant[2] >= veganFriendly:
                if restaurant[3] <= maxPrice:
                    if restaurant[4] <= maxDistance:
                        if restaurant[1] in candidates:
                            candidates[restaurant[1]].append(restaurant[0])
                        else:
                            candidates[restaurant[1]] = [restaurant[0]]
        fout = []
        for rating in sorted(candidates.keys(),reverse=True): #Sort in decreasing ratings
            candidates[rating].sort(reverse=True)
            for restaurant in candidates[rating]: #In each rating, sort in increasing id
                fout.append(restaurant)
        return fout
                        
