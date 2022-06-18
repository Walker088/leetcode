"""
Given a string s, find the length of the longest substring without repeating characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
"""

import time
import random, string

letters = string.ascii_lowercase + string.digits

class Solution:
    def lengthOfLongestSubstringV1(self, s: str) -> int:
        n = len(s)
        if n == 0: return 0

        max_length = 1
        for round in range(n):
            round_lst = []
            for char in s[round:n]:
                if not char in round_lst: round_lst.append(char)
                else: break
            max_length = max(max_length, len(round_lst))
        return max_length

    def lengthOfLongestSubstringV2(self, s: str) -> int:
        max_length, start, end = 0, 0, 0
        while(end < len(s)):
            char_len = len(s[start:end+1])
            diff_char_len = len(set(s[start:end+1]))

            if ( diff_char_len == char_len ):
                max_length = max(max_length, char_len)
                end += 1
            else: start += 1
        return max_length
    
    def lengthOfLongestSubstringV3(self, s: str) -> int:
        if len(s) == 0: return 0
        start = maxLength = 0
        usedChar = {}
        for i in range(len(s)):
            if s[i] in usedChar and start <= usedChar[s[i]]: start = usedChar[s[i]] + 1
            else: maxLength = max(maxLength, i - start + 1)
            usedChar[s[i]] = i
        return maxLength

if __name__ == "__main__":
    def loopPerf(times: int, slt):
        for _ in range(times):
            slt(''.join(random.choice(letters) for _ in range(random.randint(0, 5*104))))
    iters = 1400
    solution = Solution()
    exp_1 = "abcabcbb"
    exp_2 = "bbbbb"
    exp_3 = "pwwpew"

    s1 = time.time_ns()
    solution.lengthOfLongestSubstringV1(exp_1)
    solution.lengthOfLongestSubstringV1(exp_2)
    solution.lengthOfLongestSubstringV1(exp_3)
    loopPerf(iters, solution.lengthOfLongestSubstringV1)
    e1 = time.time_ns()
    print(f"V1 EXEC TIME: {(e1 - s1)/1000000} ms")

    s2 = time.time_ns()
    solution.lengthOfLongestSubstringV2(exp_1)
    solution.lengthOfLongestSubstringV2(exp_2)
    solution.lengthOfLongestSubstringV2(exp_3)
    loopPerf(iters, solution.lengthOfLongestSubstringV2)
    e2 = time.time_ns()
    print(f"V2 EXEC TIME: {(e2 - s2)/1000000} ms")

    s3 = time.time_ns()
    solution.lengthOfLongestSubstringV3(exp_1)
    solution.lengthOfLongestSubstringV3(exp_2)
    solution.lengthOfLongestSubstringV3(exp_3)
    loopPerf(iters, solution.lengthOfLongestSubstringV3)
    e3 = time.time_ns()
    print(f"V3 EXEC TIME: {(e3 - s3)/1000000} ms")
