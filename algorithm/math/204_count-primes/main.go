package main

import "fmt"

/* 计数质数 */

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	//var isPrime func(x int) bool
	//isPrime = func(x int) bool {
	//	for i := 2; i < x; i++ {
	//		if x%i == 0 {
	//			return false
	//		}
	//	}
	//	return true
	//}
	//var count int
	//for i := 2; i < n; i++ {
	//	if isPrime(i) {
	//		count++
	//	}
	//}
	//return count
	isPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			for j := 2 * i; j < n; j = j + i {
				isPrime[j] = false
			}
		}
	}

	var count int
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}
