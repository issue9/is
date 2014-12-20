// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"
)

// 验证函数返回的结果。相对于直接返回bool，该结构提供了
// 相应的key和message，适合有针对性的自定义验证函数。
type Result struct {
	Message, Key string
	Ok           bool
}

// 修改Key的值
func (r *Result) SetKey(key string) *Result {
	r.Key = key

	return r
}

// 修改Message的值。
func (r *Result) SetMessage(msg string, args ...interface{}) *Result {
	if len(args) == 0 {
		r.Message = msg
	} else {
		r.Message = fmt.Sprintf(msg, args...)
	}

	return r
}

// Validation相当于一个错误容器，存放从Apply()获取的错误信息。
type Validation struct {
	errs    []string
	errsMap map[string]string
}

// 声明一个新的Validation
func New() *Validation {
	return &Validation{errsMap: make(map[string]string)}
}

// 从一个Result对象中判断是否存在错误，有则保存之。
func (v *Validation) ApplyResult(r *Result) *Validation {
	return v.Apply(r.Ok, r.Message, r.Key)
}

// 判断expr的值，若是false，则保存msg和key到Validation对象中。
// 若不需要key则传递空字符串。同一key若提供了多条msg，则只有最
// 后一条会被保存。
func (v *Validation) Apply(expr bool, msg, key string) *Validation {
	if expr {
		return v
	}

	v.errs = append(v.errs, msg)

	if len(key) > 0 {
		v.errsMap[key] = msg
	}

	return v
}

// 存在错误
func (v *Validation) HasErrors() bool {
	return len(v.errs) > 0
}

// 返回所有的错误信息
func (v *Validation) Errors() []string {
	return v.errs
}

// 返回以key分类的错误信息.
func (v *Validation) ErrorsMap() map[string]string {
	return v.errsMap
}

// 清除所有的错误信息
func (v *Validation) Clear() {
	v.errs = v.errs[:0]
	v.errsMap = make(map[string]string)
}
