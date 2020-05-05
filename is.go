// SPDX-License-Identifier: MIT

// Package is 包提供了一系列的判断函数。
package is

import (
	"reflect"
	"strconv"
	"time"

	"github.com/issue9/is/luhn"
)

// Number 判断一个值是否可转换为数值
//
// NOTE: 不支持全角数值的判断
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

// Nil 是否为 nil
//
// 有类型但无具体值的也将返回 true，
// 当特定类型的变量，已经声明，但还未赋值时，也将返回 true
func Nil(val interface{}) bool {
	if nil == val {
		return true
	}

	v := reflect.ValueOf(val)
	k := v.Kind()
	return k >= reflect.Chan && k <= reflect.Slice && v.IsNil()
}

// Empty 判断当前是否为空或是零值
//
// ptr 表示当 val 是指针时，是否分析指向的值。
//
// 若是容器类型，长度为 0 也将返回 true，
// 但是 []string{""}空数组里套一个空字符串，不会被判断为空。
func Empty(val interface{}, ptr bool) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	if ptr {
		for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
			if v.IsNil() {
				return true
			}
			v = v.Elem()
		}
	}

	switch v.Kind() {
	case reflect.Slice:
		return v.Len() == 0
	}

	return zero(v)
}

// Zero 判断当前是否为空或是零值
//
// ptr 表示当 val 是指针时，是否分析指向的值。
// 在 reflect.Value.IsZero 的基础上对特写类型作为特殊处理，比如 time.IsZero()
func Zero(val interface{}, ptr bool) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	if ptr {
		for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
			if v.IsNil() {
				return true
			}
			v = v.Elem()
		}
	}
	return zero(v)
}

func zero(v reflect.Value) bool {
	if v.IsZero() {
		return true
	}

	val := v.Interface()

	// 特定类型的判断
	switch v := val.(type) {
	case time.Time:
		return v.IsZero()
	}

	return Nil(val)
}

// HexColor 判断一个字符串是否为合法的 16 进制颜色表示法
func HexColor(val interface{}) bool {
	var bs []byte
	switch v := val.(type) {
	case []byte:
		bs = v
	case []rune:
		bs = []byte(string(v))
	case string:
		bs = []byte(v)
	default:
		return false
	}

	if len(bs) != 4 && len(bs) != 7 {
		return false
	}

	if bs[0] != '#' {
		return false
	}

	for _, v := range bs[1:] {
		switch {
		case '0' <= v && v <= '9':
		case 'a' <= v && v <= 'f':
		case 'A' <= v && v <= 'F':
		default:
			return false
		}
	}
	return true
}

// BankCard 是否为正确的银行卡号
func BankCard(val interface{}) bool {
	switch v := val.(type) {
	case []byte:
		return luhn.Valid(v)
	case string:
		return luhn.ValidString(v)
	case []rune:
		return luhn.ValidString(string(v))
	default:
		return false
	}
}
