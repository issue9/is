// SPDX-License-Identifier: MIT

package gb11643

import (
	"testing"
	"time"

	"github.com/issue9/assert"
)

func TestParse(t *testing.T) {
	a := assert.New(t)

	g, err := Parse("140622209101019893")
	date, err := time.Parse(layout, "20910101")
	a.NotError(err)
	a.NotError(err).NotNil(g)
	a.Equal(g.Raw, "140622209101019893").
		Equal(g.Region, "140622").
		Equal(g.Date, date).
		True(g.IsMale)

	g, err = Parse("140622209101019883")
	date, err = time.Parse(layout, "20910101")
	a.NotError(err)
	a.NotError(err).NotNil(g)
	a.Equal(g.Raw, "140622209101019883").
		Equal(g.Region, "140622").
		Equal(g.Date, date).
		False(g.IsMale)

	g, err = Parse("140622910101989")
	date, err = time.Parse(layout, "19910101")
	a.NotError(err)
	a.NotError(err).NotNil(g)
	a.Equal(g.Raw, "140622910101989").
		Equal(g.Region, "140622").
		Equal(g.Date, date).
		True(g.IsMale)

	g, err = Parse("140622910101988")
	date, err = time.Parse(layout, "19910101")
	a.NotError(err)
	a.NotError(err).NotNil(g)
	a.Equal(g.Raw, "140622910101988").
		Equal(g.Region, "140622").
		Equal(g.Date, date).
		False(g.IsMale)

	g, err = Parse("")
	a.Error(err).Nil(g)
}
