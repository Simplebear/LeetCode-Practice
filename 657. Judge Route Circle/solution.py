
class Solution(object):
    def judgeCircle(self, moves):
        """
        :type moves: str
        :rtype: bool
        """
        originalpoint = [0,0]
        for i in moves:
            if i == "U":
                originalpoint[0] = originalpoint[0] + 1
            if (i == "D"):
                originalpoint[0] = originalpoint[0] - 1
            if (i == "L"):
                originalpoint[1] = originalpoint[1] + 1
            if (i == "R"):
                originalpoint[1] = originalpoint[1] - 1
        result = (originalpoint[0]== originalpoint[1] == 0) 
        return result         
test = Solution()
a = test.judgeCircle("UUUDDD")
print(a)