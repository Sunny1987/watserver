package rule

import (
	"golang.org/x/net/html"
)

// attributeSearch will respond if the attribute is present
func attributeSearch(attrList []html.Attribute, key string) bool {
	for _, att := range attrList {
		if att.Key == key {
			return true
		}
	}
	return false
}

// attributeCheckVal will respond if the attribute is present and check its value
func attributeCheckVal(attrList []html.Attribute, key string, val string) bool {
	for _, att := range attrList {
		if att.Key == key && att.Val == val {
			return true
		}
	}
	return false
}

// attributeCheckValEmpty will respond if the attribute is present and value is empty
func attributeCheckValEmpty(attrList []html.Attribute, key string) bool {
	for _, att := range attrList {
		if att.Key == key && att.Val != "" {
			return true
		}
	}
	return false
}
