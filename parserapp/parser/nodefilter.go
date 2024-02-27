package parser

import "golang.org/x/net/html"

// FilterAnchorNodes will give the list of links
func FilterAnchorNodes(n *html.Node) []*html.Node {
	return filterNode(n, isAnchor)
}

// filterDivNodes will give the list of divs
func filterDivNodes(n *html.Node) []*html.Node {
	return filterNode(n, isDiv)
}

// filterParaNodes will give the list of divs
func filterParaNodes(n *html.Node) []*html.Node {
	return filterNode(n, isParagraph)
}

// filterSpanNodes will give the list of spans
func filterSpanNodes(n *html.Node) []*html.Node {
	return filterNode(n, isSpan)
}

// filterH1Nodes will give the list of H1
func filterH1Nodes(n *html.Node) []*html.Node {
	return filterNode(n, isH1)
}

// filterH2Nodes will give the list of H2
func filterH2Nodes(n *html.Node) []*html.Node {
	return filterNode(n, isH2)
}

// filterH3Nodes will give the list of H3
func filterH3Nodes(n *html.Node) []*html.Node {
	return filterNode(n, isH3)
}

// filterH4Nodes will give the list of H4
func filterH4Nodes(n *html.Node) []*html.Node {
	return filterNode(n, isH4)
}

// filterH5Nodes will give the list of H5
func filterH5Nodes(n *html.Node) []*html.Node {
	return filterNode(n, isH5)
}

// filterH6Nodes will give the list of H6
func filterH6Nodes(n *html.Node) []*html.Node {
	return filterNode(n, isH6)
}

// filterImageNodes will give the list of image
func filterImageNodes(n *html.Node) []*html.Node {
	return filterNode(n, isImage)
}

// filterInputNodes will give the list of input
func filterInputNodes(n *html.Node) []*html.Node {
	return filterNode(n, isInput)
}

// filterButtonNodes will give the list of button
func filterButtonNodes(n *html.Node) []*html.Node {
	return filterNode(n, isButton)
}

// filterVideoNodes will give the list of video
func filterVideoNodes(n *html.Node) []*html.Node {
	return filterNode(n, isVideo)
}

// filterAudioNodes will give the list of audio
func filterAudioNodes(n *html.Node) []*html.Node {
	return filterNode(n, isAudio)
}

// filterSelectNodes will give the list of select
func filterSelectNodes(n *html.Node) []*html.Node {
	return filterNode(n, isSelect)
}

// filterTextAreaNodes will give the list of TextArea
func filterTextAreaNodes(n *html.Node) []*html.Node {
	return filterNode(n, isTextArea)
}

// filterIframeNodes will give the list of iframes
func filterIframeNodes(n *html.Node) []*html.Node {
	return filterNode(n, isIFrame)
}

// filterAreaNodes will give the list of areas
func filterAreaNodes(n *html.Node) []*html.Node {
	return filterNode(n, isArea)
}

// filterObjectNodes will give the list of objects
func filterObjectNodes(n *html.Node) []*html.Node {
	return filterNode(n, isObject)
}

// filterEmbedNodes will give the list of embeds
func filterEmbedNodes(n *html.Node) []*html.Node {
	return filterNode(n, isEmbed)
}

// filterTrackNodes will give the list of tracks
func filterTrackNodes(n *html.Node) []*html.Node {
	return filterNode(n, isTrack)
}

// filterAppletNodes will give the list of applet
func filterAppletNodes(n *html.Node) []*html.Node {
	return filterNode(n, isApplet)
}

// filterPreNodes will give the list of pre
func filterPreNodes(n *html.Node) []*html.Node {
	return filterNode(n, isPre)
}

// filterAbbrNodes will give the list of Abbr
func filterAbbrNodes(n *html.Node) []*html.Node {
	return filterNode(n, isAbbr)
}

// filterSvgNodes will give the list of svg
func filterSvgNodes(n *html.Node) []*html.Node {
	return filterNode(n, isSVG)
}

// filterCanvasNodes will give the list of canvas
func filterCanvasNodes(n *html.Node) []*html.Node {
	return filterNode(n, isCanvas)
}

// filterAcronymNodes will give the list of Acronym
func filterAcronymNodes(n *html.Node) []*html.Node {
	return filterNode(n, isAcronym)
}

// filterAddressNodes will give the list of Address
func filterAddressNodes(n *html.Node) []*html.Node {
	return filterNode(n, isAddress)
}

// filterArticleNodes will give the list of Article
func filterArticleNodes(n *html.Node) []*html.Node {
	return filterNode(n, isArticle)
}

// filterAsideNodes will give the list of Aside
func filterAsideNodes(n *html.Node) []*html.Node {
	return filterNode(n, isAside)
}

// filterBaseNodes will give the list of Aside
func filterBaseNodes(n *html.Node) []*html.Node {
	return filterNode(n, isBase)
}

// filterBasefontNodes will give the list of Basefont
func filterBasefontNodes(n *html.Node) []*html.Node {
	return filterNode(n, isBasefont)
}

// filterTitleNodes will give the list of Title
func filterTitleNodes(n *html.Node) []*html.Node {
	return filterNode(n, isTitle)
}

// filterTFootNodes will give the list of TFoot
func filterTFootNodes(n *html.Node) []*html.Node {
	return filterNode(n, isTFoot)
}

// filterTHeadNodes will give the list of THead
func filterTHeadNodes(n *html.Node) []*html.Node {
	return filterNode(n, isTHead)
}

// filterTBodyNodes will give the list of TBody
func filterTBodyNodes(n *html.Node) []*html.Node {
	return filterNode(n, isTbody)
}

// filterTemplateNodes will give the list of Template
func filterTemplateNodes(n *html.Node) []*html.Node {
	return filterNode(n, isTemplate)
}

// filterTableNodes will give the list of Table
func filterTableNodes(n *html.Node) []*html.Node {
	return filterNode(n, isTable)
}

// filterNavNodes will give the list of Nav
func filterNavNodes(n *html.Node) []*html.Node {
	return filterNode(n, isNav)
}

// filterMenuItemNodes will give the list of MenuItem
func filterMenuItemNodes(n *html.Node) []*html.Node {
	return filterNode(n, isMenuitem)
}

// filterMarqueeItemNodes will give the list of Marquee
func filterMarqueeItemNodes(n *html.Node) []*html.Node {
	return filterNode(n, isMarquee)
}

// filterMainItemNodes will give the list of Main
func filterMainItemNodes(n *html.Node) []*html.Node {
	return filterNode(n, isMain)
}

// filterLabelNodes will give the list of Label
func filterLabelNodes(n *html.Node) []*html.Node {
	return filterNode(n, isLabel)
}

// filterHeaderNodes will give the list of Header
func filterHeaderNodes(n *html.Node) []*html.Node { return filterNode(n, isHeader) }

// filterHeadNodes will give the list of Head
func filterHeadNodes(n *html.Node) []*html.Node { return filterNode(n, isHead) }

// filterFooterNodes will give the list of Footer
func filterFooterNodes(n *html.Node) []*html.Node { return filterNode(n, isFooter) }

// filterFrameNodes will give the list of Frame
func filterFrameNodes(n *html.Node) []*html.Node { return filterNode(n, isFrame) }

// filterFormNodes will give the list of Form
func filterFormNodes(n *html.Node) []*html.Node { return filterNode(n, isForm) }

// filterDirNodes will give the list of Dir
func filterDirNodes(n *html.Node) []*html.Node { return filterNode(n, isDir) }

// filterDialogNodes will give the list of Dialog
func filterDialogNodes(n *html.Node) []*html.Node { return filterNode(n, isDialog) }

// filterDetailsNodes will give the list of Details
func filterDetailsNodes(n *html.Node) []*html.Node { return filterNode(n, isDetails) }

// filterBodyNodes will give the list of Body
func filterBodyNodes(n *html.Node) []*html.Node { return filterNode(n, isBody) }

// filterCaptionNodes will give the list of Caption
func filterCaptionNodes(n *html.Node) []*html.Node { return filterNode(n, isCaption) }

// filterCiteNodes will give the list of Cite
func filterCiteNodes(n *html.Node) []*html.Node { return filterNode(n, isCite) }

// filterMainNodes will give the list of Main
func filterMainNodes(n *html.Node) []*html.Node { return filterNode(n, isMain) }

// filterCSSLinks will filter all the CSS based links
func filterCSSLinks(list []*html.Node) []*html.Node {
	var retLnk []*html.Node

	for _, n := range list {
		if isCSSLink(n) {
			retLnk = append(retLnk, n)
		}
	}
	return retLnk
}

// FilterLinkNodes will give the list of links
func FilterLinkNodes(n *html.Node) []*html.Node {
	return filterNode(n, isLink)
}

// filterNode is common filter method for all nodes
func filterNode(n *html.Node, f func(node *html.Node) bool) []*html.Node {
	var retLnk []*html.Node
	if f(n) {
		retLnk = append(retLnk, n)
		if hasChildren(n) {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				retLnk = append(retLnk, filterNode(c, f)...)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterNode(c, f)...)
	}
	return retLnk

}
