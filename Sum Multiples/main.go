package main

func main() {
	println(sunOfMultiples(7))  //21
	println(sunOfMultiples2(7)) //21

	println(sunOfMultiples(10))  //40
	println(sunOfMultiples2(10)) //40

	println(sunOfMultiples(9))  //30
	println(sunOfMultiples2(9)) //30
}

// Solution1
func sunOfMultiples(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			result += i
		}
	}

	return result
}

// Solution2
func sunOfMultiples2(n int) int {
	return f(n, 3) + f(n, 5) + f(n, 7) - f(n, 3*5) - f(n, 3*7) - f(n, 5*7) + f(n, 3*5*7)
}

func f(n, m int) int {
	return (m + n/m*m) * (n / m) / 2
}
