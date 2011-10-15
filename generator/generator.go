package main

import (
	"fmt"
	"html"
	"os"
	"strings"
	"path/filepath"
	"template"
	//"bytes"
)

const HtmlFileSpec = "../web/html/"
const GeneatedFile = "../web/twist/generated.go"

type visitor struct {
	Templates *[]Template
}

func (v *visitor) AppendTemplate(t *Template) {
	tem := append(*v.Templates, *t)
	v.Templates = &tem
}

func (v *visitor) VisitDir(path string, f *os.FileInfo) bool {
	return true
}

type Binder struct {
	Templates []Template
}

func main() {
	t := []Template{}
	v := visitor{&t}
	filepath.Walk(HtmlFileSpec, &v, nil)
	binder := Binder{*v.Templates}

	temp := template.New("")
	//temp.SetDelims("{{", "}}")

	p, err := temp.Parse(getTemplate())

	if err != nil {
		fmt.Println("Error!", err.String())
		return
	}
	//b := new(bytes.Buffer)
	f, _ := os.Create(GeneatedFile)
	defer f.Close()

	p.Execute(f, binder)

	fmt.Println("Done!!")

}

func (v *visitor) VisitFile(path string, fi *os.FileInfo) {

	templateName := fi.Name[0:strings.Index(fi.Name, ".")]
	s := ""
	f, _ := os.Open(path)
	t := html.NewTokenizer(f)
	items := []Item{}
	root := Tag{
		Contents: make([]Element, 0),
	}
	currentElement := &root

	for {
		tt := t.Next()

		if tt == html.ErrorToken {
			if t.Error().String() == "EOF" {
				break
			}
		}
		token := t.Token()
		switch token.Type {
		case html.TextToken:
			currentElement.Contents = append(currentElement.Contents, &Text{
				Text:   token.Data,
				parent: currentElement,
			})
			s1 := strings.Replace(token.Data, `"`, `\"`, -1)
			s2 := strings.Replace(s1, `
`, `\n`, -1)
			s += s2
		case html.StartTagToken, html.SelfClosingTagToken:
			tagPlainText := ``
			att := token.Attr
			tagId := ``
			tagPlainText += `<` + token.Data
			tagAttributes := make(map[string]string)
			tagStyles := make(map[string]string)

			for _, v := range att {
				val := v.Val
				if strings.ToLower(v.Key) == "id" {
					val = `"+id+"_` + v.Val
					newItem := Item{
						Id:            v.Val,
						ItemNameLower: toLower(v.Val),
						ItemNameUpper: toUpper(v.Val),
					}
					items = append(items, newItem)
					tagId = v.Val
				} else if strings.ToLower(v.Key) == `style` {
					tagStyles = makeMap(v.Val, `;`, `:`)
				} else {
					tagAttributes[v.Key] = v.Val
				}
				tagPlainText += ` ` + v.Key + `=\"` + val + `\"`
			}

			if token.Type == html.StartTagToken {
				tagPlainText += `>`
			} else {
				tagPlainText += ` />`
			}
			s += tagPlainText

			newTag := &Tag{
				Id:         tagId,
				Name:       token.Data,
				Attributes: tagAttributes,
				Styles:     tagStyles,
				Contents:   make([]Element, 0),
				parent:     currentElement,
			}
			currentElement.Contents = append(currentElement.Contents, newTag)
			if token.Type == html.StartTagToken {
				currentElement = newTag
			}

		case html.EndTagToken:
			tagPlainText := fmt.Sprint(`</`, token.Data, `>`)
			s += tagPlainText
			currentElement = currentElement.Parent()
		}
	}
	sOut := `<script>function template_` + templateName + `(id){return "` + s + `"}</script>`

	defs := ``
	names := ``
	sequence := 0
	namesMap := make(map[string]string)
	for _, v := range root.Contents {
		newSequence, name, def := v.Definition(sequence, namesMap)
		defs += def
		names += `&` + name + `, `
		sequence = newSequence
	}

	for i := 0; i < len(items); i++ {
		items[i].Variable = namesMap[items[i].Id]
	}

	temp := Template{
		NameUpper: toUpper(templateName),
		NameLower: toLower(templateName),
		Html:      sOut,
		Items:     items,
		Defs:      defs,
		Names:     names,
	}
	v.AppendTemplate(&temp)
}
func makeMap(data string, seperator1 string, seperator2 string) map[string]string {
	out := make(map[string]string)
	itemsA := strings.Split(data, seperator1)
	for _, item := range itemsA {
		if len(item) > 0 {
			keyVal := strings.Split(item, seperator2)
			out[keyVal[0]] = keyVal[1]
		}
	}
	return out
}
func toUpper(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}
func toLower(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}

type Text struct {
	Text   string
	parent *Tag
}
type Tag struct {
	Id         string
	Name       string
	Attributes map[string]string
	Styles     map[string]string
	Contents   []Element
	parent     *Tag
}
type Element interface {
	Parent() *Tag
	Definition(int, map[string]string) (int, string, string)
}

func (pte *Text) Definition(sequence int, names map[string]string) (int, string, string) {
	sequence++
	v := strings.Replace(pte.Text, "`", "`+\"`\"+`", -1)
	v1 := strings.Replace(v, "\n", "`+\"\\n\"+`", -1)
	s := fmt.Sprint(`
	v`, sequence, ` := Item{Text:`, "`", v1, "`", `}`)
	return sequence, fmt.Sprint(`v`, sequence), s
}
func (he *Tag) Definition(sequence int, names map[string]string) (int, string, string) {
	sequence++

	if len(he.Id) > 0 {
		names[he.Id] = fmt.Sprint(`v`, sequence)
	}

	newSequence := sequence
	s := fmt.Sprint(`
	v`, sequence, ` := Item{Name: "`, he.Name, `", `)

	if len(he.Id) > 0 {
		s += fmt.Sprint(`template: t, `)
		s += fmt.Sprint(`writer: c.Writer, `)
		s += fmt.Sprint(`id: "`, he.Id, `", `)
	}

	if len(he.Attributes) > 0 {
		s += `Attributes: map[string]string{`
		for k, v := range he.Attributes {
			s += fmt.Sprint(`"`, k, `":"`, v, `", `)
		}
		s += `}, `
	} else {
		s += `Attributes: map[string]string{}, `
	}
	if len(he.Styles) > 0 {
		s += `Styles: map[string]string{`
		for k, v := range he.Styles {
			s += fmt.Sprint(`"`, k, `":"`, v, `", `)
		}
		s += `}, `
	} else {
		s += `Styles: map[string]string{}, `
	}
	if len(he.Contents) > 0 {
		s += `Contents: []*Item{`
		for _, c := range he.Contents {
			newSequence1, name, def := c.Definition(newSequence, names)
			s += `&` + name + `, `
			s = def + s
			newSequence = newSequence1
		}
		s += `}, `
	}
	s += `}`
	return newSequence, fmt.Sprint(`v`, sequence), s
}

func (pte *Text) Parent() *Tag {
	return pte.parent
}
func (he *Tag) Parent() *Tag {
	return he.parent
}

type Item struct {
	Id            string
	ItemNameUpper string
	ItemNameLower string
	Variable      string
}
type Template struct {
	NameUpper string
	NameLower string
	Html      string
	Items     []Item
	Defs      string
	Names     string
}

func getTemplate() string {
	return `
package twist

func getTemplateByName(name string) *Template {
	switch name {
		{{range .Templates}}
		case "{{.NameLower}}" : 
			return {{.NameLower}}_Template()
		{{end}}
	}
	return nil
}

{{range .Templates}}
type {{.NameUpper}}_T struct {
	
	name string
	*Template
	{{range .Items}}
	{{.ItemNameUpper}} *Item
	{{end}}

}
func {{.NameLower}}_Template() *Template{
	return &Template {
		name     : "{{.NameLower}}",
		Html     : ` + "`" + `{{.Html}}` + "`" + `,
	}
}
func (t *{{.NameUpper}}_T) GetTemplate() *Template {
	return t.Template
}

func {{.NameUpper}}(c *Context, id string) *{{.NameUpper}}_T {
	
	t := {{.NameLower}}_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(*t)

	{{.Defs}}

	t.Contents = []*Item{ {{.Names}} }
	
	return &{{.NameUpper}}_T{
		name : t.name, 
		Template : t,
		{{range .Items}}
		{{.ItemNameUpper}} : &{{.Variable}},
		{{end}}
	}
}
{{end}}`
}
