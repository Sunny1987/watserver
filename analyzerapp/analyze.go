package analyzerapp

import (
	"log"
	"sync"
	"webserver/analyzerapp/helper"
	"webserver/analyzerapp/rule"
	"webserver/resultsapp"

	"golang.org/x/net/html"
)

var wg sync.WaitGroup
var MU sync.Mutex

// AnalyzeBundle is the object for Analyser
type AnalyzeBundle struct {
	Req           *resultsapp.MyRequest
	Logger        *log.Logger
	Person        *string //the property is currently optional
	Base          string
	Doc           *html.Node
	CollectedTags resultsapp.TagsFamily
	rules         *rule.RuleResults
	FinalResponse resultsapp.FinalResponse
}

// NewAnalyzeBundle is the constructor for AnalyzeBundle
func NewAnalyzeBundle(req *resultsapp.MyRequest, logger *log.Logger, base string, doc *html.Node, collectedTags resultsapp.TagsFamily) *AnalyzeBundle {
	return &AnalyzeBundle{Req: req, Logger: logger, Base: base, Doc: doc, CollectedTags: collectedTags}
}

func (aBundle *AnalyzeBundle) Analyze() resultsapp.FinalResponse {
	aBundle.FinalResponse.Request = aBundle.Req

	wg.Add(42)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Anchors)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Audios)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Areas)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Abbrs)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Asides)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Buttons)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Bodys)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Canvases)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Divs)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Dirs)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Embeds)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Forms)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Footers)
	go aBundle.tagAnalysis(aBundle.CollectedTags.H1s)
	go aBundle.tagAnalysis(aBundle.CollectedTags.H2s)
	go aBundle.tagAnalysis(aBundle.CollectedTags.H3s)
	go aBundle.tagAnalysis(aBundle.CollectedTags.H4s)
	go aBundle.tagAnalysis(aBundle.CollectedTags.H5s)
	go aBundle.tagAnalysis(aBundle.CollectedTags.H6s)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Headers)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Heads)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Inputs)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Imgs)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Iframes)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Links)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Labels)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Mains)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Navs)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Objects)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Paras)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Pres)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Selects)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Svgs)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Spans)
	go aBundle.tagAnalysis(aBundle.CollectedTags.TextAreas)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Tracks)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Titles)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Tables)
	go aBundle.tagAnalysis(aBundle.CollectedTags.THeads)
	go aBundle.tagAnalysis(aBundle.CollectedTags.TBodys)
	go aBundle.tagAnalysis(aBundle.CollectedTags.TFoots)
	go aBundle.tagAnalysis(aBundle.CollectedTags.Videos)

	wg.Wait()
	return aBundle.FinalResponse
}

// tagAnalysis helps in all types of tag rules analysis
func (aBundle *AnalyzeBundle) tagAnalysis(nodes []*html.Node) {
	defer wg.Done()
	if len(nodes) == 0 {
		aBundle.Logger.Println("no tags collected")
		return
	}
	t := nodes[0].Data
	aBundle.Logger.Printf("Initiating %v Analysis......", t)

	var list []resultsapp.TagResult
	for _, node := range nodes {
		var tag resultsapp.TagResult

		//build the node
		tag.Tag = helper.NodeText(node)

		//refresh ruleResult
		ruleResult := rule.NewRuleResults(aBundle.Logger)
		aBundle.rules = ruleResult

		if status, results := aBundle.rules.Execute(node); status == true {
			tag.Result = results
			if len(list) < 50 {
				list = append(list, tag)
			}
		}
	}
	MU.Lock()
	aBundle.FinalResponse.Results = append(aBundle.FinalResponse.Results, list...)
	MU.Unlock()
}

//// cssAnalysis function initiates all the CSS rule based analysis
//func (aBundle *AnalyzeBundle) cssAnalysis() {
//	aBundle.Logger.Println("Initiating CSS Analysis......")
//	defer wg.Done()
//	var list []resultsapp.CSS
//	nodes := aBundle.CollectedTags.Divs
//	cssLinks := aBundle.CollectedTags.CssLinks
//	for _, css := range cssLinks {
//		aBundle.Logger.Printf("CSS : %v ", css)
//		var tag resultsapp.CSS
//
//		//add css data
//		tag.CSS = css
//
//		//implement rule
//		aBundle.rules.Css = css
//		for _, node := range nodes {
//			aBundle.rules.Logger = aBundle.Logger
//			if status, results := aBundle.rules.Execute(node); status == true {
//				tag.Result = results
//				if len(list) < 50 {
//					list = append(list, tag)
//				}
//			}
//		}
//
//	}
//	aBundle.Response.CSSResults = list
//}
