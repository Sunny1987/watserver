package parser

import (
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html"
	"log"
	"strings"
)

// isAnchor will validate an Anchor tag
func isAnchor(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

// isDiv will validate an Div tag
func isDiv(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "div"
}

// isSpan will validate an span tag
func isSpan(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "span"
}

// isH1 will validate an H1 tag
func isH1(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h1"
}

// isH2 will validate an H2 tag
func isH2(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h2"
}

// isH3 will validate an H3 tag
func isH3(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h3"
}

// isH4 will validate an H4 tag
func isH4(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h4"
}

// isH5 will validate an H5 tag
func isH5(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h5"
}

// isH6 will validate an H6 tag
func isH6(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h6"
}

// isParagraph will validate an Paragraph tag
func isParagraph(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "p"
}

// isImage will validate an Image tag
func isImage(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "img"
}

// isUL will validate an UL tag
func isUL(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "ul"
}

// isOL will validate an OL tag
func isOL(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "ol"
}

// isLI will validate an LI tag
func isLI(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "li"
}

// isScript will validate an Script tag
func isScript(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "script"
}

// isDocTypeOrDocumentType will validate if a tag is DocType or DocumentType
func isDocTypeOrDocumentType(n *html.Node) bool {
	return n.Type == html.DocumentNode || n.Type == html.DoctypeNode
}

// isTextNode will validate if node is TextNode
func isTextNode(n *html.Node) bool {
	if n.Type == html.TextNode && n.DataAtom == 0 {
		return true
	}
	return false
}

// isInput will validate an Image tag
func isInput(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "input"
}

// isVideo will validate an Image tag
func isVideo(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "video"
}

// isVideo will validate an Image tag
func isAudio(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "audio"
}

// isSelect will validate an Image tag
func isSelect(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "select"
}

// isTextArea will validate an Image tag
func isTextArea(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "textarea"
}

// isButton will validate an Image tag
func isButton(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "button"
}

// isIFrame will validate an Image tag
func isIFrame(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "iframe"
}

// isArea will validate an Area tag
func isArea(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "area"
}

// isObject will validate an object tag
func isObject(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "object"
}

// isEmbed will validate an embed tag
func isEmbed(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "embed"
}

// isTrack will validate an track tag
func isTrack(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "track"
}

// isApplet will validate an applet tag
func isApplet(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "applet"
}

// isPre will validate pre tag
func isPre(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "pre"
}

// isSVG will validate svg tag
func isSVG(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "svg"
}

// isAbbr will validate abbr tag
func isAbbr(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "abbr"
}

// isCanvas will validate canvas tag
func isCanvas(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "canvas"
}

// isAcronym will validate Acronym tag
func isAcronym(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "acronym"
}

// isAddress will validate Address tag
func isAddress(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "address"
}

// isArticle will validate Article tag
func isArticle(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "article"
}

// isAside will validate Aside tag
func isAside(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "aside"
}

// isBase will validate base tag
func isBase(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "base"
}

// isBasefont will validate Basefont tag
func isBasefont(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "basefont"
}

// isBdi will validate Bdi tag
func isBdi(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "bdi"
}

// isBdo will validate Bdo tag
func isBdo(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "bdo"
}

// isBgsound will validate Bgsound tag
func isBgsound(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "bgsound"
}

// isBig will validate Big tag
func isBig(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "big"
}

// isBlockquote will validate Blockquote tag
func isBlockquote(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "blockquote"
}

// isBody will validate Body tag
func isBody(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "body"
}

// isCaption will validate Caption tag
func isCaption(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "caption" }

// isCite will validate Cite tag
func isCite(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "cite"
}

// isCode will validate Code tag
func isCode(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "code"
}

// isColgroup will validate Colgroup tag
func isColgroup(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "colgroup"
}

// isColumn will validate Column tag
func isColumn(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "column"
}

// isData will validate Data tag
func isData(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "data"
}

// isDatalist will validate Datalist tag
func isDatalist(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "datalist" }

// isDefine will validate Define tag
func isDefine(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "define" }

// isDelete will validate Delete tag
func isDelete(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "delete" }

// isDetails will validate Details tag
func isDetails(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "details" }

// isDialog will validate Dialog tag
func isDialog(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "dialog" }

// isDir will validate Dir tag
func isDir(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "dir" }

// isForm will validate Form tag
func isForm(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "form" }

// isFrame will validate Frame tag
func isFrame(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "frame" }

// isFooter will validate Footers tag
func isFooter(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "footer" }

// isHead will validate Head tag
func isHead(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "head" }

// isHeader will validate Headers tag
func isHeader(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "header" }

// isIns will validate Ins tag
func isIns(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "ins" }

// isLabel will validate Label tag
func isLabel(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "label" }

// isMain will validate Main tag
func isMain(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "main" }

// isMarquee will validate Marquee tag
func isMarquee(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "marquee" }

// isMenuitem will validate Menuitem tag
func isMenuitem(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "menuitem" }

// isNav will validate Nav tag
func isNav(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "nav" }

// isTable will validate Table tag
func isTable(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "table" }

// isTbody will validate Tbody tag
func isTbody(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "tbody" }

// isTemplate will validate Template tag
func isTemplate(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "template" }

// isTHead will validate THead tag
func isTHead(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "thead" }

// isTFoot will validate TFoot tag
func isTFoot(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "tfoot" }

// isTitle will validate Title tag
func isTitle(n *html.Node) bool { return n.Type == html.ElementNode && n.Data == "title" }

func isCSSLink(n *html.Node) bool {
	for _, att := range n.Attr {
		if att.Key == "rel" && att.Val == "stylesheet" {
			return true
		}
		if att.Key == "type" && att.Val == "text/css" {
			return true
		}
	}
	return false
}

// isLink will validate an Anchor tag
func isLink(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "link"
}

// HrefLinks will return the formatted links
func HrefLinks(linkNodes []*html.Node, base string, l *log.Logger) []string {

	l.Println("***Formatting the link nodes****")

	//format the correct links
	var hrefs []string
	for _, l := range linkNodes {

		var href string
		for _, att := range l.Attr {
			if att.Key == "href" {
				href = att.Val
			}
		}
		//log.Printf("href : %v",href)
		switch {
		case strings.HasPrefix(href, "/"):
			if !containsData(hrefs, base+href) {
				hrefs = append(hrefs, base+href)
			}
		case strings.HasPrefix(href, "http"):
			if !containsData(hrefs, href) {
				hrefs = append(hrefs, href)
			}
		}
	}

	//wg.Wait()
	l.Println("***Completed formatting the link nodes****")
	return hrefs
}

// containsData checks redundancy in cssList
func containsData(l []string, val string) bool {
	for _, str := range l {
		if str == val {
			return true
		}
	}
	return false
}

func readCSSLinks(link string, l *log.Logger) {
	var client *resty.Client = resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").Get(link)
	if err != nil {
		l.Println(err)

	}
	cssList = append(cssList, string(resp.Body()))

}

// hasNoChild will check if node has child
func hasNoChild(n *html.Node) bool {
	return n.FirstChild == nil
}

// hasChildren will check if node has children
func hasChildren(n *html.Node) bool {
	return n.FirstChild != nil && n.FirstChild != n.LastChild
}

// hasOneChild will check if node has only one child
func hasOneChild(n *html.Node) bool {
	return n.FirstChild != nil && n.FirstChild == n.LastChild
}
