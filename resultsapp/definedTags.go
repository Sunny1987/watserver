package resultsapp

type Tags struct {
	Applettag
	Areatag
	Anchortag
	Audiotag
	Buttontag
	Divtag
	Embedtag
	H1tag
	H2tag
	H3tag
	H4tag
	H5tag
	H6tag
	Imgtag
	Linktag
	Objecttag
	Pretag
	Paratag
	Inputtag
	Selecttag
	Textareatag
	Tracktag
	Videotag
}

// Result object will be used to simplify the tags result rules
type Result struct {
	Guideline string
	Rules     []string
}

// NewResult is the constructor for Result object
func NewResult(guideline string, rules []string) Result {
	return Result{Guideline: guideline, Rules: rules}
}

// Divtag is the Div based rule results
type Divtag struct {
	Div    string
	Result []Result
}

// Anchortag is the anchor based rule results
type Anchortag struct {
	Anchor string
	Result []Result
}

// Audiotag is the audio based rule results
type Audiotag struct {
	Audio  string
	Result []Result
}

// Buttontag is the button based rule results
type Buttontag struct {
	Button string
	Result []Result
}

// Imgtag is the img based rule results
type Imgtag struct {
	Img    string
	Result []Result
}

// Inputtag is the input based rule results
type Inputtag struct {
	Input  string
	Result []Result
}

// Selecttag is the select based rule results
type Selecttag struct {
	Select string
	Result []Result
}

// Textareatag is the textarea based rule results
type Textareatag struct {
	Textarea string
	Result   []Result
}

// Videotag is the video based rule results
type Videotag struct {
	Video  string
	Result []Result
}

// Paratag is the para based rule results
type Paratag struct {
	Para   string
	Result []Result
}

// Iframetag is the iframe based rule results
type Iframetag struct {
	Iframe string
	Result []Result
}

// Areatag is the area based rule results
type Areatag struct {
	Area   string
	Result []Result
}

// Objecttag is the object based rule results
type Objecttag struct {
	Object string
	Result []Result
}

// Embedtag is the embed based rule results
type Embedtag struct {
	Embed  string
	Result []Result
}

// Tracktag is the track based rule results
type Tracktag struct {
	Track  string
	Result []Result
}

// Applettag is the applet based rule results
type Applettag struct {
	Applet string
	Result []Result
}

// H1tag is the h1 based rule results
type H1tag struct {
	H1     string
	Result []Result
}

// H2tag is the h2 based rule results
type H2tag struct {
	H2     string
	Result []Result
}

// H3tag is the h3 based rule results
type H3tag struct {
	H3     string
	Result []Result
}

// H4tag is the h4 based rule results
type H4tag struct {
	H4     string
	Result []Result
}

// H5tag is the h5 based rule results
type H5tag struct {
	H5     string
	Result []Result
}

// H6tag is the h6 based rule results
type H6tag struct {
	H6     string
	Result []Result
}

// Pretag is the Pre based rule results
type Pretag struct {
	Pre    string
	Result []Result
}
type CSS struct {
	CSS    string
	Result []Result
}

// Linktag is the Link based rule results
type Linktag struct {
	Link   string
	Result []Result
}

// Abbrtag is the abbr based rule results
type Abbrtag struct {
	Abbr   string
	Result []Result
}

// SVGtag is the svg based rule results
type SVGtag struct {
	SVG    string
	Result []Result
}

// Canvastag is the canvas based rule results
type Canvastag struct {
	Canvas string
	Result []Result
}

// Spantag is the span based rule results
type Spantag struct {
	Span   string
	Result []Result
}

// Acronymtag is the acronym based rule results
type Acronymtag struct {
	Acronym string
	Result  []Result
}

// Addresstag is the address based rule results
type Addresstag struct {
	Address string
	Result  []Result
}

// Articletag is the article based rule results
type Articletag struct {
	Article string
	Result  []Result
}

// Asidetag is the aside based rule results
type Asidetag struct {
	Aside  string
	Result []Result
}

// Basetag is the base based rule results
type Basetag struct {
	Base   string
	Result []Result
}

// Basefonttag is the basefont based rule results
type Basefonttag struct {
	Basefont string
	Result   []Result
}

// Bditag is the bdi based rule results
type Bditag struct {
	Bdi    string
	Result []Result
}

// Bdotag is the bdo based rule results
type Bdotag struct {
	Bdo    string
	Result []Result
}

// Bgsoundtag is the Bgsound based rule results
type Bgsoundtag struct {
	Bgsound string
	Result  []Result
}

// Bigtag is the Big based rule results
type Bigtag struct {
	Big    string
	Result []Result
}

// Blockquotetag is the Blockquote based rule results
type Blockquotetag struct {
	Blockquote string
	Result     []Result
}

// Bodytag is the Body based rule results
type Bodytag struct {
	Body   string
	Result []Result
}

// Boldtag is the Bold based rule results
type Boldtag struct {
	Bold   string
	Result []Result
}

// Breaktag is the Break based rule results
type Breaktag struct {
	Break  string
	Result []Result
}

// Captiontag is the Caption based rule results
type Captiontag struct {
	Caption string
	Result  []Result
}

// Centertag is the Center based rule results
type Centertag struct {
	Center string
	Result []Result
}

// Citetag is the Cite based rule results
type Citetag struct {
	Cite   string
	Result []Result
}

// Codetag is the Code based rule results
type Codetag struct {
	Code   string
	Result []Result
}

// Colgrouptag is the Colgroup based rule results
type Colgrouptag struct {
	Colgroup string
	Result   []Result
}

// Columntag is the Column based rule results
type Columntag struct {
	Column string
	Result []Result
}

// Datatag is the Data based rule results
type Datatag struct {
	Data   string
	Result []Result
}

// Datalisttag is the Datalist based rule results
type Datalisttag struct {
	Datalist string
	Result   []Result
}

// Ddtag is the Dd based rule results
type Ddtag struct {
	Dd     string
	Result []Result
}

// Definetag is the Define based rule results
type Definetag struct {
	Define string
	Result []Result
}

// Deletetag is the Delete based rule results
type Deletetag struct {
	Delete string
	Result []Result
}

// Detailstag is the Detail based rule results
type Detailstag struct {
	Details string
	Result  []Result
}

// Dialogtag is the Dialog based rule results
type Dialogtag struct {
	Dialog string
	Result []Result
}

// Dirtag is the Dir based rule results
type Dirtag struct {
	Dir    string
	Result []Result
}

// Dltag is the Dl based rule results
type Dltag struct {
	Dl     string
	Result []Result
}

// Dttag is the Dt based rule results
type Dttag struct {
	Dt     string
	Result []Result
}

// Fieldsettag is the Fieldset based rule results
type Fieldsettag struct {
	Fieldset string
	Result   []Result
}

// Figcaptiontag is the Figcaption based rule results
type Figcaptiontag struct {
	Figcaption string
	Result     []Result
}

// Figuretag is the Figure based rule results
type Figuretag struct {
	Figure string
	Result []Result
}

// Fonttag is the Font based rule results
type Fonttag struct {
	Font   string
	Result []Result
}

// Footertag is the Footer based rule results
type Footertag struct {
	Footer string
	Result []Result
}

// Formtag is the Form based rule results
type Formtag struct {
	Form   string
	Result []Result
}

// Frametag is the Frame based rule results
type Frametag struct {
	Frame  string
	Result []Result
}

// Framesettag is the Frameset based rule results
type Framesettag struct {
	Frameset string
	Result   []Result
}

// Headtag is the Head based rule results
type Headtag struct {
	Head   string
	Result []Result
}

// Headertag is the Header based rule results
type Headertag struct {
	Header string
	Result []Result
}

// Hgrouptag is the Hgroup based rule results
type Hgrouptag struct {
	Hgroup string
	Result []Result
}

// Instag is the Ins based rule results
type Instag struct {
	Ins    string
	Result []Result
}

// Labeltag is the Label based rule results
type Labeltag struct {
	Label  string
	Result []Result
}

// Listtag is the List based rule results
type Listtag struct {
	List   string
	Result []Result
}

// Maintag is the Main based rule results
type Maintag struct {
	Main   string
	Result []Result
}

// Marqueetag is the Marquee based rule results
type Marqueetag struct {
	Marquee string
	Result  []Result
}

// Menuitemtag is the Menuitem based rule results
type Menuitemtag struct {
	Menuitem string
	Result   []Result
}

// Navtag is the Nav based rule results
type Navtag struct {
	Nav    string
	Result []Result
}

// Optgrouptag is the Optgroup based rule results
type Optgrouptag struct {
	Optgroup string
	Result   []Result
}

// Optiontag is the Option based rule results
type Optiontag struct {
	Option string
	Result []Result
}

// Paramtag is the Param based rule results
type Paramtag struct {
	Param  string
	Result []Result
}

// Tabletag is the Table based rule results
type Tabletag struct {
	Table  string
	Result []Result
}

// Tbodytag is the Tbody based rule results
type Tbodytag struct {
	Tbody  string
	Result []Result
}

// Tdtag is the Td based rule results
type Tdtag struct {
	Td     string
	Result []Result
}

// Templatetag is the Template based rule results
type Templatetag struct {
	Template string
	Result   []Result
}

// Tfoottag is the Tfoot based rule results
type Tfoottag struct {
	Tfoot  string
	Result []Result
}

// Thtag is the Th based rule results
type Thtag struct {
	Th     string
	Result []Result
}

// Theadtag is the Thead based rule results
type Theadtag struct {
	Thead  string
	Result []Result
}

// Titletag is the Title based rule results
type Titletag struct {
	Title  string
	Result []Result
}
