// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"testing"

	"github.com/issue9/assert"
)

func TestIsCnPhone(t *testing.T) {
	a := assert.New(t)

	a.True(IsCnPhone("444488888888-4444"))
	a.True(IsCnPhone("3337777777-1"))
	a.True(IsCnPhone("333-7777777-1"))
	a.True(IsCnPhone("7777777"))
	a.True(IsCnPhone("88888888"))

	a.False(IsCnPhone("333-7777777-"))      // 尾部没有分机号
	a.False(IsCnPhone("333-999999999"))     // 9位数的号码
	a.False(IsCnPhone("22-88888888"))       // 区号只有2位
	a.False(IsCnPhone("33-88888888-55555")) // 分机号超过4位
}

func TestIsCnMobile(t *testing.T) {
	a := assert.New(t)

	a.True(IsCnMobile("15011111111"))
	a.True(IsCnMobile("015011111111"))
	a.True(IsCnMobile("8615011111111"))
	a.True(IsCnMobile("+8615011111111"))

	a.False(IsCnMobile("+86150111111112")) // 尾部多个2
	a.False(IsCnMobile("50111111112"))     // 开头少1
	a.False(IsCnMobile("+8650111111112"))  // 开头少1
	a.False(IsCnMobile("8650111111112"))   // 开头少1
	a.False(IsCnMobile("154111111112"))    // 不存在的前缀150
}

func TestIsURL(t *testing.T) {
	a := assert.New(t)

	a.True(IsURL("http://www.example.com"))
	a.True(IsURL("http://www.example.com/path/?a=b"))
	a.True(IsURL("https://www.example.com:88/path"))
	a.True(IsURL("ftp://pwd:user@www.example.com/index.go?a=b"))
	a.True(IsURL("pwd:user@www.example.com/path/"))
	a.True(IsURL("pwd:user@www.example.com:80/path/"))
	a.True(IsURL("https://127.0.0.1/path/"))
	a.True(IsURL("https://fe80:0:0:0:204:61ff:fe9d:f156/path/"))
	a.True(IsURL("https://127.0.0.1/path//index.go?arg1=val1"))
	a.True(IsURL("https://::1/path/index.go?arg1=val1"))

	a.False(IsURL("https://[::1]:80/path/"))
	a.False(IsURL("https://298.1.1.1/path/index.go?arg1=val1"))
	a.False(IsURL("https://~.example.com/path/index.go?arg1=val1"))
}

func TestIsIP(t *testing.T) {
	a := assert.New(t)

	a.True(IsIP("fe80:0000:0000:0000:0204:61ff:fe9d:f156"))
	a.True(IsIP("fe80:0:0:0:204:61ff:fe9d:f156"))
	a.True(IsIP("0.0.0.0"))
	a.True(IsIP("255.255.255.255"))
	a.True(IsIP("255.0.3.255"))

	a.False(IsIP("255.0:3.255"))
	a.False(IsIP("275.0.3.255"))
}

func TestIsIP6(t *testing.T) {
	a := assert.New(t)

	a.True(IsIP6("fe80:0000:0000:0000:0204:61ff:fe9d:f156"))      // full form of IPv6
	a.True(IsIP6("fe80:0:0:0:204:61ff:fe9d:f156"))                // drop leading zeroes
	a.True(IsIP6("fe80::204:61ff:fe9d:f156"))                     // collapse multiple zeroes to :: in the IPv6 address
	a.True(IsIP6("fe80:0000:0000:0000:0204:61ff:254.157.241.86")) // IPv4 dotted quad at the end
	a.True(IsIP6("fe80:0:0:0:0204:61ff:254.157.241.86"))          // drop leading zeroes, IPv4 dotted quad at the end
	a.True(IsIP6("fe80::204:61ff:254.157.241.86"))                // dotted quad at the end, multiple zeroes collapsed
	a.True(IsIP6("::1"))                                          // localhost
	a.True(IsIP6("fe80::"))                                       // link-local prefix
	a.True(IsIP6("2001::"))                                       //global unicast prefix
}

func TestIsIP4(t *testing.T) {
	a := assert.New(t)

	a.True(IsIP4("0.0.0.0"))
	a.True(IsIP4("255.255.255.255"))
	a.True(IsIP4("255.0.3.255"))
	a.True(IsIP4("127.010.0.1"))
	a.True(IsIP4("027.01.0.1"))

	a.False(IsIP4("1127.01.0.1"))
}

func TestIsEmail(t *testing.T) {
	a := assert.New(t)

	a.True(IsEmail("email@email.com"))
	a.True(IsEmail("em2il@email.com.cn"))
	a.True(IsEmail("email.test@email.com"))
	a.True(IsEmail("email.test@email123.com"))
	a.True(IsEmail("em2il@email"))

	// 2个@
	a.False(IsEmail("em@2l@email.com"))
	// 没有@
	a.False(IsEmail("email2email.com.cn"))
}
