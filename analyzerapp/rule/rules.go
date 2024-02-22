package rule

import (
	"log"
	"reflect"
	"sync"
	"webserver/resultsapp"

	"golang.org/x/net/html"
)

const (
	Fail = "Fail"
	Pass = "Pass"
)

func NewResult(guideline string, rules []string) *resultsapp.Result {
	return &resultsapp.Result{Guideline: guideline, Rules: rules}
}

// RuleResults is the finalized rules results for tags
type RuleResults struct {
	Results []*resultsapp.Result
	Rules   Rules
	Logger  *log.Logger
	Css     string
	sync.Mutex
}

// Rules carries the rules results
type Rules struct {
	WCAG111 WCAG111
	WCAG121 WCAG121
}

func (rule Rules) ClearRules() {
	p := reflect.ValueOf(rule).Elem()
	p.Set(reflect.Zero(p.Type()))
}

// Execute method executes all the rules
func (rule *RuleResults) Execute(node *html.Node) bool {
	rule.Lock()
	defer rule.Unlock()
	var results []*resultsapp.Result
	var status bool

	//Execute WCAG111 guideline
	guideline, techniques := rule.ExecuteWCAG111(node)
	//status = UpdateRuleList(guideline, techniques, &results)
	if len(techniques) > 0 {
		result := NewResult(guideline, techniques)
		results = append(results, result)
		status = true
		techniques = nil
		guideline = ""
	}

	//Execute WCAG121 guideline
	guideline, techniques = rule.ExecuteWCAG121(node)
	//status = UpdateRuleList(guideline, techniques, &results)
	if len(techniques) > 0 {
		result := NewResult(guideline, techniques)
		results = append(results, result)
		status = true
		techniques = nil
		guideline = ""
	}

	rule.Results = results
	return status
}

//func UpdateRuleList(guideline string, techniques []string, results *[]*resultsapp.Result) bool {
//	if len(techniques) > 0 {
//		result := NewResult(guideline, techniques)
//		*results = append(*results, result)
//		techniques = nil
//		guideline = ""
//		return true
//	}
//	return false
//}
