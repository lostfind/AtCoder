package main

import "fmt"

func main() {
	var length, max, sumTop, sumBottom int
	fmt.Scan(&length)

	top := make([]int, length)
	bottom := make([]int, length)

	for i := 0; i < length; i++ {
		fmt.Scan(&top[i])
	}
	for i := 0; i < length; i++ {
		fmt.Scan(&bottom[i])
	}

	for i := 0; i < length; i++ {
		sumTop += top[i]
		sumBottom = 0
		for j := i; j < length; j++ {
			sumBottom += bottom[j]
		}
		if max < sumTop+sumBottom {
			max = sumTop + sumBottom
		}
	}

	fmt.Print(max)
}
