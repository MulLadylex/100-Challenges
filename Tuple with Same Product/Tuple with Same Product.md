# Description
Given an array nums of distinct positive integers, return the number of tuples (a, b, c, d) such that a * b = c * d where a, b, c, and d are elements of nums, and a != b != c != d.

## Test Case
Case1:

    Input: nums = [2,3,4,6]

    Output: 8
    Explanation：(2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3) , (3,4,2,6) , (4,3,2,6) , (3,4,6,2) , (4,3,6,2)
---
Case2:

    Input: nums = [1,2,4,5,10]

    Output: 16
    Explanation：(2,5,1,10) , (2,5,10,1) , (5,2,1,10) , (5,2,10,1) , 
                 (1,10,2,5) , (10,1,2,5) , (1,10,5,2) , (10,1,5,2) , 
                 (2,10,4,5) , (10,2,4,5) , (2,10,5,4) , (10,2,5,4) ,  
                 (4,5,2,10) , (4,5,10,2) , (5,4,2,10) , (5,4,10,2) 
---

## tips
* 1 <= nums.length <= 1000
* 1 <= nums[i] <= 104
* All elements in nums are distinct.

## Solution
### 哈希统计
根据题意：元组(a,b,c,d)满足 a × b = c × d，且满足 a ≠ b ≠ c ≠ d，则该元组根据排列组合共有8种不同的元组：
- (a, b, c, d),(a, b, d, c)；
- (b, a, c, d),(b, a, d, c)；
- (c, d, a, b),(c, d, b, a)；
- (d, c, a, b),(d, c, b, a)。

因为数组nums是由不同整数组成，此时只需在乘积为 a×b 的数对中任意选择一个不同的数对 (c,d) ，一定满足同积元组(a,b,c,d)

因此，我们只需统计数组中所有不同元素的乘积组合，然后计算能组成同积元组的数对的数目，即可得到最终结果


> 复杂度分析  
> 时间复杂度：O(n^2)  
> n为数组长度，求两数之积以及遍历所有的两数之积均为O(n^2)
> 空间复杂度：O(n^2) 
