package analyzerapp

import (
	"golang.org/x/net/html"
	"strings"
)

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

// nodeText will retrieve node text
func nodeText(n *html.Node) string {
	var strbuilder strings.Builder
	//var str *string
	for _, a := range n.Attr {
		var sbdr strings.Builder
		sbdr.WriteString(a.Key)
		sbdr.WriteString("=\\")
		sbdr.WriteString(a.Val)
		sbdr.WriteString("\\ ")
		//s := a.Key + "=\"" + a.Val + "\" "
		strbuilder.WriteString(sbdr.String())
		//str += s
	}
	var resBuilder strings.Builder
	//var res string

	switch n.Data {
	case "div":
		if hasChildren(n) || text(n) != "" {
			resBuilder.WriteString("<div ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(text(n))
			resBuilder.WriteString("</div>")
			//res = "<div " + str + ">" + text(n) + "</div>"
		} else {
			resBuilder.WriteString("<div ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString("/>")

			//res = "<div " + str + "/>"
		}

	case "button":
		if hasChildren(n) || text(n) != "" {
			resBuilder.WriteString("<button ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(">")
			resBuilder.WriteString(text(n))
			resBuilder.WriteString("</button>")
			//res = "<button " + str + ">" + text(n) + "</button>"
		} else {
			resBuilder.WriteString("<button ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString("/>")
			//res = "<button " + str + "/>"
		}

	case "input":
		if hasChildren(n) || text(n) != "" {
			resBuilder.WriteString("<button ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(">")
			resBuilder.WriteString(text(n))
			resBuilder.WriteString("</button>")
			//res = "<button " + str + ">" + text(n) + "</button>"
		} else {
			resBuilder.WriteString("<input ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString("/>")
			//res = "<input " + str + "/>"
		}

	case "img":
		resBuilder.WriteString("<img ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString("/>")
		//res = "<img " + str + "/>"
	case "select":
		resBuilder.WriteString("<select ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString("/>")
		//res = "<select " + str + "/>"
	case "textarea":
		if hasChildren(n) || text(n) != "" {
			resBuilder.WriteString("<textarea ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(">")
			resBuilder.WriteString(text(n))
			resBuilder.WriteString("</textarea>")
			//res = "<textarea " + str + ">" + text(n) + "</textarea>"
		} else {
			resBuilder.WriteString("<textarea ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString("/>")
			//res = "<textarea " + str + "/>"
		}
	case "a":
		if hasChildren(n) || text(n) != "" {
			resBuilder.WriteString("<a ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(">")
			resBuilder.WriteString(text(n))
			resBuilder.WriteString("</a>")
			//res = "<a " + str + ">" + text(n) + "</a>"
		} else {
			//res = "<a " + str + "/>"
			resBuilder.WriteString("<a ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString("/>")
		}
	case "span":
		if hasChildren(n) || text(n) != "" {
			resBuilder.WriteString("<span ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(">")
			resBuilder.WriteString(text(n))
			resBuilder.WriteString("</span>")
			//res = "<span " + str + ">" + text(n) + "</span>"
		} else {
			//res = "<span " + str + "/>"
			resBuilder.WriteString("<span ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString("/>")
		}
	case "h1":
		//res = "<h1 " + str + ">" + text(n) + "</h1>"
		resBuilder.WriteString("<h1 ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString(">")
		resBuilder.WriteString(text(n))
		resBuilder.WriteString("</h1>")
	case "h2":
		//res = "<h2 " + str + ">" + text(n) + "</h2>"
		resBuilder.WriteString("<h2 ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString(">")
		resBuilder.WriteString(text(n))
		resBuilder.WriteString("</h2>")
	case "h3":
		//res = "<h3 " + str + ">" + text(n) + "</h3>"
		resBuilder.WriteString("<h3 ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString(">")
		resBuilder.WriteString(text(n))
		resBuilder.WriteString("</h3>")
	case "h4":
		//res = "<h4 " + str + ">" + text(n) + "</h4>"
		resBuilder.WriteString("<h4 ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString(">")
		resBuilder.WriteString(text(n))
		resBuilder.WriteString("</h4>")
	case "h5":
		//res = "<h5 " + str + ">" + text(n) + "</h5>"
		resBuilder.WriteString("<h5 ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString(">")
		resBuilder.WriteString(text(n))
		resBuilder.WriteString("</h5>")
	case "h6":
		//res = "<h6 " + str + ">" + text(n) + "</h6>"
		resBuilder.WriteString("<h6 ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString(">")
		resBuilder.WriteString(text(n))
		resBuilder.WriteString("</h6>")
	case "p":
		if hasChildren(n) || text(n) != "" {
			//res = "<p " + str + ">" + text(n) + "</p>"
			resBuilder.WriteString("<p ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(">")
			resBuilder.WriteString(text(n))
			resBuilder.WriteString("</p>")
		} else {
			//res = "<p " + str + "/>"
			resBuilder.WriteString("<p ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString("/>")
		}

	case "video":
		//res = "<video " + str + "/>"

	case "audio":
		//res = "<audio " + str + "/>"
		resBuilder.WriteString("<audio ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString("/>")

	}
	res := resBuilder.String()
	return res
}

// isTextNode will validate if node is TextNode
func isTextNode(n *html.Node) bool {
	if n.Type == html.TextNode {
		return true
	}
	return false
}

// text will get the text value of the node
func text(n *html.Node) string {
	if isTextNode(n) {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = text(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}
