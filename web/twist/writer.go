package twist

import (
	"http"
	"fmt"
)

type Writer struct {
	Output    http.ResponseWriter
	Buffer    string
	Templates []Template
	IsRoot    bool
}

func Root(w *Writer) *Item {

	return &Item{
		id:       "root",
		template: nil,
		writer:   w,
	}

}

func NewWriter(o http.ResponseWriter, isRoot bool) *Writer {
	return &Writer{
		Output:    o,
		Buffer:    "",
		Templates: make([]Template, 0),
		IsRoot:    isRoot}
}

func (w *Writer) RegisterTemplate(t Template) {

	for i := 0; i < len(w.Templates); i++ {
		if w.Templates[i].name == t.name {
			return
		}
	}
	w.Templates = append(w.Templates, t)

}

func (c *Context) Send() {
	if c.Writer.IsRoot {
		c.Writer.sendHtml(c.Root)
	} else {
		c.Writer.sendFragment()
	}

}

func (w *Writer) sendHtml(item *Item) {

	fmt.Fprint(w.Output, item.RenderHtml())

}
func (w *Writer) sendPage() {

	root := `
<script src="/static/jquery.js"></script>
<script src="/static/json.js"></script>
<script src="/static/helpers.js"></script>
<div id="head"></div>
<div id="root"></div>
<script>`

	templates := ""
	script := ""

	if len(w.Templates) > 0 {
		templates = `
var templatesToLoad = ` + fmt.Sprint(len(w.Templates)) + `;
var templatesLoaded = 0;`
		for i := 0; i < len(w.Templates); i++ {
			templates += `
$("#head").append($("<div>").load("/template_` + w.Templates[i].name + `", function() {templatesLoaded++;if(templatesLoaded == templatesToLoad){runScript();}}));`
		}

		script = `
function runScript()
{` + w.Buffer + `
}
</script>`

	} else {
		script = w.Buffer + `
</script>`
	}

	fmt.Fprint(w.Output, root+templates+script)
	w.Buffer = ""

}
func (w *Writer) sendFragment() {

	templates := ``
	script := ``
	if len(w.Templates) > 0 {
		templates = `
var templatesToLoad = ` + fmt.Sprint(len(w.Templates)) + `;
var templatesLoaded = 0;`
		for i := 0; i < len(w.Templates); i++ {
			templates += `
$("#head").append($("<div>").load("/template_` + w.Templates[i].name + `", function() {templatesLoaded++;if(templatesLoaded == templatesToLoad){runScript();}}));`
		}
		script = `
function runScript()
{` + w.Buffer + `
}`
	} else {
		script = w.Buffer
	}

	fmt.Fprint(w.Output,
		`<script>`+templates+script+`</script>`)
	w.Buffer = ""

}
