// SPDX-License-Identifier: MIT

package is

import (
	"testing"

	"github.com/issue9/assert"
)

func TestCNPhone(t *testing.T) {
	a := assert.New(t)

	a.True(CNPhone("444488888888-4444"))
	a.True(CNPhone("3337777777-1"))
	a.True(CNPhone("333-7777777-1"))
	a.True(CNPhone("7777777"))
	a.True(CNPhone("88888888"))

	a.False(CNPhone("333-7777777-"))      // 尾部没有分机号
	a.False(CNPhone("22-88888888"))       // 区号只有2位
	a.False(CNPhone("33-88888888-55555")) // 分机号超过4位
}

func TestCNMobile(t *testing.T) {
	a := assert.New(t)

	a.True(CNMobile("15011111111"))
	a.True(CNMobile("015011111111"))
	a.True(CNMobile("8615011111111"))
	a.True(CNMobile("+8615011111111"))
	a.True(CNMobile("+8619911111111"))

	a.False(CNMobile("+86150111111112")) // 尾部多个2
	a.False(CNMobile("50111111112"))     // 开头少1
	a.False(CNMobile("+8650111111112"))  // 开头少1
	a.False(CNMobile("8650111111112"))   // 开头少1
	a.False(CNMobile("154111111112"))    // 不存在的前缀154
	a.True(CNMobile("15411111111"))
}

func TestCNTel(t *testing.T) {
	a := assert.New(t)

	a.True(CNTel("444488888888-4444"))
	a.True(CNTel("3337777777-1"))
	a.True(CNTel("333-7777777-1"))
	a.True(CNTel("7777777"))
	a.True(CNTel("88888888"))
	a.True(CNTel("15011111111"))
	a.True(CNTel("015011111111"))
	a.True(CNTel("8615011111111"))
	a.True(CNTel("+8615011111111"))
}

func TestURL(t *testing.T) {
	a := assert.New(t)

	a.True(URL("www.example.com"))
	a.True(URL("http://www.example.com"))
	a.True(URL([]byte("http://example.com")))
	a.True(URL("http://www.example.com/"))
	a.True(URL("http://www.example.com/path/?a=b"))
	a.True(URL("https://www.example.com:88/path1/path2"))
	a.True(URL("ftp://pwd:user@www.example.com/index.go?a=b"))
	a.True(URL([]byte("ftp://pwd:user@www.example.com/index.go?a=b")))
	a.True(URL("pwd:user@www.example.com/path/"))
	a.True(URL("pwd:user@www.example.com:80/path/"))
	a.True(URL("https://127.0.0.1/path/"))
	a.True(URL("https://fe80:0:0:0:204:61ff:fe9d:f156/path/"))
	a.True(URL("https://127.0.0.1/path//index.go?arg1=val1&arg2=val/2"))
	a.True(URL("https://::1/path/index.go?arg1=val1"))

	a.False(URL("https://[::1]:80/path/"))
	a.False(URL("https://298.1.1.1/path/index.go?arg1=val1"))
	a.False(URL("https://~.example.com/path/index.go?arg1=val1"))
}

func TestIP(t *testing.T) {
	a := assert.New(t)

	a.True(IP("fe80:0000:0000:0000:0204:61ff:fe9d:f156"))
	a.True(IP("fe80:0:0:0:204:61ff:fe9d:f156"))
	a.True(IP("0.0.0.0"))
	a.True(IP("255.255.255.255"))
	a.True(IP("255.0.3.255"))

	a.False(IP("255.0:3.255"))
	a.False(IP("275.0.3.255"))
}

func TestIP6(t *testing.T) {
	a := assert.New(t)

	a.True(IP6("fe80:0000:0000:0000:0204:61ff:fe9d:f156"))      // full form of IPv6
	a.True(IP6("fe80:0:0:0:204:61ff:fe9d:f156"))                // drop leading zeroes
	a.True(IP6("fe80::204:61ff:fe9d:f156"))                     // collapse multiple zeroes to :: in the IPv6 address
	a.True(IP6("fe80:0000:0000:0000:0204:61ff:254.157.241.86")) // IPv4 dotted quad at the end
	a.True(IP6("fe80:0:0:0:0204:61ff:254.157.241.86"))          // drop leading zeroes, IPv4 dotted quad at the end
	a.True(IP6("fe80::204:61ff:254.157.241.86"))                // dotted quad at the end, multiple zeroes collapsed
	a.True(IP6("::1"))                                          // localhost
	a.True(IP6("fe80::"))                                       // link-local prefix
	a.True(IP6("2001::"))                                       // global unicast prefix
}

func TestIP4(t *testing.T) {
	a := assert.New(t)

	a.True(IP4("0.0.0.0"))
	a.True(IP4("255.255.255.255"))
	a.True(IP4("255.0.3.255"))
	a.True(IP4("127.010.0.1"))
	a.True(IP4("027.01.0.1"))

	a.False(IP4("1127.01.0.1"))
}

func TestEmail(t *testing.T) {
	a := assert.New(t)

	a.True(Email("email@email.com"))
	a.True(Email("em2il@email.com.cn"))
	a.True(Email("12345@qq.com"))
	a.True(Email("email.test@email.com"))
	a.True(Email("email.test@email123.com"))
	a.True(Email("em2il@email"))

	// 2个@
	a.False(Email("em@2l@email.com"))
	// 没有@
	a.False(Email("email2email.com.cn"))
}
