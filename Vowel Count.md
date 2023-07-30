# 要求
返回给定字符串中的元音字母（a，e，i，o，u）个数

输入字符只包含字母或空格

# 分析
只需要通过一个循环遍历字符串中的每个字符

然后检查该字符是否为元音字母

# 示例
```go
package kata

func GetCount(str string) (count int) {
  for _, c := range str {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
      count++
    }
  }
  return count
}
```
通过` strings `的` Count `或` Contains `方法实现：
```go
package kata

import "strings"

func GetCount(strn string) int {
	count := 0

	vowels := []string{"a", "e", "i", "o", "u"}

	for _, vowel := range vowels {
		count += strings.Count(strn, vowel)
	}

	return count
}
```
```go
package main

import "strings"

func countVowels(s string) int {
	vowels := "aeiou"
	count := 0

	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
```