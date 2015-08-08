// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// is包提供了一系列的判断函数。
package is

import (
	"strconv"

	"github.com/issue9/assert"
)

// 判断一个值是否可转换为数值。不支持全角数值的判断。
func Number(val interface{}) bool {
	switch v := val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	case []byte:
		_, err := strconv.ParseFloat(string(v), 32)
		return err == nil
	case string:
		_, err := strconv.ParseFloat(v, 32)
		return err == nil
	case []rune:
		_, err := strconv.ParseFloat(string(v), 32)
		return err == nil
	default:
		return false
	}
}

// 是否为nil，有类型但无具体值的也将返回true，具体可参考
// github.com/assert.IsNil()函数。
func Nil(val interface{}) bool {
	return assert.IsNil(val)
}

// 是否为空，若是窝器类型，长度为0也将返回true，具体可参考
// github.com/assert.IsEmpty()函数。
func Empty(val interface{}) bool {
	return assert.IsEmpty(val)
}
