package resultsapp

import "golang.org/x/net/html"

type TagsFamily struct {
	Anchors   []*html.Node
	Divs      []*html.Node
	Paras     []*html.Node
	Spans     []*html.Node
	H1s       []*html.Node
	H2s       []*html.Node
	H3s       []*html.Node
	H4s       []*html.Node
	H5s       []*html.Node
	H6s       []*html.Node
	Imgs      []*html.Node
	Buttons   []*html.Node
	Videos    []*html.Node
	Audios    []*html.Node
	Selects   []*html.Node
	TextAreas []*html.Node
	Iframes   []*html.Node
	Areas     []*html.Node
	Objects   []*html.Node
	Embeds    []*html.Node
	Tracks    []*html.Node
	Applets   []*html.Node
	Pres      []*html.Node
	Links     []*html.Node
	Inputs    []*html.Node
	Abbrs     []*html.Node
	Svgs      []*html.Node
	Canvases  []*html.Node
	CssLinks  []string
}

func NewTagsFamily(nodeMap map[string][]*html.Node, cssList []string) TagsFamily {
	return TagsFamily{
		Anchors:   nodeMap["Anchors"],
		Divs:      nodeMap["Divs"],
		Paras:     nodeMap["Paras"],
		Spans:     nodeMap["Spans"],
		H1s:       nodeMap["H1s"],
		H2s:       nodeMap["H2s"],
		H3s:       nodeMap["H3s"],
		H4s:       nodeMap["H4s"],
		H5s:       nodeMap["H5s"],
		H6s:       nodeMap["H6s"],
		Imgs:      nodeMap["Imgs"],
		Buttons:   nodeMap["Buttons"],
		Videos:    nodeMap["Videos"],
		Audios:    nodeMap["Audios"],
		Selects:   nodeMap["Selects"],
		TextAreas: nodeMap["TextAreas"],
		Iframes:   nodeMap["Iframes"],
		Areas:     nodeMap["Areas"],
		Objects:   nodeMap["Objects"],
		Embeds:    nodeMap["Embeds"],
		Tracks:    nodeMap["Tracks"],
		Applets:   nodeMap["Applets"],
		Pres:      nodeMap["Pres"],
		Links:     nodeMap["Links"],
		Inputs:    nodeMap["Inputs"],
		Abbrs:     nodeMap["Abbrs"],
		Svgs:      nodeMap["Svgs"],
		Canvases:  nodeMap["Canvases"],
		CssLinks:  cssList,
	}
}
