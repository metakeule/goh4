package goh4

type Style struct{ Property, Value string }

func (ø Style) Val(v string) Style { ø.Value = v; return ø }
func (ø Style) String() string     { return ø.Property + ": " + ø.Value + ";" }
func (ø Style) Style() string      { return ø.String() }

func Styles(styles ...Style) (s []Style) {
	return styles
}
