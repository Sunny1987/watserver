package resultsapp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response is the response object for json
type Response struct {
	Request         interface{} `json:"request"`
	Person          *string     `json:"person"`
	DivResults      *[]Div      `json:"divs"`
	ButtonResults   *[]Button   `json:"buttons"`
	InputResults    *[]Input    `json:"inputs"`
	ImageResults    *[]Img      `json:"images"`
	VideoResults    *[]Video    `json:"videos"`
	AudioResults    *[]Audio    `json:"audios"`
	TextareaResults *[]Textarea `json:"textAreas"`
	SelectResults   *[]Select   `json:"selects"`
	ParaResults     *[]Para     `json:"paras"`
	IframeResults   *[]Iframe   `json:"iframes"`
	AnchorResults   *[]Anchor   `json:"anchors"`
	AreaResults     *[]Area     `json:"areas"`
	ObjectResults   *[]Object   `json:"objects"`
	EmbedResults    *[]Embed    `json:"embeds"`
	TrackResults    *[]Track    `json:"tracks"`
	H1Results       *[]H1       `json:"h1s"`
	H2Results       *[]H2       `json:"h2s"`
	H3Results       *[]H3       `json:"h3s"`
	H4Results       *[]H4       `json:"h4s"`
	H5Results       *[]H5       `json:"h5s"`
	H6Results       *[]H6       `json:"h6s"`
	PreResults      *[]Pre      `json:"pres"`
	LinkResults     *[]Link     `json:"links"`
	AbbrResults     *[]Abbr     `json:"abbrs"`
	SvgResults      *[]SVG      `json:"svgs"`
	CanvasResults   *[]Canvas   `json:"canvases"`
	SpanResults     *[]Span     `json:"spans"`
	NavResults      *[]Nav      `json:"navs"`
	MainResults     *[]Main     `json:"mains"`
	AsideResults    *[]Aside    `json:"asides"`
	CSSResults      *[]CSS      `json:"CSS"`
}

func PrintResponse(rw http.ResponseWriter, l *log.Logger, results []Response) {
	l.Println("Initiating the response....")
	var filteredRes []Response
	if len(results) > 100 {
		filteredRes = results[:99]
	} else {
		filteredRes = results
	}
	rep, err := json.MarshalIndent(filteredRes, "", " ")
	if err != nil {
		l.Println(err)
	}
	fresp := string(rep)
	_, err = fmt.Fprintln(rw, fresp)
	if err != nil {
		l.Printf("Error : %v", err)
	}
}
