// Copyright 2019 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package luhn 模 10 校验算法
//
// https://en.wikipedia.org/wiki/Luhn_algorithm
//
// 1. 从右往左，偶数位数字乘以 2，如果是两位数，将其个数数和十位数相加；
// 2. 将以上的所有数值相加得到值 n1；
// 3. 从右往左，奇数位的数值加相加午到值 n2；
// 4. (n1+n2) % 10 如果值为 0，表示正确。
package luhn

// Bytes 传入 []byte 验证是否正确
func Bytes(v []byte) bool {
	var sum int

	// 奇数位
	for i := len(v) - 1; i >= 0; i -= 2 {
		sum += int(v[i] - '0')
	}

	// 偶数位
	for i := len(v) - 2; i >= 0; i -= 2 {
		n := int(v[i]-'0') * 2
		if n > 9 {
			n = n%10 + 1
		}
		sum += n
	}

	return 0 == (sum % 10)
}

// String 传入 string 验证是否正确
func String(v string) bool {
	return Bytes([]byte(v))
}
