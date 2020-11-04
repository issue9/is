// SPDX-License-Identifier: MIT

package is

import "github.com/issue9/validation/is"

// GB11643 判断一个身份证是否符合 gb11643 标准
//
// 若是 15 位则当作一代身份证，仅简单地判断各位是否都是数字；
// 若是 18 位则当作二代身份证，会计算校验位是否正确；
// 其它位数都返回 false。
func GB11643(val interface{}) bool {
	return is.GB11643(val)
}
