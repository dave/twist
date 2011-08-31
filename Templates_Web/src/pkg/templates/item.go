package templates

import (
	"fmt"
	"reflect"
	//"http"
	"json"
)

type Value string
type ValueEncrypted string
type Item struct {
	id       string
	template *Template
	writer   *Writer
	Value    string
}

func newItemFromAction(id string, writer *Writer) *Item {
	return &Item{
		id:     id,
		writer: writer,
	}
}
func newValueFromAction(value string) Value {
	return Value(value)
}

type ItemStub struct {
	N string
	I string
	V string
}
type ValueStub struct {
	N string
	V string
	E bool
}
type AllStubs struct {
	Func string
	Items []ItemStub
	Values []ValueStub
}

func (i *Item) Click(handlerFunc interface{}, values interface{}) {
//	fmt.Fprint(wr, "hello!", "<br>")

//	itemsJson := ""
//	valuesJson := ""
	itemStubs := make([]ItemStub, 0)
	valueStubs := make([]ValueStub, 0)

//	fmt.Fprint(wr, reflect.ValueOf(values).Type().Name(), " --------<br>")
	val := reflect.ValueOf(values)
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		name := typ.Field(i).Name
		switch o := val.FieldByName(name).Interface().(type) {
		case *Item:
			itemStubs = append(itemStubs, ItemStub{N: name, I: o.FullId(), V: ""}) 
			
		//	if len(itemsJson) > 0 {
		//		itemsJson += ","
		//	}
		//	itemsJson += fmt.Sprint(`{"n":"` + name + `","i":"` + o.FullId() + `","v":""}`)
		case Value:
			
			valueStubs = append(valueStubs, ValueStub{N: name, V: string(o), E: false})
			
		//	if len(valuesJson) > 0 {
		//		valuesJson += ","
		//	}
		//	valuesJson += fmt.Sprint(`{"n":"` + name + `","v":"` + strings.Replace(strings.Replace(string(o), "\n", `\n`, -1), `"`, `\"`, -1) + `"}`)
		case ValueEncrypted:
			panic("TODO")
			//fmt.Fprint(wr, name, " (encrypted) ", o, "<br>")
		default:
			panic("Incorrect value " + name)
			//fmt.Fprint(wr, "failed... ", o, "<br>")
		}
	}
	//fmt.Fprint(wr, "Items: ", itemsJson, "<br>")
	//fmt.Fprint(wr, "Values: ", valuesJson, "<br>")
	
	stubs := AllStubs{Func: getFunctionName(handlerFunc), Items: itemStubs, Values: valueStubs}
	
	b, _ := json.Marshal(stubs)
	
//	fmt.Fprint(wr, string(b))
	
	//val := `{"items":[`+itemsJson+`],"values":[`+valuesJson+`]}`

	//	for i := 0; i < val.NumField(); i++ {
	//		fmt.Fprint(wr, val.Field(i), "<br>")
	//	}
	
	i.writer.Buffer += `
$("#` + i.FullId() + `").click(function(){var j = ` + string(b) + `; getValues(j.Items); $.post("/function", JSON.stringify(j), function(data){$("#head").append($("<div>").html(data))}, "html")});`

}
func getFunctionName(input interface{}) string {
	v := reflect.ValueOf(input)
	t := v.Type().In(0)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		if v.Pointer() == m.Func.Pointer() {
			return m.Name
		}
	}
	return "not found"
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
			i.templateGeneric(replace && j == 0, t)
		case string:
			i.htmlGeneric(replace && j == 0, t)
		case int:
			i.htmlGeneric(replace && j == 0, fmt.Sprint(t))
		case float32:
			i.htmlGeneric(replace && j == 0, fmt.Sprint(t))
		case float64:
			i.htmlGeneric(replace && j == 0, fmt.Sprint(t))
		case bool:
			i.htmlGeneric(replace && j == 0, fmt.Sprint(t))
		case Value:
			i.htmlGeneric(replace && j == 0, string(t))
		default:
			if t, ok := o1.(Templater); ok {
				i.templateGeneric(replace && j == 0, t.GetTemplate())
			} else {
				panic("Html / Append failed: " + reflect.TypeOf(o1).String())
			}
		}
	}
}
func (i *Item) htmlGeneric(replace bool, s string) {
	command := ""
	if replace {
		command = "html"
	} else {
		command = "append"
	}
	i.writer.Buffer += `
$("#` + i.FullId() + `").` + command + `("` + s + `");`
}

func (i *Item) templateGeneric(replace bool, t *Template) {
	t.parentId = i.FullId()
	command := ""
	if replace {
		command = "html"
	} else {
		command = "append"
	}
	i.writer.Buffer += `
$("#` + i.FullId() + `").` + command + `(template_` + t.name + `("` + t.FullId() + `"));`
}
func (i *Item) Attr(attrib string, o interface{}) {
	i.attrCss("attr", attrib, o)
}
func (i *Item) Css(attrib string, o interface{}) {
	i.attrCss("css", attrib, o)
}
func (i *Item) attrCss(command string, attrib string, o interface{}) {
	switch t := o.(type) {
	case string:
		i.attrCssGeneric(command, attrib, t)
	case int:
		i.attrCssGeneric(command, attrib, fmt.Sprint(t))
	case float32:
		i.attrCssGeneric(command, attrib, fmt.Sprint(t))
	case float64:
		i.attrCssGeneric(command, attrib, fmt.Sprint(t))
	case bool:
		i.attrCssGeneric(command, attrib, fmt.Sprint(t))
	}
}
func (i *Item) attrCssGeneric(command string, attrib string, val string) {
	i.writer.Buffer += `
$("#` + i.FullId() + `").` + command + `("` + attrib + `", "` + val + `");`
}


func (i *Item) FullId() string {

	if i.template == nil {
		return i.id
	} else {
		return i.template.FullId() + `_` + i.id
	}
	return ``
}
