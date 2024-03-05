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

	//sample source #1 : https://uxplanet.org/deep-dive-into-wcag-2-1-with-html-css-examples-59845b2174c4#:~:text=It's%20important%20to%20note%20that%20sufficient%20contrast%20between%20text%20and,14pt%20bold%20or%2018pt%20regular).
	//sample source #1 : https://www.w3.org/TR/2016/NOTE-WCAG20-TECHS-20161007/H95
	if (node.Parent.Data == "video" || node.Parent.Data == "object" || node.Parent.Data == "embed") &&
		node.Data == "track" &&
		!helper.IsAttributeKeyValueMatching(node.Attr, "kind", "caption") {
		rule.Rules.WCAG122.H95 = Fail
		AddErrorAttribute(node, Wcag122, "H95")
	}
}
