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
// Deprecated: 请使用 Zero 代替
func Empty(val interface{}) bool {
	return Zero(val, false)
}

// Zero 判断当前是否为空或是零值
//
// ptr 表示当 val 是指针时，是否分析指向的值。
//
// 若是容器类型，长度为 0 也将返回 true，
// 但是 []string{""}空数组里套一个空字符串，不会被判断为空。
//
// NOTE: 如果是指针，并不会判断指针指向的值，只判断指针是否为 nil。
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

	// 特定类型的判断
	switch val := v.Interface().(type) {
	case time.Time:
		return val.IsZero()
	}

	// 内置类型的判断
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	}

	// 空值的判断
	if reflect.Zero(v.Type()).Interface() == v.Interface() {
		return true
	}

	// 符合 Nil 条件的，都为 Empty
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
