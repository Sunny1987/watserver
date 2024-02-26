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
	MyTags        resultsapp.Tags
	Response      resultsapp.Response
	rules         *rule.RuleResults
}

// NewAnalyzeBundle is the constructor for AnalyzeBundle
func NewAnalyzeBundle(req *resultsapp.MyRequest, logger *log.Logger, base string, doc *html.Node, collectedTags resultsapp.TagsFamily) *AnalyzeBundle {
	return &AnalyzeBundle{Req: req, Logger: logger, Base: base, Doc: doc, CollectedTags: collectedTags}
}

func (aBundle *AnalyzeBundle) Analyze() resultsapp.Response {
	aBundle.Response.Request = aBundle.Req

	wg.Add(27)
	go aBundle.anchorAnalysis()
	go aBundle.audioAnalysis()
	go aBundle.areaAnalysis()
	go aBundle.AbbrAnalysis()
	go aBundle.buttonAnalysis()
	go aBundle.canvasAnalysis()
	//go aBundle.cssAnalysis()
	go aBundle.divAnalysis()
	go aBundle.embedAnalysis()
	go aBundle.h1Analysis()
	go aBundle.h2Analysis()
	go aBundle.h3Analysis()
	go aBundle.h4Analysis()
	go aBundle.h5Analysis()
	go aBundle.h6Analysis()
	go aBundle.inputAnalysis()
	go aBundle.imagesAnalysis()
	go aBundle.iframeAnalysis()
	go aBundle.linkAnalysis()
	go aBundle.objectAnalysis()
	go aBundle.paraAnalysis()
	go aBundle.preAnalysis()
	go aBundle.selectAnalysis()
	go aBundle.SVGAnalysis()
	go aBundle.spanAnalysis()
	go aBundle.textareaAnalysis()
	go aBundle.trackAnalysis()
	go aBundle.videoAnalysis()

	wg.Wait()
	return aBundle.Response
}

// divAnalysis function initiates all the div rule based analysis
func (aBundle *AnalyzeBundle) divAnalysis() {
	aBundle.Logger.Println("Initiating div tag Analysis......")
	defer wg.Done()
	nodes := aBundle.CollectedTags.Divs

	var list []resultsapp.Divtag
	for _, node := range nodes {
		var tag resultsapp.Divtag

		//build the node
		tag.Div = helper.NodeText(node)

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

	var list []resultsapp.Buttontag
	for _, node := range nodes {
		var tag resultsapp.Buttontag

		//build the node
		tag.Button = helper.NodeText(node)

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

	var list []resultsapp.Inputtag
	for _, node := range nodes {
		var tag resultsapp.Inputtag

		//build the node
		tag.Input = helper.NodeText(node)

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

	var list []resultsapp.Imgtag
	for _, node := range nodes {
		var tag resultsapp.Imgtag

		//build the node
		tag.Img = helper.NodeText(node)

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

	var list []resultsapp.Videotag
	for _, node := range nodes {
		var tag resultsapp.Videotag

		//build the node
		tag.Video = helper.NodeText(node)

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

	var list []resultsapp.Audiotag
	for _, node := range nodes {
		var tag resultsapp.Audiotag

		//build the node
		tag.Audio = helper.NodeText(node)

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

	var list []resultsapp.Textareatag
	for _, node := range nodes {
		var tag resultsapp.Textareatag

		//build the node
		tag.Textarea = helper.NodeText(node)

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

	var list []resultsapp.Selecttag
	for _, node := range nodes {
		var tag resultsapp.Selecttag

		//build the node
		tag.Select = helper.NodeText(node)

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

	var list []resultsapp.Iframetag
	for _, node := range nodes {
		var tag resultsapp.Iframetag

		//build the node
		tag.Iframe = helper.NodeText(node)

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

	var list []resultsapp.Linktag
	for _, node := range nodes {
		var tag resultsapp.Linktag

		//build the node
		tag.Link = helper.NodeText(node)

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

	var list []resultsapp.Anchortag
	for _, node := range nodes {
		var tag resultsapp.Anchortag

		//build the node
		tag.Anchor = helper.NodeText(node)

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

	var list []resultsapp.Areatag
	for _, node := range nodes {
		var tag resultsapp.Areatag

		//build the node
		tag.Area = helper.NodeText(node)

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

	var list []resultsapp.Objecttag
	for _, node := range nodes {
		var tag resultsapp.Objecttag

		//build the node
		tag.Object = helper.NodeText(node)

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

	var list []resultsapp.Embedtag
	for _, node := range nodes {
		var tag resultsapp.Embedtag

		//build the node
		tag.Embed = helper.NodeText(node)

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

	var list []resultsapp.Tracktag
	for _, node := range nodes {
		var tag resultsapp.Tracktag

		//build the node
		tag.Track = helper.NodeText(node)

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

	var list []resultsapp.H1tag
	for _, node := range nodes {
		var tag resultsapp.H1tag

		//build the node
		tag.H1 = helper.NodeText(node)

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

	var list []resultsapp.H2tag
	for _, node := range nodes {
		var tag resultsapp.H2tag

		//build the node
		tag.H2 = helper.NodeText(node)

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

	var list []resultsapp.H3tag
	for _, node := range nodes {
		var tag resultsapp.H3tag

		//build the node
		tag.H3 = helper.NodeText(node)

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

	var list []resultsapp.H4tag
	for _, node := range nodes {
		var tag resultsapp.H4tag

		//build the node
		tag.H4 = helper.NodeText(node)

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

	var list []resultsapp.H5tag
	for _, node := range nodes {
		var tag resultsapp.H5tag

		//build the node
		tag.H5 = helper.NodeText(node)

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

	var list []resultsapp.H6tag
	for _, node := range nodes {
		var tag resultsapp.H6tag

		//build the node
		tag.H6 = helper.NodeText(node)

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

	var list []resultsapp.Paratag
	for _, node := range nodes {
		var tag resultsapp.Paratag

		//build the node
		tag.Para = helper.NodeText(node)

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

	var list []resultsapp.Pretag
	for _, node := range nodes {
		var tag resultsapp.Pretag

		//build the node
		tag.Pre = helper.NodeText(node)

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

	var list []resultsapp.Abbrtag
	for _, node := range nodes {
		var tag resultsapp.Abbrtag

		//build the node
		tag.Abbr = helper.NodeText(node)

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

	var list []resultsapp.SVGtag
	for _, node := range nodes {
		var tag resultsapp.SVGtag

		//build the node
		tag.SVG = helper.NodeText(node)

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

	var list []resultsapp.Canvastag
	for _, node := range nodes {
		var tag resultsapp.Canvastag

		//build the node
		tag.Canvas = helper.NodeText(node)

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

	var list []resultsapp.Spantag
	for _, node := range nodes {
		var tag resultsapp.Spantag

		//build the node
		tag.Span = helper.NodeText(node)

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
