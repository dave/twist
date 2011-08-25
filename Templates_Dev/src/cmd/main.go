package main

import (
	"fmt"
	"html"
	"os"
	"strings"
	"path/filepath"
	"template"
)

const HtmlFileSpec = "/Users/d.brophy/Projects/Templates/Templates_Web/src/pkg/html/"
const GeneatedFile = "/Users/d.brophy/Projects/Templates/Templates_Web/src/pkg/templates/generated.go"

type visitor struct {
	Templates *[]Template
}

func (v *visitor) AppendTemplate (t *Template) {
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
	
	temp := template.New(nil)
	temp.SetDelims("{{", "}}")
    err := temp.Parse(getTemplate())
    
    if err != nil {
    	fmt.Println("Error!", err.String())
    	return
    }
    
    f, _ := os.Create(GeneatedFile)
    defer f.Close()
    
    temp.Execute(f, binder)
    fmt.Println("Done!!")

}

func (v *visitor) VisitFile(path string, fi *os.FileInfo) {

	templateName := fi.Name[0:strings.Index(fi.Name, ".")]
	s := "" // main output
	f, _ := os.Open(path)
	t := html.NewTokenizer(f)
	items := []Item{}
//	u, _ := ioutil.ReadAll(f)
//	s := string(u)

	for {
		tt := t.Next()
		
		if tt == html.ErrorToken {
			if t.Error().String() == "EOF" {
				break
			}
		}
		token := t.Token()
		//fmt.Println(token.Type.String() + ": " + token.Data)
		switch token.Type {
			case html.TextToken :
				s1 := strings.Replace(token.Data, `"`, `\"`, -1)
				s2 := strings.Replace(s1, `
`, `\n`, -1)
				s += s2
			case html.StartTagToken, html.SelfClosingTagToken :
				att := token.Attr
				s += `<` + token.Data
				for _, v := range att {
					val := v.Val
					if v.Key == "id" {
						val = `"+id+"_` + v.Val
						items = append(items, Item{ItemNameLower:toLower(v.Val), ItemNameUpper:toUpper(v.Val)})
					}
					s += ` ` + v.Key + `=\"` + val + `\"`
				}
				if token.Type == html.StartTagToken {
					s += `>`
				} else {
					s += ` />`
				}
			case html.EndTagToken :
				s += `</` + token.Data + `>`
		}
		
		//if (t.Token().Type == 
		//fmt.Println(t.Token().Type.String() + " "  + t.Token().String())
		//fmt.Println(t.Token().String() + " " + t.Token().Type.String())
	}
	sOut := `<script>function template_`+templateName+`(id){return "` + s + `"}</script>`
	//fmt.Println(sOut)
//	for _, v := range items {
//		fmt.Println(strings.ToUpper(v[0:1]) + v[1:])
//	}
	
	temp := Template{
		NameUpper: toUpper(templateName),
		NameLower: toLower(templateName),
		Html: sOut,
		Items: items,
	}
	v.AppendTemplate(&temp)
}

func toUpper(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}
func toLower(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}
type Item struct {
	ItemNameUpper string
	ItemNameLower string
}
type Template struct {
	NameUpper string
	NameLower string
	Html string
	Items []Item
}

func getTemplate() string {
	return `
package templates

func GetTemplateByName(name string) *Template {
	switch name {
		{{.repeated section Templates}}
		case "{{NameLower}}" : 
			return {{NameLower}}_Template()
		{{.end}}
	}
	return nil
}

{{.repeated section Templates}}
type {{NameUpper}}_T struct {
	
	name string
	*Template
	{{.repeated section Items}}
	{{ItemNameUpper}} *Item
	{{.end}}

}
func {{NameLower}}_Template() *Template{
	return &Template {
		name : "{{NameLower}}",
		Html : `+"`"+`{{Html}}`+"`"+`,
	}
}
func (t *{{NameUpper}}_T) GetTemplate() *Template {
	return t.Template
}

func {{NameUpper}}(w *Writer, id string) *{{NameUpper}}_T {
	
	t := {{NameLower}}_Template()
	t.Writer = w
	t.Id = id
	
	w.RegisterTemplate(t.name)
	
	return &{{NameUpper}}_T{
		name : t.name, 
		Template : t,
		{{.repeated section Items}}
		{{ItemNameUpper}} : &Item{
			id : "{{ItemNameLower}}",
			template : t, 
			writer : w,
		},
		{{.end}}
	}
}
{{.end}}`
}
