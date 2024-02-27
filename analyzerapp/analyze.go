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
	//MyTags        resultsapp.Tags
	Response resultsapp.Response
	rules    *rule.RuleResults
}

// NewAnalyzeBundle is the constructor for AnalyzeBundle
func NewAnalyzeBundle(req *resultsapp.MyRequest, logger *log.Logger, base string, doc *html.Node, collectedTags resultsapp.TagsFamily) *AnalyzeBundle {
	return &AnalyzeBundle{Req: req, Logger: logger, Base: base, Doc: doc, CollectedTags: collectedTags}
}

func (aBundle *AnalyzeBundle) Analyze() resultsapp.Response {
	aBundle.Response.Request = aBundle.Req

	wg.Add(42)
	go aBundle.anchorAnalysis()
	go aBundle.audioAnalysis()
	go aBundle.areaAnalysis()
	go aBundle.AbbrAnalysis()
	go aBundle.asideAnalysis()
	go aBundle.buttonAnalysis()
	go aBundle.bodyAnalysis()
	go aBundle.canvasAnalysis()
	//go aBundle.cssAnalysis()
	go aBundle.divAnalysis()
	go aBundle.dirAnalysis()
	go aBundle.embedAnalysis()
	go aBundle.footerAnalysis()
	go aBundle.formAnalysis()
	go aBundle.h1Analysis()
	go aBundle.h2Analysis()
	go aBundle.h3Analysis()
	go aBundle.h4Analysis()
	go aBundle.h5Analysis()
	go aBundle.h6Analysis()
	go aBundle.headerAnalysis()
	go aBundle.headAnalysis()
	go aBundle.inputAnalysis()
	go aBundle.imagesAnalysis()
	go aBundle.iframeAnalysis()
	go aBundle.linkAnalysis()
	go aBundle.labelAnalysis()
	go aBundle.mainAnalysis()
	go aBundle.navAnalysis()
	go aBundle.objectAnalysis()
	go aBundle.paraAnalysis()
	go aBundle.preAnalysis()
	go aBundle.selectAnalysis()
	go aBundle.SVGAnalysis()
	go aBundle.spanAnalysis()
	go aBundle.textareaAnalysis()
	go aBundle.trackAnalysis()
	go aBundle.titleAnalysis()
	go aBundle.tableAnalysis()
	go aBundle.theadAnalysis()
	go aBundle.tbodyAnalysis()
	go aBundle.tfootAnalysis()
	go aBundle.videoAnalysis()

	wg.Wait()
	return aBundle.Response
}

// divAnalysis function initiates all the div rule based analysis
func (aBundle *AnalyzeBundle) divAnalysis() {
	aBundle.Logger.Println("Initiating div tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Divs

	var list []resultsapp.Div
	for _, node := range nodes {
		var tag resultsapp.Div

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
	aBundle.Response.DivResults = &list
	MU.Unlock()
}

// buttonAnalysis function initiates all the button rule based analysis
func (aBundle *AnalyzeBundle) buttonAnalysis() {
	aBundle.Logger.Println("Initiating button tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Buttons

	var list []resultsapp.Button
	for _, node := range nodes {
		var tag resultsapp.Button

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
	aBundle.Response.ButtonResults = &list
	MU.Unlock()
}

// inputAnalysis function initiates all the input rule based analysis
func (aBundle *AnalyzeBundle) inputAnalysis() {
	aBundle.Logger.Println("Initiating input tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Inputs

	var list []resultsapp.Input
	for _, node := range nodes {
		var tag resultsapp.Input

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
	aBundle.Response.InputResults = &list
	MU.Unlock()
}

// imagesAnalysis function initiates all the images rule based analysisx`
func (aBundle *AnalyzeBundle) imagesAnalysis() {
	aBundle.Logger.Println("Initiating images tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Imgs

	var list []resultsapp.Img
	for _, node := range nodes {
		var tag resultsapp.Img

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
	aBundle.Response.ImageResults = &list
	MU.Unlock()
}

// videoAnalysis function initiates all the videos rule based analysis
func (aBundle *AnalyzeBundle) videoAnalysis() {
	aBundle.Logger.Println("Initiating videos tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Videos

	var list []resultsapp.Video
	for _, node := range nodes {
		var tag resultsapp.Video

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
	aBundle.Response.VideoResults = &list
	MU.Unlock()
}

// audioAnalysis function initiates all the audios rule based analysis
func (aBundle *AnalyzeBundle) audioAnalysis() {
	aBundle.Logger.Println("Initiating audios tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Audios

	var list []resultsapp.Audio
	for _, node := range nodes {
		var tag resultsapp.Audio

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
	aBundle.Response.AudioResults = &list
	MU.Unlock()
}

// textareaAnalysis function initiates all the textarea rule based analysis
func (aBundle *AnalyzeBundle) textareaAnalysis() {
	aBundle.Logger.Println("Initiating textarea tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.TextAreas

	var list []resultsapp.Textarea
	for _, node := range nodes {
		var tag resultsapp.Textarea

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
	aBundle.Response.TextareaResults = &list
	MU.Unlock()
}

// selectAnalysis function initiates all the select rule based analysis
func (aBundle *AnalyzeBundle) selectAnalysis() {
	aBundle.Logger.Println("Initiating select tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Selects

	var list []resultsapp.Select
	for _, node := range nodes {
		var tag resultsapp.Select

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
	aBundle.Response.SelectResults = &list
	MU.Unlock()
}

// iframeAnalysis function initiates all the iframe rule based analysis
func (aBundle *AnalyzeBundle) iframeAnalysis() {
	aBundle.Logger.Println("Initiating iframe tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Iframes

	var list []resultsapp.Iframe
	for _, node := range nodes {
		var tag resultsapp.Iframe

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
	aBundle.Response.IframeResults = &list
	MU.Unlock()
}

// linkAnalysis function initiates all the link rule based analysis
func (aBundle *AnalyzeBundle) linkAnalysis() {
	aBundle.Logger.Println("Initiating link tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Links

	var list []resultsapp.Link
	for _, node := range nodes {
		var tag resultsapp.Link

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
	aBundle.Response.LinkResults = &list
	MU.Unlock()
}

// anchorAnalysis function initiates all the anchor rule based analysis
func (aBundle *AnalyzeBundle) anchorAnalysis() {
	aBundle.Logger.Println("Initiating anchor tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Anchors

	var list []resultsapp.Anchor
	for _, node := range nodes {
		var tag resultsapp.Anchor

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
	aBundle.Response.AnchorResults = &list
	MU.Unlock()
}

// areaAnalysis function initiates all the area rule based analysis
func (aBundle *AnalyzeBundle) areaAnalysis() {
	aBundle.Logger.Println("Initiating area tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Areas

	var list []resultsapp.Area
	for _, node := range nodes {
		var tag resultsapp.Area

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
	aBundle.Response.AreaResults = &list
	MU.Unlock()
}

// objectAnalysis function initiates all the object rule based analysis
func (aBundle *AnalyzeBundle) objectAnalysis() {
	aBundle.Logger.Println("Initiating object tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Objects

	var list []resultsapp.Object
	for _, node := range nodes {
		var tag resultsapp.Object

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
	aBundle.Response.ObjectResults = &list
	MU.Unlock()
}

// embedAnalysis function initiates all the embed rule based analysis
func (aBundle *AnalyzeBundle) embedAnalysis() {
	aBundle.Logger.Println("Initiating embed tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Embeds

	var list []resultsapp.Embed
	for _, node := range nodes {
		var tag resultsapp.Embed

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
	aBundle.Response.EmbedResults = &list
	MU.Unlock()
}

// trackAnalysis function initiates all the track rule based analysis
func (aBundle *AnalyzeBundle) trackAnalysis() {
	aBundle.Logger.Println("Initiating track tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Tracks

	var list []resultsapp.Track
	for _, node := range nodes {
		var tag resultsapp.Track

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
	aBundle.Response.TrackResults = &list
	MU.Unlock()
}

// h1Analysis function initiates all the h1 rule based analysis
func (aBundle *AnalyzeBundle) h1Analysis() {
	aBundle.Logger.Println("Initiating h1 tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.H1s

	var list []resultsapp.H1
	for _, node := range nodes {
		var tag resultsapp.H1

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
	aBundle.Response.H1Results = &list
	MU.Unlock()
}

// h2Analysis function initiates all the h2 rule based analysis
func (aBundle *AnalyzeBundle) h2Analysis() {
	aBundle.Logger.Println("Initiating h2 tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.H2s

	var list []resultsapp.H2
	for _, node := range nodes {
		var tag resultsapp.H2

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
	aBundle.Response.H2Results = &list
	MU.Unlock()
}

// h3Analysis function initiates all the h3 rule based analysis
func (aBundle *AnalyzeBundle) h3Analysis() {
	aBundle.Logger.Println("Initiating h3 tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.H3s

	var list []resultsapp.H3
	for _, node := range nodes {
		var tag resultsapp.H3

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
	aBundle.Response.H3Results = &list
	MU.Unlock()
}

// h4Analysis function initiates all the h4 rule based analysis
func (aBundle *AnalyzeBundle) h4Analysis() {
	aBundle.Logger.Println("Initiating h4 tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.H4s

	var list []resultsapp.H4
	for _, node := range nodes {
		var tag resultsapp.H4

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
	aBundle.Response.H4Results = &list
	MU.Unlock()
}

// h5Analysis function initiates all the h5 rule based analysis
func (aBundle *AnalyzeBundle) h5Analysis() {
	aBundle.Logger.Println("Initiating h5 tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.H5s

	var list []resultsapp.H5
	for _, node := range nodes {
		var tag resultsapp.H5

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
	aBundle.Response.H5Results = &list
	MU.Unlock()
}

// h6Analysis function initiates all the h6 rule based analysis
func (aBundle *AnalyzeBundle) h6Analysis() {
	aBundle.Logger.Println("Initiating h6 tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.H6s

	var list []resultsapp.H6
	for _, node := range nodes {
		var tag resultsapp.H6

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
	aBundle.Response.H6Results = &list
	MU.Unlock()
}

// paraAnalysis function initiates all the para rule based analysis
func (aBundle *AnalyzeBundle) paraAnalysis() {
	aBundle.Logger.Println("Initiating para tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Paras

	var list []resultsapp.Para
	for _, node := range nodes {
		var tag resultsapp.Para

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
	aBundle.Response.ParaResults = &list
	MU.Unlock()
}

// preAnalysis function initiates all the pre rule based analysis
func (aBundle *AnalyzeBundle) preAnalysis() {
	aBundle.Logger.Println("Initiating pre tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Pres

	var list []resultsapp.Pre
	for _, node := range nodes {
		var tag resultsapp.Pre

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
	aBundle.Response.PreResults = &list
	MU.Unlock()
}

// AbbrAnalysis function initiates all the abbr rule based analysis
func (aBundle *AnalyzeBundle) AbbrAnalysis() {
	aBundle.Logger.Println("Initiating abbr tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Abbrs

	var list []resultsapp.Abbr
	for _, node := range nodes {
		var tag resultsapp.Abbr

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
	aBundle.Response.AbbrResults = &list
	MU.Unlock()
}

// SVGAnalysis function initiates all the svg rule based analysis
func (aBundle *AnalyzeBundle) SVGAnalysis() {
	aBundle.Logger.Println("Initiating svg tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Svgs

	var list []resultsapp.SVG
	for _, node := range nodes {
		var tag resultsapp.SVG

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
	aBundle.Response.SvgResults = &list
	MU.Unlock()
}

// canvasAnalysis function initiates all the canvas rule based analysis
func (aBundle *AnalyzeBundle) canvasAnalysis() {
	aBundle.Logger.Println("Initiating canvas tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Canvases

	var list []resultsapp.Canvas
	for _, node := range nodes {
		var tag resultsapp.Canvas

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
	aBundle.Response.CanvasResults = &list
	MU.Unlock()
}

// spanAnalysis function initiates all the spans rule based analysis
func (aBundle *AnalyzeBundle) spanAnalysis() {
	aBundle.Logger.Println("Initiating span tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Spans

	var list []resultsapp.Span
	for _, node := range nodes {
		var tag resultsapp.Span

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
	aBundle.Response.SpanResults = &list
	MU.Unlock()
}

// navAnalysis function initiates all the nav rule based analysis
func (aBundle *AnalyzeBundle) navAnalysis() {
	aBundle.Logger.Println("Initiating nav tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Navs

	var list []resultsapp.Nav
	for _, node := range nodes {
		var tag resultsapp.Nav

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
	aBundle.Response.NavResults = &list
	MU.Unlock()
}

// asideAnalysis function initiates all the aside rule based analysis
func (aBundle *AnalyzeBundle) asideAnalysis() {
	aBundle.Logger.Println("Initiating aside tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Navs

	var list []resultsapp.Aside
	for _, node := range nodes {
		var tag resultsapp.Aside

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
	aBundle.Response.AsideResults = &list
	MU.Unlock()
}

// mainAnalysis function initiates all the main rule based analysis
func (aBundle *AnalyzeBundle) mainAnalysis() {
	aBundle.Logger.Println("Initiating main tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Mains

	var list []resultsapp.Main
	for _, node := range nodes {
		var tag resultsapp.Main

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
	aBundle.Response.MainResults = &list
	MU.Unlock()
}

// headerAnalysis function initiates all the header rule based analysis
func (aBundle *AnalyzeBundle) headerAnalysis() {
	aBundle.Logger.Println("Initiating header tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Headers

	var list []resultsapp.Header
	for _, node := range nodes {
		var tag resultsapp.Header

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
	aBundle.Response.HeaderResults = &list
	MU.Unlock()
}

// footerAnalysis function initiates all the footer rule based analysis
func (aBundle *AnalyzeBundle) footerAnalysis() {
	aBundle.Logger.Println("Initiating footer tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Footers

	var list []resultsapp.Footer
	for _, node := range nodes {
		var tag resultsapp.Footer

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
	aBundle.Response.FooterResults = &list
	MU.Unlock()
}

// headAnalysis function initiates all the head rule based analysis
func (aBundle *AnalyzeBundle) headAnalysis() {
	aBundle.Logger.Println("Initiating head tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Heads

	var list []resultsapp.Head
	for _, node := range nodes {
		var tag resultsapp.Head

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
	aBundle.Response.HeadResults = &list
	MU.Unlock()
}

// labelAnalysis function initiates all the label rule based analysis
func (aBundle *AnalyzeBundle) labelAnalysis() {
	aBundle.Logger.Println("Initiating label tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Labels

	var list []resultsapp.Label
	for _, node := range nodes {
		var tag resultsapp.Label

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
	aBundle.Response.LabelResults = &list
	MU.Unlock()
}

// formAnalysis function initiates all the form rule based analysis
func (aBundle *AnalyzeBundle) formAnalysis() {
	aBundle.Logger.Println("Initiating form tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Forms

	var list []resultsapp.Form
	for _, node := range nodes {
		var tag resultsapp.Form

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
	aBundle.Response.FormResults = &list
	MU.Unlock()
}

// dirAnalysis function initiates all the dir rule based analysis
func (aBundle *AnalyzeBundle) dirAnalysis() {
	aBundle.Logger.Println("Initiating dir tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Dirs

	var list []resultsapp.Dir
	for _, node := range nodes {
		var tag resultsapp.Dir

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
	aBundle.Response.DirResults = &list
	MU.Unlock()
}

// bodyAnalysis function initiates all the body rule based analysis
func (aBundle *AnalyzeBundle) bodyAnalysis() {
	aBundle.Logger.Println("Initiating body tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Bodys

	var list []resultsapp.Body
	for _, node := range nodes {
		var tag resultsapp.Body

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
	aBundle.Response.BodyResults = &list
	MU.Unlock()
}

// titleAnalysis function initiates all the title rule based analysis
func (aBundle *AnalyzeBundle) titleAnalysis() {
	aBundle.Logger.Println("Initiating title tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Titles

	var list []resultsapp.Title
	for _, node := range nodes {
		var tag resultsapp.Title

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
	aBundle.Response.TitleResults = &list
	MU.Unlock()
}

// tableAnalysis function initiates all the table rule based analysis
func (aBundle *AnalyzeBundle) tableAnalysis() {
	aBundle.Logger.Println("Initiating table tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Tables

	var list []resultsapp.Table
	for _, node := range nodes {
		var tag resultsapp.Table

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
	aBundle.Response.TableResults = &list
	MU.Unlock()
}

// theadAnalysis function initiates all the thead rule based analysis
func (aBundle *AnalyzeBundle) theadAnalysis() {
	aBundle.Logger.Println("Initiating thead tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.THeads

	var list []resultsapp.Thead
	for _, node := range nodes {
		var tag resultsapp.Thead

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
	aBundle.Response.THeadResults = &list
	MU.Unlock()
}

// tbodyAnalysis function initiates all the tbody rule based analysis
func (aBundle *AnalyzeBundle) tbodyAnalysis() {
	aBundle.Logger.Println("Initiating tbody tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.TBodys

	var list []resultsapp.Tbody
	for _, node := range nodes {
		var tag resultsapp.Tbody

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
	aBundle.Response.TbodyResults = &list
	MU.Unlock()
}

// tfootAnalysis function initiates all the tfoot rule based analysis
func (aBundle *AnalyzeBundle) tfootAnalysis() {
	aBundle.Logger.Println("Initiating tfoot tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.TFoots

	var list []resultsapp.Tfoot
	for _, node := range nodes {
		var tag resultsapp.Tfoot

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
	aBundle.Response.TfootResults = &list
	MU.Unlock()
}

// cssAnalysis function initiates all the CSS rule based analysis
func (aBundle *AnalyzeBundle) cssAnalysis() {
	aBundle.Logger.Println("Initiating CSS Analysis......")
	defer wg.Done()
	var list []resultsapp.CSS
	nodes := aBundle.CollectedTags.Divs
	cssLinks := aBundle.CollectedTags.CssLinks
	for _, css := range cssLinks {
		aBundle.Logger.Printf("CSS : %v ", css)
		var tag resultsapp.CSS

		//add css data
		tag.CSS = css

		//implement rule
		aBundle.rules.Css = css
		for _, node := range nodes {
			aBundle.rules.Logger = aBundle.Logger
			if status, results := aBundle.rules.Execute(node); status == true {
				tag.Result = results
				if len(list) < 50 {
					list = append(list, tag)
				}
			}
		}

	}
	aBundle.Response.CSSResults = &list
}
