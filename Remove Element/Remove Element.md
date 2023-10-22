# Description
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.

Return k.
## Test Case
Case1:

    Input: nums = [3,2,2,3], val = 3

    Output: 2, nums = [2,2,_,_]
    Explanation：
---
Case2:

    Input: nums = [0,1,2,2,3,0,4,2], val = 2

    Output: 5, nums = [0,1,4,0,3,_,_,_]
    Explanation：
---

## tips
* 0 <= nums.length <= 100
* 0 <= nums[i] <= 50
* 0 <= val <= 100

## Solution
### 双指针
使用双指针，把输出的数组直接写在输入数组上：右指针 right 指向当前将要处理的元素，左指针 left 指向下一个将要赋值的位置。

* 如果右指针指向的元素不等于 val，它一定是输出数组的一个元素，我们就将右指针指向内容复制到左指针位置；左指针右移一位。
* 如果右指针指向的元素等于 val，它就不需要进入输出数组，此时左指针不动，右指针右移一位。

> 复杂度分析  
> 时间复杂度：O(n) 遍历该数组最多两次完成  
> 空间复杂度：O(1)  

### 双指针优化
优化场景：如果需要移除的元素恰好在数组的开头，我们需要把每一个元素都左移一位，但题目不要求元素顺序，因此将数组最后一个元素移动到序列开头仍满足要求。

* 如果左指针 left 指向的元素等于 val，此时将右指针 right 指向的元素复制到左指针 left 的位置，然后右指针 right 左移一位。
* 如果赋值过来的元素恰好也等于 val，可以继续把右指针 right 指向的元素的值赋值过来（左指针 left 指向的等于 val 的元素的位置继续被覆盖），直到左指针指向的元素的值不等于 val 为止。

当左指针 left 和右指针 right 重合的时候，左右指针遍历完数组中所有的元素。

这样在最坏的情况下只遍历了数组一次，避免了需要保留的元素的重复赋值操作。

复杂度仍是一样的
