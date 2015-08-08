// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package is

import (
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
