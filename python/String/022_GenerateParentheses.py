"""
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

## Example 1:
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

## Example 2:
Input: n = 1
Output: ["()"]

----------------------------------------------------------------------------------------------------------
Constraints:

1 <= n <= 8
"""

from typing import List
import time
import random

class Solution:
    def generateParenthesisV1(self, n: int) -> List[str]:
        if 1 == n: return ["()"]
        total = "(" * n + ")" * n

        def is_parenthes_valid(prth: str) -> bool:
            num = 0
            for p in prth[::-1]:
                if "(" == p: num -= 1
                if ")" == p: num += 1
                if num < 0: return False
            return True

        def recursive_append(crr_prth_opts: str):
            n = len(crr_prth_opts)
            if (n == 0): return [""]
            acc = []
            for prth_idx in range(n):
                if crr_prth_opts[prth_idx] in crr_prth_opts[:prth_idx] : continue

                for children in recursive_append(crr_prth_opts[:prth_idx] + crr_prth_opts[prth_idx+1:]):
                    candidate = crr_prth_opts[prth_idx] + children
                    if(is_parenthes_valid(candidate)): acc.append(candidate)
            return acc
        return recursive_append(total)

    def generateParenthesisV2(self, n: int) -> List[str]:
        # ref: https://leetcode.com/problems/generate-parentheses/discuss/10369/Clean-Python-DP-Solution
        dp = [[] for i in range(n + 1)]
        dp[0].append('')
        for i in range(n + 1):
            for j in range(i):
                dp[i] += ['(' + x + ')' + y for x in dp[j] for y in dp[i - j - 1]]
        return dp[n]
    
    def generateParenthesisV3(self, n: int) -> List[str]:
        # stack, ref: https://leetcode.com/problems/generate-parentheses/discuss/10283/Python-simple-stack-solution-without-recursion
        res = []
        s = [("(", 1, 0)]
        while s:
            x, l, r = s.pop()
            if l - r < 0 or l > n or r > n:
                continue
            if l == n and r == n:
                res.append(x)
            s.append((x+"(", l+1, r))
            s.append((x+")", l, r+1))
        return res



if __name__ == "__main__":
    def loopPerf(times: int, slt):
        for _ in range(times): slt(random.randint(1, 8))
    iters = 8
    solution = Solution()
    exp_1 = 7
    exp_2 = 8
    exp_3 = 6

    s1 = time.time_ns()
    loopPerf(iters, solution.generateParenthesisV1)
    e1 = time.time_ns()
    print(f"V1 EXEC TIME: {(e1 - s1)/1000000} ms")

    s2 = time.time_ns()
    loopPerf(iters, solution.generateParenthesisV2)
    e2 = time.time_ns()
    print(f"V2 EXEC TIME: {(e2 - s2)/1000000} ms")
