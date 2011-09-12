package twist

import (
	"fmt"
	"reflect"
	"json"
	"crypto/md5"
	"encoding/hex"
	"http"
	"strconv"
)

type String string
type StringHashed string
type StringEncrypted string

type Int int
type IntHashed int
type IntEncrypted int

type Item struct {
	id       string
	template *Template
	writer   *Writer
	Value    string
}

func (v String) Value() string {
	return string(v)
}
func (v StringHashed) Value() string {
	return string(v)
}
func (v StringEncrypted) Value() string {
	return string(v)
}
func (v String) String() string {
	return string(v)
}
func (v StringHashed) String() string {
	return string(v)
}
func (v StringEncrypted) String() string {
	return string(v)
}
func (v Int) Value() int {
	return int(v)
}
func (v IntHashed) Value() int {
	return int(v)
}
func (v IntEncrypted) Value() int {
	return int(v)
}
func (v Int) String() string {
	return strconv.Itoa(int(v))
}
func (v IntHashed) String() string {
	return strconv.Itoa(int(v))
}
func (v IntEncrypted) String() string {
	return strconv.Itoa(int(v))
}

func newItemFromAction(id string, writer *Writer) *Item {
	return &Item{
		id:     id,
		writer: writer,
	}
}

type itemStub struct {
	N string
	I string
	V string
}
type valueStub struct {
	N string
	V interface{}
	T int
}
type allStubs struct {
	Func string
	Hash string
	Items []itemStub
	Values []valueStub
}


func (i *Item) Click(handlerFunc interface{}, values interface{}) {

	valueStubs, itemStubs, _ := makeStubs(values, true)
/*
	val := reflect.ValueOf(values)
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		name := typ.Field(i).Name
		switch o := val.FieldByName(name).Interface().(type) {
		case *Item:
			itemStubs = append(itemStubs, itemStub{N: name, I: o.FullId(), V: ""})
		case String:
			valueStubs = append(valueStubs, valueStub{N: name, V: string(o), T:1})
		case StringHashed:
			valueStubs = append(valueStubs, valueStub{N: name, V: string(o), T:2})
		case StringEncrypted:
			panic("TODO")
		case Int:
			valueStubs = append(valueStubs, valueStub{N: name, V: int(o), T:4})
		case IntHashed:
			valueStubs = append(valueStubs, valueStub{N: name, V: int(o), T:5})
		case IntEncrypted:
			panic("TODO")
		default:
			panic("Incorrect value " + name)
		}
	}
*/
	
	stubs := allStubs{Func: getFunctionName(handlerFunc), Items: itemStubs, Values: valueStubs}
	
	hash := getHash(stubs)
    stubs.Hash = hash
	
	marshalled, _ := json.Marshal(stubs)
	
	i.writer.Buffer += `
$("#` + i.FullId() + `").click(function(){var j = ` + string(marshalled) + `; getValues(j.Items); $.post("/function", JSON.stringify(j), function(data){$("#head").append($("<div>").html(data))}, "html");return false;});`

}

func makeStubs(values interface{}, isClick bool) (valueStubs []valueStub, itemStubs []itemStub, needsHash bool) {
	
	valueStubs = make([]valueStub, 0)
	itemStubs = make([]itemStub, 0)
	needsHash = false

	val := reflect.ValueOf(values)
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		name := typ.Field(i).Name
		switch o := val.FieldByName(name).Interface().(type) {
		case String:
			valueStubs = append(valueStubs, valueStub{N: name, V: string(o), T:1})
		case StringHashed:
			needsHash = true
			valueStubs = append(valueStubs, valueStub{N: name, V: string(o), T:2})
		case StringEncrypted:
			needsHash = true //???
			panic("TODO")
		case Int:
			valueStubs = append(valueStubs, valueStub{N: name, V: int(o), T:4})
		case IntHashed:
			needsHash = true
			valueStubs = append(valueStubs, valueStub{N: name, V: int(o), T:5})
		case IntEncrypted:
			needsHash = true //???
			panic("TODO")
		case *Item:
			needsHash = true
			if isClick {
				itemStubs = append(itemStubs, itemStub{N: name, I: o.FullId(), V: ""})
			} else {
				panic("We can't have Items in a Link - name:" + name)
			}
		default:
			panic("Incorrect value " + name)
		}
	}
	return valueStubs, itemStubs, needsHash
}

func (i *Item) Link(handlerFunc interface{}, values interface{}) {

	valueStubs, _, needsHash := makeStubs(values, false)
	/*
	val := reflect.ValueOf(values)
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		name := typ.Field(i).Name
		switch o := val.FieldByName(name).Interface().(type) {
		case String:
			valueStubs = append(valueStubs, valueStub{N: name, V: string(o), T:1})
		case StringHashed:
			needsHash = true
			valueStubs = append(valueStubs, valueStub{N: name, V: string(o), T:2})
		case StringEncrypted:
			needsHash = true //???
			panic("TODO")
		case Int:
			valueStubs = append(valueStubs, valueStub{N: name, V: int(o), T:4})
		case IntHashed:
			needsHash = true
			valueStubs = append(valueStubs, valueStub{N: name, V: int(o), T:5})
		case IntEncrypted:
			needsHash = true //???
			panic("TODO")
		case *Item:
			panic("We can't have Items in a Link - name:" + name)
		default:
			panic("Incorrect value " + name)
		}
	}
	*/
	stubs := allStubs{Func: getFunctionName(handlerFunc), Values: valueStubs}
	
	hashQuery := ""
	if needsHash {
		hash := getHash(stubs)
    	stubs.Hash = hash
    	hashQuery = "&_hash=" + hash
	}

	marshalled, _ := json.Marshal(stubs)
	
	qstring := ""
	for _, v := range stubs.Values {
		if len(qstring) == 0 {
			qstring += "?"
		} else {
			qstring += "&"
		}
		
		qstring += v.N + "=" + http.URLEscape(toString(v.V))
	}
	href := "/" + stubs.Func + qstring + hashQuery

	i.writer.Buffer += `
$("#` + i.FullId() + `").attr("href", "` + href + `");
$("#` + i.FullId() + `").click(function(){var j = ` + string(marshalled) + `; getValues(j.Items); $.post("/function", JSON.stringify(j), function(data){$("#head").append($("<div>").html(data))}, "html");return false;});`

}


//We use this when we have to clear values out in the hash function
func (self *allStubs) Copy() *allStubs { 

	r := new(allStubs)
	*r = *self
	r.Items = make([]itemStub, len(r.Items))
	r.Values = make([]valueStub, len(r.Values))
	copy(r.Items, self.Items)
	copy(r.Values, self.Values)
	return r

}

func getHash(stubs1 allStubs) string {

	var stubs allStubs = *stubs1.Copy()

	//clear variant data out of stubs before checking hash
	stubs.Hash = ""
	for i, _ := range stubs.Items {
		stubs.Items[i].V = ""
	}

	//clear data out of non-hashed items
	for i, _ := range stubs.Values {
		if (stubs.Values[i].T == 1) {
			stubs.Values[i].V = ""
		} else if stubs.Values[i].T == 4 {
			stubs.Values[i].V = 0
		}
	}

	stubs.Hash = "oiheworkvnxcvwetrytknmxznuihkfnknvkcskjnsjdnanjvdskjsvnmzxbc" // this works as a crude salt.
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
		case String, StringHashed, StringEncrypted:
			i.htmlGeneric(replace && j == 0, toString(t))
		case Int, IntHashed, IntEncrypted:
			i.htmlGeneric(replace && j == 0, fmt.Sprint(toInt(t)))
		default:
			if t, ok := o1.(Templater); ok {
				i.templateGeneric(replace && j == 0, t.GetTemplate())
			} else {
				panic("Html / Append failed: " + reflect.TypeOf(o1).String())
			}
		}
	}
}
func toString(input interface{}) string {
	s := ""
	if st, ok := input.(string); ok {
		s = st
	} else if in, ok := input.(int); ok {
		s = strconv.Itoa(in)
	}
	return s
}
func toInt(input interface{}) int {
	i := 0
	if in, ok := input.(int); ok {
		i = in
	} else if fl, ok := input.(float64); ok {
		in := int(fl)
		i = in
	} else if st, ok := input.(string); ok {
		in, _ := strconv.Atoi(st)
		i = in
	}
	return i
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
