package templates

import (
	"fmt"
	"strings"
	"reflect"
	"runtime"
)
type Item struct {
	id string
	template *Template
	writer *Writer
	Value string
}
func NewItemStub(id string, writer *Writer) *Item{
	return &Item{
		id : id,
		writer : writer,
	}
}
func (i *Item) Click(handlerFunc interface{}, items ...interface{}) {
	itemsString := ``
	for i:=0; i<len(items); i++ {
		if item, ok := items[i].(*Item); ok {
			if len(itemsString) > 0 { itemsString += "," }
			itemsString += `"`+item.FullId()+`" : ""`
		}
	}
	v := reflect.ValueOf(handlerFunc)
	f := runtime.FuncForPC(v.Pointer()).Name()
	handler := f[strings.Index(f, "·") + 2:]
	i.writer.Buffer += `
$("#` + i.FullId() + `").click(function(){var itemsMap = {`+itemsString+`}; getValues(itemsMap); itemsMap["function"]="`+handler+`"; $.post("/function", itemsMap, function(data){$("#head").append($("<div>").html(data))}, "html")});`
}
func (i *Item) Html(o ...interface{}) {
	i.generic(true, o)
}
func (i *Item) Append(o ...interface{}) {
	i.generic(false, o)
}
func (i *Item) generic(replace bool, o []interface{}) {
	for j := 0; j < len(o); j++ {
		o1 := o[j]
		
		switch t := o1.(type) {
			case *Template:
				i.templateGeneric(replace && j==0, t)
			case string:
				i.htmlGeneric(replace && j==0, t)
			case int:
				i.htmlGeneric(replace && j==0, fmt.Sprint(t))
			case float32:
				i.htmlGeneric(replace && j==0, fmt.Sprint(t))
			case float64:
				i.htmlGeneric(replace && j==0, fmt.Sprint(t))
			case bool:
				i.htmlGeneric(replace && j==0, fmt.Sprint(t))
			default:
				if t, ok := o1.(Templater); ok {
					i.templateGeneric(replace && j==0, t.GetTemplate())
				} else {
					panic("Html / Append failed")
				}
		}
	}
}
func (i *Item) htmlGeneric(replace bool, s string) {
	command := ""
	if replace { command = "html" } else { command = "append" }
	i.writer.Buffer += `
$("#` + i.FullId() + `").`+command+`("` + s + `");`
}

func (i *Item) templateGeneric(replace bool, t *Template) {
	t.parentId = i.FullId()
	command := ""
	if replace { command = "html" } else { command = "append" }
	i.writer.Buffer += `
$("#` + i.FullId() + `").`+command+`(template_` + t.name + `("` + t.FullId() + `"));`
}

func (i *Item) Attr(attrib string, o interface{}) {
		switch t := o.(type) {
			case string:
				i.attr(attrib, t)
			case int:
				i.attr(attrib, fmt.Sprint(t))
			case float32:
				i.attr(attrib, fmt.Sprint(t))
			case float64:
				i.attr(attrib, fmt.Sprint(t))
			case bool:
				i.attr(attrib, fmt.Sprint(t))
		}
}
func (i *Item) attr(attrib string, val string) {
	i.writer.Buffer += `
$("#` + i.FullId() + `").attr("` + attrib + `", "` + val + `");`
}


func (i *Item) FullId() string {
	
	if i.template == nil {
		return i.id
	} else {
		return i.template.FullId() + `_` + i.id
	}
	return ``
}

