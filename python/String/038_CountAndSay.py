"""
The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

countAndSay(1) = "1"
countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.
To determine how you "say" a digit string, split it into the minimal number of substrings such that each substring contains exactly one unique digit. Then for each substring, say the number of digits, then say the digit. Finally, concatenate every said digit.

For example, the saying and conversion for digit string "3322251":
Given a positive integer n, return the nth term of the count-and-say sequence.

## Example 1:
Input: n = 1
Output: "1"
Explanation: This is the base case.

## Example 2:
Input: n = 4
Output: "1211"
Explanation:
countAndSay(1) = "1"
countAndSay(2) = say "1" = one 1 = "11"
countAndSay(3) = say "11" = two 1's = "21"
countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"

----------------------------------------------------------------------------------------------------------
Constraints:
- 1 <= n <= 30

----------------------------------------------------------------------------------------------------------
## Personal Explaination 
Given:
countAndSay(1) = "1"
countAndSay(2) = Say(countAndSay(1))
countAndSay(3) = Say(countAndSay(2))
...
Hence, one of the solutions can be:
```countAndSayV1
def countAndSayV1(n: int) -> str:
    # base case
    result = "1"

    # from the base case, convert the number n-1 times
    # e.g, 
    # - if n=2, we need to execute the say_function once
    # - if n=3, we need to execute the say_function twice, and so on.
    for _ in range(n-1):
        result = say(result)
    return result
```
Time complexity: O((n-1)*m), where n is the input, m is the string length of countAndSayV1(n-1)
if you wanna cheat, you can increase the base case from 1 to say, 10, 
then you can get the following:
(Note that the first loop can be reduced to n-10, since we start from the 10th loop.)
```countAndSayV2
def countAndSayV2(n: int) -> str:
    if (1 == n): return "1"
        if (2 == n): return "11"
        if (3 == n): return "21"
        if (4 == n): return "1211"
        if (5 == n): return "111221"
        if (6 == n): return "312211"
        if (7 == n): return "13112221"
        if (8 == n): return "1113213211"
        if (9 == n): return "31131211131221"
        if (10 == n): return "13211311123113112211"

        result = "13211311123113112211"
        for _ in range(n-10):
            result = say(result)
        return result
```
"""

import time

def say(query: str):
    if (query=="0"): return "1"

    result = ""
    base = query[0]
    base_repeated_times = 1
    for i in range(len(query)):
        if i == len(query)-1:
            return result + f"{base_repeated_times}{base}"

        if (query[i+1] == base):
            base_repeated_times += 1
        else:
            result += f"{base_repeated_times}{base}"
            base = query[i+1]
            base_repeated_times = 1

class Solution:
    def countAndSayV1(self, n: int) -> str:
        result = "1"
        for _ in range(n-1):
            result = say(result)
        return result

    def countAndSayV2(self, n: int) -> str:
        if (1 == n): return "1"
        if (2 == n): return "11"
        if (3 == n): return "21"
        if (4 == n): return "1211"
        if (5 == n): return "111221"
        if (6 == n): return "312211"
        if (7 == n): return "13112221"
        if (8 == n): return "1113213211"
        if (9 == n): return "31131211131221"
        if (10 == n): return "13211311123113112211"

        result = "13211311123113112211"
        for _ in range(n-10):
            result = say(result)
        return result


if __name__ == "__main__":
    def loopPerf(times: int, slt):
        for _ in range(times):
            for n in range(1, 15):
                s = slt(n)
                #print(f"countAndSay({n})=['{s}']")
    solution = Solution()
    times = 100

    s1 = time.time_ns()
    loopPerf(1, solution.countAndSayV1)
    e1 = time.time_ns()
    print(f"V1 EXEC TIME: {(e1 - s1)/1000000} ms")

    s2 = time.time_ns()
    loopPerf(1, solution.countAndSayV2)
    e2 = time.time_ns()
    print(f"V2 EXEC TIME: {(e2 - s2)/1000000} ms")
