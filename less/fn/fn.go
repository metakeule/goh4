package fn

import (
	"gopkg.in/metakeule/goh4.v5/less"
)

func fn(f string, i ...string) string {
	return less.Call(f, i...).String()
}

/*
escape(@string);               // URL encodes a string
e(@string);                    // escape string content
%(@string, values...);         // formats a string
*/

var ESCAPE_ = func(i ...string) string {
	return fn("escape", i...)
}
var E_ = func(i ...string) string {
	return fn("e", i...)
}
var FORMAT_ = func(i ...string) string {
	return fn("%", i...)
}

/*
unit(@dimension, [@unit: ""]); // remove or change the unit of a dimension
color(@string);                // parses a string to a color
data-uri([mimetype,] url);       // * inlines a resource and falls back to url()
*/
var UNIT_ = func(i ...string) string {
	return fn("unit", i...)
}
var COLOR_ = func(i ...string) string {
	return fn("color", i...)
}
var DATA_URI_ = func(i ...string) string {
	return fn("data-uri", i...)
}

/*
ceil(@number);                 // rounds up to an integer
floor(@number);                // rounds down to an integer
percentage(@number);           // converts to a %, e.g. 0.5 -> 50%
round(number, [places: 0]);    // rounds a number to a number of places
sqrt(number);                  // * calculates square root of a number
abs(number);                   // * absolute value of a number
sin(number);                   // * sine function
asin(number);                  // * arcsine - inverse of sine function
cos(number);                   // * cosine function
acos(number);                  // * arccosine - inverse of cosine function
tan(number);                   // * tangent function
atan(number);                  // * arctangent - inverse of tangent function
pi();                          // * returns pi
pow(@base, @exponent);     // * first argument raised to the power of the second argument
mod(number, number);       // * first argument modulus second argument
*/
var CEIL_ = func(i ...string) string {
	return fn("ceil", i...)
}
var FLOOR_ = func(i ...string) string {
	return fn("floor", i...)
}
var PERCENTAGE_ = func(i ...string) string {
	return fn("percentage", i...)
}
var ROUND_ = func(i ...string) string {
	return fn("round", i...)
}
var SQRT_ = func(i ...string) string {
	return fn("sqrt", i...)
}
var ABS_ = func(i ...string) string {
	return fn("abs", i...)
}
var SIN_ = func(i ...string) string {
	return fn("sin", i...)
}
var ASIN_ = func(i ...string) string {
	return fn("asin", i...)
}
var COS_ = func(i ...string) string {
	return fn("cos", i...)
}
var ACOS_ = func(i ...string) string {
	return fn("acos", i...)
}
var TAN_ = func(i ...string) string {
	return fn("tan", i...)
}
var ATAN_ = func(i ...string) string {
	return fn("atan", i...)
}
var PI_ = func(i ...string) string {
	return fn("pi", i...)
}
var POW_ = func(i ...string) string {
	return fn("pow", i...)
}
var MOD_ = func(i ...string) string {
	return fn("mod", i...)
}

/*
convert(number, units);    // * converts between number types
*/
var CONVERT_ = func(i ...string) string {
	return fn("convert", i...)
}

/*
rgb(@r, @g, @b);                             // converts to a color
rgba(@r, @g, @b, @a);                        // converts to a color
argb(@color);                                // creates a #AARRGGBB
hsl(@hue, @saturation, @lightness);          // creates a color
hsla(@hue, @saturation, @lightness, @alpha); // creates a color
hsv(@hue, @saturation, @value);              // creates a color
hsva(@hue, @saturation, @value, @alpha);     // creates a color
*/
var RGB_ = func(i ...string) string {
	return fn("rgb", i...)
}
var RGBA_ = func(i ...string) string {
	return fn("rgba", i...)
}
var ARGB_ = func(i ...string) string {
	return fn("argb", i...)
}
var HSL_ = func(i ...string) string {
	return fn("hsl", i...)
}
var HSLA_ = func(i ...string) string {
	return fn("hsla", i...)
}
var HSV_ = func(i ...string) string {
	return fn("hsv", i...)
}
var HSVA_ = func(i ...string) string {
	return fn("hsva", i...)
}

/*
hue(@color);           // returns the `hue` channel of @color in the HSL space
saturation(@color);    // returns the `saturation` channel of @color in the HSL space
lightness(@color);     // returns the 'lightness' channel of @color in the HSL space
hsvhue(@color);        // * returns the `hue` channel of @color in the HSV space
hsvsaturation(@color); // * returns the `saturation` channel of @color in the HSV space
hsvvalue(@color);      // * returns the 'value' channel of @color in the HSV space
red(@color);           // returns the 'red' channel of @color
green(@color);         // returns the 'green' channel of @color
blue(@color);          // returns the 'blue' channel of @color
alpha(@color);         // returns the 'alpha' channel of @color
luma(@color);          // returns the 'luma' value (perceptual brightness) of @color
*/
var HUE_ = func(i ...string) string {
	return fn("hue", i...)
}
var SATURATION_ = func(i ...string) string {
	return fn("saturation", i...)
}
var LIGHTNESS_ = func(i ...string) string {
	return fn("lightness", i...)
}
var HSV_HUE_ = func(i ...string) string {
	return fn("hsvhue", i...)
}
var HSV_SATURATION_ = func(i ...string) string {
	return fn("hsvsaturation", i...)
}
var HSV_VALUE_ = func(i ...string) string {
	return fn("hsvvalue", i...)
}
var RED_ = func(i ...string) string {
	return fn("red", i...)
}
var GREEN_ = func(i ...string) string {
	return fn("green", i...)
}
var BLUE_ = func(i ...string) string {
	return fn("blue", i...)
}
var ALPHA_ = func(i ...string) string {
	return fn("alpha", i...)
}
var LUMA_ = func(i ...string) string {
	return fn("luma", i...)
}

/*
saturate(@color, 10%);                  // return a color 10% points *more* saturated
desaturate(@color, 10%);                // return a color 10% points *less* saturated
lighten(@color, 10%);                   // return a color 10% points *lighter*
darken(@color, 10%);                    // return a color 10% points *darker*
fadein(@color, 10%);                    // return a color 10% points *less* transparent
fadeout(@color, 10%);                   // return a color 10% points *more* transparent
fade(@color, 50%);                      // return @color with 50% transparency
spin(@color, 10);                       // return a color with a 10 degree larger in hue
mix(@color1, @color2, [@weight: 50%]);  // return a mix of @color1 and @color2
greyscale(@color);                      // returns a grey, 100% desaturated color
contrast(@color1, [@darkcolor: black], [@lightcolor: white], [@threshold: 43%]); 
                                        // return @darkcolor if @color1 is > 43% luma
                                        // otherwise return @lightcolor, see notes
*/
var SATURATE_ = func(i ...string) string {
	return fn("saturate", i...)
}
var DESATURATE_ = func(i ...string) string {
	return fn("desaturate", i...)
}
var LIGHTEN_ = func(i ...string) string {
	return fn("lighten", i...)
}
var DARKEN_ = func(i ...string) string {
	return fn("darken", i...)
}
var FADE_IN_ = func(i ...string) string {
	return fn("fadein", i...)
}
var FADE_OUT_ = func(i ...string) string {
	return fn("fadeout", i...)
}
var FADE_ = func(i ...string) string {
	return fn("fade", i...)
}
var SPIN_ = func(i ...string) string {
	return fn("spin", i...)
}
var MIX_ = func(i ...string) string {
	return fn("mix", i...)
}
var GREYSCALE_ = func(i ...string) string {
	return fn("greyscale", i...)
}
var CONTRAST_ = func(i ...string) string {
	return fn("contrast", i...)
}

/*
multiply(@color1, @color2);
screen(@color1, @color2);
overlay(@color1, @color2);
softlight(@color1, @color2);
hardlight(@color1, @color2);
difference(@color1, @color2);
exclusion(@color1, @color2);
average(@color1, @color2);
negation(@color1, @color2);
*/
var MULTIPLY_ = func(i ...string) string {
	return fn("multiply", i...)
}
var SCREEN_ = func(i ...string) string {
	return fn("screen", i...)
}
var OVERLAY_ = func(i ...string) string {
	return fn("overlay", i...)
}
var SOFTLIGHT_ = func(i ...string) string {
	return fn("softlight", i...)
}
var HARDLIGHT_ = func(i ...string) string {
	return fn("hardlight", i...)
}
var DIFFERENCE_ = func(i ...string) string {
	return fn("difference", i...)
}
var EXCLUSION_ = func(i ...string) string {
	return fn("exclusion", i...)
}
var AVERAGE_ = func(i ...string) string {
	return fn("average", i...)
}
var NEGATION_ = func(i ...string) string {
	return fn("negation", i...)
}

/*
iscolor(@colorOrAnything);              // returns true if passed a color, including keyword colors
isnumber(@numberOrAnything);            // returns true if a number of any unit
isstring(@stringOrAnything);            // returns true if it is passed a string
iskeyword(@keywordOrAnything);          // returns true if it is passed keyword
isurl(@urlOrAnything);                  // returns true if it is a string and a url
ispixel(@pixelOrAnything);              // returns true if it is a number and a px
ispercentage(@percentageOrAnything);    // returns true if it is a number and a %
isem(@emOrAnything);                    // returns true if it is a number and an em
isunit(@numberOrAnything, "rem");       // * returns if a parameter is a number and is in a particular unit
*/
var IS_COLOR_ = func(i ...string) string {
	return fn("iscolor", i...)
}
var IS_NUMBER_ = func(i ...string) string {
	return fn("isnumber", i...)
}
var IS_STRING_ = func(i ...string) string {
	return fn("isstring", i...)
}
var IS_KEYWORD_ = func(i ...string) string {
	return fn("iskeyword", i...)
}
var IS_URL_ = func(i ...string) string {
	return fn("isurl", i...)
}
var IS_PIXEL_ = func(i ...string) string {
	return fn("ispixel", i...)
}
var IS_PERCENTAGE_ = func(i ...string) string {
	return fn("ispercentage", i...)
}
var IS_EM_ = func(i ...string) string {
	return fn("isem", i...)
}
var IS_UNIT_ = func(i ...string) string {
	return fn("isunit", i...)
}
