# Description
Given an integer array nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once. You can return the answer in any order.

You must write an algorithm that runs in linear runtime complexity and uses only constant extra space.

# Test Case
Case1:  

    Input: nums = [1,2,1,3,2,5]

    Output: [3, 5] or [5, 3]
---
Case2:

    Input: nums = [-1,0]

    Output: [-1,0]
---
Case3:

    Input: nums = [0,1]

    Output: [0,1]
# tips
* `2 <= nums.length <= 3 * 10^4`
* `-2^31 <= nums[i] <= 2^31 - 1`
* 除两个只出现一次的整数外，nums 中的其他数字都出现两次

# 