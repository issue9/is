// SPDX-License-Identifier: MIT

package is

import (
	"bytes"
	"math"
	"testing"

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

	a.True(Empty(0))
	a.True(Empty(nil))

	var x interface{}
	a.True(Empty(x))
	x = []string{}
	a.True(Empty(x))
	x = []string{""}
	a.False(Empty(x))
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
