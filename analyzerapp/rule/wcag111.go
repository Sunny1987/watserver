package rule

import (
	"reflect"
	"strings"
	"webserver/analyzerapp/helper"

	"golang.org/x/net/html"
)

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
func (rule *RuleResults) ExecuteWCAG111(node *html.Node) (string, []string) {
	rule.Logger.Printf("...intiating WCAG111 for %v ", node.Data)

	//Refresh struct
	wcag111 := WCAG111{}
	rule.Rules.WCAG111 = wcag111

	//implement the techniques
	rule.Logger.Println("....Execute ARIA6")
	rule.Aria6Technique(node)
	rule.Logger.Println("....Execute ARIA10")
	rule.Aria10Technique(node)
	rule.Logger.Println("....Execute ARIA15")
	rule.ARIA15Technique(node)
	rule.Logger.Println("....Execute G94")
	rule.G94Technique(node)
	rule.Logger.Println("....Execute H2")
	rule.H2Technique(node)
	rule.Logger.Println("....Execute H35")
	rule.H35Technique(node)
	rule.Logger.Println("....Execute H45")
	rule.H45Technique(node)
	rule.Logger.Println("....Execute H86")
	rule.H86Technique(node)

	return Wcag111, rule.Rules.WCAG111.GetRuleFailures()
}

// GetRuleFailures will get the list of Techniques failures
func (rule WCAG111) GetRuleFailures() []string {
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

// Aria6Technique analysis is invoked for all tags
func (rule *RuleResults) Aria6Technique(node *html.Node) {
	rule.Logger.Println("...Initiating ARIA6 analysis")

	//logic implementation
	if helper.IsAttributePresent(node.Attr, "role") && helper.IsAttributeValueEmpty(node.Attr, "aria-label") {
		rule.Rules.WCAG111.Aria6 = Fail
		AddErrorAttribute(node, Wcag111, "Aria6")
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
		if helper.IsAttributeValueEmpty(node.Attr, "aria-labelledby") {
			for c := node.FirstChild; c != nil; c = c.NextSibling {
				if helper.IsAttributePresent(c.Attr, "src") {
					if helper.IsAttributeValueEmpty(node.Attr, "alt") {
						rule.Rules.WCAG111.Aria10 = Fail
						AddErrorAttribute(node, Wcag111, "Aria10")
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
		if helper.IsAttributeKeyValueMatching(node.Attr, "class", className) {
			if helper.IsAttributePresent(node.Attr, "alt") {
				rule.Rules.WCAG111.H86CSS = Fail
				AddErrorAttribute(node, Wcag111, "H86CSS")
			}
		}

	}
}

// H45Technique analysis for all tags
func (rule *RuleResults) H45Technique(node *html.Node) {
	if helper.IsAttributePresent(node.Attr, "longdesc") {
		if helper.IsAttributeValueEmpty(node.Attr, "longdesc") {
			rule.Rules.WCAG111.H45 = Fail
			AddErrorAttribute(node, Wcag111, "H45")
		}
	}
}

// ARIA15Technique analysis for all tags
func (rule *RuleResults) ARIA15Technique(node *html.Node) {
	if helper.IsAttributePresent(node.Attr, "aria-describedby") {
		if helper.IsAttributeValueEmpty(node.Attr, "aria-describedby") {
			rule.Rules.WCAG111.Aria15 = Fail
			AddErrorAttribute(node, Wcag111, "Aria15")
		}
	}
}

// H2Technique analysis for all tags
func (rule *RuleResults) H2Technique(node *html.Node) {
	if node.Parent.Data == "a" {
		if helper.IsAttributeValueEmpty(node.Attr, "alt") {
			rule.Rules.WCAG111.H2 = Fail
			AddErrorAttribute(node, Wcag111, "H2")
		}
	}
}

// H35Technique analysis for all tags
func (rule *RuleResults) H35Technique(node *html.Node) {
	if helper.HasOneChild(node) {
		if helper.IsTextNode(node.FirstChild) {
			if helper.IsAttributeValueEmpty(node.Attr, "alt") {
				rule.Rules.WCAG111.H35 = Fail
				AddErrorAttribute(node, Wcag111, "H35")
			}

		}
	}
}

// G94Technique analysis for all tags
func (rule *RuleResults) G94Technique(node *html.Node) {
	rule.Logger.Println(node.Data)
	if node.Data == "img" && helper.IsAttributeValueEmpty(node.Attr, "alt") ||
		helper.IsAttributeKeyValueMatching(node.Attr, "role", "img") && helper.IsAttributeValueEmpty(node.Attr, "aria-labelledby") ||
		helper.IsAttributeKeyValueMatching(node.Attr, "role", "img") && helper.IsAttributeValueEmpty(node.Attr, "aria-label") ||
		node.Data == "img" && helper.IsAttributeValueEmpty(node.Attr, "title") {
		rule.Rules.WCAG111.G94 = Fail
		AddErrorAttribute(node, Wcag111, "G94")
	}

	if node.Data == "area" && helper.IsAttributeValueEmpty(node.Attr, "alt") {
		rule.Rules.WCAG111.G94 = Fail
		AddErrorAttribute(node, Wcag111, "G94")
	}

	if node.Data == "svg" && helper.IsAttributePresent(node.Attr, "aria-label") && helper.IsAttributeKeyValueMatching(node.Attr, "role", "img") {
		if helper.IsAttributeValueEmpty(node.Attr, "aria-label") {
			rule.Rules.WCAG111.G94 = Fail
			AddErrorAttribute(node, Wcag111, "G94")
		}
	}

	if node.Data == "canvas" && helper.IsAttributePresent(node.Attr, "aria-label") {
		if helper.IsAttributeValueEmpty(node.Attr, "aria-label") {
			rule.Rules.WCAG111.G94 = Fail
			AddErrorAttribute(node, Wcag111, "G94")

		}
	}

}
