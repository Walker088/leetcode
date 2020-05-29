# -*- coding: utf-8 -*-

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def flatten(self, root: TreeNode) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        if not root:
            return None
        # 1
        flatten_left = self.flatten(root.left)
        flatten_right = self.flatten(root.right)

        # 2
        root.left = None
        root.right = flatten_left

        # 3
        tail = root
        while tail.right:
            tail = tail.right
        tail.right = flatten_right

        return root

    def preorder(self, root: TreeNode) -> None:
        if None == root: return
        print(root.val)
        self.preorder(root.left)
        self.preorder(root.right)

t5 = TreeNode(val=4)
t4 = TreeNode(val=3)
t3 = TreeNode(val=5, left=t4, right=t5)
t2 = TreeNode(val=2, left=t3)
t7 = TreeNode(val=6)
t1 = TreeNode(val=1, left=t2, right=t7)

s = Solution()
nt = s.flatten(t1)
s.preorder(nt)