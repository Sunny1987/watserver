package rule

import (
	"log"
	"webserver/resultsapp"

	"golang.org/x/net/html"
)

const (
	Fail = "Fail"
)

// RuleResults is the finalized rules results for tags
type RuleResults struct {
	Results []resultsapp.Result
	Rules   Rules
	Logger  *log.Logger
	Css     string
	status  bool
}

func NewRuleResults(logger *log.Logger) *RuleResults {
	return &RuleResults{Logger: logger}
}

// Rules carries the rules results
type Rules struct {
	WCAG111 WCAG111
	WCAG121 WCAG121
	WCAG122 WCAG122
}

// Execute method executes all the rules
func (rule *RuleResults) Execute(node *html.Node) bool {

	var results []resultsapp.Result

	//Execute WCAG111 guideline
	guideline, techniques := rule.ExecuteWCAG111(node)
	results = append(results, rule.UpdateRuleList(guideline, techniques)...)

	//Execute WCAG121 guideline
	guideline, techniques = rule.ExecuteWCAG121(node)
	results = append(results, rule.UpdateRuleList(guideline, techniques)...)

	//Execute WCAG122 guideline
	guideline, techniques = rule.ExecuteWCAG122(node)
	results = append(results, rule.UpdateRuleList(guideline, techniques)...)

	rule.Results = results
	return rule.status

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
