package resultsapp

import "golang.org/x/net/html"

const (
	Anchors   = "Anchors"
	Divs      = "Divs"
	Paras     = "Paras"
	Spans     = "Spans"
	H1s       = "H1s"
	H2s       = "H2s"
	H3s       = "H3s"
	H4s       = "H4s"
	H5s       = "H5s"
	H6s       = "H6s"
	Imgs      = "Imgs"
	Buttons   = "Buttons"
	Videos    = "Videos"
	Audios    = "Audios"
	Selects   = "Selects"
	TextAreas = "TextAreas"
	Iframes   = "Iframes"
	Areas     = "Areas"
	Objects   = "Objects"
	Embeds    = "Embeds"
	Tracks    = "Tracks"
	Applets   = "Applets"
	Pres      = "Pres"
	Links     = "Links"
	Inputs    = "Inputs"
	Abbrs     = "Abbrs"
	Svgs      = "Svgs"
	Canvases  = "Canvases"
	Asides    = "Asides"
	Mains     = "Mains"
	Navs      = "Navs"
	Headers   = "Headers"
	Footers   = "Footers"
	Heads     = "Heads"
	Labels    = "Labels"
	Forms     = "Forms"
	Dirs      = "Dirs"
	Bodys     = "Bodys"
	Titles    = "Titles"
	Tables    = "Tables"
	THeads    = "THeads"
	TBodys    = "TBodys"
	TFoots    = "TFoots"
)

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
	Asides    []*html.Node
	Mains     []*html.Node
	Navs      []*html.Node
	Headers   []*html.Node
	Footers   []*html.Node
	Heads     []*html.Node
	Labels    []*html.Node
	Forms     []*html.Node
	Dirs      []*html.Node
	Bodys     []*html.Node
	Titles    []*html.Node
	Tables    []*html.Node
	THeads    []*html.Node
	TBodys    []*html.Node
	TFoots    []*html.Node

	CssLinks []string
}

func NewTagsFamily(nodeMap map[string][]*html.Node, cssList []string) TagsFamily {
	return TagsFamily{
		Anchors:   nodeMap[Anchors],
		Divs:      nodeMap[Divs],
		Paras:     nodeMap[Paras],
		Spans:     nodeMap[Spans],
		H1s:       nodeMap[H1s],
		H2s:       nodeMap[H2s],
		H3s:       nodeMap[H3s],
		H4s:       nodeMap[H4s],
		H5s:       nodeMap[H5s],
		H6s:       nodeMap[H6s],
		Imgs:      nodeMap[Imgs],
		Buttons:   nodeMap[Buttons],
		Videos:    nodeMap[Videos],
		Audios:    nodeMap[Audios],
		Selects:   nodeMap[Selects],
		TextAreas: nodeMap[TextAreas],
		Iframes:   nodeMap[Iframes],
		Areas:     nodeMap[Areas],
		Objects:   nodeMap[Objects],
		Embeds:    nodeMap[Embeds],
		Tracks:    nodeMap[Tracks],
		Applets:   nodeMap[Applets],
		Pres:      nodeMap[Pres],
		Links:     nodeMap[Links],
		Inputs:    nodeMap[Inputs],
		Abbrs:     nodeMap[Abbrs],
		Svgs:      nodeMap[Svgs],
		Canvases:  nodeMap[Canvases],
		Navs:      nodeMap[Navs],
		Asides:    nodeMap[Asides],
		Mains:     nodeMap[Mains],
		Headers:   nodeMap[Headers],
		Footers:   nodeMap[Footers],
		Heads:     nodeMap[Heads],
		Labels:    nodeMap[Labels],
		Forms:     nodeMap[Forms],
		Dirs:      nodeMap[Dirs],
		Bodys:     nodeMap[Bodys],
		Titles:    nodeMap[Titles],
		Tables:    nodeMap[Tables],
		THeads:    nodeMap[THeads],
		TBodys:    nodeMap[TBodys],
		TFoots:    nodeMap[TFoots],
		CssLinks:  cssList,
	}
}
