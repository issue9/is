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

	a.True(String("6259650871772098"))
	a.True(String("79927398713"))
	a.False(String("79927398710"))
}
