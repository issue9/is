// SPDX-License-Identifier: MIT

package is

import (
	"bytes"
	"math"
	"testing"
	"time"

	"github.com/issue9/assert"
)

func TestNumber(t *testing.T) {
	a := assert.New(t)

	a.True(Number("123"))
	a.True(Number("+123"))
	a.True(Number("-123"))
	a.True(Number(".123"))
	a.True(Number("12.3"))
	a.True(Number("+12.3"))
	a.True(Number("-123."))

	a.False(Number("1-23"))
	a.False(Number("+abc"))
	a.False(Number(".12.3"))
	a.False(Number("１２３")) // 全角

	a.True(Number(123))
	a.True(Number(123.1))
	a.True(Number(0))
	a.True(Number(math.Inf(1)))
	a.True(Number(math.Inf(-1)))
	a.True(Number([]rune("123.3")))
	a.True(Number([]byte("123.3")))
}

func TestNil(t *testing.T) {
	a := assert.New(t)

	a.True(Nil(nil))

	var x interface{}
	a.True(Nil(x))

	var y *bytes.Buffer
	a.True(Nil(y))

	a.False(Nil(0))

	x = 5
	a.False(Nil(x))

	y = new(bytes.Buffer)
	a.False(Nil(y))
}

func TestEmpty(t *testing.T) {
	a := assert.New(t)

	a.True(Empty(0, false))
	a.True(Empty(nil, false))
	a.True(Empty(nil, true))
	a.True(Empty(0.0, false))
	a.False(Empty(-0.0001, true))

	var i interface{}
	a.True(Empty(i, false))
	a.True(Empty(i, true))
	a.False(Empty(&i, false))
	a.True(Empty(&i, true))

	i = []string{}
	a.True(Empty(i, false))
	a.True(Empty(i, true))
	a.False(Empty(&i, false))
	a.True(Empty(&i, true))

	var x []string
	a.True(Empty(x, false))
	a.True(Empty(x, true))
	a.False(Empty(&x, false))
	a.True(Empty(&x, true))

	x = []string{""}
	a.False(Empty(x, false))
	a.False(Empty(x, true))

	var ii interface{} = []string{}
	a.True(Empty(ii, false))
	a.True(Empty(ii, true))
	a.False(Empty(&ii, false))
	a.True(Empty(&ii, true))

	var m map[string]string
	a.True(Empty(m, false))
	a.True(Empty(m, true))
	a.False(Empty(&m, false))
	a.True(Empty(&m, true))

	m = map[string]string{}
	a.True(Empty(m, false))
	a.True(Empty(m, true))
	a.False(Empty(&m, false))
	a.True(Empty(&m, true))

	i = m
	a.True(Empty(m, false))
	a.True(Empty(m, true))
	a.False(Empty(&m, false))
	a.True(Empty(&m, true))

	a.True(Empty(time.Time{}, false))
	a.True(Empty(time.Time{}, true))
	a.False(Empty(&time.Time{}, false))
	a.True(Empty(&time.Time{}, true))
	a.False(Empty(time.Now(), false))
	a.False(Empty(time.Now(), true))

	type obj struct{ Int int }
	a.True(Empty(obj{Int: 0}, false))
	a.False(Empty(&obj{Int: 0}, false))
	a.True(Empty(&obj{Int: 0}, true))
	a.False(Empty(obj{Int: 1}, false))
}

func TestZero(t *testing.T) {
	a := assert.New(t)

	a.True(Zero(0, false))
	a.True(Zero(nil, false))
	a.True(Zero(nil, true))
	a.True(Zero(0.0, false))
	a.False(Zero(-0.0001, true))

	var i interface{}
	a.True(Zero(i, false))
	a.True(Zero(i, true))
	a.False(Zero(&i, false))
	a.True(Zero(&i, true))

	i = []string{}
	a.False(Zero(i, false))
	a.False(Zero(i, true))
	a.False(Zero(&i, false))
	a.False(Zero(&i, true))

	var x []string
	a.True(Zero(x, false))
	a.True(Zero(x, true))
	a.False(Zero(&x, false))
	a.True(Zero(&x, true))

	x = []string{""}
	a.False(Zero(x, false))
	a.False(Zero(x, true))

	var ii interface{} = []string{}
	a.False(Zero(ii, false))
	a.False(Zero(ii, true))
	a.False(Zero(&ii, false))
	a.False(Zero(&ii, true))

	var m map[string]string
	a.True(Zero(m, false))
	a.True(Zero(m, true))
	a.False(Zero(&m, false))
	a.True(Zero(&m, true))

	m = map[string]string{}
	a.False(Zero(m, false))
	a.False(Zero(m, true))
	a.False(Zero(&m, false))
	a.False(Zero(&m, true))

	i = m
	a.False(Zero(m, false))
	a.False(Zero(m, true))
	a.False(Zero(&m, false))
	a.False(Zero(&m, true))

	a.True(Zero(time.Time{}, false))
	a.True(Zero(time.Time{}, true))
	a.False(Zero(&time.Time{}, false))
	a.True(Zero(&time.Time{}, true))
	a.False(Zero(time.Now(), false))
	a.False(Zero(time.Now(), true))

	type obj struct{ Int int }
	a.True(Zero(obj{Int: 0}, false))
	a.False(Zero(&obj{Int: 0}, false))
	a.True(Zero(&obj{Int: 0}, true))
	a.False(Zero(obj{Int: 1}, false))
}

func TestHexColor(t *testing.T) {
	a := assert.New(t)

	a.True(HexColor("#123"))
	a.True(HexColor("#fff"))
	a.True(HexColor("#f0f0f0"))
	a.True(HexColor("#fafafa"))
	a.True(HexColor("#224422"))

	a.False(HexColor("#2244"))
	a.False(HexColor("#hhh"))
	a.False(HexColor("#asdf"))
	a.False(HexColor("#ffff"))
}
