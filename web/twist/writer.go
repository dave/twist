package twist

import (
	"http"
	"fmt"
)

type Writer struct {
	Output    http.ResponseWriter
	Buffer    string
	Templates []*Template
	SendRoot  bool
	SendHtml  bool
}

func getRoot(w *Writer) *Item {

	return &Item{
		id:         "root",
		template:   nil,
		writer:     w,
		Attributes: make(map[string]string),
		Styles:     make(map[string]string),
	}

}

func newWriter(o http.ResponseWriter, sendRoot bool, sendHtml bool) *Writer {
	return &Writer{
		Output:    o,
		Buffer:    "",
		Templates: make([]*Template, 0),
		SendRoot:  sendRoot,
		SendHtml:  sendHtml,
	}
}

func (w *Writer) RegisterTemplate(t *Template) {

	for i := 0; i < len(w.Templates); i++ {
		if w.Templates[i].Name == t.Name {
			return
		}
	}
	w.Templates = append(w.Templates, t)

}

func (c *Context) Send() {

	c.Root.RunCommands()

	for _, inner := range c.itemsInRequest {
		inner.RunCommands()
	}

	if c.Writer.SendRoot {
		c.Writer.sendPage(c.Root)
	} else {
		c.Writer.sendFragment()
	}

}

func (w *Writer) sendPage(item *Item) {

	root := `
<script src="/static/jquery.js"></script>
<script src="/static/json.js"></script>
<script src="/static/helpers.js"></script>
<script src="/static/native.history.js"></script>
<div id="head"></div>
<div id="root">`
	root += item.RenderHtml()
	root += `
</div>
<script>
	var ignoreNextStateChange = false;
	(function(window,undefined){
	    var History = window.History;
	    if ( !History.enabled ) {
	        return false;
	    }
	    History.Adapter.bind(window,'statechange',function(){
	    	if (!ignoreNextStateChange)
	    	{
				var State = History.getState();
				//History.log(State.data, State.title, State.url);
				$.post(State.url, null, function(data){$("#head").append($("<div>").html(data))}, "html");
			}
			else
				ignoreNextStateChange = false
	    });
	})(window);

`

	templates := ""
	script := ""

	if len(w.Templates) > 0 {
		templates = `
var templatesToLoad = ` + fmt.Sprint(len(w.Templates)) + `;
var templatesLoaded = 0;`
		for i := 0; i < len(w.Templates); i++ {
			templates += `
$("#head").append($("<div>").load("/template_` + w.Templates[i].Path + `", function() {templatesLoaded++;if(templatesLoaded == templatesToLoad){runScript();}}));`
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
$("#head").append($("<div>").load("/template_` + w.Templates[i].Path + `", function() {templatesLoaded++;if(templatesLoaded == templatesToLoad){runScript();}}));`
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
