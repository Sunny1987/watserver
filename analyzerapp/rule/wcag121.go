package rule

import (
	"reflect"
	"strings"
	"webserver/analyzerapp/helper"

	"golang.org/x/net/html"
)

// WCAG121 rule object for Analysis
type WCAG121 struct {
	H96  string
	G166 string
}

// ExecuteWCAG121 executes the WCAG121 techniques
func (rule *RuleResults) ExecuteWCAG121(node *html.Node) (string, []string) {
	rule.Logger.Printf("...intiating WCAG121 for %v ", node.Data)

	//Refresh struct
	wcag121 := WCAG121{}
	rule.Rules.WCAG121 = wcag121

	//implement the techniques
	rule.Logger.Println("....Execute H96")
	rule.H96Technique(node)
	rule.Logger.Println("....Execute G166")
	rule.G166Technique(node)

	return "WCAG121", rule.Rules.WCAG121.GetRuleFailures()
}

// GetRuleFailures will get the list of Techniques failures
func (rule WCAG121) GetRuleFailures() []string {
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

// H96Technique analysis for video and audio tags
func (rule *RuleResults) H96Technique(node *html.Node) {
	if node.Parent.Data == "audio" || node.Parent.Data == "video" {
		if helper.AttributeCheckVal(node.Attr, "kind", "captions") ||
			helper.AttributeCheckVal(node.Attr, "kind", "descriptions") &&
				helper.AttributeCheckValEmpty(node.Attr, "label") {
			rule.Rules.WCAG121.H96 = Fail
		}
	}
}

var listEx = []string{".mpg", ".mpeg", ".avi", ".wmv", ".mov", ".rm", ".ram", ".swf", ".flv", ".ogg", ".mp4"}

// G166Technique analysis for all tags
func (rule *RuleResults) G166Technique(node *html.Node) {
	for _, attr := range node.Attr {
		if attr.Key == "data" {
			for _, item := range listEx {
				if strings.Contains(attr.Val, item) {
					rule.Rules.WCAG121.G166 = Fail
				}
			}
		}
	}
}
