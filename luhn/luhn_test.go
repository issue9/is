// Copyright 2019 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package luhn

import (
	"testing"

	"github.com/issue9/assert"
)

func TestString(t *testing.T) {
	a := assert.New(t)

	a.True(ValidString("6259650871772098"))
	a.True(ValidString("79927398713"))
	a.False(ValidString("79927398710"))
}

func TestGenerateWithPrefix(t *testing.T) {
	a := assert.New(t)

	a.Equal("6259650871772098", string(GenerateWithPrefix([]byte("625965087177209"))))
	a.True(Valid(GenerateWithPrefix([]byte("625965087177209"))))
}
