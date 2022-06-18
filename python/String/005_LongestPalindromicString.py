"""
Given a string s, return the longest palindromic substring in s.
--------------------------------------------
Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
--------------------------------------------
Example 2:

Input: s = "cbbd"
Output: "bb"
--------------------------------------------
Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.

"""
import time
import random, string

letters = string.ascii_lowercase + string.digits

class Solution:
    def longestPalindromeV1(self, s: str) -> str:
        """
        Brute-Force (Time Limit Exceeded on LeetCode lol)
        """
        max_palindrome = ""
        sub_strings = [s[i: j] for i in range(len(s)) for j in range(i + 1, len(s) + 1)]
        for sub in sub_strings:
            if sub == sub[::-1] and len(sub) > len(max_palindrome): max_palindrome = sub
        return max_palindrome

    def longestPalindromeV2(self, s: str) -> str:
        """
        Dynamic Programming
        ref: https://leetcode.com/problems/longest-palindromic-substring/discuss/900639/Python-Solution-%3A-with-detailed-explanation-%3A-using-DP
        """
        longest_palindrom = ''
        dp = [[0]*len(s) for _ in range(len(s))]
        #filling out the diagonal by 1
        for i in range(len(s)):
            dp[i][i] = True
            longest_palindrom = s[i]
			
        # filling the dp table
        for i in range(len(s)-1,-1,-1):
				# j starts from the i location : to only work on the upper side of the diagonal 
            for j in range(i+1,len(s)):  
                if s[i] == s[j]:  #if the chars mathces
                    # if len slicied sub_string is just one letter if the characters are equal, we can say they are palindomr dp[i][j] =True 
                    #if the slicied sub_string is longer than 1, then we should check if the inner string is also palindrom (check dp[i+1][j-1] is True)
                    if j-i ==1 or dp[i+1][j-1] is True:
                        dp[i][j] = True
                        # we also need to keep track of the maximum palindrom sequence 
                        if len(longest_palindrom) < len(s[i:j+1]):
                            longest_palindrom = s[i:j+1]
                
        return longest_palindrom

    def longestPalindromeV3(self, s: str) -> str:
        """
        Expand Around Center
        ref: https://leetcode.com/problems/longest-palindromic-substring/discuss/2954/Python-easy-to-understand-solution-with-comments-(from-middle-to-two-ends).
        """
        def helper(_s,l,r):
            """
            1. l, r in [s.index]
            2. Will always enter the first iteration when it is executed either by helper(s, i, i),
               or helper(s, i, i+1) so l will be deducted by that loop, to recover it we need to return _s[l+1:r]
            """
            while 0 <= l and r < len(_s) and _s[l] == _s[r]:
                l -= 1; r += 1
            return _s[l+1:r]
        max_palindrome = ""
        for i in range(len(s)):
            max_palindrome = max(
                helper(s, i, i), # odd case, like "aba"
                helper(s, i, i+1), # even case like "abba"
                max_palindrome, key=len
            )
        return max_palindrome
    
    def longestPalindromeV4(self, s: str) -> str:
        """
        A better version of Expand Around Center
        Basic thought is simple. when you increase s by 1 character, you could only increase maxPalindromeLen by 1 or 2, and that new maxPalindrome includes this new character
        ref: https://leetcode.com/problems/longest-palindromic-substring/discuss/2925/Python-O(n2)-method-with-some-optimization-88ms.
        """
        if len(s)==0: return 0
        maxLen=1
        start=0
        # scan from beginning to the end, adding one character at a time, keeping track of maxPalindromeLen
        # Note: since in the worst case, s[start:start+maxLen] = s[0:0+1], we do not need i=0
        for i in range(1, len(s)):
            # Check if the substrings ending with this new character, with length P+2, [i-maxLen-1:i+1], is palindromes
            if i-maxLen >=1 and s[i-maxLen-1:i+1]==s[i-maxLen-1:i+1][::-1]:
                start=i-maxLen-1
                maxLen+=2
                continue
            # Check if the substrings ending with this new character, with length P+1, [i-maxLen:i+1], is palindromes
            if i-maxLen >=0 and s[i-maxLen:i+1]==s[i-maxLen:i+1][::-1]:
                start=i-maxLen
                maxLen+=1
        return s[start:start+maxLen]

    def longestPalindromeV5(self, s: str) -> str:
        """
        Manacher algorithm in Python O(n)
        ref: https://leetcode.com/problems/longest-palindromic-substring/discuss/3337/Manacher-algorithm-in-Python-O(n)
             https://en.wikipedia.org/wiki/Longest_palindromic_substring
        """
        # Transform S into T.
        # For example, S = "abba", T = "^#a#b#b#a#$".
        # ^ and $ signs are sentinels appended to each end to avoid bounds checking
        T = '#'.join('^{}$'.format(s))
        n = len(T)
        P = [0] * n
        C = R = 0
        for i in range (1, n-1):
            P[i] = (R > i) and min(R - i, P[2*C - i]) # equals to i' = C - (i-C)
            # Attempt to expand palindrome centered at i
            while T[i + 1 + P[i]] == T[i - 1 - P[i]]:
                P[i] += 1
    
            # If palindrome centered at i expand past R,
            # adjust center based on expanded palindrome.
            if i + P[i] > R:
                C, R = i, i + P[i]
    
        # Find the maximum element in P.
        maxLen, centerIndex = max((n, i) for i, n in enumerate(P))
        return s[(centerIndex  - maxLen)//2: (centerIndex  + maxLen)//2]


if __name__ == "__main__":
    def loopPerf(times: int, slt):
        for _ in range(times):
            slt(''.join(random.choice(letters) for _ in range(random.randint(0, 1000))))
    iters = 1500
    solution = Solution()
    exp_1 = "babad"
    exp_2 = "a"
    exp_3 = "pwwpew"

    s2 = time.time_ns()
    solution.longestPalindromeV2(exp_1)
    solution.longestPalindromeV2(exp_2)
    solution.longestPalindromeV2(exp_3)
    loopPerf(iters, solution.longestPalindromeV2)
    e2 = time.time_ns()
    print(f"V2 [DP] EXEC TIME: {(e2 - s2)/1000000} ms")

    s3 = time.time_ns()
    solution.longestPalindromeV3(exp_1)
    solution.longestPalindromeV3(exp_2)
    solution.longestPalindromeV3(exp_3)
    loopPerf(iters, solution.longestPalindromeV3)
    e3 = time.time_ns()
    print(f"V3 [Expand Around Center] EXEC TIME: {(e3 - s3)/1000000} ms")

    s4 = time.time_ns()
    solution.longestPalindromeV4(exp_1)
    solution.longestPalindromeV4(exp_2)
    solution.longestPalindromeV4(exp_3)
    loopPerf(iters, solution.longestPalindromeV4)
    e4 = time.time_ns()
    print(f"V4 [Opt. Expand Around Center] EXEC TIME: {(e4 - s4)/1000000} ms")

    s5 = time.time_ns()
    solution.longestPalindromeV5(exp_1)
    solution.longestPalindromeV5(exp_2)
    solution.longestPalindromeV5(exp_3)
    loopPerf(iters, solution.longestPalindromeV5)
    e5 = time.time_ns()
    print(f"V5 [Manacher algorithm] EXEC TIME: {(e5 - s5)/1000000} ms")

    """
    s1 = time.time_ns()
    solution.longestPalindromeV1(exp_1)
    solution.longestPalindromeV1(exp_2)
    solution.longestPalindromeV1(exp_3)
    #loopPerf(iters, solution.longestPalindromeV1)
    e1 = time.time_ns()
    print(f"V1 EXEC TIME: {(e1 - s1)/1000000} ms")
    """

