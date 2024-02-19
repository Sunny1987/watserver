package rule

import (
	"golang.org/x/net/html"
	"reflect"
)

//type WCAG111rule interface {
//	Aria6Technique(node *html.Node) (string, error)
//	Aria10Technique(node *html.Node) (string, error)
//	G94Technique(node *html.Node) (string, error)
//	H2Technique(node *html.Node) (string, error)
//	H35Technique(node *html.Node) (string, error)
//	H53Technique(node *html.Node) (string, error)
//	Aria15Technique(node *html.Node) (string, error)
//	H86Technique(node *html.Node) (string, error)
//	H86CSSTechnique(css string, node *html.Node) (string, error)
//}

// WCAG111 rule object for Analysis
type WCAG111 struct {
	Aria6  string
	Aria10 string
	G94    string
	H2     string
	H35    string
	H53    string
	Aria15 string
	H86    string
	H86CSS string
}

// ExecuteWCAG111 executes the WCAG111 techniques
func (rule *RuleResults) ExecuteWCAG111(node *html.Node) (string, []*string) {
	rule.Logger.Printf("...intiating WCAG111 for %v ", node.Data)

	//implement the techniques
	rule.Logger.Println("....Execute ARIA6")
	rule.Aria6Technique(node)

	return "WCAG111", rule.Rules.WCAG111.GetRuleFailures()
}

// GetRuleFailures will get the list of Techniques failures
func (rule WCAG111) GetRuleFailures() []*string {
	var techniques []*string
	structVal := reflect.ValueOf(rule)
	for i := 0; i < structVal.NumField(); i++ {
		field := structVal.Field(i)
		name := structVal.Type().Field(i).Name
		value := field.Interface()

		if value == Fail {
			techniques = append(techniques, &name)
		}
	}
	return techniques
}

// Aria6Technique analysis is invoked for all tags
func (rule *RuleResults) Aria6Technique(node *html.Node) {
	rule.Logger.Println("...Initiating ARIA6 analysis")

	//logic implementation
	if attributeSearch(node.Attr, "role") && attributeCheckValEmpty(node.Attr, "aria-label") {
		rule.Rules.WCAG111.Aria6 = Fail
	} else {
		rule.Rules.WCAG111.Aria6 = Pass
	}
}
