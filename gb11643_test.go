// SPDX-License-Identifier: MIT

package is

import (
	"math"
	"testing"

	"github.com/issue9/assert"
)

// 计算各个数值位对应的系数值。
func getWeight() []int {
	l := 17
	ret := make([]int, 17, 17)
	for i := 0; i < l; i++ {
		k := int(math.Pow(2, float64((l - i)))) // k值足够大，不能用byte保存
		ret[i] = k % 11
	}
	return ret
}

func TestGetWeight(t *testing.T) {
	assert.Equal(t, gb11643Weight, getWeight())
}

func TestGB11643(t *testing.T) {
	a := assert.New(t)

	// 网上扒来的身份证，不与现实中的对应。
	a.True(GB11643("350303199002033073"))
	a.True(GB11643("350303900203307"))
	a.True(GB11643("331122197905116239"))
	a.True(GB11643("513330199111066159"))
	a.True(GB11643("33050219880702447x"))
	a.True(GB11643("33050219880702447X"))
	a.True(GB11643("330502880702447"))

	a.False(GB11643("111111111111111111"))
	a.False(GB11643("330502198807024471")) // 最后一位不正确
	a.False(GB11643("33050288070244"))     // 14位
	a.False(GB11643("3305028807024411"))   // 16位
}
