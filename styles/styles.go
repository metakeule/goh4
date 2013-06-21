package styles

import (
	. "github.com/metakeule/goh4"
)

// type Color string

// "background-attachment"=scroll | fixed | inherit
func BackgroundAttachment() (ba Style) { return Style{Property: "background-attachment"} }

// "background-position"=left | center | right | top | bottom | inherit
func BackgroundPosition() (ba Style) { return Style{Property: "background-position"} }

// "background-color"=<color> | inherit
func BackgroundColor() (ba Style) { return Style{Property: "background-color"} }

// "background-image"=<uri> | none | inherit
func BackgroundImage() (ba Style) { return Style{Property: "background-image"} }

// "background-repeat"=repeat | repeat-x | repeat-y | no-repeat | inherit
func BackgroundRepeat() (ba Style) { return Style{Property: "background-repeat"} }

// "border-collapse"=collapse | separate | inherit
func BorderCollapse() (ba Style) { return Style{Property: "border-collapse"} }

// "border-color"=<color> | inherit
func BorderColor() (ba Style) { return Style{Property: "border-color"} }

// "border-spacing"=inherit
func BorderSpacing() (ba Style) { return Style{Property: "border-spacing"} }

// "border-style"=<border-style> | inherit
func BorderStyle() (ba Style) { return Style{Property: "border-style"} }

// "border-top" "border-right" "border-bottom" "border-left"=<border-width> | <border-style> | <color> | inherit
func BorderTop() (ba Style)    { return Style{Property: "border-top"} }
func BorderRight() (ba Style)  { return Style{Property: "border-right"} }
func BorderBottom() (ba Style) { return Style{Property: "border-bottom"} }
func BorderLeft() (ba Style)   { return Style{Property: "border-left"} }

// "border-top-color" "border-right-color" "border-bottom-color" "border-left-color"=<color> | inherit
func BorderTopColor() (ba Style) { return Style{Property: "border-top-color"} }

// "border-top-style" "border-right-style" "border-bottom-style" "border-left-style"=<border-style> | inherit
func BorderTopStyle() (ba Style) { return Style{Property: "border-top-style"} }

// "border-top-width" "border-right-width" "border-bottom-width" "border-left-width"=<border-width> | inherit
func BorderTopWidth() (ba Style)    { return Style{Property: "border-top-width"} }
func BorderRightColor() (ba Style)  { return Style{Property: "border-right-color"} }
func BorderRightStyle() (ba Style)  { return Style{Property: "border-right-style"} }
func BorderRightWidth() (ba Style)  { return Style{Property: "border-right-width"} }
func BorderBottomColor() (ba Style) { return Style{Property: "border-bottom-color"} }
func BorderBottomStyle() (ba Style) { return Style{Property: "border-bottom-style"} }
func BorderBottomWidth() (ba Style) { return Style{Property: "border-bottom-width"} }
func BorderLeftColor() (ba Style)   { return Style{Property: "border-left-color"} }
func BorderLeftStyle() (ba Style)   { return Style{Property: "border-left-style"} }
func BorderLeftWidth() (ba Style)   { return Style{Property: "border-left-width"} }

// "border-width"=<border-width> | inherit
func BorderWidth() (ba Style) { return Style{Property: "border-width"} }

// "bottom"=<length> | <percentage> | auto | inherit
func Bottom() (ba Style) { return Style{Property: "bottom"} }

// "caption-side"=top | bottom | inherit
func CaptionSide() (ba Style) { return Style{Property: "caption-side"} }

// "clear"=none | left | right | both | inherit
func Clear() (ba Style) { return Style{Property: "clear"} }

// "clip"=<shape> | auto | inherit
func Clip() (ba Style) { return Style{Property: "clip"} }

// "color"=<color> | inherit
func Color() (ba Style) { return Style{Property: "color"} }

// "content"=normal | none | <uri> | open-quote | close-quote | no-open-quote | no-close-quote | inherit
func Content() (ba Style) { return Style{Property: "content"} }

// "counter-increment"=none | inherit
func CounterIncrement() (ba Style) { return Style{Property: "counter-increment"} }

// "counter-reset"=none | inherit
func CounterReset() (ba Style) { return Style{Property: "counter-reset"} }

// "cursor"=<uri> | auto | crosshair | default | pointer | move | e-resize | ne-resize | nw-resize | n-resize | se-resize | sw-resize | s-resize | w-resize | text | wait | help | progress | inherit
func Cursor() (ba Style) { return Style{Property: "cursor"} }

// "direction"=ltr | rtl | inherit
func Direction() (ba Style) { return Style{Property: "direction"} }

// "display"=inline | block | list-item | inline-block | table | inline-table | table-row-group | table-header-group | table-footer-group | table-row | table-column-group | table-column | table-cell | table-caption | none | inherit
func Display() (ba Style) { return Style{Property: "display"} }

// "empty-cells"=show | hide | inherit
func EmptyCells() (ba Style) { return Style{Property: "empty-cells"} }

// "float"=left | right | none | inherit
func Float() (ba Style) { return Style{Property: "float"} }

// "font-family"=<generic-family>| inherit
func FontFamily() (ba Style) { return Style{Property: "font-family"} }

// "font-size"=inherit
func FontSize() (ba Style) { return Style{Property: "font-size"} }

// "font-style"=normal | italic | oblique | inherit
func FontStyle() (ba Style) { return Style{Property: "font-style"} }

// "font-variant"=normal | small-caps | inherit
func FontVariant() (ba Style) { return Style{Property: "font-variant"} }

// "font-weight"=normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | inherit
func FontWeight() (ba Style) { return Style{Property: "font-weight"} }

// "height"=<length> | <percentage> | auto | inherit
func Height() (ba Style) { return Style{Property: "height"} }

// "left"=<length> | <percentage> | auto | inherit
func Left() (ba Style) { return Style{Property: "left"} }

// "letter-spacing"=normal | <length> | inherit
func LetterSpacing() (ba Style) { return Style{Property: "letter-spacing"} }

// "line-height"=normal | <number> | <length> | <percentage> | inherit
func LineHeight() (ba Style) { return Style{Property: "line-height"} }

// "list-style-image"=<uri> | none | inherit
func ListStyleImage() (ba Style) { return Style{Property: "list-style-image"} }

// "list-style-position"=inside | outside | inherit
func ListStylePosition() (ba Style) { return Style{Property: "list-style-position"} }

// "list-style-type"=disc | circle | square | decimal | decimal-leading-zero | lower-roman | upper-roman | lower-greek | lower-latin | upper-latin | armenian | georgian | lower-alpha | upper-alpha | none | inherit
func ListStyleType() (ba Style) { return Style{Property: "list-style-type"} }

// "list-style"=disc | circle | square | decimal | decimal-leading-zero | lower-roman | upper-roman | lower-greek | lower-latin | upper-latin | armenian | georgian | lower-alpha | upper-alpha | none | inside | outside | <uri> | inherit
func ListStyle() (ba Style) { return Style{Property: "list-style"} }

// "margin-right" "margin-left"=<margin-width> | inherit
func MarginRight() (ba Style) { return Style{Property: "margin-right"} }
func MarginLeft() (ba Style)  { return Style{Property: "margin-left"} }

// "margin-top" "margin-bottom"=<margin-width> | inherit
func MarginTop() (ba Style)    { return Style{Property: "margin-top"} }
func MarginBottom() (ba Style) { return Style{Property: "margin-bottom"} }

// "margin"=<margin-width> | inherit
func Margin() (ba Style) { return Style{Property: "margin"} }

// "max-height"=<length> | <percentage> | none | inherit
func MaxHeight() (ba Style) { return Style{Property: "max-height"} }

// "max-width"=<length> | <percentage> | none | inherit
func MaxWidth() (ba Style) { return Style{Property: "max-width"} }

// "min-height"=<length> | <percentage> | inherit
func MinHeight() (ba Style) { return Style{Property: "min-height"} }

// "min-width"=<length> | <percentage> | inherit
func MinWidth() (ba Style) { return Style{Property: "min-width"} }

// "opacity"=<number> | inherit
func Opacity() (ba Style) { return Style{Property: "opacity"} }

// "orphans"=<integer> | inherit
func Orphans() (ba Style) { return Style{Property: "orphans"} }

// "outline-color"=<color> | invert | inherit
func OutlineColor() (ba Style) { return Style{Property: "outline-color"} }

// "outline-style"=<border-style> | inherit
func OutlineStyle() (ba Style) { return Style{Property: "outline-style"} }

// "outline-width"=<border-width> | inherit
func OutlineWidth() (ba Style) { return Style{Property: "outline-width"} }

// "outline"=<color> | <border-style> | <border-width> | inherit
func Outline() (ba Style) { return Style{Property: "outline"} }

// "overflow"=visible | hidden | scroll | auto | inherit
func Overflow() (ba Style) { return Style{Property: "min-width"} }

// "padding-top" "padding-right" "padding-bottom" "padding-left"=<padding-width> | inherit
func PaddingTop() (ba Style)    { return Style{Property: "padding-top"} }
func PaddingRight() (ba Style)  { return Style{Property: "padding-right"} }
func PaddingBottom() (ba Style) { return Style{Property: "padding-bottom"} }
func PaddingLeft() (ba Style)   { return Style{Property: "padding-left"} }

// "padding"=<padding-width> | inherit
func Padding() (ba Style) { return Style{Property: "padding"} }

// "page-break-after"=auto | always | avoid | left | right | inherit
func PageBreakAfter() (ba Style) { return Style{Property: "page-break-after"} }

// "page-break-before"=auto | always | avoid | left | right | inherit
func PageBreakBefore() (ba Style) { return Style{Property: "page-break-before"} }

// "page-break-inside"=avoid | auto | inherit
func PageBreakInside() (ba Style) { return Style{Property: "page-break-inside"} }

// "position"=static | relative | absolute | fixed | inherit
func position() (ba Style) { return Style{Property: "position"} }

// "quotes"=none | inherit
func Quotes() (ba Style) { return Style{Property: "quotes"} }

// "right"=<length> | <percentage> | auto | inherit
func Right() (ba Style) { return Style{Property: "right"} }

// "table-layout"=auto | fixed | inherit
func TableLayout() (ba Style) { return Style{Property: "table-layout"} }

// "text-align"=left | right | center | justify | inherit
func TextAlign() (ba Style) { return Style{Property: "text-align"} }

// "text-decoration"=none | underline | overline | line-through | blink | inherit | none
func TextDecoration() (ba Style) { return Style{Property: "text-decoration"} }

// "text-indent"=<length> | <percentage> | inherit
func TextIndent() (ba Style) { return Style{Property: "text-indent"} }

// "text-transform"=capitalize | uppercase | lowercase | none | inherit
func TextTransform() (ba Style) { return Style{Property: "text-transform"} }

//"top"=<length> | <percentage> | auto | inherit
func Top() (ba Style) { return Style{Property: "top"} }

//"unicode-bidi"=normal | embed | bidi-override | inherit
func UnicodeBidi() (ba Style) { return Style{Property: "unicode-bidi"} }

//"vertical-align"=baseline | sub | super | top | text-top | middle | bottom | text-bottom | <percentage> | <length> | inherit
func VerticalAlign() (ba Style) { return Style{Property: "vertical-align"} }

//"visibility"=visible | hidden | collapse | inherit
func Visibility() (ba Style) { return Style{Property: "visibility"} }

//"white-space"=normal | pre | nowrap | pre-wrap | pre-line | inherit
func WhiteSpace() (ba Style) { return Style{Property: "white-space"} }

//"widows"=<integer> | inherit
func Widows() (ba Style) { return Style{Property: "widows"} }

//"width"=<length> | <percentage> | auto | inherit
func Width() (ba Style) { return Style{Property: "width"} }

//"word-spacing"=normal | <length> | inherit
func WordSpacing() (ba Style) { return Style{Property: "word-spacing"} }

//"z-index"=auto | <integer> | inherit
func ZIndex() (ba Style) { return Style{Property: "z-index"} }

//"background-clip"=<box>
func BackgroundClip() (ba Style) { return Style{Property: "background-clip"} }

//"background-origin"=<box>
func BackgroundOrigin() (ba Style) { return Style{Property: "background-origin"} }

//"background-size"=<bg-size>
func BackgroundSize() (ba Style) { return Style{Property: "background-size"} }

// "border"=<border-width> | <border-style> | <color>
func Border() (ba Style) { return Style{Property: "border"} }

// "border-image"=<border-image-source> | <border-image-slice> | <border-image-width> | <border-image-width> | <border-image-outset> | <border-image-repeat>
func BorderImage() (ba Style) { return Style{Property: "border-image"} }

// "border-image-outset"=<length> | <number>
func BorderImageOutset() (ba Style) { return Style{Property: "border-image-outset"} }

// "border-image-repeat"=stretch | repeat | round | space
func BorderImageRepeat() (ba Style) { return Style{Property: "border-image-repeat"} }

// "border-image-slice"=<number> | <percentage>
func BorderImageSlice() (ba Style) { return Style{Property: "border-image-slice"} }

// "border-image-source"=none | <image>
func BorderImageSource() (ba Style) { return Style{Property: "border-image-source"} }

// "border-image-width"=<length> | <percentage> | <number> | auto
func BorderImageWidth() (ba Style) { return Style{Property: "border-image-width"} }

// "border-radius"=<length> | <percentage>
func BorderRadius() (ba Style) { return Style{Property: "border-radius"} }

// "border-top-left-radius" "border-top-right-radius" "border-bottom-right-radius" "border-bottom-left-radius"=<length> | <percentage>
func BorderTopLeftRadius() (ba Style) { return Style{Property: "border-top-left-radius"} }

// "box-decoration-break"=slice | clone
func BoxDecorationBreak() (ba Style) { return Style{Property: "box-decoration-break"} }

// "box-shadow"=none | <shadow> | none
func BoxShadow() (ba Style) { return Style{Property: "box-shadow"} }

// "background"=<color> | <uri> | repeat | repeat-x | repeat-y | no-repeat | scroll | fixed | left | center | right | top | bottom | inherit
func Background() (ba Style) { return Style{Property: "background"} }

// "font"=normal | italic | oblique | normal | small-caps | normal | bold | bolder | lighter | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | normal | <generic-family> | caption | icon | menu | message-box | small-caption | status-bar | inherit
func Font() (ba Style) { return Style{Property: "font"} }

// caption | icon | menu | message-box | small-caption | status-bar
