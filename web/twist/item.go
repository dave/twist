package twist

import (
	"fmt"
	"reflect"
	"json"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"url"
)

type String string
type StringHashed string
type StringEncrypted string

type Int int
type IntHashed int
type IntEncrypted int

type Item struct {
	id         string
	template   *Template
	writer     *Writer
	value      string
	name       string
	Attributes map[string]string
	Styles     map[string]string
	Contents   []*Item
	Text       string
	commands   []func()
}

func NewItemId(name string, template *Template, writer *Writer, id string, attributes map[string]string, styles map[string]string, contents []*Item) *Item {
	i := Item{
		name:       name,
		template:   template,
		writer:     writer,
		id:         id,
		Attributes: attributes,
		Styles:     styles,
		Contents:   contents,
	}
	return &i
}
func NewItem(name string, attributes map[string]string, styles map[string]string, contents []*Item) *Item {
	i := Item{
		name:       name,
		Attributes: attributes,
		Styles:     styles,
		Contents:   contents,
	}
	return &i
}
func NewTextItem(text string) *Item {
	i := Item{
		Text: text,
	}
	return &i
}

func (i *Item) Value() string {
	return i.value
}
func (i *Item) Int() int {
	v, _ := strconv.Atoi(i.value)
	return v
}
func (i *Item) String() string {
	return i.Value()
}

func (he *Item) RunCommands() {

	for _, command := range he.commands {
		command()
	}

	for _, inner := range he.Contents {
		inner.RunCommands()
	}

}

func (he *Item) RenderHtml() string {

	s := ``
	if len(he.name) > 0 {
		s += `<`
		s += he.name
		if len(he.id) > 0 {
			s += fmt.Sprint(` id="`, he.fullId(), `"`)
		}
		for attKey, attVal := range he.Attributes {
			s += fmt.Sprint(` `, attKey, `="`, attVal, `"`)
		}
		if len(he.Styles) > 0 {
			s += ` style="`
			for k, v := range he.Styles {
				s += fmt.Sprint(k, `:`, v, `;`)
			}
			s += `"`
		}
		if len(he.Contents) == 0 {
			s += ` />`
		} else {
			s += `>`
			for _, inner := range he.Contents {
				s += inner.RenderHtml()
			}
			s += fmt.Sprint(`</`, he.name, `>`)
		}
	} else if len(he.Text) > 0 {
		s += he.Text
	} else {
		for _, inner := range he.Contents {
			s += inner.RenderHtml()
		}
	}
	return s
}

func (v String) Value() string           { return string(v) }
func (v StringHashed) Value() string     { return string(v) }
func (v StringEncrypted) Value() string  { return string(v) }
func (v String) String() string          { return string(v) }
func (v StringHashed) String() string    { return string(v) }
func (v StringEncrypted) String() string { return string(v) }
func (v Int) Value() int                 { return int(v) }
func (v IntHashed) Value() int           { return int(v) }
func (v IntEncrypted) Value() int        { return int(v) }
func (v Int) String() string             { return strconv.Itoa(int(v)) }
func (v IntHashed) String() string       { return strconv.Itoa(int(v)) }
func (v IntEncrypted) String() string    { return strconv.Itoa(int(v)) }
func (v String) Int() int                { i, _ := strconv.Atoi(string(v)); return i }
func (v StringHashed) Int() int          { i, _ := strconv.Atoi(string(v)); return i }
func (v StringEncrypted) Int() int       { i, _ := strconv.Atoi(string(v)); return i }

func newItemFromAction(id string, writer *Writer) *Item {
	return &Item{
		id:         id,
		writer:     writer,
		Attributes: make(map[string]string),
		Styles:     make(map[string]string),
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
	Func   string
	Hash   string
	Items  []itemStub
	Values []valueStub
}

func (i *Item) Click(handlerFunc interface{}, values interface{}) {
	i.commands = append(i.commands, func() { i.clickAtRender(handlerFunc, values) })
}
func (i *Item) clickAtRender(handlerFunc interface{}, values interface{}) {

	valueStubs, itemStubs, _ := makeStubs(values, true)

	stubs := allStubs{Func: getFunctionName(handlerFunc), Items: itemStubs, Values: valueStubs}

	hash := getHash(stubs)
	stubs.Hash = hash

	marshalled, _ := json.Marshal(stubs)

	i.writer.Buffer += `
$("#` + i.fullId() + `").unbind('click');
$("#` + i.fullId() + `").click(function(){var j = ` + string(marshalled) + `; getValues(j.Items); $.post("/function", JSON.stringify(j), function(data){("#head").append($("<div>").html(data))}, "html");return false;});`

}
func (i *Item) getLinkUrl(handlerFunc interface{}, values interface{}) string {

	valueStubs, _, needsHash := makeStubs(values, false)

	stubs := allStubs{Func: getFunctionName(handlerFunc), Values: valueStubs}

	hashQuery := ""
	if needsHash {
		hash := getHash(stubs)
		stubs.Hash = hash
		hashQuery = "&_hash=" + hash
	}

	qstring := ""
	for _, v := range stubs.Values {
		if len(qstring) == 0 {
			qstring += "?"
		} else {
			qstring += "&"
		}

		qstring += v.N + "=" + url.QueryEscape(toString(v.V))
	}
	href := "/" + stubs.Func + qstring + hashQuery

	return href

}
func (i *Item) Link(handlerFunc interface{}, values interface{}) {

	href := i.getLinkUrl(handlerFunc, values)

	i.Attributes["href"] = href

	i.commands = append(i.commands, func() { i.linkAtRender(href) })

}
func (i *Item) linkAtRender(href string) {

	if !i.writer.SendHtml {
		i.writer.Buffer += `
$("#` + i.fullId() + `").attr("href", "` + href + `");`
	}

	i.writer.Buffer += `
$("#` + i.fullId() + `").unbind('click');
$("#` + i.fullId() + `").click(function(){History.pushState(null, null, "` + href + `");return false;});`

}

func makeStubs(values interface{}, isClick bool) (valueStubs []valueStub, itemStubs []itemStub, needsHash bool) {

	valueStubs = make([]valueStub, 0)
	itemStubs = make([]itemStub, 0)
	needsHash = false

	if values != nil {
		val := reflect.ValueOf(values)
		typ := val.Type()
		for i := 0; i < typ.NumField(); i++ {
			name := typ.Field(i).Name
			switch o := val.FieldByName(name).Interface().(type) {
			case String:
				valueStubs = append(valueStubs, valueStub{N: name, V: string(o), T: 1})
			case StringHashed:
				needsHash = true
				valueStubs = append(valueStubs, valueStub{N: name, V: string(o), T: 2})
			case StringEncrypted:
				needsHash = true //???
				panic("TODO")
			case Int:
				valueStubs = append(valueStubs, valueStub{N: name, V: int(o), T: 4})
			case IntHashed:
				needsHash = true
				valueStubs = append(valueStubs, valueStub{N: name, V: int(o), T: 5})
			case IntEncrypted:
				needsHash = true //???
				panic("TODO")
			case *Item:
				needsHash = true
				if isClick {
					itemStubs = append(itemStubs, itemStub{N: name, I: o.fullId(), V: ""})
				} else {
					panic("We can't have Items in a Link - name:" + name)
				}
			default:
				panic("Incorrect value " + name)
			}
		}
	}
	return valueStubs, itemStubs, needsHash
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
		if stubs.Values[i].T == 1 {
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
			if m.Name == "Root" {
				return ""
			} else {
				return m.Name
			}
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

	if replace {
		i.Contents = make([]*Item, 0)
	}
	i.Contents = append(i.Contents, &Item{Text: s, Attributes: make(map[string]string), Styles: make(map[string]string)})

	i.commands = append(i.commands, func() { i.htmlGenericAtRender(replace, s) })

}
func (i *Item) htmlGenericAtRender(replace bool, s string) {

	if !i.writer.SendHtml {
		command := ""
		if replace {
			command = "html"
		} else {
			command = "append"
		}
		i.writer.Buffer += `
$("#` + i.fullId() + `").` + command + `("` + s + `");`
	}

}

func (i *Item) templateGeneric(replace bool, t *Template) {

	if replace {
		i.Contents = make([]*Item, 0)
	}
	for _, item := range t.Contents {
		i.Contents = append(i.Contents, item)
	}

	i.commands = append(i.commands, func() { i.templateGenericAtRender(replace, t) })

}
func (i *Item) templateGenericAtRender(replace bool, t *Template) {

	t.parentId = i.fullId()

	if !i.writer.SendHtml {
		command := ""
		if replace {
			command = "html"
		} else {
			command = "append"
		}
		i.writer.Buffer += `
$("#` + i.fullId() + `").` + command + `(template_` + t.Name + `("` + t.fullId() + `"));`
	}

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

	if command == "attr" {
		i.Attributes[attrib] = val
	} else {
		i.Styles[attrib] = val
	}

	i.commands = append(i.commands, func() { i.attrCssGenericAtRender(command, attrib, val) })

}
func (i *Item) attrCssGenericAtRender(command string, attrib string, val string) {
	if !i.writer.SendHtml {
		i.writer.Buffer += `
$("#` + i.fullId() + `").` + command + `("` + attrib + `", "` + val + `");`
	}
}

func (i *Item) fullId() string {

	if len(i.id) == 0 {
		return ``
	} else if i.template == nil {
		return i.id
	} else {
		return i.template.fullId() + `_` + i.id
	}
	return ``
}
