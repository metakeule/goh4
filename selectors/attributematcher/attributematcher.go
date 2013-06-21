package attributematcher

import "fmt"

type attrMatcher struct {
	attr      string
	operator  string
	compareTo string
}

func (ø attrMatcher) Selector() string {
	if ø.operator == "" && ø.compareTo == "" {
		return fmt.Sprintf(`[%s]`, ø.attr)
	}

	return fmt.Sprintf(`[%s%s"%s"]`, ø.attr, ø.operator, ø.compareTo)
}

// E[foo]	an E element with a "foo" attribute	Attribute selectors	2
func Exists(attr string) attrMatcher { return attrMatcher{attr, "", ""} }

// E[foo="bar"]	an E element whose "foo" attribute value is exactly equal to "bar"	Attribute selectors	2
func Equals(attr string, compareTo string) attrMatcher { return attrMatcher{attr, "=", compareTo} }

// E[foo~="bar"]	an E element whose "foo" attribute value is a list of whitespace-separated values, one of which is exactly equal to "bar"	Attribute selectors	2
func Includes(attr string, compareTo string) attrMatcher { return attrMatcher{attr, "~=", compareTo} }

// E[foo^="bar"]	an E element whose "foo" attribute value begins exactly with the string "bar"	Attribute selectors	3
func BeginsWith(attr string, compareTo string) attrMatcher { return attrMatcher{attr, "^=", compareTo} }

// E[foo$="bar"]	an E element whose "foo" attribute value ends exactly with the string "bar"	Attribute selectors	3
func EndsWith(attr string, compareTo string) attrMatcher { return attrMatcher{attr, "$=", compareTo} }

// E[foo*="bar"]	an E element whose "foo" attribute value contains the substring "bar"	Attribute selectors	3
func Contains(attr string, compareTo string) attrMatcher { return attrMatcher{attr, "*=", compareTo} }

// E[foo|="en"]	an E element whose "foo" attribute has a hyphen-separated list of values beginning (from the left) with "en"
func HyphenBeginsWith(attr string, compareTo string) attrMatcher {
	return attrMatcher{attr, "|=", compareTo}
}
