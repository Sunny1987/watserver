package resultsapp

// Result object will be used to simplify the tags result rules
type Result struct {
	Guideline string   `json:"guideline"`
	Rules     []string `json:"rules"`
}

// NewResult is the constructor for Result object
func NewResult(guideline string, rules []string) Result {
	return Result{Guideline: guideline, Rules: rules}
}

// TagResult is the generic Tag for all Html tags
type TagResult struct {
	Tag    string   `json:"tag"`
	Result []Result `json:"result"`
}
