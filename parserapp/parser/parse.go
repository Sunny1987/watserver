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

// ParseBundle object is the DataBundle for future parsing
type ParseBundle struct {
	Req        *resultsapp.MyRequest
	Logger     *log.Logger
	Person     *string //the property is currently optional
	Base       string
	Doc        *html.Node
	TagsFamily resultsapp.TagsFamily
}

// NewParseBundle is the constructor for ParseBundle
func NewParseBundle(Req *resultsapp.MyRequest, log *log.Logger, Base string) DataBundle {
	return &ParseBundle{
		Req:    Req,
		Logger: log,
		Base:   Base,
	}
}

// SetDocValue is the Setter for Doc property
func (pBundle *ParseBundle) SetDocValue(doc *html.Node) {
	pBundle.Doc = doc
}

// Parse method will parse the incoming http response body/ html file body
func (pBundle *ParseBundle) Parse(responseBody io.Reader) resultsapp.Response {
	doc, err := html.Parse(responseBody)
	if err != nil {
		pBundle.Logger.Printf("Error parsing the html : %v", err)
	}

	//Set Doc value for ParseBundle
	pBundle.SetDocValue(doc)

	//Collect all nodes from html doc
	nodeMap := pBundle.collectNode()

	//Update Tags in ParseBundle
	pBundle.TagsFamily = resultsapp.NewTagsFamily(nodeMap, cssList)
	pBundle.Logger.Println("Initiating WCAG 2.1 analysis.....")

	//Call AnalyzeBundle constructor
	analyzerBundle := analyzerapp.NewAnalyzeBundle(pBundle.Req, pBundle.Logger, pBundle.Base, pBundle.Doc, pBundle.TagsFamily)
	return analyzerBundle.Analyze()
}

// collectNode will extract the html nodes
func (pBundle *ParseBundle) collectNode() map[string][]*html.Node {

	//initialize a map
	nodeMap := make(map[string][]*html.Node)

	wg.Add(24)
	go pBundle.getNode(FilterAnchorNodes, "Anchors", nodeMap)
	go pBundle.getNode(filterDivNodes, "Divs", nodeMap)
	go pBundle.getNode(filterParaNodes, "Paras", nodeMap)
	go pBundle.getNode(filterSpanNodes, "Spans", nodeMap)
	go pBundle.getNode(filterH1Nodes, "H1s", nodeMap)
	go pBundle.getNode(filterH2Nodes, "H2s", nodeMap)
	go pBundle.getNode(filterH3Nodes, "H3s", nodeMap)
	go pBundle.getNode(filterH4Nodes, "H4s", nodeMap)
	go pBundle.getNode(filterH5Nodes, "H5s", nodeMap)
	go pBundle.getNode(filterH6Nodes, "H6s", nodeMap)
	go pBundle.getNode(filterImageNodes, "Imgs", nodeMap)
	go pBundle.getNode(filterInputNodes, "Inputs", nodeMap)
	go pBundle.getNode(filterButtonNodes, "Buttons", nodeMap)
	go pBundle.getNode(filterVideoNodes, "Videos", nodeMap)
	go pBundle.getNode(filterAudioNodes, "Audios", nodeMap)
	go pBundle.getNode(filterSelectNodes, "Selects", nodeMap)
	go pBundle.getNode(filterTextAreaNodes, "Textareas", nodeMap)
	go pBundle.getNode(filterIframeNodes, "Iframes", nodeMap)
	go pBundle.getNode(filterAreaNodes, "Areas", nodeMap)
	go pBundle.getNode(filterObjectNodes, "Objects", nodeMap)
	go pBundle.getNode(filterEmbedNodes, "Embeds", nodeMap)
	go pBundle.getNode(filterTrackNodes, "Tracks", nodeMap)
	go pBundle.getNode(filterAppletNodes, "Applets", nodeMap)
	go pBundle.getNode(filterPreNodes, "Pres", nodeMap)
	wg.Add(1)
	go func(base string) {
		pBundle.Logger.Println("Collecting all links...")
		defer wg.Done()
		if base != "" {
			linkNodes := FilterLinkNodes(pBundle.Doc)
			pBundle.Logger.Printf("base link: %v", base)

			//collect CSS links
			cssLinks := HrefLinks(filterCSSLinks(linkNodes), base, pBundle.Logger)
			for _, link := range cssLinks {
				pBundle.Logger.Println(link)
				readCSSLinks(link, pBundle.Logger)
			}
			if len(linkNodes) > 0 {
				mu.Lock()
				nodeMap["Links"] = linkNodes
				mu.Unlock()
			}

		}
	}(pBundle.Base)
	wg.Wait()
	return nodeMap
}

// getNode function is a common function to retrieve the node
func (pBundle *ParseBundle) getNode(fn func(node *html.Node) []*html.Node, nodeName string, nodeMap map[string][]*html.Node) {
	defer wg.Done()
	pBundle.Logger.Printf("Collecting all %v...", nodeName)
	nodes := fn(pBundle.Doc)
	if len(nodes) > 0 {
		mu.Lock()
		nodeMap[nodeName] = nodes
		mu.Unlock()
	}
}
