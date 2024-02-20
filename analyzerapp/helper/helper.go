package helper

import (
	"strings"

	"golang.org/x/net/html"
)

// AttributeSearch will respond if the attribute is present
func AttributeSearch(attrList []html.Attribute, key string) bool {
	for _, att := range attrList {
		if att.Key == key {
			return true
		}
	}
	return false
}

// AttributeCheckVal will respond if the attribute is present and check its value
func AttributeCheckVal(attrList []html.Attribute, key string, val string) bool {
	for _, att := range attrList {
		if att.Key == key && att.Val == val {
			return true
		}
	}
	return false
}

// AttributeCheckValEmpty will respond if the attribute is present and value is empty
func AttributeCheckValEmpty(attrList []html.Attribute, key string) bool {
	for _, att := range attrList {
		if att.Key == key && att.Val != "" {
			return true
		}
	}
	return false
}

// hasNoChild will check if node has child
func HasNoChild(n *html.Node) bool {
	return n.FirstChild == nil
}

// hasChildren will check if node has children
func HasChildren(n *html.Node) bool {
	return n.FirstChild != nil && n.FirstChild != n.LastChild
}

// hasOneChild will check if node has only one child
func HasOneChild(n *html.Node) bool {
	return n.FirstChild != nil && n.FirstChild == n.LastChild
}

// nodeText will retrieve node text
func NodeText(n *html.Node) string {
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
		if HasChildren(n) || Text(n) != "" {
			resBuilder.WriteString("<div ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(Text(n))
			resBuilder.WriteString("</div>")
			//res = "<div " + str + ">" + text(n) + "</div>"
		} else {
			resBuilder.WriteString("<div ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString("/>")

			//res = "<div " + str + "/>"
		}

	case "button":
		if HasChildren(n) || Text(n) != "" {
			resBuilder.WriteString("<button ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(">")
			resBuilder.WriteString(Text(n))
			resBuilder.WriteString("</button>")
			//res = "<button " + str + ">" + text(n) + "</button>"
		} else {
			resBuilder.WriteString("<button ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString("/>")
			//res = "<button " + str + "/>"
		}

	case "input":
		if HasChildren(n) || Text(n) != "" {
			resBuilder.WriteString("<button ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(">")
			resBuilder.WriteString(Text(n))
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
		if HasChildren(n) || Text(n) != "" {
			resBuilder.WriteString("<textarea ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(">")
			resBuilder.WriteString(Text(n))
			resBuilder.WriteString("</textarea>")
			//res = "<textarea " + str + ">" + text(n) + "</textarea>"
		} else {
			resBuilder.WriteString("<textarea ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString("/>")
			//res = "<textarea " + str + "/>"
		}
	case "a":
		if HasChildren(n) || Text(n) != "" {
			resBuilder.WriteString("<a ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(">")
			resBuilder.WriteString(Text(n))
			resBuilder.WriteString("</a>")
			//res = "<a " + str + ">" + text(n) + "</a>"
		} else {
			//res = "<a " + str + "/>"
			resBuilder.WriteString("<a ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString("/>")
		}
	case "span":
		if HasChildren(n) || Text(n) != "" {
			resBuilder.WriteString("<span ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(">")
			resBuilder.WriteString(Text(n))
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
		resBuilder.WriteString(Text(n))
		resBuilder.WriteString("</h1>")
	case "h2":
		//res = "<h2 " + str + ">" + text(n) + "</h2>"
		resBuilder.WriteString("<h2 ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString(">")
		resBuilder.WriteString(Text(n))
		resBuilder.WriteString("</h2>")
	case "h3":
		//res = "<h3 " + str + ">" + text(n) + "</h3>"
		resBuilder.WriteString("<h3 ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString(">")
		resBuilder.WriteString(Text(n))
		resBuilder.WriteString("</h3>")
	case "h4":
		//res = "<h4 " + str + ">" + text(n) + "</h4>"
		resBuilder.WriteString("<h4 ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString(">")
		resBuilder.WriteString(Text(n))
		resBuilder.WriteString("</h4>")
	case "h5":
		//res = "<h5 " + str + ">" + text(n) + "</h5>"
		resBuilder.WriteString("<h5 ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString(">")
		resBuilder.WriteString(Text(n))
		resBuilder.WriteString("</h5>")
	case "h6":
		//res = "<h6 " + str + ">" + text(n) + "</h6>"
		resBuilder.WriteString("<h6 ")
		resBuilder.WriteString(strbuilder.String())
		resBuilder.WriteString(">")
		resBuilder.WriteString(Text(n))
		resBuilder.WriteString("</h6>")
	case "p":
		if HasChildren(n) || Text(n) != "" {
			//res = "<p " + str + ">" + text(n) + "</p>"
			resBuilder.WriteString("<p ")
			resBuilder.WriteString(strbuilder.String())
			resBuilder.WriteString(">")
			resBuilder.WriteString(Text(n))
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
func IsTextNode(n *html.Node) bool {
	return n.Type == html.TextNode
}

// text will get the text value of the node
func Text(n *html.Node) string {
	if IsTextNode(n) {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = Text(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}
