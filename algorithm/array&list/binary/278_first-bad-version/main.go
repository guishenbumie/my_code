package main

import "sort"

/* 第一个错误的版本 */

func main() {

}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	return sort.Search(n, func(i int) bool {
		return isBadVersion(i)
	})
}
