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

// Valid 传入 []byte 验证是否正确
func Valid(v []byte) bool {
	return 0 == (checksum(v, false) % 10)
}

// ValidString 传入 string 验证是否正确
func ValidString(v string) bool {
	return Valid([]byte(v))
}

// GenerateWithPrefix 给定前缀，添加最后一位校验位
func GenerateWithPrefix(prefix []byte) []byte {
	sum := checksum(prefix, true)
	m := 10 - sum%10
	return append(prefix, byte(m)+'0')
}

// v 需要计算工的字符串
// alt 倒数第一位是否为偶数位，如果是验证，则是 false，如果是计算验证位，则为 true。
func checksum(v []byte, alt bool) (sum int) {
	for i := len(v) - 1; i >= 0; i-- {
		n := int(v[i] - '0')
		if alt {
			n *= 2
			if n > 9 {
				n = n%10 + 1
			}
		}

		sum += n
		alt = !alt
	}

	return sum
}
