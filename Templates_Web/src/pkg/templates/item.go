package templates

import (
	"fmt"
	"reflect"
	"json"
	"crypto/md5"
	"encoding/hex"
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
type allStubs struct {
	Func string
	Hash string
	Items []ItemStub
	Values []ValueStub
}

func (i *Item) Click(handlerFunc interface{}, values interface{}) {

	itemStubs := make([]ItemStub, 0)
	valueStubs := make([]ValueStub, 0)

	val := reflect.ValueOf(values)
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		name := typ.Field(i).Name
		switch o := val.FieldByName(name).Interface().(type) {
		case *Item:
			itemStubs = append(itemStubs, ItemStub{N: name, I: o.FullId(), V: ""})
		case Value:
			valueStubs = append(valueStubs, ValueStub{N: name, V: string(o), E: false})
		case ValueEncrypted:
			panic("TODO")
		default:
			panic("Incorrect value " + name)
		}
	}
	
	stubs := allStubs{Func: getFunctionName(handlerFunc), Items: itemStubs, Values: valueStubs}
	
	hash := getHash(stubs)
    stubs.Hash = hash
	
	marshalled, _ := json.Marshal(stubs)
	
	i.writer.Buffer += `
$("#` + i.FullId() + `").click(function(){var j = ` + string(marshalled) + `; getValues(j.Items); $.post("/function", JSON.stringify(j), function(data){$("#head").append($("<div>").html(data))}, "html")});`

}
func getHash(stubs allStubs) string {
	
	b, _ := json.Marshal(stubs)
	h := md5.New()
	h.Write([]byte(b))
	
	return hex.EncodeToString([]byte(h.Sum()))
	
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
