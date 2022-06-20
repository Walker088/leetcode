"""
Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).

The algorithm for myAtoi(string s) is as follows:

Read in and ignore any leading whitespace.
Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range. Specifically, integers less than -231 should be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
Return the integer as the final result.
Note:

Only the space character ' ' is considered a whitespace character.
Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.

-----------------------------------------------------------------------------------------------------------
## Example 1
Input: s = "42"
Output: 42
Explanation: The underlined characters are what is read in, the caret is the current reader position.
Step 1: "42" (no characters read because there is no leading whitespace)
         ^
Step 2: "42" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "42" ("42" is read in)
           ^
The parsed integer is 42.
Since 42 is in the range [-231, 231 - 1], the final result is 42.

## Example 2
Input: s = "   -42"
Output: -42
Explanation:
Step 1: "   -42" (leading whitespace is read and ignored)
            ^
Step 2: "   -42" ('-' is read, so the result should be negative)
             ^
Step 3: "   -42" ("42" is read in)
               ^
The parsed integer is -42.
Since -42 is in the range [-231, 231 - 1], the final result is -42.

## Example 3
Input: s = "4193 with words"
Output: 4193
Explanation:
Step 1: "4193 with words" (no characters read because there is no leading whitespace)
         ^
Step 2: "4193 with words" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "4193 with words" ("4193" is read in; reading stops because the next character is a non-digit)
             ^
The parsed integer is 4193.
Since 4193 is in the range [-231, 231 - 1], the final result is 4193.

-----------------------------------------------------------------------------------------------------------
## Constraints:

0 <= s.length <= 200
s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.

"""

import re
import time
import random, string

class Solution:
    def myAtoiV1(self, s: str) -> int:
        if (len(s) == 0): return 0
        
        min, max = -2147483648, 2147483647
        s = s.lstrip()
        if not s: return 0

        alphabet = string.ascii_lowercase + string.ascii_uppercase + " ."
        result = ""
        sign = ""
        for c in s:
            if c not in string.digits and result != "":
                break
            elif c not in string.digits and result == "":
                if c in alphabet: return 0
                if c in "+-" and sign != "": return 0
                if c in "+-" and sign == "": sign = c
            else:
                result += c
        if (result == ""): return 0
        convered = int(sign + result)
        if (convered < min): return min
        if (convered > max): return max
        return convered

    def myAtoiV2(self, s: str) -> int:
        if (len(s) == 0): return 0
        s = s.lstrip() # trim the leading white space first
        if not s: return 0

        MIN_NUM, MAX_NUM = -2147483648, 2147483647
        sign = 1
        index = 0
        num = 0
        
        if s[0] == '-':
            sign = -1
            index += 1
        elif s[0] == '+':
            index += 1
        
        while index < len(s) and s[index].isdigit():
            curr_digit = ord(s[index]) - ord('0')
            # here we do the check before adding current digit
            if num > MAX_NUM // 10 or (num == MAX_NUM // 10 and curr_digit > 7):
                return MAX_NUM if sign == 1 else MIN_NUM
            num = num * 10 + curr_digit
            index += 1

        num = sign * num
        return num

    def myAtoiV3(self, s: str) -> int:
        # Simply use regex, ref: https://leetcode.com/problems/string-to-integer-atoi/discuss/4653/Python-solution-based-on-RegEx
        if (len(s) == 0): return 0
        s = s.lstrip()
        if not s: return 0

        s = re.findall('^[+\-]?\d+', s)
        try:
            res = int(''.join(s))
            MAX, MIN = 2147483647, -2147483648
            if res > MAX: return MAX
            if res < MIN: return MIN
            return res
        except: 
            return 0

letters = string.ascii_lowercase + string.ascii_uppercase + string.digits + '.+- '

if __name__ == "__main__":
    def loopPerf(times: int, slt):
        for _ in range(times):
            leading_spaces = ''.join(' ' for _ in range(random.randint(1, 99)))
            sign = ''.join(random.choice('+- ') for _ in range(1))
            digits = ''.join(random.choice(string.digits) for _ in range(random.randint(1, 10))) # 2^31 - 1 = 2,147,483,648
            used_lenth = len(leading_spaces) + len(sign) + len(digits)
            slt(
                leading_spaces + sign + digits + ''.join(random.choice(letters) for _ in range(random.randint(1, 200-used_lenth)))
            )
    iters = 10000
    solution = Solution()
    exp_1 = "words and 987"
    exp_2 = "+-12"
    exp_3 = "00000-42a1234"

    s1 = time.time_ns()
    solution.myAtoiV1(exp_1)
    solution.myAtoiV1(exp_2)
    solution.myAtoiV1(exp_3)
    loopPerf(iters, solution.myAtoiV1)
    e1 = time.time_ns()
    print(f"V1 EXEC TIME: {(e1 - s1)/1000000} ms")

    s2 = time.time_ns()
    solution.myAtoiV2(exp_1)
    solution.myAtoiV2(exp_2)
    solution.myAtoiV2(exp_3)
    loopPerf(iters, solution.myAtoiV2)
    e2 = time.time_ns()
    print(f"V1 EXEC TIME: {(e2 - s2)/1000000} ms")

    s3 = time.time_ns()
    solution.myAtoiV3(exp_1)
    solution.myAtoiV3(exp_2)
    solution.myAtoiV3(exp_3)
    loopPerf(iters, solution.myAtoiV3)
    e3 = time.time_ns()
    print(f"V1 EXEC TIME: {(e3 - s3)/1000000} ms")
