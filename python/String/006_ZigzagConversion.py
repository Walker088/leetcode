"""
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows
like this: (you may want to display this pattern in a fixed font for better legibility)
P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:
string convert(string s, int numRows);
 
----------------------------------------------------------------------------------------
## Example 1:
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

## Example 2:
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"

Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I

## Example 3:
Input: s = "A", numRows = 1
Output: "A"
 
----------------------------------------------------------------------------------------
Constraints:

1 <= s.length <= 1000
s consists of English letters (lower-case and upper-case), ',' and '.'.
1 <= numRows <= 1000

"""

import time
import random, string

letters = string.ascii_lowercase + string.ascii_uppercase + string.digits + ',.'

class Solution:
    def convertV1(self, s: str, numRows: int) -> str:
        """
        Problem: This solution builds an array of the columns of the zigzag, then convert it to the string.
        However, if we could construct the rows instead of the columns, we'd be able to remove the last two for-loops.
        """
        if (1 == len(s) or 1 == numRows or len(s) < numRows): return s
        if (2 == numRows): return s[::2] + s[1::2]
        if (3 == numRows): return s[::4] + s[1::2] + s[2::4]

        zigzag = ""
        cacheArr = []

        ptr = 0
        while ptr < len(s):
            # Add the first zigzag column substring to the cache arr
            if (ptr+numRows <= len(s)):
                cacheArr.append(s[ptr:ptr+numRows])
                ptr += numRows
            else:
                cacheArr.append(s[ptr:len(s)].ljust(numRows, "*"))
                break
            # Add the middle part of the zigzag column, padding with *
            for middle in range(numRows - 2):
                if (ptr < len(s)):
                    cacheArr.append(("".rjust(numRows - middle - 2, "*") + s[ptr]).ljust(numRows, "*"))
                    ptr += 1
                else:
                    break
        
        # Build the result
        for row in range(numRows):
            for col in range(len(cacheArr)):
                if ("*" != cacheArr[col][row]): 
                    zigzag += cacheArr[col][row]
        return zigzag

    def convertV2(self, s: str, numRows: int) -> str:
        """
        zigzag_rows: stores substrings of each rows of zigzag result.
        row_ptr: since we're gonna iterate the input string, we use this pointer to decide to which row should we add this character.
        nxt_step: the row_ptr is gonna start from 0, the first row of the zigzag result,
            adding by 1 until reaches the bottom, row_ptr == numRows - 1,
            than minus by 1 until reaches the top, row_ptr == 0.

        inspired by https://leetcode.com/problems/zigzag-conversion/discuss/3404/Python-O(n)-Solution-in-96ms-(99.43)
        """
        if (1 == len(s) or 1 == numRows or len(s) < numRows): return s
        if (2 == numRows): return s[::2] + s[1::2]
        if (3 == numRows): return s[::4] + s[1::2] + s[2::4]

        zigzag_rows = [''] * numRows # used to store the rows of the zigzag output
        row_ptr, nxt_step = 0, 1

        for c in s:
            zigzag_rows[row_ptr] += c
            if row_ptr == 0 : nxt_step = 1
            if row_ptr == numRows - 1: nxt_step = -1
            row_ptr += nxt_step

        return "".join(zigzag_rows)


if __name__ == "__main__":
    def loopPerf(times: int, slt):
        for _ in range(times):
            slt(''.join(random.choice(letters) for _ in range(random.randint(1, 1000))), random.randint(1, 1000))
    iters = 1500
    solution = Solution()
    exp_1 = {"s": "PAYPALISHIRING", "nr": 2}
    exp_2 = {"s": "PAYPALISHIRING", "nr": 3}
    exp_3 = {"s": "PAYPALISHIRING", "nr": 4}
    exp_4 = {"s": "A", "nr": 1}

    s1 = time.time_ns()
    solution.convertV1(exp_1["s"], exp_1["nr"])
    solution.convertV1(exp_2["s"], exp_2["nr"])
    solution.convertV1(exp_3["s"], exp_3["nr"])
    loopPerf(iters, solution.convertV1)
    e1 = time.time_ns()
    print(f"V1 EXEC TIME: {(e1 - s1)/1000000} ms")

    s2 = time.time_ns()
    solution.convertV2(exp_1["s"], exp_1["nr"])
    solution.convertV2(exp_2["s"], exp_2["nr"])
    solution.convertV2(exp_3["s"], exp_3["nr"])
    loopPerf(iters, solution.convertV2)
    e2 = time.time_ns()
    print(f"V3 EXEC TIME: {(e2 - s2)/1000000} ms")
