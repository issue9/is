// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package validator

import (
	"testing"

	"github.com/issue9/assert"
)

func TestIsNumber(t *testing.T) {
	a := assert.New(t)

	a.True(IsNumber("123"))
	a.True(IsNumber("+123"))
	a.True(IsNumber("-123"))
	a.True(IsNumber(".123"))
	a.True(IsNumber("12.3"))
	a.True(IsNumber("+12.3"))
	a.True(IsNumber("-123."))

	a.False(IsNumber("1-23"))
	a.False(IsNumber("+abc"))
	a.False(IsNumber(".12.3"))
	a.False(IsNumber("１２３")) // 全角
}
