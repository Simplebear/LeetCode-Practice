# -*- coding: utf-8 -*-

class Solution:
    def numJewelsInStones(self, J, S):
        count = 0
        for a in J:
            for b in S:
                if a == b:
                    count = count +1
        return count
# test code
v = Solution()
print(v.numJewelsInStones("a","aaaaaaaqaaaaaaaaaaaaf"))