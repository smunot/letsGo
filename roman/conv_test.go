package roman_test

import (
	"testing"

	"github.com/smunot/letsGo/roman"
	"github.com/stretchr/testify/assert"
)

func TestToRoman1(t *testing.T) {
	// Go tuple using struct
	uu := []struct {
		number int
		glyph  string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{5, "V"},
	}

	for _, u := range uu {
		// e := u.glyph
		// n := roman.ToRoman(u.number)

		// // if e != n {
		// // 	t.Fatalf("Expecting %s but got %s", e, n)
		// // }

		assert.Equal(t, u.glyph, roman.ToRoman(u.number))

	}
}
