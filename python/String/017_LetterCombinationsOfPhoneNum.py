"""
Given a string containing digits from 2-9 inclusive, 
return all possible letter combinations that the number could represent. 
Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. 
Note that 1 does not map to any letters.

https://assets.leetcode.com/uploads/2022/03/15/1200px-telephone-keypad2svg.png

-------------------------------------------------------------------------------------
## Example 1:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

## Example 2:
Input: digits = ""
Output: []

## Example 3:
Input: digits = "2"
Output: ["a","b","c"]

-------------------------------------------------------------------------------------
## Constraints:

0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].

"""
from functools import reduce
import time
import random
from typing import List

class Solution:
    def letterCombinationsV1(self, digits: str) -> List[str]:
        # Brute-Force
        if not digits: return []
        table = {
            "2": ["a", "b", "c"],
            "3": ["d", "e", "f"],
            "4": ["g", "h", "i"],
            "5": ["j", "k", "l"],
            "6": ["m", "n", "o"],
            "7": ["p", "q", "r", "s"],
            "8": ["t", "u", "v"],
            "9": ["w", "x", "y", "z"],
        }
        mapped = [table[s] for s in digits if s in ["2", "3", "4", "5", "6", "7", "8", "9"]]
        total = len(mapped)
        if total == 0: return []
        if total == 1: return mapped[0]
        
        resp = []
        if total == 2:
            for l_1 in mapped[0]:
                for l_2 in mapped[1]:
                    resp.append(l_1 + l_2)
            return resp
        if total == 3:
            for l_1 in mapped[0]:
                for l_2 in mapped[1]:
                    for l_3 in mapped[2]:
                        resp.append(l_1 + l_2 + l_3)
            return resp
        if total == 4:
            for l_1 in mapped[0]:
                for l_2 in mapped[1]:
                    for l_3 in mapped[2]:
                        for l_4 in mapped[3]:
                            resp.append(l_1 + l_2 + l_3 + l_4)
            return resp
    def letterCombinationsV2(self, digits: str) -> List[str]:
        # DFS, REF: https://leetcode.com/problems/letter-combinations-of-a-phone-number/discuss/8067/Python-easy-to-understand-backtracking-solution.
        
        if not digits: return []
        mapper = {"2":"abc", '3':"def", '4':"ghi", '5':"jkl", '6':"mno", '7':"pqrs", '8':"tuv", '9':"wxyz"}
        def dfs(nums: str, str_builder: str, combinations: list):
            # 2. Once reach the end of the 
            if not nums: 
                combinations.append(str_builder)
                return 
            # 1. For each mapped chars, add it to its corresponded str_builder till the last digit
            for c in mapper[nums[0]]:
                dfs(nums[1:], str_builder + c, combinations)
        combinations = []
        dfs(digits, "", combinations)
        return combinations
    def letterCombinationsV3(self, digits: str) -> List[str]:
        # REF: https://leetcode.com/problems/letter-combinations-of-a-phone-number/discuss/8070/One-line-python-solution
        if '' == digits: return []
        kvmaps = {
            '2': 'abc',
            '3': 'def',
            '4': 'ghi',
            '5': 'jkl',
            '6': 'mno',
            '7': 'pqrs',
            '8': 'tuv',
            '9': 'wxyz'
        }
        return reduce(lambda acc, digit: [x + y for x in acc for y in kvmaps[digit]], digits, [''])


if __name__ == "__main__":
    def loopPerf(times: int, slt):
        for _ in range(times):
            ''.join(random.choice("23456789") for _ in range(random.randint(1, 4)))
    iters = 100000
    solution = Solution()
    exp_1 = "23"
    exp_2 = "368"
    exp_3 = "258"

    s1 = time.time_ns()
    solution.letterCombinationsV1(exp_1)
    solution.letterCombinationsV1(exp_2)
    solution.letterCombinationsV1(exp_3)
    loopPerf(iters, solution.letterCombinationsV1)
    e1 = time.time_ns()
    print(f"V1 EXEC TIME: {(e1 - s1)/1000000} ms")

    s2 = time.time_ns()
    solution.letterCombinationsV2(exp_1)
    solution.letterCombinationsV2(exp_2)
    solution.letterCombinationsV2(exp_3)
    loopPerf(iters, solution.letterCombinationsV2)
    e2 = time.time_ns()
    print(f"V2 EXEC TIME: {(e2 - s2)/1000000} ms")
