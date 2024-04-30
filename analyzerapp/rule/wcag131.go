package rule

import (
	"golang.org/x/net/html"
	"reflect"
	"webserver/analyzerapp/helper"
)

type WCAG131 struct {
	H44   string
	G117  string
	H51   string
	H42   string
	ARIA6 string
}

// ExecuteWCAG131 executes the WCAG131 techniques
func (rule *RuleResults) ExecuteWCAG131(input Inputs) (string, []string) {
	node := input.node
	rule.Logger.Printf("...intiating WCAG131 for %v ", node.Data)

	//Refresh struct
	wcag131 := WCAG131{}
	rule.Rules.WCAG131 = wcag131

	//implement the techniques
	rule.Logger.Println("....Execute H44")
	rule.H44Technique(node)

	return Wcag131, rule.Rules.WCAG131.GetRuleFailures()
}

// GetRuleFailures will get the list of Techniques failures
func (rule WCAG131) GetRuleFailures() []string {
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

// H44Technique analysis for label tags
func (rule *RuleResults) H44Technique(node *html.Node) {
	//<!-- Incorrect: Missing label for input -->
	//<input type="text" id="username" name="username">
	//
	//<!-- Correct: Input with associated label -->
	//<label for="username">Username:</label>
	//<input type="text" id="username" name="username">

	if node.Data == "label" {
		if helper.IsAttributePresent(node.Attr, "for") {
			idVal := helper.GetAttribute(node.Attr, "for")
			if !helper.IsAttributePresent(node.NextSibling.Attr, "id") ||
				!helper.IsAttributeKeyValueMatching(node.NextSibling.Attr, "id", idVal) {
				rule.Rules.WCAG131.H44 = Fail
				AddErrorAttribute(node, Wcag131, "H44")
			}
		}
	}

}

// G117Technique analysis for div/nav tags
func (rule *RuleResults) G117Technique(node *html.Node) {
	//<!-- Incorrect: Using divs instead of semantic elements -->
	//<div class="nav">
	//<div><a href="/">Home</a></div>
	//<div><a href="/about">About</a></div>
	//<div><a href="/contact">Contact</a></div>
	//</div>
	//
	//<!-- Correct: Using semantic elements -->
	//<nav>
	//<ul>
	//<li><a href="/">Home</a></li>
	//<li><a href="/about">About</a></li>
	//<li><a href="/contact">Contact</a></li>
	//</ul>
	//</nav>

	if node.Data == "div" && helper.IsAttributeKeyValueMatching(node.Attr, "class", "nav") {
		rule.Rules.WCAG131.G117 = Fail
		AddErrorAttribute(node, Wcag131, "G117")
	}

}

// H42Technique analysis for headers tags
func (rule *RuleResults) H42Technique(node *html.Node) {
	//<!-- Incorrect: Skipping heading levels -->
	//<h1>Main Heading</h1>
	//<h3>Subheading</h3>
	//
	//<!-- Correct: Using sequential heading levels -->
	//<h1>Main Heading</h1>
	//<h2>Subheading</h2>

	if node.Data == "h1" {
		if node.NextSibling.Data == "h3" || node.NextSibling.Data == "h4" || node.NextSibling.Data == "h5" || node.NextSibling.Data == "h6" {
			rule.Rules.WCAG131.H42 = Fail
			AddErrorAttribute(node, Wcag131, "H42")
		}
	}

	if node.Data == "h2" {
		if node.NextSibling.Data == "h1" || node.NextSibling.Data == "h4" || node.NextSibling.Data == "h5" || node.NextSibling.Data == "h6" {
			rule.Rules.WCAG131.H42 = Fail
			AddErrorAttribute(node, Wcag131, "H42")
		}
	}

	if node.Data == "h3" {
		if node.NextSibling.Data == "h1" || node.NextSibling.Data == "h2" || node.NextSibling.Data == "h5" || node.NextSibling.Data == "h6" {
			rule.Rules.WCAG131.H42 = Fail
			AddErrorAttribute(node, Wcag131, "H42")
		}
	}

	if node.Data == "h4" {
		if node.NextSibling.Data == "h1" || node.NextSibling.Data == "h2" || node.NextSibling.Data == "h3" || node.NextSibling.Data == "h6" {
			rule.Rules.WCAG131.H42 = Fail
			AddErrorAttribute(node, Wcag131, "H42")
		}
	}

	if node.Data == "h5" {
		if node.NextSibling.Data == "h1" || node.NextSibling.Data == "h2" || node.NextSibling.Data == "h3" || node.NextSibling.Data == "h4" {
			rule.Rules.WCAG131.H42 = Fail
			AddErrorAttribute(node, Wcag131, "H42")
		}
	}

}
