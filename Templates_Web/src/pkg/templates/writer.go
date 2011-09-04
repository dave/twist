package templates

import (
	"http"
	"fmt"
)

func Root(w *Writer) *Item {

	return &Item{
		id : "root",
		template : nil, 
		writer : w,
	}

}
type Writer struct {
	Output    http.ResponseWriter
	Buffer    string
	Templates []string
	IsRoot    bool
}

func NewWriter(o http.ResponseWriter, isRoot bool) *Writer {
	return &Writer{
		Output:    o,
		Buffer:    "",
		Templates: make([]string, 0),
		IsRoot : isRoot}
}

func (w *Writer) RegisterTemplate(name string) {

	for i := 0; i < len(w.Templates); i++ {
		if w.Templates[i] == name {
			return
		}
	}
	w.Templates = append(w.Templates, name)

}

func (w *Writer) Send() {
	if (w.IsRoot) {
		w.sendPage()
	} else {
		w.sendFragment()
	}
		
}

func (w *Writer) sendPage() {

	root := `
<script src="/static/jquery.js"></script>
<script src="/static/json.js"></script>
<script>
function getValues(items) {
	$.each(items, function(i,n){try{items[i].V = $("#" + items[i].I).val()}catch(ex){}})
}
</script>
<div id="head"></div>
<div id="root"></div>
<script>`

	templates := `
var templatesToLoad = ` + fmt.Sprint(len(w.Templates)) + `;
var templatesLoaded = 0;`
	for i := 0; i < len(w.Templates); i++ {
		templates += `
$("#head").append($("<div>").load("/template_` + w.Templates[i] + `", function() {templatesLoaded++;if(templatesLoaded == templatesToLoad){runScript();}}));`
	}
	
	script := `
function runScript()
{`+w.Buffer+`
}
</script>`

	fmt.Fprint(w.Output,
		root + templates + script)
	w.Buffer = ""

}
func (w *Writer) sendFragment() {

	templates := ``
	script := ``
	if len(w.Templates) > 0 {
		templates = `
var templatesToLoad_1 = ` + fmt.Sprint(len(w.Templates)) + `;
var templatesLoaded_1 = 0;`
		for i := 0; i < len(w.Templates); i++ {
			templates += `
$("#head").append($("<div>").load("/template_` + w.Templates[i] + `", function() {templatesLoaded_1++;if(templatesLoaded_1 == templatesToLoad_1){runScript_1();}}));`
		}
		script = `
function runScript_1()
{`+w.Buffer+`
}`
	} else {
		script = w.Buffer
	}
	

	fmt.Fprint(w.Output,
		`<script>` + templates + script + `</script>`)
	w.Buffer = ""

}