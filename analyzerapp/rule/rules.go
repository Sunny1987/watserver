package rule

import (
	"log"
	"webserver/resultsapp"

	"golang.org/x/net/html"
)

const (
	Fail    = "Fail"
	Wcag111 = "WCAG111"
	Wcag121 = "WCAG121"
	Wcag122 = "WCAG122"
	Wcag131 = "WCAG131"
)

// RuleResults is the finalized rules results for tags
type RuleResults struct {
	//Results *[]resultsapp.Result
	Rules  Rules
	Logger *log.Logger
	Css    string
	status bool
}

// Inputs struct will handle parameters for execute for optional css scenario
type Inputs struct {
	node *html.Node
	css  string
}

// NewInputs is the constructor for Inputs
func NewInputs(node *html.Node) *Inputs {
	return &Inputs{node: node}
}

// NewInputsWithCSS is the constructor for Inputs
func NewInputsWithCSS(node *html.Node, css string) *Inputs {
	return &Inputs{node: node, css: css}
}

func NewRuleResults(logger *log.Logger) *RuleResults {
	return &RuleResults{Logger: logger}
}

// Rules carries the rules results
type Rules struct {
	WCAG111 WCAG111
	WCAG121 WCAG121
	WCAG122 WCAG122
	WCAG131 WCAG131
}

// Execute method executes all the rules
func (rule *RuleResults) Execute(input Inputs) (bool, []resultsapp.Result) {

	var results []resultsapp.Result

	//Execute WCAG111 guideline
	guideline, techniques := rule.ExecuteWCAG111(input)
	results = append(results, rule.UpdateRuleList(guideline, techniques)...)

	//Execute WCAG121 guideline
	guideline, techniques = rule.ExecuteWCAG121(input)
	results = append(results, rule.UpdateRuleList(guideline, techniques)...)

	//Execute WCAG122 guideline
	guideline, techniques = rule.ExecuteWCAG122(input)
	results = append(results, rule.UpdateRuleList(guideline, techniques)...)

	//Execute WCAG131 guideline
	guideline, techniques = rule.ExecuteWCAG131(input)
	results = append(results, rule.UpdateRuleList(guideline, techniques)...)

	return rule.status, results

}

// UpdateRuleList will filter and update the result slice
func (rule *RuleResults) UpdateRuleList(guideline string, techniques []string) []resultsapp.Result {
	var results []resultsapp.Result
	if len(techniques) > 0 {
		result := resultsapp.NewResult(guideline, techniques)
		results = append(results, result)
		techniques = nil
		guideline = ""
		rule.status = true
	}
	return results
}

func AddErrorAttribute(node *html.Node, key, value string) {
	node.Attr = append(node.Attr, html.Attribute{
		Key: key,
		Val: value,
	})
}
