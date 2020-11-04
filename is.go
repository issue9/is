// SPDX-License-Identifier: MIT

// Package is 包提供了一系列的判断函数
package is

import (
	"github.com/issue9/validation/is"
)

// Number 判断一个值是否可转换为数值
//
// NOTE: 不支持全角数值的判断
func Number(val interface{}) bool {
	return is.Number(val)
}

// Nil 是否为 nil
//
// 有类型但无具体值的也将返回 true，
// 当特定类型的变量，已经声明，但还未赋值时，也将返回 true
func Nil(val interface{}) bool {
	return is.Nil(val)
}

// Empty 判断当前是否为空或是零值
//
// ptr 表示当 val 是指针时，是否分析指向的值。
//
// 若是容器类型，长度为 0 也将返回 true，
// 但是 []string{""}空数组里套一个空字符串，不会被判断为空。
func Empty(val interface{}, ptr bool) bool {
	return is.Empty(val, ptr)
}

// Zero 判断当前是否为空或是零值
//
// ptr 表示当 val 是指针时，是否分析指向的值。
// 在 reflect.Value.IsZero 的基础上对特写类型作为特殊处理，比如 time.IsZero()
func Zero(val interface{}, ptr bool) bool {
	return is.Zero(val, ptr)
}

// HexColor 判断一个字符串是否为合法的 16 进制颜色表示法
func HexColor(val interface{}) bool {
	return is.HexColor(val)
}

// BankCard 是否为正确的银行卡号
func BankCard(val interface{}) bool {
	return is.BankCard(val)
}
