
# Definition for a binary tree node.
class TreeNode(object):
     def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
# solution
class Solution(object):
    def constructMaximumBinaryTree(self, nums):
        if nums:
            max_vaule = max(nums)
            max_vaule_index = nums.index(max_vaule)
            result = TreeNode(max_vaule)
            result.left = self.constructMaximumBinaryTree(nums[:max_vaule_index])
            result.right = self.constructMaximumBinaryTree(nums[max_vaule_index+1:])
            return result
#test
a = Solution()
print(a.constructMaximumBinaryTree([1,23,4,5,3]).val)
   