package main

import "fmt"

func main() {
	var n, a, b, result int

	fmt.Scanf("%d %d %d", &n, &a, &b)

	for i := 1; i <= n; i++ {
		var tmp, sum int
		tmp = i

		for tmp != 0 {
			sum += tmp % 10
			tmp /= 10
		}

		if sum >= a && sum <= b {
			result += i
		}
	}
	fmt.Println(result)
}
