# Description
给你一个整数数组` nums `，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按*任意顺序*返回答案。

你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。
# Test Case
Case1:
    Input: nums = {1,2,1,3,2,5}
    Output: {3, 5} or {5, 3}
===
Case2:
    Input: nums = {-1,0}
    Output: {-1,0}
===
Case3:
    Input: nums = {0,1}
    Output: {0,1}
# tips
* `2 <= nums.length <= 3 * 10^4`
* `-2^31 <= nums[i] <= 2^31 - 1`
* 除两个只出现一次的整数外，nums 中的其他数字都出现两次

# 