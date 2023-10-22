# Description
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

## Test Case
Case1:

    Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3

    Output: [1,2,2,3,5,6]
    Explanation： The arrays we are merging are [1,2,3] and [2,5,6].  
    The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
---
Case2:

    Input: nums1 = [0], m = 0, nums2 = [1], n = 1

    Output: [1]
    Explanation： The arrays we are merging are [] and [1].  
    The result of the merge is [1].  
    Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.
---

## tips
* nums1.length == m + n
* nums2.length == n
* 0 <= m, n <= 200
* 1 <= m + n <= 200
* -10^9 <= nums1[i], nums2[j] <= 10^9

## Solution
### 直接合并后排序
将nums2放进nums1的尾部，然后对整个数组进行排序

### 逆向双指针
数组 nums1 与 nums2 已经被排序了，因此可以将其看做两个队列，采用双指针，每次将取出一个数字放入数组中。

由于从头部开始合并，需要使用额外的临时变量存储结果，否则会将未取出的元素覆盖掉。因此可以利用nums1尾部的空余空间。

将指针设置为从后向前遍历，每次取两者之中的较大者放进 nums1 的最后面。

> 复杂度分析  
> 时间复杂度：O(m+n) 指针最多移动 m+n 次
> 空间复杂度：O(1) 