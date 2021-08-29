// Copyright (c) 2017-2019 Sergey Kamardin <gobwas@gmail.com>
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package internal

const (
	// ^uint(0) 64为无符号整形取反，每一位都是1
	// 左移63位，剩下1， 32 << 1 = 64相当于乘以2
	bitsize = 32 << (^uint(0) >> 63)
	// 1 << 62
	maxintHeadBit = 1 << (bitsize - 2)
)

// IsPowerOfTwo reports whether given integer is a power of two.
// 说明这个数是不是2的指数
func IsPowerOfTwo(n int) bool {
	return n&(n-1) == 0
}

// CeilToPowerOfTwo returns the least power of two integer value greater than
// or equal to n.
// 近似输出一个2的指数的值
func CeilToPowerOfTwo(n int) int {
	// n & maxintHeadBit !=0 说明两数在63位相同，至少这个数要不小于maxintHeadBit
	if n&maxintHeadBit != 0 && n > maxintHeadBit {
		panic("argument is too large")
	}
	if n <= 2 {
		return 2
	}
	n--
	n = fillBits(n)
	n++
	return n
}

// FloorToPowerOfTwo returns the greatest power of two integer value less than
// or equal to n.
// 相当于向下取2的指数值
func FloorToPowerOfTwo(n int) int {
	if n <= 2 {
		return 2
	}
	n = fillBits(n)
	n >>= 1
	n++
	return n
}

// 将n涉及的所有位填满1
func fillBits(n int) int {
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	return n
}
