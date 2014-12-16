package styl

import (
	"fmt"
	. "gopkg.in/metakeule/goh4.v5"
	"strings"
)

// values
const (
	Scroll_             = "scroll" // "background-attachment"=scroll | fixed | inherit
	Fixed_              = "fixed"
	Inherit_            = "inherit"
	Left_               = "left" // "background-position"=left | center | right | top | bottom | inherit
	Center_             = "center"
	Right_              = "right"
	Top_                = "top"
	Bottom_             = "bottom"
	Repeat_             = "repeat" // "background-repeat"=repeat | repeat-x | repeat-y | no-repeat | inherit
	RepeatX_            = "repeat-x"
	RepeatY_            = "repeat-y"
	NoRepeat_           = "no-repeat"
	Collapse_           = "collapse" // "border-collapse"=collapse | separate | inherit
	Separate_           = "separate"
	Auto_               = "auto"
	Both_               = "both"   // "clear"=none | left | right | both | inherit	
	Normal_             = "normal" // "content"=normal | none | <uri> | open-quote | close-quote | no-open-quote | no-close-quote | inherit
	OpenQuote_          = "open-quote"
	CloseQuote_         = "close-quote"
	NoOpenQuote_        = "no-open-quote"
	NoCloseQuote_       = "no-close-quote"
	Default_            = "default" // "cursor"=<uri> | auto | crosshair | default | pointer | move | e-resize | ne-resize | nw-resize | n-resize | se-resize | sw-resize | s-resize | w-resize | text | wait | help | progress | inherit
	Crosshair_          = "crosshair"
	Pointer_            = "pointer"
	Move_               = "move"
	E_Resize_           = "e-resize"
	NE_Resize_          = "ne-resize"
	NW_Resize_          = "nw-resize"
	N_Resize_           = "n-resize"
	SE_Resize_          = "se-resize"
	SW_Resize_          = "sw-resize"
	S_Resize_           = "s-resize"
	W_Resize_           = "w-resize"
	Text_               = "text"
	Wait_               = "wait"
	Help_               = "help"
	Progress_           = "progress"
	Ltr_                = "ltr" // "direction"=ltr | rtl | inherit
	Rtl_                = "rtl"
	Inline_             = "inline" // "display"=inline | block | list-item | inline-block | table | inline-table | table-row-group | table-header-group | table-footer-group | table-row | table-column-group | table-column | table-cell | table-caption | none | inherit
	Block_              = "block"
	ListItem_           = "list-item"
	InlineBlock_        = "inline-block"
	Table_              = "table"
	InlineTable_        = "inline-table"
	TableRowGroup_      = "table-row-group "
	TableHeaderGroup_   = "table-header-group"
	TableFooterGroup_   = "table-footer-group"
	TableRow_           = "table-row"
	TableColumnGroup_   = "table-column-group"
	TableColumn_        = "table-column"
	TableCell_          = "table-cell"
	TableCaption_       = "table-caption"
	Show_               = "show" // "empty-cells"=show | hide | inherit
	Hide_               = "hide"
	Invert_             = "invert"  // "outline-color"=<color> | invert | inherit
	Visible_            = "visible" // "overflow"=visible | hidden | scroll | auto | inherit
	Hidden_             = "hidden"
	Always_             = "always" // "page-break-after"=auto | always | avoid | left | right | inherit
	Avoid_              = "avoid"
	Static_             = "static" // "position"=static | relative | absolute | fixed | inherit
	Relative_           = "relative"
	Absolute_           = "absolute"
	Justify_            = "justify"   // "text-align"=left | right | center | justify | inherit
	Underline_          = "underline" // "text-decoration"=none | underline | overline | line-through | blink | inherit | none
	Overline_           = "overline"
	LineThrough_        = "line-through"
	Blink_              = "blink"
	Capitalize_         = "capitalize" // "text-transform"=capitalize | uppercase | lowercase | none | inherit
	Uppercase_          = "uppercase"
	Lowercase_          = "lowercase"
	Embed_              = "embed" //"unicode-bidi"=normal | embed | bidi-override | inherit
	BidiOverride_       = "bidi-override"
	Baseline_           = "baseline" //"vertical-align"=baseline | sub | super | top | text-top | middle | bottom | text-bottom | <percentage> | <length> | inherit
	Sub_                = "sub"
	Super_              = "super"
	TextTop_            = "text-top"
	Middle_             = "middle"
	TextBottom_         = "text-bottom"
	Pre_                = "pre" //"white-space"=normal | pre | nowrap | pre-wrap | pre-line | inherit
	Nowrap_             = "nowrap"
	PreWrap_            = "pre-wrap"
	PreLine_            = "pre-line"
	Dashed_             = "dashed" // "border-image-width"=<length> | <percentage> | <number> | auto
	Dotted_             = "dotted"
	Solid_              = "solid"
	Caption_            = "caption" // caption | icon | menu | message-box | small-caption | status-bar
	Icon_               = "icon"
	Menu_               = "menu"
	MessageBox_         = "message-box"
	SmallCaption_       = "small-caption"
	StatusBar_          = "status-bar"
	None_               = "none"
	Italic_             = "italic" // "font-style"=normal | italic | oblique | inherit
	Oblique_            = "oblique"
	Bold_               = "bold" // "font-weight"=normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | inherit
	Bolder_             = "bolder"
	Lighter_            = "lighter"
	SmallCaps_          = "small-caps" // "font-variant"=normal | small-caps | inherit
	Inside_             = "inside"     // "list-style-position"=inside | outside | inherit
	Outside_            = "outside"
	Disc_               = "disc" // "list-style-type"=disc | circle | square | decimal | decimal-leading-zero | lower-roman | upper-roman | lower-greek | lower-latin | upper-latin | armenian | georgian | lower-alpha | upper-alpha | none | inherit
	Circle_             = "circle"
	Square_             = "square"
	Decimal_            = "decimal"
	DecimalLeadingZero_ = "decimal-leading-zero"
	LowerRoman_         = "lower-roman"
	UpperRoman_         = "upper-roman"
	LowerGreek_         = "lower-greek"
	LowerLatin_         = "lower-latin"
	UpperLatin_         = "upper-latin"
	Armenian_           = "armenian"
	Georgian_           = "georgian"
	LowerAlpha_         = "lower-alpha"
	UpperAlpha_         = "upper-alpha"
)

func j(s []string) string { return strings.Join(s, " ") }

func Url(url string) string    { return fmt.Sprintf(`url('%s')`, url) }
func Percent(f float64) string { return fmt.Sprintf(`%v`, f) + "%" }
func Pc(f float64) string      { return Percent(f) }
func Px(f float64) string      { return fmt.Sprintf(`%vpx`, f) }
func Em(f float64) string      { return fmt.Sprintf(`%vem`, f) }
func Pt(f float64) string      { return fmt.Sprintf(`%vpt`, f) }

func BackgroundAttachment(s string) Style  { return Style{"background-attachment", s} }
func BorderTopLeftRadius(s string) Style   { return Style{"border-top-left-radius", s} }
func Position(s string) Style              { return Style{"position", s} }
func BackgroundPosition(s ...string) Style { return Style{"background-position", j(s)} }
func BackgroundColor(s string) Style       { return Style{"background-color", s} }
func BackgroundImage(s string) Style       { return Style{"background-image", s} }
func BackgroundRepeat(s string) Style      { return Style{"background-repeat", s} }
func BorderCollapse(s string) Style        { return Style{"border-collapse", s} }
func BorderColor(s string) Style           { return Style{"border-color", s} }
func BorderSpacing(s string) Style         { return Style{"border-spacing", s} }
func BorderStyle(s string) Style           { return Style{"border-style", s} }
func BorderTop(s ...string) Style          { return Style{"border-top", j(s)} }
func BorderRight(s ...string) Style        { return Style{"border-right", j(s)} }
func BorderBottom(s ...string) Style       { return Style{"border-bottom", j(s)} }
func BorderLeft(s ...string) Style         { return Style{"border-left", j(s)} }
func BorderTopColor(s string) Style        { return Style{"border-top-color", s} }
func BorderTopStyle(s string) Style        { return Style{"border-top-style", s} }
func BorderTopWidth(s string) Style        { return Style{"border-top-width", s} }
func BorderRightColor(s string) Style      { return Style{"border-right-color", s} }
func BorderRightStyle(s string) Style      { return Style{"border-right-style", s} }
func BorderRightWidth(s string) Style      { return Style{"border-right-width", s} }
func BorderBottomColor(s string) Style     { return Style{"border-bottom-color", s} }
func BorderBottomStyle(s string) Style     { return Style{"border-bottom-style", s} }
func BorderBottomWidth(s string) Style     { return Style{"border-bottom-width", s} }
func BorderLeftColor(s string) Style       { return Style{"border-left-color", s} }
func BorderLeftStyle(s string) Style       { return Style{"border-left-style", s} }
func BorderLeftWidth(s string) Style       { return Style{"border-left-width", s} }
func BorderWidth(s string) Style           { return Style{"border-width", s} }
func Bottom(s string) Style                { return Style{"bottom", s} }
func CaptionSide(s string) Style           { return Style{"caption-side", s} }
func Clear(s string) Style                 { return Style{"clear", s} }
func Clip(s string) Style                  { return Style{"clip", s} }
func Color(s string) Style                 { return Style{"color", s} }
func Content(s string) Style               { return Style{"content", s} }
func CounterIncrement(s string) Style      { return Style{"counter-increment", s} }
func CounterReset(s string) Style          { return Style{"counter-reset", s} }
func Cursor(s string) Style                { return Style{"cursor", s} }
func Direction(s string) Style             { return Style{"direction", s} }
func Display(s string) Style               { return Style{"display", s} }
func EmptyCells(s string) Style            { return Style{"empty-cells", s} }
func Float(s string) Style                 { return Style{"float", s} }
func FontFamily(s string) Style            { return Style{"font-family", s} }
func FontSize(s string) Style              { return Style{"font-size", s} }
func FontStyle(s string) Style             { return Style{"font-style", s} }
func FontVariant(s string) Style           { return Style{"font-variant", s} }
func FontWeight(s string) Style            { return Style{"font-weight", s} }
func Height(s string) Style                { return Style{"height", s} }
func Left(s string) Style                  { return Style{"left", s} }
func LetterSpacing(s string) Style         { return Style{"letter-spacing", s} }
func LineHeight(s string) Style            { return Style{"line-height", s} }
func ListStyleImage(s string) Style        { return Style{"list-style-image", s} }
func ListStylePosition(s string) Style     { return Style{"list-style-position", s} }
func ListStyleType(s string) Style         { return Style{"list-style-type", s} }
func ListStyle(s string) Style             { return Style{"list-style", s} }
func MarginRight(s string) Style           { return Style{"margin-right", s} }
func MarginLeft(s string) Style            { return Style{"margin-left", s} }
func MarginTop(s string) Style             { return Style{"margin-top", s} }
func MarginBottom(s string) Style          { return Style{"margin-bottom", s} }
func Margin(s ...string) Style             { return Style{"margin", j(s)} }
func MaxHeight(s string) Style             { return Style{"max-height", s} }
func MaxWidth(s string) Style              { return Style{"max-width", s} }
func MinHeight(s string) Style             { return Style{"min-height", s} }
func MinWidth(s string) Style              { return Style{"min-width", s} }
func Opacity(s string) Style               { return Style{"opacity", s} }
func Orphans(s string) Style               { return Style{"orphans", s} }
func OutlineColor(s string) Style          { return Style{"outline-color", s} }
func OutlineStyle(s string) Style          { return Style{"outline-style", s} }
func OutlineWidth(s string) Style          { return Style{"outline-width", s} }
func Outline(s string) Style               { return Style{"outline", s} }
func Overflow(s string) Style              { return Style{"overflow", s} }
func PaddingTop(s string) Style            { return Style{"padding-top", s} }
func PaddingRight(s string) Style          { return Style{"padding-right", s} }
func PaddingBottom(s string) Style         { return Style{"padding-bottom", s} }
func PaddingLeft(s string) Style           { return Style{"padding-left", s} }
func Padding(s ...string) Style            { return Style{"padding", j(s)} }
func PageBreakAfter(s string) Style        { return Style{"page-break-after", s} }
func PageBreakBefore(s string) Style       { return Style{"page-break-before", s} }
func PageBreakInside(s string) Style       { return Style{"page-break-inside", s} }
func position(s string) Style              { return Style{"position", s} }
func Quotes(s string) Style                { return Style{"quotes", s} }
func Right(s string) Style                 { return Style{"right", s} }
func Resize(s string) Style                { return Style{"resize", s} }
func TableLayout(s string) Style           { return Style{"table-layout", s} }
func TextAlign(s string) Style             { return Style{"text-align", s} }
func TextDecoration(s string) Style        { return Style{"text-decoration", s} }
func TextIndent(s string) Style            { return Style{"text-indent", s} }
func TextTransform(s string) Style         { return Style{"text-transform", s} }
func Top(s string) Style                   { return Style{"top", s} }
func UnicodeBidi(s string) Style           { return Style{"unicode-bidi", s} }
func VerticalAlign(s string) Style         { return Style{"vertical-align", s} }
func Visibility(s string) Style            { return Style{"visibility", s} }
func WhiteSpace(s string) Style            { return Style{"white-space", s} }
func Widows(s string) Style                { return Style{"widows", s} }
func WordSpacing(s string) Style           { return Style{"word-spacing", s} }
func ZIndex(s string) Style                { return Style{"z-index", s} }
func BackgroundClip(s string) Style        { return Style{"background-clip", s} }
func BackgroundOrigin(s string) Style      { return Style{"background-origin", s} }
func Width(s string) Style                 { return Style{"width", s} }
func BackgroundSize(s string) Style        { return Style{"background-size", s} }
func Border(s ...string) Style             { return Style{"border", j(s)} }
func BorderImage(s ...string) Style        { return Style{"border-image", j(s)} }
func BorderImageOutset(s string) Style     { return Style{"border-image-outset", s} }
func BorderImageRepeat(s string) Style     { return Style{"border-image-repeat", s} }
func BorderImageSlice(s string) Style      { return Style{"border-image-slice", s} }
func BorderImageSource(s string) Style     { return Style{"border-image-source", s} }
func BorderImageWidth(s string) Style      { return Style{"border-image-width", s} }
func BorderRadius(s string) Style          { return Style{"border-radius", s} }
func BoxDecorationBreak(s string) Style    { return Style{"box-decoration-break", s} }
func BoxShadow(s ...string) Style          { return Style{"box-shadow", j(s)} }
func Background(s ...string) Style         { return Style{"background", j(s)} }
func Font(s ...string) Style               { return Style{"font", j(s)} }

// "margin-right" "margin-left"=<margin-width> | inherit
// "margin-top" "margin-bottom"=<margin-width> | inherit
// "margin"=<margin-width> | inherit
// "max-height"=<length> | <percentage> | none | inherit
// "max-width"=<length> | <percentage> | none | inherit
// "min-height"=<length> | <percentage> | inherit
// "min-width"=<length> | <percentage> | inherit
// "opacity"=<number> | inherit
// "orphans"=<integer> | inherit
// "outline-color"=<color> | invert | inherit
// "outline-style"=<border-style> | inherit
// "outline-width"=<border-width> | inherit
// "outline"=<color> | <border-style> | <border-width> | inherit
// "overflow"=visible | hidden | scroll | auto | inherit
// "padding-top" "padding-right" "padding-bottom" "padding-left"=<padding-width> | inherit
// "padding"=<padding-width> | inherit
// "page-break-after"=auto | always | avoid | left | right | inherit
// "page-break-before"=auto | always | avoid | left | right | inherit
// "page-break-inside"=avoid | auto | inherit
// "position"=static | relative | absolute | fixed | inherit
// "quotes"=none | inherit
// "right"=<length> | <percentage> | auto | inherit
// "table-layout"=auto | fixed | inherit
// "text-align"=left | right | center | justify | inherit
// "text-decoration"=none | underline | overline | line-through | blink | inherit | none
// "text-indent"=<length> | <percentage> | inherit
// "text-transform"=capitalize | uppercase | lowercase | none | inherit
// "top"=<length> | <percentage> | auto | inherit
// "unicode-bidi"=normal | embed | bidi-override | inherit
// "vertical-align"=baseline | sub | super | top | text-top | middle | bottom | text-bottom | <percentage> | <length> | inherit
// "visibility"=visible | hidden | collapse | inherit
// "white-space"=normal | pre | nowrap | pre-wrap | pre-line | inherit
// "widows"=<integer> | inherit
// "width"=<length> | <percentage> | auto | inherit
// "word-spacing"=normal | <length> | inherit
// "z-index"=auto | <integer> | inherit
// "background-clip"=<box>
// "background-origin"=<box>
// "background-size"=<bg-size>
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
// "background-position"=left | center | right | top | bottom | inherit
// caption | icon | menu | message-box | small-caption | status-bar
// "text-indent"=<length> | <percentage> | inherit
// "border-color"=<color> | inherit
// "border-spacing"=inherit
// "bottom"=<length> | <percentage> | auto | inherit
// "caption-side"=top | bottom | inherit
// "color"=<color> | inherit
// "counter-increment"=none | inherit
// "counter-reset"=none | inherit
// "orphans"=<integer> | inherit
// "outline-style"=<border-style> | inherit
// "outline-width"=<border-width> | inherit
// "outline"=<color> | <border-style> | <border-width> | inherit
// "padding-top" "padding-right" "padding-bottom" "padding-left"=<padding-width> | inherit
// "padding"=<padding-width> | inherit
// "page-break-before"=auto | always | avoid | left | right | inherit
// "page-break-inside"=avoid | auto | inherit
// "quotes"=none | inherit
// "right"=<length> | <percentage> | auto | inherit
// "table-layout"=auto | fixed | inherit
// "top"=<length> | <percentage> | auto | inherit
// "visibility"=visible | hidden | collapse | inherit
// "widows"=<integer> | inherit
// "width"=<length> | <percentage> | auto | inherit
// "word-spacing"=normal | <length> | inherit
// "z-index"=auto | <integer> | inherit
// "background-clip"=<box>
// "background-origin"=<box>
// "background-size"=<bg-size>
// "border"=<border-width> | <border-style> | <color>
// "border-image"=<border-image-source> | <border-image-slice> | <border-image-width> | <border-image-width> | <border-image-outset> | <border-image-repeat>
// "border-image-outset"=<length> | <number>
// "border-image-repeat"=stretch | repeat | round | space
// "border-image-slice"=<number> | <percentage>
// "border-image-source"=none | <image>
// "border-radius"=<length> | <percentage>
// "border-top-left-radius" "border-top-right-radius" "border-bottom-right-radius" "border-bottom-left-radius"=<length> | <percentage>
// "box-decoration-break"=slice | clone
// "box-shadow"=none | <shadow> | none
// "background"=<color> | <uri> | repeat | repeat-x | repeat-y | no-repeat | scroll | fixed | left | center | right | top | bottom | inherit
// "font"=normal | italic | oblique | normal | small-caps | normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | normal | <generic-family> | caption | icon | menu | message-box | small-caption | status-bar | inherit
// colors
// "clip"=<shape> | auto | inherit
// func Shape(s string) Style  = s
// "background-color"=<color> | inherit
// func Color(c string) Style  = c
// "background-image"=<uri> | none | inherit
// "border-style"=<border-style> | inherit
// func Style(s string) Style  = s
// "border-top" "border-right" "border-bottom" "border-left"=<border-width> | <border-style> | <color> | inherit
// func Width(w string) Style  = w
// "border-top-color" "border-right-color" "border-bottom-color" "border-left-color"=<color> | inherit
// "border-top-style" "border-right-style" "border-bottom-style" "border-left-style"=<border-style> | inherit
// "border-top-width" "border-right-width" "border-bottom-width" "border-left-width"=<border-width> | inherit
// "border-width"=<border-width> | inherit
// type Color string
// "float"=left | right | none | inherit
// "font-family"=<generic-family>| inherit
// FontFamily(f string) Style  = f
// "font-size"=inherit
// Weight(i int) Style      = fmt.Sprintf(`%v`, i)
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
// "background-color"=<color> | inherit
// "background-image"=<uri> | none | inherit
// "background-repeat"=repeat | repeat-x | repeat-y | no-repeat | inherit
// "border-collapse"=collapse | separate | inherit
// "border-color"=<color> | inherit
// "border-spacing"=inherit
// "border-style"=<border-style> | inherit
// "border-top" "border-right" "border-bottom" "border-left"=<border-width> | <border-style> | <color> | inherit
// "border-top-color" "border-right-color" "border-bottom-color" "border-left-color"=<color> | inherit
// "border-top-style" "border-right-style" "border-bottom-style" "border-left-style"=<border-style> | inherit
// "border-top-width" "border-right-width" "border-bottom-width" "border-left-width"=<border-width> | inherit
// "border-width"=<border-width> | inherit
// "bottom"=<length> | <percentage> | auto | inherit
// "caption-side"=top | bottom | inherit
// "clear"=none | left | right | both | inherit
// "clip"=<shape> | auto | inherit
// "color"=<color> | inherit
// "content"=normal | none | <uri> | open-quote | close-quote | no-open-quote | no-close-quote | inherit
// "counter-increment"=none | inherit
// "counter-reset"=none | inherit
// "cursor"=<uri> | auto | crosshair | default | pointer | move | e-resize | ne-resize | nw-resize | n-resize | se-resize | sw-resize | s-resize | w-resize | text | wait | help | progress | inherit
// "direction"=ltr | rtl | inherit
// "display"=inline | block | list-item | inline-block | table | inline-table | table-row-group | table-header-group | table-footer-group | table-row | table-column-group | table-column | table-cell | table-caption | none | inherit
// "empty-cells"=show | hide | inherit
// "float"=left | right | none | inherit
// "font-family"=<generic-family>| inherit
// "font-size"=inherit
// "font-style"=normal | italic | oblique | inherit
// "font-variant"=normal | small-caps | inherit
// "font-weight"=normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | inherit
// "height"=<length> | <percentage> | auto | inherit
// "left"=<length> | <percentage> | auto | inherit
// "letter-spacing"=normal | <length> | inherit
// "line-height"=normal | <number> | <length> | <percentage> | inherit
// "list-style-image"=<uri> | none | inherit
// "list-style-position"=inside | outside | inherit
// "list-style-type"=disc | circle | square | decimal | decimal-leading-zero | lower-roman | upper-roman | lower-greek | lower-latin | upper-latin | armenian | georgian | lower-alpha | upper-alpha | none | inherit
// "list-style"=disc | circle | square | decimal | decimal-leading-zero | lower-roman | upper-roman | lower-greek | lower-latin | upper-latin | armenian | georgian | lower-alpha | upper-alpha | none | inside | outside | <uri> | inherit
