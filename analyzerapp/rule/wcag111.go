package rule

import (
	"reflect"
	"strings"
	"webserver/analyzerapp/helper"

	"golang.org/x/net/html"
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
	H45    string
}

// ExecuteWCAG111 executes the WCAG111 techniques
func (rule *RuleResults) ExecuteWCAG111(node *html.Node) (string, []*string) {
	rule.Logger.Printf("...intiating WCAG111 for %v ", node.Data)

	//implement the techniques
	rule.Logger.Println("....Execute ARIA6")
	rule.Aria6Technique(node)
	rule.Aria10Technique(node)

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
	if helper.AttributeSearch(node.Attr, "role") && helper.AttributeCheckValEmpty(node.Attr, "aria-label") {
		rule.Rules.WCAG111.Aria6 = Fail
	}
}

// Aria10Technique analysis is invoked for all tags
func (rule *RuleResults) Aria10Technique(node *html.Node) {
	rule.Logger.Println("...Initiating ARIA10 analysis")

	//*******************logic implementation**********//
	//The implementation is as per https://www.w3.org/WAI/WCAG21/Techniques/aria/ARIA10
	//This example shows how to use the aria-labelledby attribute to provide a short text description for a read-only complex graphic of a star rating pattern; the graphic is composed of several image elements. The text alternative for the graphic is the label, visible on the page beneath the star pattern.
	//
	//<div role="img" aria-labelledby="star_id">
	//<img src="fullstar.png" alt=""/>
	//<img src="fullstar.png" alt=""/>
	//<img src="fullstar.png" alt=""/>
	//<img src="fullstar.png" alt=""/>
	//<img src="emptystar.png" alt=""/>
	//</div>
	//
	//<div id="star_id">4 of 5</div>
	if helper.HasChildren(node) {
		if helper.AttributeCheckValEmpty(node.Attr, "aria-labelledby") {
			for c := node.FirstChild; c != nil; c = c.NextSibling {
				if helper.AttributeSearch(c.Attr, "src") {
					if helper.AttributeCheckValEmpty(node.Attr, "alt") {
						rule.Rules.WCAG111.Aria10 = Fail
					}
				}
			}
		}
	}
}

// H86Technique analysis invoked for div tag
func (rule *RuleResults) H86Technique(node *html.Node) {
	if strings.Contains(rule.Css, "white-space: pre") {
		name := strings.Split(rule.Css, " ")
		className := strings.Trim(name[0], "@")
		if helper.AttributeCheckVal(node.Attr, "class", className) {
			if helper.AttributeSearch(node.Attr, "alt") {
				rule.Rules.WCAG111.H86CSS = Fail
			}
		}

	}
}

// H45Technique analysis for all tags
func (rule *RuleResults) H45Technique(node *html.Node) {
	if helper.AttributeSearch(node.Attr, "longdesc") {
		if helper.AttributeCheckValEmpty(node.Attr, "longdesc") {
			rule.Rules.WCAG111.H45 = Fail
		}
	}
}
