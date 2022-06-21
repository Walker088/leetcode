"""
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II.
The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII.
Instead, the number four is written as IV. Because the one is before the five we subtract it making four.
The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9. 
X can be placed before L (50) and C (100) to make 40 and 90. 
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral.

---------------------------------------------------------------------------------------------------------------

## Example 1:

Input: num = 3
Output: "III"
Explanation: 3 is represented as 3 ones.

## Example 2:

Input: num = 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.

## Example 3:

Input: num = 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

---------------------------------------------------------------------------------------------------------------

Constraints:

1 <= num <= 3999
"""

import time
import random

class Solution:
    def intToRomanV1(self, num: int) -> str:
        # Ref: https://leetcode.com/problems/integer-to-roman/discuss/292605/Python-Solution-44ms
        numDict = {
            1000 : 'M',
            900 : 'CM',
            500 : 'D',
            400 : 'CD',
            100 : 'C',
            90 : 'XC',
            50 : 'L',
            40 : 'XL',
            10 : 'X',
            9 : 'IX',
            5 : 'V',
            4 : 'IV',
            1 : 'I',
        }
        romanNum = ''
        while num > 0:
            for i, r in numDict.items():
                if i <= num:
                    romanNum += r
                    num -= i
                    break
        return romanNum

    def intToRomanV2(self, num: int) -> str:
        # Ref: https://leetcode.com/problems/integer-to-roman/discuss/6304/Python-simple-solution.
        roman = {1000:"M", 900:"CM", 500:"D", 400:"CD", 100:"C", 90:"XC", 50:"L", 40:"XL", 10:"X", 9:"IX", 5:"V", 4:"IV", 1:"I"}
        ans = []
        for k, v in roman.items(): 
            ans.append(num//k * v)
            num %= k
        return "".join(ans)

    def intToRomanV3(self, num: int) -> str:
        # Ref: https://leetcode.com/problems/integer-to-roman/discuss/1293176/C%2B%2B-easy-solution-O(1)
        ones = ["","I","II","III","IV","V","VI","VII","VIII","IX"]
        tens = ["","X","XX","XXX","XL","L","LX","LXX","LXXX","XC"]
        hrns = ["","C","CC","CCC","CD","D","DC","DCC","DCCC","CM"]
        ths = ["","M","MM","MMM"]
        return ths[num//1000] + hrns[num%1000//100] + tens[num%100//10] + ones[num%10]

if __name__ == "__main__":
    def loopPerf(times: int, slt):
        for _ in range(times): slt(random.randint(1, 3999))
    iters = 100000
    solution = Solution()
    exp_1 = 1800
    exp_2 = 901
    exp_3 = 10

    s1 = time.time_ns()
    solution.intToRomanV1(exp_1)
    solution.intToRomanV1(exp_2)
    solution.intToRomanV1(exp_3)
    loopPerf(iters, solution.intToRomanV1)
    e1 = time.time_ns()
    print(f"V1 EXEC TIME: {(e1 - s1)/1000000} ms")

    s2 = time.time_ns()
    solution.intToRomanV2(exp_1)
    solution.intToRomanV2(exp_2)
    solution.intToRomanV2(exp_3)
    loopPerf(iters, solution.intToRomanV2)
    e2 = time.time_ns()
    print(f"V2 EXEC TIME: {(e2 - s2)/1000000} ms")

    s3 = time.time_ns()
    solution.intToRomanV3(exp_1)
    solution.intToRomanV3(exp_2)
    solution.intToRomanV3(exp_3)
    loopPerf(iters, solution.intToRomanV3)
    e3 = time.time_ns()
    print(f"V2 EXEC TIME: {(e3 - s3)/1000000} ms")
