# Definition for a binary tree node.
class TreeNode(object):
     def __init__(self, x):
         self.val = x
         self.left = None
         self.right = None

class Solution(object):
    def mergeTrees(self, t1, t2):
        """
        :type t1: TreeNode
        :type t2: TreeNode
        :rtype: TreeNode
        """
        if(t1!=None and t2!=None):
            result =  TreeNode(t1.val+t2.val)
            result.left = self.mergeTrees(t1.left,t2.left)
            result.right = self.mergeTrees(t1.right,t2.right)
        if(t1 == None and t2!=None):
            result =  TreeNode(t2.val)
            result.left = self.mergeTrees(None,t2.left)
            result.right = self.mergeTrees(None,t2.right)
        if(t1 != None and t2 ==None):
            result =  TreeNode(t1.val)
            result.left = self.mergeTrees(t1.left,None)
            result.right = self.mergeTrees(t1.right,None)
        if (t1 == None and t2 ==None):
             result = None
        return result
#test
a = Solution()
b = TreeNode(1)
b.left = TreeNode(2)
c = TreeNode(3)
c.left = TreeNode(5)
c.right = TreeNode(8)
d = a.mergeTrees(b,c)
print(d.val)
print(d.left.val)
print(d.right.val)