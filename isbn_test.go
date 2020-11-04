// SPDX-License-Identifier: MIT

package is

import (
	"testing"

	"github.com/issue9/assert"
)

func TestISBN(t *testing.T) {
	a := assert.New(t)

	a.True(ISBN("1-919876-03-0"))
	a.True(ISBN("0-471-00084-1"))
	a.True(ISBN("978-7-301-04815-3"))
	a.True(ISBN("978-986-181-728-6"))
}
