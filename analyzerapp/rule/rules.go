package rule

import (
	"golang.org/x/net/html"
	"log"
	"reflect"
	"webserver/resultsapp"
)

const (
	Fail = "Fail"
	Pass = "Pass"
)

func NewResult(guideline *string, rules []*string) *resultsapp.Result {
	return &resultsapp.Result{Guideline: guideline, Rules: rules}
}

// RuleResults is the finalized rules results for tags
type RuleResults struct {
	Results []*resultsapp.Result
	Rules   Rules
	Logger  *log.Logger
	Css     string
}

// Rules carries the rules results
type Rules struct {
	WCAG111 WCAG111
}

func (rule Rules) ClearRules() {
	p := reflect.ValueOf(rule).Elem()
	p.Set(reflect.Zero(p.Type()))
}

// Execute method executes all the rules
func (rule *RuleResults) Execute(node *html.Node) bool {
	var results []*resultsapp.Result
	var status bool

	//Execute WCAG111 guideline
	guideline, techniques := rule.ExecuteWCAG111(node)
	status, results = CompileResults(guideline, techniques, results)

	//update Results in RuleResults
	rule.Results = results

	return status
}

// CompileResults will consolidate the results
func CompileResults(guideline string, techniques []*string, results []*resultsapp.Result) (bool, []*resultsapp.Result) {
	var status bool
	if len(techniques) > 0 {
		result := NewResult(&guideline, techniques)
		results = append(results, result)
		status = true
	}
	return status, results
}
