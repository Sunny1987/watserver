package rule

import (
	"golang.org/x/net/html"
	"reflect"
	"webserver/analyzerapp/helper"
)

type WCAG122 struct {
	H95 string
}

// ExecuteWCAG122 executes the WCAG122 techniques
func (rule *RuleResults) ExecuteWCAG122(node *html.Node) (string, []string) {
	rule.Logger.Printf("...intiating WCAG122 for %v ", node.Data)

	//Refresh struct
	wcag122 := WCAG122{}
	rule.Rules.WCAG122 = wcag122

	//implement the techniques
	rule.Logger.Println("....Execute H95")
	rule.H95Technique(node)

	return Wcag122, rule.Rules.WCAG122.GetRuleFailures()
}

// GetRuleFailures will get the list of Techniques failures
func (rule WCAG122) GetRuleFailures() []string {
	var techniques []string
	structVal := reflect.ValueOf(rule)
	for i := 0; i < structVal.NumField(); i++ {
		field := structVal.Field(i)
		name := structVal.Type().Field(i).Name
		value := field.Interface()

		if value == Fail {
			techniques = append(techniques, name)
		}
	}
	return techniques
}

// H95Technique analysis for track tags
func (rule *RuleResults) H95Technique(node *html.Node) {
	if (node.Parent.Data == "video" || node.Parent.Data == "object" || node.Parent.Data == "embed") &&
		node.Data == "track" &&
		helper.IsAttributeKeyValueMatching(node.Attr, "kind", "caption") {
		rule.Rules.WCAG122.H95 = Fail
	}

	if (node.Parent.Data == "video" || node.Parent.Data == "object" || node.Parent.Data == "embed") &&
		helper.HasNoChild(node) &&
		helper.IsAttributeValueContaining(node.Attr, "src", "caption") {
		rule.Rules.WCAG122.H95 = Fail
	}

	if (node.Parent.Data == "video" || node.Parent.Data == "object" || node.Parent.Data == "embed") &&
		helper.IsAttributePresent(node.Attr, "ariadescribedby") {
		attval := helper.GetAttribute(node.Attr, "ariadescribedby")
		for c := node.Parent.FirstChild; c != nil; c = c.NextSibling {
			if helper.IsAttributePresent(c.Attr, "id") {
				if helper.IsAttributeValueContaining(c.Attr, "id", attval) {
					rule.Rules.WCAG122.H95 = Fail
					break
				}
			}
		}
	}
}
