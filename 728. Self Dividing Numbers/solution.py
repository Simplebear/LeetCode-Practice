class Solution(object):
    def selfDividingNumbers(self, left, right):
        """
        :type left: int
        :type right: int
        :rtype: List[int]
        """
        judge_filter = lambda num: '0' not in str(num) and all([num % int(element) == 0 for element in str(num)])
        x = filter(judge_filter, range(left, right + 1))
        return list(x)
test = Solution()
for i in test.selfDividingNumbers(1,33):
    print(i)