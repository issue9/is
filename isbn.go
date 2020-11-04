// SPDX-License-Identifier: MIT

package is

import (
	"github.com/issue9/validation/is"
)

// 有关 ISBN 的算法及其它相关内容，可参照http://zh.wikipedia.org/wiki/%E5%9B%BD%E9%99%85%E6%A0%87%E5%87%86%E4%B9%A6%E5%8F%B7

// ISBN 判断是否为合法的 ISBN 串号。可以同时判断 ISBN10 和 ISBN13
func ISBN(val interface{}) bool {
	return is.ISBN(val)
}

// ISBN10 判断是否为合法的 ISBN10
func ISBN10(val []byte) bool {
	return is.ISBN10(val)
}

// ISBN13 判断是否为合法的 ISBN13
func ISBN13(val []byte) bool {
	return is.ISBN13(val)
}
