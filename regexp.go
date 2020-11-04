// SPDX-License-Identifier: MIT

package is

import "github.com/issue9/validation/is"

// CNPhone 验证中国大陆的电话号码
//
// 支持如下格式：
//  0578-12345678-1234
//  057812345678-1234
// 若存在分机号，则分机号的连接符不能省略。
func CNPhone(val interface{}) bool {
	return is.CNPhone(val)
}

// CNMobile 验证中国大陆的手机号码
func CNMobile(val interface{}) bool {
	return is.CNMobile(val)
}

// CNTel 验证手机和电话类型
func CNTel(val interface{}) bool {
	return is.CNTel(val)
}

// URL 验证一个值是否标准的URL格式
//
// 支持 IP 和域名等格式
func URL(val interface{}) bool {
	return is.URL(val)
}

// IP 验证一个值是否为 IP
//
// 可验证 IP4 和 IP6
func IP(val interface{}) bool {
	return is.IP(val)
}

// IP6 验证一个值是否为 IP6
func IP6(val interface{}) bool {
	return is.IP6(val)
}

// IP4 验证一个值是滞为 IP4
func IP4(val interface{}) bool {
	return is.IP4(val)
}

// Email 验证一个值是否匹配一个邮箱
func Email(val interface{}) bool {
	return is.Email(val)
}
