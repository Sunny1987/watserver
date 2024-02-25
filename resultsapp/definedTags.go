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

type Abbrtag struct {
	Abbr   string
	Result []Result
}
