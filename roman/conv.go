package roman

type (
	glyph struct {
		limit int
		glyph string
	}

	glyphs []glyph
)

var romanGlyphs = glyphs{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRoman converts int num to Roman
func ToRoman(n int) (r string) {

	for _, g := range romanGlyphs {
		for ; n >= g.limit; n -= g.limit {
			r += g.glyph
		}
	}

	return
}
