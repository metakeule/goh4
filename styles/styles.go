package styles

import (
	"fmt"
	. "github.com/metakeule/goh4"
)

const (
	// "background-attachment"=scroll | fixed | inherit
	SCROLL  = "scroll"
	FIXED   = "fixed"
	INHERIT = "inherit"

	// "background-position"=left | center | right | top | bottom | inherit
	LEFT   = "left"
	CENTER = "center"
	RIGHT  = "right"
	TOP    = "top"
	BOTTOM = "bottom"

	// "background-repeat"=repeat | repeat-x | repeat-y | no-repeat | inherit
	REPEAT    = "repeat"
	REPEAT_X  = "repeat-x"
	REPEAT_Y  = "repeat-y"
	NO_REPEAT = "no-repeat"

	// "border-collapse"=collapse | separate | inherit
	COLLAPSE = "collapse"
	SEPARATE = "separate"

	// "border-color"=<color> | inherit

	// "border-spacing"=inherit

	// "bottom"=<length> | <percentage> | auto | inherit
	AUTO = "auto"

	// "caption-side"=top | bottom | inherit

	// "clear"=none | left | right | both | inherit
	BOTH = "both"

	// "color"=<color> | inherit

	// "content"=normal | none | <uri> | open-quote | close-quote | no-open-quote | no-close-quote | inherit
	NORMAL         = "normal"
	OPEN_QUOTE     = "open-quote"
	CLOSE_QUOTE    = "close-quote"
	NO_OPEN_QUOTE  = "no-open-quote"
	NO_CLOSE_QUOTE = "no-close-quote"

	// "counter-increment"=none | inherit

	// "counter-reset"=none | inherit

	// "cursor"=<uri> | auto | crosshair | default | pointer | move | e-resize | ne-resize | nw-resize | n-resize | se-resize | sw-resize | s-resize | w-resize | text | wait | help | progress | inherit
	DEFAULT   = "default"
	CROSSHAIR = "crosshair"
	POINTER   = "pointer"
	MOVE      = "move"
	E_RESIZE  = "e-resize"
	NE_RESIZE = "ne-resize"
	NW_RESIZE = "nw-resize"
	N_RESIZE  = "n-resize"
	SE_RESIZE = "se-resize"
	SW_RESIZE = "sw-resize"
	S_RESIZE  = "s-resize"
	W_RESIZE  = "w-resize"
	TEXT      = "text"
	WAIT      = "wait"
	HELP      = "help"
	PROGRESS  = "progress"

	// "direction"=ltr | rtl | inherit
	LTR = "ltr"
	RTL = "rtl"

	// "display"=inline | block | list-item | inline-block | table | inline-table | table-row-group | table-header-group | table-footer-group | table-row | table-column-group | table-column | table-cell | table-caption | none | inherit
	INLINE             = "inline"
	BLOCK              = "block"
	LIST_ITEM          = "list-item"
	INLINE_BLOCK       = "inline-block"
	TABLE              = "table"
	INLINE_TABLE       = "inline-table"
	TABLE_ROW_GROUP    = "table-row-group "
	TABLE_HEADER_GROUP = "table-header-group"
	TABLE_FOOTER_GROUP = "table-footer-group"
	TABLE_ROW          = "table-row"
	TABLE_COLUMN_GROUP = "table-column-group"
	TABLE_COLUMN       = "table-column"
	TABLE_CELL         = "table-cell"
	TABLE_CAPTION      = "table-caption"

	// "empty-cells"=show | hide | inherit
	SHOW = "show"
	HIDE = "hide"

	// "orphans"=<integer> | inherit

	// "outline-color"=<color> | invert | inherit
	INVERT = "invert"

	// "outline-style"=<border-style> | inherit

	// "outline-width"=<border-width> | inherit

	// "outline"=<color> | <border-style> | <border-width> | inherit

	// "overflow"=visible | hidden | scroll | auto | inherit
	VISIBLE = "visible"
	HIDDEN  = "hidden"

	// "padding-top" "padding-right" "padding-bottom" "padding-left"=<padding-width> | inherit

	// "padding"=<padding-width> | inherit

	// "page-break-after"=auto | always | avoid | left | right | inherit
	ALWAYS = "always"
	AVOID  = "avoid"

	// "page-break-before"=auto | always | avoid | left | right | inherit

	// "page-break-inside"=avoid | auto | inherit

	// "position"=static | relative | absolute | fixed | inherit
	STATIC   = "static"
	RELATIVE = "relative"
	ABSOLUTE = "absolute"

	// "quotes"=none | inherit

	// "right"=<length> | <percentage> | auto | inherit

	// "table-layout"=auto | fixed | inherit

	// "text-align"=left | right | center | justify | inherit
	JUSTIFY = "justify"

	// "text-decoration"=none | underline | overline | line-through | blink | inherit | none
	UNDERLINE    = "underline"
	OVERLINE     = "overline"
	LINE_THROUGH = "line-through"
	BLINK        = "blink"

	// "text-indent"=<length> | <percentage> | inherit

	// "text-transform"=capitalize | uppercase | lowercase | none | inherit
	CAPITALIZE = "capitalize"
	UPPERCASE  = "uppercase"
	LOWERCASE  = "lowercase"

	//"top"=<length> | <percentage> | auto | inherit

	//"unicode-bidi"=normal | embed | bidi-override | inherit
	EMBED         = "embed"
	BIDI_OVERRIDE = "bidi-override"

	//"vertical-align"=baseline | sub | super | top | text-top | middle | bottom | text-bottom | <percentage> | <length> | inherit
	BASELINE   = "baseline"
	SUB        = "sub"
	SUPER      = "super"
	TEXTTOP    = "text-top"
	MIDDLE     = "middle"
	TEXTBOTTOM = "text-bottom"

	//"visibility"=visible | hidden | collapse | inherit

	//"white-space"=normal | pre | nowrap | pre-wrap | pre-line | inherit
	PRE      = "pre"
	NO_WRAP  = "nowrap"
	PRE_WRAP = "pre-wrap"
	PRE_LINE = "pre-line"

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
	CAPTION       = "caption"
	ICON          = "icon"
	MENU          = "menu"
	MESSAGE_BOX   = "message-box"
	SMALL_CAPTION = "small-caption"
	STATUS_BAR    = "status-bar"
	NONE          = "none"

	// "font-style"=normal | italic | oblique | inherit
	ITALIC  = "italic"
	OBLIQUE = "oblique"

	// "font-weight"=normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | inherit
	BOLD    = "bold"
	BOLDER  = "bolder"
	LIGHTER = "lighter"
	// "font-variant"=normal | small-caps | inherit
	SMALLCAPS = "small-caps"
	// "list-style-position"=inside | outside | inherit
	INSIDE  = "inside"
	OUTSIDE = "outside"

	// "list-style-type"=disc | circle | square | decimal | decimal-leading-zero | lower-roman | upper-roman | lower-greek | lower-latin | upper-latin | armenian | georgian | lower-alpha | upper-alpha | none | inherit
	DISC                 = "disc"
	CIRCLE               = "circle"
	SQUARE               = "square"
	DECIMAL              = "decimal"
	DECIMAL_LEADING_ZERO = "decimal-leading-zero"
	LOWER_ROMAN          = "lower-roman"
	UPPER_ROMAN          = "upper-roman"
	LOWER_GREEK          = "lower-greek"
	LOWER_LATIN          = "lower-latin"
	UPPER_LATIN          = "upper-latin"
	ARMENIAN             = "armenian"
	GEORGIAN             = "georgian"
	LOWER_ALPHA          = "lower-alpha"
	UPPER_ALPHA          = "upper-alpha"
)

// colors

const (
	GREEN  = "green"
	CYAN   = "cyan"
	BLACK  = "black"
	WHITE  = "white"
	BLUE   = "blue"
	RED    = "red"
	YELLOW = "yellow"
)

// "clip"=<shape> | auto | inherit
//func Shape(s string) Style  = s

// "background-color"=<color> | inherit
//func Color(c string) Style  = c

// "background-image"=<uri> | none | inherit
func URL(url string) string {
	return fmt.Sprintf(`url('%s')`, url)
}

// "border-style"=<border-style> | inherit
//func Style(s string) Style  = s

// "border-top" "border-right" "border-bottom" "border-left"=<border-width> | <border-style> | <color> | inherit
//func Width(w string) Style  = w

// "border-top-color" "border-right-color" "border-bottom-color" "border-left-color"=<color> | inherit

// "border-top-style" "border-right-style" "border-bottom-style" "border-left-style"=<border-style> | inherit

// "border-top-width" "border-right-width" "border-bottom-width" "border-left-width"=<border-width> | inherit

// "border-width"=<border-width> | inherit
func PERCENT(f float64) string { return fmt.Sprintf(`%v`, f) + "%" }
func PX(f float64) string      { return fmt.Sprintf(`%vpx`, f) }
func EM(f float64) string      { return fmt.Sprintf(`%vem`, f) }
func PT(f float64) string      { return fmt.Sprintf(`%vpt`, f) }

// type Color string

// "float"=left | right | none | inherit

// "font-family"=<generic-family>| inherit
// FontFamily(f string) Style  = f

// "font-size"=inherit

//Weight(i int) Style      = fmt.Sprintf(`%v`, i)

// "height"=<length> | <percentage> | auto | inherit

// "left"=<length> | <percentage> | auto | inherit

// "letter-spacing"=normal | <length> | inherit

// "line-height"=normal | <number> | <length> | <percentage> | inherit

// "list-style-image"=<uri> | none | inherit

// "list-style"=disc | circle | square | decimal | decimal-leading-zero | lower-roman | upper-roman | lower-greek | lower-latin | upper-latin | armenian | georgian | lower-alpha | upper-alpha | none | inside | outside | <uri> | inherit

// "margin-right" "margin-left"=<margin-width> | inherit

// "margin-top" "margin-bottom"=<margin-width> | inherit

// "margin"=<margin-width> | inherit

// "max-height"=<length> | <percentage> | none | inherit

// "max-width"=<length> | <percentage> | none | inherit

// "min-height"=<length> | <percentage> | inherit

// "min-width"=<length> | <percentage> | inherit

// "opacity"=<number> | inherit

// "background-attachment"=scroll | fixed | inherit
func BackgroundAttachment(s string) (ba Style) {
	return Style{Property: "background-attachment", Value: s}
}

func Position(s string) (ba Style) { return Style{Property: "position", Value: s} }

// "background-position"=left | center | right | top | bottom | inherit
func BackgroundPosition(s string) (ba Style) { return Style{Property: "background-position", Value: s} }

// "background-color"=<color> | inherit
func BackgroundColor(s string) (ba Style) { return Style{Property: "background-color", Value: s} }

// "background-image"=<uri> | none | inherit
func BackgroundImage(s string) (ba Style) { return Style{Property: "background-image", Value: s} }

// "background-repeat"=repeat | repeat-x | repeat-y | no-repeat | inherit
func BackgroundRepeat(s string) (ba Style) { return Style{Property: "background-repeat", Value: s} }

// "border-collapse"=collapse | separate | inherit
func BorderCollapse(s string) (ba Style) { return Style{Property: "border-collapse", Value: s} }

// "border-color"=<color> | inherit
func BorderColor(s string) (ba Style) { return Style{Property: "border-color", Value: s} }

// "border-spacing"=inherit
func BorderSpacing(s string) (ba Style) { return Style{Property: "border-spacing", Value: s} }

// "border-style"=<border-style> | inherit
func BorderStyle(s string) (ba Style) { return Style{Property: "border-style", Value: s} }

// "border-top" "border-right" "border-bottom" "border-left"=<border-width> | <border-style> | <color> | inherit
func BorderTop(s string) (ba Style)    { return Style{Property: "border-top", Value: s} }
func BorderRight(s string) (ba Style)  { return Style{Property: "border-right", Value: s} }
func BorderBottom(s string) (ba Style) { return Style{Property: "border-bottom", Value: s} }
func BorderLeft(s string) (ba Style)   { return Style{Property: "border-left", Value: s} }

// "border-top-color" "border-right-color" "border-bottom-color" "border-left-color"=<color> | inherit
func BorderTopColor(s string) (ba Style) { return Style{Property: "border-top-color", Value: s} }

// "border-top-style" "border-right-style" "border-bottom-style" "border-left-style"=<border-style> | inherit
func BorderTopStyle(s string) (ba Style) { return Style{Property: "border-top-style", Value: s} }

// "border-top-width" "border-right-width" "border-bottom-width" "border-left-width"=<border-width> | inherit
func BorderTopWidth(s string) (ba Style)    { return Style{Property: "border-top-width", Value: s} }
func BorderRightColor(s string) (ba Style)  { return Style{Property: "border-right-color", Value: s} }
func BorderRightStyle(s string) (ba Style)  { return Style{Property: "border-right-style", Value: s} }
func BorderRightWidth(s string) (ba Style)  { return Style{Property: "border-right-width", Value: s} }
func BorderBottomColor(s string) (ba Style) { return Style{Property: "border-bottom-color", Value: s} }
func BorderBottomStyle(s string) (ba Style) { return Style{Property: "border-bottom-style", Value: s} }
func BorderBottomWidth(s string) (ba Style) { return Style{Property: "border-bottom-width", Value: s} }
func BorderLeftColor(s string) (ba Style)   { return Style{Property: "border-left-color", Value: s} }
func BorderLeftStyle(s string) (ba Style)   { return Style{Property: "border-left-style", Value: s} }
func BorderLeftWidth(s string) (ba Style)   { return Style{Property: "border-left-width", Value: s} }

// "border-width"=<border-width> | inherit
func BorderWidth(s string) (ba Style) { return Style{Property: "border-width", Value: s} }

// "bottom"=<length> | <percentage> | auto | inherit
func Bottom(s string) (ba Style) { return Style{Property: "bottom", Value: s} }

// "caption-side"=top | bottom | inherit
func CaptionSide(s string) (ba Style) { return Style{Property: "caption-side", Value: s} }

// "clear"=none | left | right | both | inherit
func Clear(s string) (ba Style) { return Style{Property: "clear", Value: s} }

// "clip"=<shape> | auto | inherit
func Clip(s string) (ba Style) { return Style{Property: "clip", Value: s} }

// "color"=<color> | inherit
func Color(s string) (ba Style) { return Style{Property: "color", Value: s} }

// "content"=normal | none | <uri> | open-quote | close-quote | no-open-quote | no-close-quote | inherit
func Content(s string) (ba Style) { return Style{Property: "content", Value: s} }

// "counter-increment"=none | inherit
func CounterIncrement(s string) (ba Style) { return Style{Property: "counter-increment", Value: s} }

// "counter-reset"=none | inherit
func CounterReset(s string) (ba Style) { return Style{Property: "counter-reset", Value: s} }

// "cursor"=<uri> | auto | crosshair | default | pointer | move | e-resize | ne-resize | nw-resize | n-resize | se-resize | sw-resize | s-resize | w-resize | text | wait | help | progress | inherit
func Cursor(s string) (ba Style) { return Style{Property: "cursor", Value: s} }

// "direction"=ltr | rtl | inherit
func Direction(s string) (ba Style) { return Style{Property: "direction", Value: s} }

// "display"=inline | block | list-item | inline-block | table | inline-table | table-row-group | table-header-group | table-footer-group | table-row | table-column-group | table-column | table-cell | table-caption | none | inherit
func Display(s string) (ba Style) { return Style{Property: "display", Value: s} }

// "empty-cells"=show | hide | inherit
func EmptyCells(s string) (ba Style) { return Style{Property: "empty-cells", Value: s} }

// "float"=left | right | none | inherit
func Float(s string) (ba Style) { return Style{Property: "float", Value: s} }

// "font-family"=<generic-family>| inherit
func FontFamily(s string) (ba Style) { return Style{Property: "font-family", Value: s} }

// "font-size"=inherit
func FontSize(s string) (ba Style) { return Style{Property: "font-size", Value: s} }

// "font-style"=normal | italic | oblique | inherit
func FontStyle(s string) (ba Style) { return Style{Property: "font-style", Value: s} }

// "font-variant"=normal | small-caps | inherit
func FontVariant(s string) (ba Style) { return Style{Property: "font-variant", Value: s} }

// "font-weight"=normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | inherit
func FontWeight(s string) (ba Style) { return Style{Property: "font-weight", Value: s} }

// "height"=<length> | <percentage> | auto | inherit
func Height(s string) (ba Style) { return Style{Property: "height", Value: s} }

// "left"=<length> | <percentage> | auto | inherit
func Left(s string) (ba Style) { return Style{Property: "left", Value: s} }

// "letter-spacing"=normal | <length> | inherit
func LetterSpacing(s string) (ba Style) { return Style{Property: "letter-spacing", Value: s} }

// "line-height"=normal | <number> | <length> | <percentage> | inherit
func LineHeight(s string) (ba Style) { return Style{Property: "line-height", Value: s} }

// "list-style-image"=<uri> | none | inherit
func ListStyleImage(s string) (ba Style) { return Style{Property: "list-style-image", Value: s} }

// "list-style-position"=inside | outside | inherit
func ListStylePosition(s string) (ba Style) { return Style{Property: "list-style-position", Value: s} }

// "list-style-type"=disc | circle | square | decimal | decimal-leading-zero | lower-roman | upper-roman | lower-greek | lower-latin | upper-latin | armenian | georgian | lower-alpha | upper-alpha | none | inherit
func ListStyleType(s string) (ba Style) { return Style{Property: "list-style-type", Value: s} }

// "list-style"=disc | circle | square | decimal | decimal-leading-zero | lower-roman | upper-roman | lower-greek | lower-latin | upper-latin | armenian | georgian | lower-alpha | upper-alpha | none | inside | outside | <uri> | inherit
func ListStyle(s string) (ba Style) { return Style{Property: "list-style", Value: s} }

// "margin-right" "margin-left"=<margin-width> | inherit
func MarginRight(s string) (ba Style) { return Style{Property: "margin-right", Value: s} }
func MarginLeft(s string) (ba Style)  { return Style{Property: "margin-left", Value: s} }

// "margin-top" "margin-bottom"=<margin-width> | inherit
func MarginTop(s string) (ba Style)    { return Style{Property: "margin-top", Value: s} }
func MarginBottom(s string) (ba Style) { return Style{Property: "margin-bottom", Value: s} }

// "margin"=<margin-width> | inherit
func Margin(s string) (ba Style) { return Style{Property: "margin", Value: s} }

// "max-height"=<length> | <percentage> | none | inherit
func MaxHeight(s string) (ba Style) { return Style{Property: "max-height", Value: s} }

// "max-width"=<length> | <percentage> | none | inherit
func MaxWidth(s string) (ba Style) { return Style{Property: "max-width", Value: s} }

// "min-height"=<length> | <percentage> | inherit
func MinHeight(s string) (ba Style) { return Style{Property: "min-height", Value: s} }

// "min-width"=<length> | <percentage> | inherit
func MinWidth(s string) (ba Style) { return Style{Property: "min-width", Value: s} }

// "opacity"=<number> | inherit
func Opacity(s string) (ba Style) { return Style{Property: "opacity", Value: s} }

// "orphans"=<integer> | inherit
func Orphans(s string) (ba Style) { return Style{Property: "orphans", Value: s} }

// "outline-color"=<color> | invert | inherit
func OutlineColor(s string) (ba Style) { return Style{Property: "outline-color", Value: s} }

// "outline-style"=<border-style> | inherit
func OutlineStyle(s string) (ba Style) { return Style{Property: "outline-style", Value: s} }

// "outline-width"=<border-width> | inherit
func OutlineWidth(s string) (ba Style) { return Style{Property: "outline-width", Value: s} }

// "outline"=<color> | <border-style> | <border-width> | inherit
func Outline(s string) (ba Style) { return Style{Property: "outline", Value: s} }

// "overflow"=visible | hidden | scroll | auto | inherit
func Overflow(s string) (ba Style) { return Style{Property: "min-width", Value: s} }

// "padding-top" "padding-right" "padding-bottom" "padding-left"=<padding-width> | inherit
func PaddingTop(s string) (ba Style)    { return Style{Property: "padding-top", Value: s} }
func PaddingRight(s string) (ba Style)  { return Style{Property: "padding-right", Value: s} }
func PaddingBottom(s string) (ba Style) { return Style{Property: "padding-bottom", Value: s} }
func PaddingLeft(s string) (ba Style)   { return Style{Property: "padding-left", Value: s} }

// "padding"=<padding-width> | inherit
func Padding(s string) (ba Style) { return Style{Property: "padding", Value: s} }

// "page-break-after"=auto | always | avoid | left | right | inherit
func PageBreakAfter(s string) (ba Style) { return Style{Property: "page-break-after", Value: s} }

// "page-break-before"=auto | always | avoid | left | right | inherit
func PageBreakBefore(s string) (ba Style) { return Style{Property: "page-break-before", Value: s} }

// "page-break-inside"=avoid | auto | inherit
func PageBreakInside(s string) (ba Style) { return Style{Property: "page-break-inside", Value: s} }

// "position"=static | relative | absolute | fixed | inherit
func position(s string) (ba Style) { return Style{Property: "position", Value: s} }

// "quotes"=none | inherit
func Quotes(s string) (ba Style) { return Style{Property: "quotes", Value: s} }

// "right"=<length> | <percentage> | auto | inherit
func Right(s string) (ba Style) { return Style{Property: "right", Value: s} }

func Resize(s string) (ba Style) { return Style{Property: "resize", Value: s} }

// "table-layout"=auto | fixed | inherit
func TableLayout(s string) (ba Style) { return Style{Property: "table-layout", Value: s} }

// "text-align"=left | right | center | justify | inherit
func TextAlign(s string) (ba Style) { return Style{Property: "text-align", Value: s} }

// "text-decoration"=none | underline | overline | line-through | blink | inherit | none
func TextDecoration(s string) (ba Style) { return Style{Property: "text-decoration", Value: s} }

// "text-indent"=<length> | <percentage> | inherit
func TextIndent(s string) (ba Style) { return Style{Property: "text-indent", Value: s} }

// "text-transform"=capitalize | uppercase | lowercase | none | inherit
func TextTransform(s string) (ba Style) { return Style{Property: "text-transform", Value: s} }

//"top"=<length> | <percentage> | auto | inherit
func Top(s string) (ba Style) { return Style{Property: "top", Value: s} }

//"unicode-bidi"=normal | embed | bidi-override | inherit
func UnicodeBidi(s string) (ba Style) { return Style{Property: "unicode-bidi", Value: s} }

//"vertical-align"=baseline | sub | super | top | text-top | middle | bottom | text-bottom | <percentage> | <length> | inherit
func VerticalAlign(s string) (ba Style) { return Style{Property: "vertical-align", Value: s} }

//"visibility"=visible | hidden | collapse | inherit
func Visibility(s string) (ba Style) { return Style{Property: "visibility", Value: s} }

//"white-space"=normal | pre | nowrap | pre-wrap | pre-line | inherit
func WhiteSpace(s string) (ba Style) { return Style{Property: "white-space", Value: s} }

//"widows"=<integer> | inherit
func Widows(s string) (ba Style) { return Style{Property: "widows", Value: s} }

//"width"=<length> | <percentage> | auto | inherit
func Width(s string) (ba Style) { return Style{Property: "width", Value: s} }

//"word-spacing"=normal | <length> | inherit
func WordSpacing(s string) (ba Style) { return Style{Property: "word-spacing", Value: s} }

//"z-index"=auto | <integer> | inherit
func ZIndex(s string) (ba Style) { return Style{Property: "z-index", Value: s} }

//"background-clip"=<box>
func BackgroundClip(s string) (ba Style) { return Style{Property: "background-clip", Value: s} }

//"background-origin"=<box>
func BackgroundOrigin(s string) (ba Style) { return Style{Property: "background-origin", Value: s} }

//"background-size"=<bg-size>
func BackgroundSize(s string) (ba Style) { return Style{Property: "background-size", Value: s} }

// "border"=<border-width> | <border-style> | <color>
func Border(s string) (ba Style) { return Style{Property: "border", Value: s} }

// "border-image"=<border-image-source> | <border-image-slice> | <border-image-width> | <border-image-width> | <border-image-outset> | <border-image-repeat>
func BorderImage(s string) (ba Style) { return Style{Property: "border-image", Value: s} }

// "border-image-outset"=<length> | <number>
func BorderImageOutset(s string) (ba Style) { return Style{Property: "border-image-outset", Value: s} }

// "border-image-repeat"=stretch | repeat | round | space
func BorderImageRepeat(s string) (ba Style) { return Style{Property: "border-image-repeat", Value: s} }

// "border-image-slice"=<number> | <percentage>
func BorderImageSlice(s string) (ba Style) { return Style{Property: "border-image-slice", Value: s} }

// "border-image-source"=none | <image>
func BorderImageSource(s string) (ba Style) { return Style{Property: "border-image-source", Value: s} }

// "border-image-width"=<length> | <percentage> | <number> | auto
func BorderImageWidth(s string) (ba Style) { return Style{Property: "border-image-width", Value: s} }

// "border-radius"=<length> | <percentage>
func BorderRadius(s string) (ba Style) { return Style{Property: "border-radius", Value: s} }

// "border-top-left-radius" "border-top-right-radius" "border-bottom-right-radius" "border-bottom-left-radius"=<length> | <percentage>
func BorderTopLeftRadius(s string) (ba Style) {
	return Style{Property: "border-top-left-radius", Value: s}
}

// "box-decoration-break"=slice | clone
func BoxDecorationBreak(s string) (ba Style) { return Style{Property: "box-decoration-break", Value: s} }

// "box-shadow"=none | <shadow> | none
func BoxShadow(s string) (ba Style) { return Style{Property: "box-shadow", Value: s} }

// "background"=<color> | <uri> | repeat | repeat-x | repeat-y | no-repeat | scroll | fixed | left | center | right | top | bottom | inherit
func Background(s string) (ba Style) { return Style{Property: "background", Value: s} }

// "font"=normal | italic | oblique | normal | small-caps | normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | normal | <generic-family> | caption | icon | menu | message-box | small-caption | status-bar | inherit
func Font(s string) (ba Style) { return Style{Property: "font", Value: s} }

// caption | icon | menu | message-box | small-caption | status-bar
