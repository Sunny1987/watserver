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
	G158 string
}

// ExecuteWCAG121 executes the WCAG121 techniques
func (rule *RuleResults) ExecuteWCAG121(input Inputs) (string, []string) {
	node := input.node
	rule.Logger.Printf("...intiating WCAG121 for %v ", node.Data)

	//Refresh struct
	wcag121 := WCAG121{}
	rule.Rules.WCAG121 = wcag121

	//implement the techniques
	rule.Logger.Println("....Execute H96")
	rule.H96Technique(node)
	rule.Logger.Println("....Execute G158")
	rule.G158Technique(node)

	return Wcag121, rule.Rules.WCAG121.GetRuleFailures()
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

// sample source : https://www.w3.org/TR/2016/NOTE-WCAG20-TECHS-20161007/H96

// H96Technique analysis for video and audio tags
func (rule *RuleResults) H96Technique(node *html.Node) {
	if node.Parent.Data == "audio" || node.Parent.Data == "video" && node.Data == "track" {
		if !helper.IsAttributeKeyValueMatching(node.Attr, "kind", "captions") &&
			!helper.IsAttributeKeyValueMatching(node.Attr, "kind", "descriptions") {
			rule.Rules.WCAG121.H96 = Fail
			AddErrorAttribute(node, Wcag121, "H96")
		}
	}
}

// https://act-rules.github.io/rules/2eb176
// Implementing this example
//<html lang="en">
//	<audio src="/test-assets/moon-audio/moon-speech.mp3" controls></audio>
//	<a href="/test-assets/moon-audio/moon-speech-transcript.txt">Transcript</a>
//</html>

// G158Technique for audio tags
func (rule *RuleResults) G158Technique(node *html.Node) {
	if node.Data == "audio" &&
		node.NextSibling.Data == "a" &&
		helper.IsAttributeKeyValueMatching(node.Attr, "controls", "") {
		if !helper.IsAttributeValueContaining(node.NextSibling.Attr, "href", "transcript") &&
			!strings.Contains(helper.Text(node.NextSibling), "transcript") {
			rule.Rules.WCAG121.G158 = Fail
			AddErrorAttribute(node, Wcag121, "G158")
		}
	}
}
