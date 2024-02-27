package resultsapp

type Tags struct {
	Applet
	Area
	Anchor
	Audio
	Button
	Div
	Embed
	H1
	H2
	H3
	H4
	H5
	H6
	Img
	Link
	Object
	Pre
	Para
	Input
	Select
	Textarea
	Track
	Video
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

// Div is the Div based rule results
type Div struct {
	Tag    string
	Result []Result
}

// Anchor is the anchor based rule results
type Anchor struct {
	Tag    string
	Result []Result
}

// Audio is the audio based rule results
type Audio struct {
	Tag    string
	Result []Result
}

// Button is the button based rule results
type Button struct {
	Tag    string
	Result []Result
}

// Img is the img based rule results
type Img struct {
	Tag    string
	Result []Result
}

// Input is the input based rule results
type Input struct {
	Tag    string
	Result []Result
}

// Select is the select based rule results
type Select struct {
	Tag    string
	Result []Result
}

// Textarea is the textarea based rule results
type Textarea struct {
	Tag    string
	Result []Result
}

// Video is the video based rule results
type Video struct {
	Tag    string
	Result []Result
}

// Para is the para based rule results
type Para struct {
	Tag    string
	Result []Result
}

// Iframe is the iframe based rule results
type Iframe struct {
	Tag    string
	Result []Result
}

// Area is the area based rule results
type Area struct {
	Tag    string
	Result []Result
}

// Object is the object based rule results
type Object struct {
	Tag    string
	Result []Result
}

// Embed is the embed based rule results
type Embed struct {
	Tag    string
	Result []Result
}

// Track is the track based rule results
type Track struct {
	Tag    string
	Result []Result
}

// Applet is the applet based rule results
type Applet struct {
	Tag    string
	Result []Result
}

// H1 is the h1 based rule results
type H1 struct {
	Tag    string
	Result []Result
}

// H2 is the h2 based rule results
type H2 struct {
	Tag    string
	Result []Result
}

// H3 is the h3 based rule results
type H3 struct {
	Tag    string
	Result []Result
}

// H4 is the h4 based rule results
type H4 struct {
	Tag    string
	Result []Result
}

// H5 is the h5 based rule results
type H5 struct {
	Tag    string
	Result []Result
}

// H6 is the h6 based rule results
type H6 struct {
	Tag    string
	Result []Result
}

// Pre is the Pre based rule results
type Pre struct {
	Tag    string
	Result []Result
}
type CSS struct {
	CSS    string
	Result []Result
}

// Link is the Link based rule results
type Link struct {
	Tag    string
	Result []Result
}

// Abbr is the abbr based rule results
type Abbr struct {
	Tag    string
	Result []Result
}

// SVG is the svg based rule results
type SVG struct {
	Tag    string
	Result []Result
}

// Canvas is the canvas based rule results
type Canvas struct {
	Tag    string
	Result []Result
}

// Span is the span based rule results
type Span struct {
	Tag    string
	Result []Result
}

// Acronym is the acronym based rule results
type Acronym struct {
	Tag    string
	Result []Result
}

// Address is the address based rule results
type Address struct {
	Tag    string
	Result []Result
}

// Article is the article based rule results
type Article struct {
	Tag    string
	Result []Result
}

// Aside is the aside based rule results
type Aside struct {
	Tag    string
	Result []Result
}

// Base is the base based rule results
type Base struct {
	Tag    string
	Result []Result
}

// Basefont is the basefont based rule results
type Basefont struct {
	Tag    string
	Result []Result
}

// Bdi is the bdi based rule results
type Bdi struct {
	Tag    string
	Result []Result
}

// Bdo is the bdo based rule results
type Bdo struct {
	Tag    string
	Result []Result
}

// Bgsound is the Bgsound based rule results
type Bgsound struct {
	Tag    string
	Result []Result
}

// Big is the Big based rule results
type Big struct {
	Tag    string
	Result []Result
}

// Blockquote is the Blockquote based rule results
type Blockquote struct {
	Tag    string
	Result []Result
}

// Body is the Body based rule results
type Body struct {
	Tag    string
	Result []Result
}

// Bold is the Bold based rule results
type Bold struct {
	Tag    string
	Result []Result
}

// Break is the Break based rule results
type Break struct {
	Tag    string
	Result []Result
}

// Caption is the Caption based rule results
type Caption struct {
	Tag    string
	Result []Result
}

// Center is the Center based rule results
type Center struct {
	Tag    string
	Result []Result
}

// Cite is the Cite based rule results
type Cite struct {
	Tag    string
	Result []Result
}

// Code is the Code based rule results
type Code struct {
	Tag    string
	Result []Result
}

// Colgroup is the Colgroup based rule results
type Colgroup struct {
	Tag    string
	Result []Result
}

// Column is the Column based rule results
type Column struct {
	Tag    string
	Result []Result
}

// Data is the Data based rule results
type Data struct {
	Tag    string
	Result []Result
}

// Datalist is the Datalist based rule results
type Datalist struct {
	Tag    string
	Result []Result
}

// Dd is the Dd based rule results
type Dd struct {
	Tag    string
	Result []Result
}

// Define is the Define based rule results
type Define struct {
	Tag    string
	Result []Result
}

// Delete is the Delete based rule results
type Delete struct {
	Tag    string
	Result []Result
}

// Details is the Detail based rule results
type Details struct {
	Tag    string
	Result []Result
}

// Dialog is the Dialog based rule results
type Dialog struct {
	Tag    string
	Result []Result
}

// Dir is the Dir based rule results
type Dir struct {
	Tag    string
	Result []Result
}

// Dl is the Dl based rule results
type Dl struct {
	Tag    string
	Result []Result
}

// Dt is the Dt based rule results
type Dt struct {
	Tag    string
	Result []Result
}

// Fieldset is the Fieldset based rule results
type Fieldset struct {
	Tag    string
	Result []Result
}

// Figcaption is the Figcaption based rule results
type Figcaption struct {
	Tag    string
	Result []Result
}

// Figure is the Figure based rule results
type Figure struct {
	Tag    string
	Result []Result
}

// Font is the Font based rule results
type Font struct {
	Tag    string
	Result []Result
}

// Footer is the Footer based rule results
type Footer struct {
	Tag    string
	Result []Result
}

// Form is the Form based rule results
type Form struct {
	Tag    string
	Result []Result
}

// Frame is the Frame based rule results
type Frame struct {
	Tag    string
	Result []Result
}

// Frameset is the Frameset based rule results
type Frameset struct {
	Tag    string
	Result []Result
}

// Head is the Head based rule results
type Head struct {
	Tag    string
	Result []Result
}

// Header is the Header based rule results
type Header struct {
	Tag    string
	Result []Result
}

// Hgroup is the Hgroup based rule results
type Hgroup struct {
	Tag    string
	Result []Result
}

// Ins is the Ins based rule results
type Ins struct {
	Tag    string
	Result []Result
}

// Label is the Label based rule results
type Label struct {
	Tag    string
	Result []Result
}

// List is the List based rule results
type List struct {
	Tag    string
	Result []Result
}

// Main is the Main based rule results
type Main struct {
	Tag    string
	Result []Result
}

// Marquee is the Marquee based rule results
type Marquee struct {
	Tag    string
	Result []Result
}

// Menuitem is the Menuitem based rule results
type Menuitem struct {
	Tag    string
	Result []Result
}

// Nav is the Nav based rule results
type Nav struct {
	Tag    string
	Result []Result
}

// Optgroup is the Optgroup based rule results
type Optgroup struct {
	Tag    string
	Result []Result
}

// Option is the Option based rule results
type Option struct {
	Tag    string
	Result []Result
}

// Param is the Param based rule results
type Param struct {
	Tag    string
	Result []Result
}

// Table is the Table based rule results
type Table struct {
	Tag    string
	Result []Result
}

// Tbody is the Tbody based rule results
type Tbody struct {
	Tag    string
	Result []Result
}

// Td is the Td based rule results
type Td struct {
	Tag    string
	Result []Result
}

// Template is the Template based rule results
type Template struct {
	Tag    string
	Result []Result
}

// Tfoot is the Tfoot based rule results
type Tfoot struct {
	Tag    string
	Result []Result
}

// Th is the Th based rule results
type Th struct {
	Tag    string
	Result []Result
}

// Thead is the Thead based rule results
type Thead struct {
	Tag    string
	Result []Result
}

// Title is the Title based rule results
type Title struct {
	Tag    string
	Result []Result
}
