// SPDX-License-Identifier: MIT

// Package luhn 模 10 校验算法
//
// https://en.wikipedia.org/wiki/Luhn_algorithm
//
// 1. 从右往左，偶数位数字乘以 2，如果是两位数，将其个数数和十位数相加；
// 2. 将以上的所有数值相加得到值 n1；
// 3. 从右往左，奇数位的数值加相加午到值 n2；
// 4. (n1+n2) % 10 如果值为 0，表示正确。
package luhn

import "github.com/issue9/validation/is/luhn"

// Valid 传入 []byte 验证是否正确
func Valid(v []byte) bool {
	return luhn.IsValid(v)
}

// ValidString 传入 string 验证是否正确
func ValidString(v string) bool {
	return Valid([]byte(v))
}

// GenerateWithPrefix 给定前缀，添加最后一位校验位
func GenerateWithPrefix(prefix []byte) []byte {
	return luhn.GenerateWithPrefix(prefix)
}
