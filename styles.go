package goh4

import (
	"fmt"
)

type Style struct{ Property, Value string }

// type Color string

func (ø Style) Val(v string) Style  { ø.Value = v; return ø }
func (ø Style) Var(v scssVar) Style { ø.Value = v.Name(); return ø }
func (ø Style) Call(fn string, args ...string) Style {
	ø.Value = ScssCall(fn, args...).String()
	return ø
}

func (ø Style) String() string   { return ø.Property + ": " + ø.Value }
func (ø Style) StyleCmd() string { return ø.String() }

// "background-attachment"=scroll | fixed | inherit
func (ø Style) Scroll() Style  { ø.Value = "scroll"; return ø }
func (ø Style) Fixed() Style   { ø.Value = "fixed"; return ø }
func (ø Style) Inherit() Style { ø.Value = "inherit"; return ø }

// "background-position"=left | center | right | top | bottom | inherit
func (ø Style) Left() Style   { ø.Value = "left"; return ø }
func (ø Style) Center() Style { ø.Value = "center"; return ø }
func (ø Style) Right() Style  { ø.Value = "right"; return ø }
func (ø Style) Top() Style    { ø.Value = "top"; return ø }
func (ø Style) Bottom() Style { ø.Value = "bottom"; return ø }

// "background-color"=<color> | inherit
func (ø Style) Color(c string) Style { ø.Value = c; return ø }

// "background-image"=<uri> | none | inherit
func (ø Style) None() Style          { ø.Value = "none"; return ø }
func (ø Style) Url(url string) Style { ø.Value = fmt.Sprintf(`url(%s)`, url); return ø }

// "background-repeat"=repeat | repeat-x | repeat-y | no-repeat | inherit
func (ø Style) Repeat() Style   { ø.Value = "repeat"; return ø }
func (ø Style) RepeatX() Style  { ø.Value = "repeat-x"; return ø }
func (ø Style) RepeatY() Style  { ø.Value = "repeat-y"; return ø }
func (ø Style) NoRepeat() Style { ø.Value = "no-repeat"; return ø }

// "border-collapse"=collapse | separate | inherit
func (ø Style) Collapse() Style { ø.Value = "collapse"; return ø }
func (ø Style) Separate() Style { ø.Value = "separate"; return ø }

// "border-color"=<color> | inherit

// "border-spacing"=inherit

// "border-style"=<border-style> | inherit
func (ø Style) Style(s string) Style { ø.Value = s; return ø }

// "border-top" "border-right" "border-bottom" "border-left"=<border-width> | <border-style> | <color> | inherit
func (ø Style) Width(w string) Style { ø.Value = w; return ø }

// "border-top-color" "border-right-color" "border-bottom-color" "border-left-color"=<color> | inherit

// "border-top-style" "border-right-style" "border-bottom-style" "border-left-style"=<border-style> | inherit

// "border-top-width" "border-right-width" "border-bottom-width" "border-left-width"=<border-width> | inherit

// "border-width"=<border-width> | inherit
func (ø Style) Percent(f float64) Style { ø.Value = fmt.Sprintf(`%v\%`, f); return ø }
func (ø Style) Px(f float64) Style      { ø.Value = fmt.Sprintf(`%vpx`, f); return ø }
func (ø Style) Em(f float64) Style      { ø.Value = fmt.Sprintf(`%vem`, f); return ø }
func (ø Style) Pt(f float64) Style      { ø.Value = fmt.Sprintf(`%vpt`, f); return ø }

// "bottom"=<length> | <percentage> | auto | inherit
func (ø Style) Auto() Style { ø.Value = "auto"; return ø }

// "caption-side"=top | bottom | inherit

// "clear"=none | left | right | both | inherit
func (ø Style) Both() Style { ø.Value = "both"; return ø }

// "clip"=<shape> | auto | inherit
func (ø Style) Shape(s string) Style { ø.Value = s; return ø }

// "color"=<color> | inherit

// "content"=normal | none | <uri> | open-quote | close-quote | no-open-quote | no-close-quote | inherit
func (ø Style) Normal() Style       { ø.Value = "normal"; return ø }
func (ø Style) OpenQuote() Style    { ø.Value = "open-quote"; return ø }
func (ø Style) CloseQuote() Style   { ø.Value = "close-quote"; return ø }
func (ø Style) NoOpenQuote() Style  { ø.Value = "no-open-quote"; return ø }
func (ø Style) NoCloseQuote() Style { ø.Value = "no-close-quote"; return ø }

// "counter-increment"=none | inherit

// "counter-reset"=none | inherit

// "cursor"=<uri> | auto | crosshair | default | pointer | move | e-resize | ne-resize | nw-resize | n-resize | se-resize | sw-resize | s-resize | w-resize | text | wait | help | progress | inherit
func (ø Style) Default() Style   { ø.Value = "default"; return ø }
func (ø Style) Crosshair() Style { ø.Value = "crosshair"; return ø }
func (ø Style) Pointer() Style   { ø.Value = "pointer"; return ø }
func (ø Style) Move() Style      { ø.Value = "move"; return ø }
func (ø Style) EResize() Style   { ø.Value = "e-resize"; return ø }
func (ø Style) NEResize() Style  { ø.Value = "ne-resize"; return ø }
func (ø Style) NWResize() Style  { ø.Value = "nw-resize"; return ø }
func (ø Style) NResize() Style   { ø.Value = "n-resize"; return ø }
func (ø Style) SEResize() Style  { ø.Value = "se-resize"; return ø }
func (ø Style) SWResize() Style  { ø.Value = "sw-resize"; return ø }
func (ø Style) SResize() Style   { ø.Value = "s-resize"; return ø }
func (ø Style) WResize() Style   { ø.Value = "w-resize"; return ø }
func (ø Style) Text() Style      { ø.Value = "text"; return ø }
func (ø Style) Wait() Style      { ø.Value = "wait"; return ø }
func (ø Style) Help() Style      { ø.Value = "help"; return ø }
func (ø Style) Progress() Style  { ø.Value = "progress"; return ø }

// "direction"=ltr | rtl | inherit
func (ø Style) Ltr() Style { ø.Value = "ltr"; return ø }
func (ø Style) Rtl() Style { ø.Value = "rtl"; return ø }

// "display"=inline | block | list-item | inline-block | table | inline-table | table-row-group | table-header-group | table-footer-group | table-row | table-column-group | table-column | table-cell | table-caption | none | inherit
func (ø Style) Inline() Style           { ø.Value = "inline"; return ø }
func (ø Style) Block() Style            { ø.Value = "block"; return ø }
func (ø Style) ListItem() Style         { ø.Value = "list-item"; return ø }
func (ø Style) InlineBlock() Style      { ø.Value = "inline-block"; return ø }
func (ø Style) Table() Style            { ø.Value = "table"; return ø }
func (ø Style) InlineTable() Style      { ø.Value = "inline-table"; return ø }
func (ø Style) TableRowGroup() Style    { ø.Value = "table-row-group "; return ø }
func (ø Style) TableHeaderGroup() Style { ø.Value = "table-header-group"; return ø }
func (ø Style) TableFooterGroup() Style { ø.Value = "table-footer-group"; return ø }
func (ø Style) TableRow() Style         { ø.Value = "table-row"; return ø }
func (ø Style) TableColumnGroup() Style { ø.Value = "table-column-group"; return ø }
func (ø Style) TableColumn() Style      { ø.Value = "table-column"; return ø }
func (ø Style) TableCell() Style        { ø.Value = "table-cell"; return ø }
func (ø Style) TableCaption() Style     { ø.Value = "table-caption"; return ø }

// "empty-cells"=show | hide | inherit
func (ø Style) Show() Style { ø.Value = "show"; return ø }
func (ø Style) Hide() Style { ø.Value = "hide"; return ø }

// "float"=left | right | none | inherit

// "font-family"=<generic-family>| inherit
func (ø Style) FontFamily(f string) Style { ø.Value = f; return ø }

// "font-size"=inherit

// "font-style"=normal | italic | oblique | inherit
func (ø Style) Italic(f string) Style  { ø.Value = "italic"; return ø }
func (ø Style) Oblique(f string) Style { ø.Value = "oblique"; return ø }

// "font-variant"=normal | small-caps | inherit
func (ø Style) SmallCaps(f string) Style { ø.Value = "small-caps"; return ø }

// "font-weight"=normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | inherit
func (ø Style) Bold(f string) Style    { ø.Value = "bold"; return ø }
func (ø Style) Bolder(f string) Style  { ø.Value = "bolder"; return ø }
func (ø Style) Lighter(f string) Style { ø.Value = "lighter"; return ø }
func (ø Style) Weight(i int) Style     { ø.Value = fmt.Sprintf(`%v`, i); return ø }

// "height"=<length> | <percentage> | auto | inherit

// "left"=<length> | <percentage> | auto | inherit

// "letter-spacing"=normal | <length> | inherit

// "line-height"=normal | <number> | <length> | <percentage> | inherit

// "list-style-image"=<uri> | none | inherit

// "list-style-position"=inside | outside | inherit
func (ø Style) Inside(f string) Style  { ø.Value = "inside"; return ø }
func (ø Style) Outside(f string) Style { ø.Value = "outside"; return ø }

// "list-style-type"=disc | circle | square | decimal | decimal-leading-zero | lower-roman | upper-roman | lower-greek | lower-latin | upper-latin | armenian | georgian | lower-alpha | upper-alpha | none | inherit
func (ø Style) Disc(f string) Style               { ø.Value = "disc"; return ø }
func (ø Style) Circle(f string) Style             { ø.Value = "circle"; return ø }
func (ø Style) Square(f string) Style             { ø.Value = "square"; return ø }
func (ø Style) Decimal(f string) Style            { ø.Value = "decimal"; return ø }
func (ø Style) DecimalLeadingZero(f string) Style { ø.Value = "decimal-leading-zero"; return ø }
func (ø Style) LowerRoman(f string) Style         { ø.Value = "lower-roman"; return ø }
func (ø Style) upperRoman(f string) Style         { ø.Value = "upper-roman"; return ø }
func (ø Style) LowerGreek(f string) Style         { ø.Value = "lower-greek"; return ø }
func (ø Style) LowerLatin(f string) Style         { ø.Value = "lower-latin"; return ø }
func (ø Style) UpperLatin(f string) Style         { ø.Value = "upper-latin"; return ø }
func (ø Style) Armenian(f string) Style           { ø.Value = "armenian"; return ø }
func (ø Style) Georgian(f string) Style           { ø.Value = "georgian"; return ø }
func (ø Style) LowerAlpha(f string) Style         { ø.Value = "lower-alpha"; return ø }
func (ø Style) UpperAlpha(f string) Style         { ø.Value = "upper-alpha"; return ø }

// "list-style"=disc | circle | square | decimal | decimal-leading-zero | lower-roman | upper-roman | lower-greek | lower-latin | upper-latin | armenian | georgian | lower-alpha | upper-alpha | none | inside | outside | <uri> | inherit

// "margin-right" "margin-left"=<margin-width> | inherit

// "margin-top" "margin-bottom"=<margin-width> | inherit

// "margin"=<margin-width> | inherit

// "max-height"=<length> | <percentage> | none | inherit

// "max-width"=<length> | <percentage> | none | inherit

// "min-height"=<length> | <percentage> | inherit

// "min-width"=<length> | <percentage> | inherit

// "opacity"=<number> | inherit
func (ø Style) Number(f float64) Style { ø.Value = fmt.Sprintf(`%v\%`, f); return ø }

// "orphans"=<integer> | inherit

// "outline-color"=<color> | invert | inherit
func (ø Style) Invert() Style { ø.Value = "invert"; return ø }

// "outline-style"=<border-style> | inherit

// "outline-width"=<border-width> | inherit

// "outline"=<color> | <border-style> | <border-width> | inherit

// "overflow"=visible | hidden | scroll | auto | inherit
func (ø Style) Visible() Style { ø.Value = "visible"; return ø }
func (ø Style) Hidden() Style  { ø.Value = "hidden"; return ø }

// "padding-top" "padding-right" "padding-bottom" "padding-left"=<padding-width> | inherit

// "padding"=<padding-width> | inherit

// "page-break-after"=auto | always | avoid | left | right | inherit
func (ø Style) Always() Style { ø.Value = "always"; return ø }
func (ø Style) Avoid() Style  { ø.Value = "avoid"; return ø }

// "page-break-before"=auto | always | avoid | left | right | inherit

// "page-break-inside"=avoid | auto | inherit

// "position"=static | relative | absolute | fixed | inherit
func (ø Style) Static() Style   { ø.Value = "static"; return ø }
func (ø Style) Relative() Style { ø.Value = "relative"; return ø }
func (ø Style) Absolute() Style { ø.Value = "absolute"; return ø }

// "quotes"=none | inherit

// "right"=<length> | <percentage> | auto | inherit

// "table-layout"=auto | fixed | inherit

// "text-align"=left | right | center | justify | inherit
func (ø Style) Justify() Style { ø.Value = "justify"; return ø }

// "text-decoration"=none | underline | overline | line-through | blink | inherit | none
func (ø Style) Underline() Style   { ø.Value = "underline"; return ø }
func (ø Style) Overline() Style    { ø.Value = "overline"; return ø }
func (ø Style) LineThrough() Style { ø.Value = "line-through"; return ø }
func (ø Style) Blink() Style       { ø.Value = "blink"; return ø }

// "text-indent"=<length> | <percentage> | inherit

// "text-transform"=capitalize | uppercase | lowercase | none | inherit
func (ø Style) Capitalize() Style { ø.Value = "capitalize"; return ø }
func (ø Style) Uppercase() Style  { ø.Value = "uppercase"; return ø }
func (ø Style) Lowercase() Style  { ø.Value = "lowercase"; return ø }

//"top"=<length> | <percentage> | auto | inherit

//"unicode-bidi"=normal | embed | bidi-override | inherit
func (ø Style) Embed() Style        { ø.Value = "embed"; return ø }
func (ø Style) BidiOverride() Style { ø.Value = "bidi-override"; return ø }

//"vertical-align"=baseline | sub | super | top | text-top | middle | bottom | text-bottom | <percentage> | <length> | inherit
func (ø Style) Baseline() Style   { ø.Value = "baseline"; return ø }
func (ø Style) Sub() Style        { ø.Value = "sub"; return ø }
func (ø Style) Super() Style      { ø.Value = "super"; return ø }
func (ø Style) TextTop() Style    { ø.Value = "text-top"; return ø }
func (ø Style) Middle() Style     { ø.Value = "middle"; return ø }
func (ø Style) TextBottom() Style { ø.Value = "text-bottom"; return ø }

//"visibility"=visible | hidden | collapse | inherit

//"white-space"=normal | pre | nowrap | pre-wrap | pre-line | inherit
func (ø Style) Pre() Style     { ø.Value = "pre"; return ø }
func (ø Style) Nowrap() Style  { ø.Value = "nowrap"; return ø }
func (ø Style) PreWrap() Style { ø.Value = "pre-wrap"; return ø }
func (ø Style) PreLine() Style { ø.Value = "pre-line"; return ø }

//"widows"=<integer> | inherit

//"width"=<length> | <percentage> | auto | inherit

//"word-spacing"=normal | <length> | inherit

//"z-index"=auto | <integer> | inherit

//"background-clip"=<box>

//"background-origin"=<box>

//"background-size"=<bg-size>

// "border"=<border-width> | <border-style> | <color>

// "border-image"=<border-image-source> | <border-image-slice> | <border-image-width> | <border-image-width> | <border-image-outset> | <border-image-repeat>

// "border-image-outset"=<length> | <number>

// "border-image-repeat"=stretch | repeat | round | space

// "border-image-slice"=<number> | <percentage>

// "border-image-source"=none | <image>

// "border-image-width"=<length> | <percentage> | <number> | auto

// "border-radius"=<length> | <percentage>

// "border-top-left-radius" "border-top-right-radius" "border-bottom-right-radius" "border-bottom-left-radius"=<length> | <percentage>

// "box-decoration-break"=slice | clone

// "box-shadow"=none | <shadow> | none

// "background"=<color> | <uri> | repeat | repeat-x | repeat-y | no-repeat | scroll | fixed | left | center | right | top | bottom | inherit

// "font"=normal | italic | oblique | normal | small-caps | normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | normal | <generic-family> | caption | icon | menu | message-box | small-caption | status-bar | inherit

// caption | icon | menu | message-box | small-caption | status-bar
func (ø Style) Caption() Style      { ø.Value = "caption"; return ø }
func (ø Style) Icon() Style         { ø.Value = "icon"; return ø }
func (ø Style) Menu() Style         { ø.Value = "menu"; return ø }
func (ø Style) MessageBox() Style   { ø.Value = "message-box"; return ø }
func (ø Style) SmallCaption() Style { ø.Value = "small-caption"; return ø }
func (ø Style) StatusBar() Style    { ø.Value = "status-bar"; return ø }
