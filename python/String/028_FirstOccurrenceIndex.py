"""
Given two strings needle and haystack, 
return the index of the first occurrence of needle in haystack, 
or -1 if needle is not part of haystack.

## Example 1:
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

## Example 2:
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.

----------------------------------------------------------------------------------------------------------
Constraints:
- 1 <= haystack.length, needle.length <= 10^4
- haystack and needle consist of only lowercase English characters.
"""

import time
import random, string
letters = string.ascii_lowercase

class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        haystack_len = len(haystack)
        needle_len = len(needle)

        if needle_len > haystack_len: return -1
        if needle_len == haystack_len and needle == haystack: return 0
        if needle_len == haystack_len and needle != haystack: return -1


        for i in range(0, haystack_len - needle_len + 1):
            if (haystack[i:i + needle_len] == needle): return i
        return -1


if __name__ == "__main__":
    def loopPerf(times: int, slt):
        for _ in range(times):
            haystack = ''.join(random.choice(letters) for _ in range(random.randint(1, 10**4)))
            needle = ''.join(random.choice(letters) for _ in range(random.randint(1, 10**4)))
            s = slt(haystack, needle)
            #print(f"haystack: {haystack}, needle: {needle}, idx: {s}\n")

    iters = 8
    solution = Solution()

    haystack_1 = "sadbutsad"
    needle_1 = "sad"
    print(solution.strStr(haystack_1, needle_1))

    haystack_2 = "abc"
    needle_2 = "c"
    print(solution.strStr(haystack_2, needle_2))

    s1 = time.time_ns()
    loopPerf(iters, solution.strStr)
    e1 = time.time_ns()
    print(f"V1 EXEC TIME: {(e1 - s1)/1000000} ms")

