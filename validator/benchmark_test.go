// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"regexp"
	"strconv"
	"testing"
)

var expr = regexp.MustCompile("(^\\.|\\+|\\-)?\\d+\\.?\\d*$")

// 正则表达式方式实现
func isNumberRegexp(val interface{}) bool {
	switch v := val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	case []byte:
		return expr.Match(v)
	case string:
		return expr.MatchString(v)
	default:
		return false
	}
}

// 通过判断strconv.Parse返回值的方式实现
func isNumberStrconv(val interface{}) bool {
	switch v := val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	case []byte:
		if _, err := strconv.ParseFloat(string(v), 64); err != nil {
			return false
		}
		return true
	case string:
		if _, err := strconv.ParseFloat(v, 64); err != nil {
			return false
		}
		return true
	default:
		return false
	}
}

func BenchmarkIsNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsNumber("+1.2.1")
		IsNumber("+1abcd.2.1")
	}
}

func BenchmarkIsNumberRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isNumberRegexp("+1.2.1")
		isNumberRegexp("+1abcd.2.1")
	}
}

func BenchmarkIsNumberStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isNumberStrconv("+1.2.1")
		isNumberStrconv("+1abcd.2.1")
	}
}
