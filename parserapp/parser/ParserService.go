package parser

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"sync"
	"webserver/analyzerapp"
	"webserver/resultsapp"
)

var (
	wg      sync.WaitGroup
	mu      sync.RWMutex
	cssList []string
)

// ParserService has parser contracts
type ParserService interface {
	Parse(responseBody io.Reader) resultsapp.FinalResponse
}

// ParseBundle object is the DataBundle for future parsing
type ParseBundle struct {
	req        *resultsapp.MyRequest
	logger     *log.Logger
	person     *string //the property is currently optional
	base       string
	doc        *html.Node
	tagsFamily resultsapp.TagsFamily
	analyzer   analyzerapp.AnalyzerService
}

// NewParseBundle is the constructor for ParseBundle
func NewParseBundle(Req *resultsapp.MyRequest, log *log.Logger, Base string, analyzer analyzerapp.AnalyzerService) ParserService {
	return &ParseBundle{
		req:      Req,
		logger:   log,
		base:     Base,
		analyzer: analyzer,
	}
}

// SetDocValue is the Setter for doc property
func (pBundle *ParseBundle) setDocValue(doc *html.Node) {
	pBundle.doc = doc
}

// Parse method will parse the incoming http response body/ html file body
func (pBundle *ParseBundle) Parse(responseBody io.Reader) resultsapp.FinalResponse {
	doc, err := html.Parse(responseBody)
	if err != nil {
		pBundle.logger.Printf("Error parsing the html : %v", err)
	}

	//Set doc value for ParseBundle
	pBundle.setDocValue(doc)

	//Collect all nodes from html doc
	nodeMap := pBundle.collectNode()

	//Update Tags in ParseBundle
	pBundle.tagsFamily = resultsapp.NewTagsFamily(nodeMap, cssList)

	pBundle.logger.Println(".....Initiating WCAG 2.1 analysis.....")

	//Call AnalyzeBundle constructor
	//analyzer := analyzerapp.NewAnalyzeBundle(pBundle.req, pBundle.logger, pBundle.base, pBundle.doc, pBundle.tagsFamily)
	pBundle.analyzer.AddProperties(pBundle.base, pBundle.doc, pBundle.tagsFamily)
	return pBundle.analyzer.Analyze()
}

// collectNode will extract the html nodes
func (pBundle *ParseBundle) collectNode() map[string][]*html.Node {

	//initialize a map
	nodeMap := make(map[string][]*html.Node)

	wg.Add(42)
	go pBundle.getNode(FilterAnchorNodes, resultsapp.Anchors, nodeMap)
	go pBundle.getNode(filterDivNodes, resultsapp.Divs, nodeMap)
	go pBundle.getNode(filterParaNodes, resultsapp.Paras, nodeMap)
	go pBundle.getNode(filterSpanNodes, resultsapp.Spans, nodeMap)
	go pBundle.getNode(filterH1Nodes, resultsapp.H1s, nodeMap)
	go pBundle.getNode(filterH2Nodes, resultsapp.H2s, nodeMap)
	go pBundle.getNode(filterH3Nodes, resultsapp.H3s, nodeMap)
	go pBundle.getNode(filterH4Nodes, resultsapp.H4s, nodeMap)
	go pBundle.getNode(filterH5Nodes, resultsapp.H5s, nodeMap)
	go pBundle.getNode(filterH6Nodes, resultsapp.H6s, nodeMap)
	go pBundle.getNode(filterImageNodes, resultsapp.Imgs, nodeMap)
	go pBundle.getNode(filterInputNodes, resultsapp.Inputs, nodeMap)
	go pBundle.getNode(filterButtonNodes, resultsapp.Buttons, nodeMap)
	go pBundle.getNode(filterVideoNodes, resultsapp.Videos, nodeMap)
	go pBundle.getNode(filterAudioNodes, resultsapp.Audios, nodeMap)
	go pBundle.getNode(filterSelectNodes, resultsapp.Selects, nodeMap)
	go pBundle.getNode(filterTextAreaNodes, resultsapp.TextAreas, nodeMap)
	go pBundle.getNode(filterIframeNodes, resultsapp.Iframes, nodeMap)
	go pBundle.getNode(filterAreaNodes, resultsapp.Areas, nodeMap)
	go pBundle.getNode(filterObjectNodes, resultsapp.Objects, nodeMap)
	go pBundle.getNode(filterEmbedNodes, resultsapp.Embeds, nodeMap)
	go pBundle.getNode(filterTrackNodes, resultsapp.Tracks, nodeMap)
	go pBundle.getNode(filterAppletNodes, resultsapp.Applets, nodeMap)
	go pBundle.getNode(filterPreNodes, resultsapp.Pres, nodeMap)
	go pBundle.getNode(filterAbbrNodes, resultsapp.Abbrs, nodeMap)
	go pBundle.getNode(filterSvgNodes, resultsapp.Svgs, nodeMap)
	go pBundle.getNode(filterCanvasNodes, resultsapp.Canvases, nodeMap)
	go pBundle.getNode(filterNavNodes, resultsapp.Navs, nodeMap)
	go pBundle.getNode(filterAsideNodes, resultsapp.Asides, nodeMap)
	go pBundle.getNode(filterMainNodes, resultsapp.Mains, nodeMap)
	go pBundle.getNode(filterHeaderNodes, resultsapp.Headers, nodeMap)
	go pBundle.getNode(filterFooterNodes, resultsapp.Footers, nodeMap)
	go pBundle.getNode(filterHeadNodes, resultsapp.Heads, nodeMap)
	go pBundle.getNode(filterLabelNodes, resultsapp.Labels, nodeMap)
	go pBundle.getNode(filterFormNodes, resultsapp.Forms, nodeMap)
	go pBundle.getNode(filterDirNodes, resultsapp.Dirs, nodeMap)
	go pBundle.getNode(filterBodyNodes, resultsapp.Bodys, nodeMap)
	go pBundle.getNode(filterTitleNodes, resultsapp.Titles, nodeMap)
	go pBundle.getNode(filterTableNodes, resultsapp.Tables, nodeMap)
	go pBundle.getNode(filterTHeadNodes, resultsapp.THeads, nodeMap)
	go pBundle.getNode(filterTBodyNodes, resultsapp.TBodys, nodeMap)
	go pBundle.getNode(filterTFootNodes, resultsapp.TFoots, nodeMap)

	wg.Add(1)
	go func(base string) {
		defer wg.Done()
		pBundle.logger.Println("Collecting all links...")
		if base != "" {
			linkNodes := FilterLinkNodes(pBundle.doc)
			pBundle.logger.Printf("base link: %v", base)

			//collect CSS links
			cssLinks := HrefLinks(filterCSSLinks(linkNodes), base, pBundle.logger)
			for _, link := range cssLinks {
				pBundle.logger.Println(link)
				readCSSLinks(link, pBundle.logger)
			}
			if len(linkNodes) > 0 {
				mu.Lock()
				nodeMap["Links"] = linkNodes
				mu.Unlock()
			}

		}
	}(pBundle.base)
	wg.Wait()
	return nodeMap
}

// getNode function is a common function to retrieve the node
func (pBundle *ParseBundle) getNode(fn func(node *html.Node) []*html.Node, nodeName string, nodeMap map[string][]*html.Node) {
	defer wg.Done()
	pBundle.logger.Printf("Collecting all %v...", nodeName)
	nodes := fn(pBundle.doc)
	if len(nodes) > 0 {
		mu.Lock()
		nodeMap[nodeName] = nodes
		mu.Unlock()
	}
}
