# Description
Given a positive integer n, find the sum of all integers in the range [1, n] inclusive that are divisible by 3, 5, or 7.

Return an integer denoting the sum of all numbers in the given range satisfying the constraint.
# Test Case
Case1:

    Input: n = 7
    Output: 21

    Explanation: Numbers in the range [1, 7] that are divisible by 3, 5, or 7 are 3, 5, 6, 7. The sum of these numbers is 21.
---
Case2:

    Input: n = 10
    Output: 40

    Explanation: Numbers in the range [1, 10] that are divisible by 3, 5, or 7 are 3, 5, 6, 7, 9, 10. The sum of these numbers is 40.
---
Case3:

    Input: n = 9
    Output: 30

    Explanation: Numbers in the range [1, 9] that are divisible by 3, 5, or 7 are 3, 5, 6, 7, 9. The sum of these numbers is 30.
# Solution
## 枚举
列出`[1,n]`的所有整数，记当前枚举的数为`i`，如果`i mod 3 = 0`或`i mod 5 = 0`或`i mod 7 = 0`那么将`i`加到总和`result`。

> 复杂度分析  
> 时间复杂度：`O(n)`。  
> 空间复杂度：`O(1)`。
## 容斥原理**
抽象一个方法：在区间` [1,n] `内能被数` m `整除的整数的和为：

f(n,m) = (m + m * [n/m]) * [n/m]/2
> [n/m]：n除以m向下取整

根据 容斥原理，在区间` [1,n] `内，能被`3`、`5`和`7`整除的整数之和为：

f(n,3)+f(n,5)+f(n,7)−f(n,3×5)−f(n,3×7)−f(n,5×7)+f(n,3×5×7)
> 复杂度分析  
> 时间复杂度：`O(1)`。  
> 空间复杂度：`O(1)`。