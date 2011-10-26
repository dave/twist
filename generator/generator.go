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

const HtmlDir = "../web/html/"
const IndexFile = "../web/twist/generated.go"

type visitor struct {
	Templates map[string][]Template
}

func (v *visitor) AppendTemplate(namespace string, t Template) {
	if v.Templates[namespace] == nil {
		t := make([]Template, 0)
		v.Templates[namespace] = t
	}
	tem := append(v.Templates[namespace], t)
	v.Templates[namespace] = tem
}

func (v *visitor) VisitDir(path string, f *os.FileInfo) bool {
	//fmt.Println("DIR! ", path)
	return true
}

type Binder struct {
	Package   string
	Templates []Template
}
type IndexBinder struct {
	Packages  []Package
	Templates []Template
}

func main() {
	t := make(map[string][]Template)
	v := visitor{t}

	filepath.Walk(HtmlDir, &v, nil)

	indexBinder := IndexBinder{make([]Package, 0), make([]Template, 0)}

	for packagePath, templates := range v.Templates {

		filename := packagePath[strings.LastIndex(packagePath, "/")+1:]

		binder := Binder{filename, templates}

		temp := template.New("")

		p, err := temp.Parse(getTemplate())

		if err != nil {
			fmt.Println("Error!", err.String())
			return
		}
		htmlDirWithoutSlash := HtmlDir
		if strings.HasSuffix(htmlDirWithoutSlash, "/") {
			htmlDirWithoutSlash = htmlDirWithoutSlash[0 : len(htmlDirWithoutSlash)-1]
		}
		htmlDirWithoutBaseDir := htmlDirWithoutSlash[0:strings.LastIndex(htmlDirWithoutSlash, "/")]
		fullFilename := fmt.Sprint(htmlDirWithoutBaseDir, "/", packagePath, "/", filename, ".go")
		fmt.Println("fullFilename: ", fullFilename)
		f, _ := os.Create(fullFilename)
		defer f.Close()

		p.Execute(f, binder)

		indexBinder.Packages = append(indexBinder.Packages, Package{PackageName: filename, PackagePath: packagePath, PackagePathUnderscores: strings.Replace(packagePath, "/", "_", -1)})

		for _, template := range templates {
			indexBinder.Templates = append(indexBinder.Templates, template)
		}
	}

	tempIndex := template.New("")

	tempIndexParsed, err := tempIndex.Parse(getTemplateIndex())

	if err != nil {
		fmt.Println("Error!", err.String())
		return
	}
	fIndex, _ := os.Create(IndexFile)
	defer fIndex.Close()

	tempIndexParsed.Execute(fIndex, indexBinder)

	fmt.Println("Done!!")

}

func (v *visitor) VisitFile(path string, fi *os.FileInfo) {

	baseDir := HtmlDir
	if strings.HasSuffix(HtmlDir, "/") {
		baseDir = HtmlDir[0 : len(HtmlDir)-1]
	}
	baseDirName := baseDir[strings.LastIndex(baseDir, "/")+1:]
	packageAndFilename := path[len(baseDir)+1:]
	if !strings.Contains(packageAndFilename, "/") {
		return
	}
	packageName := packageAndFilename[0:strings.LastIndex(packageAndFilename, "/")]

	packagePath := baseDirName + "/" + packageName

	fmt.Print(packagePath, " ", fi.Name)
	if !strings.HasSuffix(fi.Name, ".html") {
		fmt.Println(" (skipping)")
		return
	} else {
		fmt.Println(" (processing)")
	}

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
		names += name + `, `
		sequence = newSequence
	}

	for i := 0; i < len(items); i++ {
		items[i].Variable = namesMap[items[i].Id]
	}

	temp := Template{
		PackageName:            packageName,
		PackagePath:            packagePath,
		PackagePathUnderscores: strings.Replace(packagePath, "/", "_", -1),
		NameUpper:              toUpper(templateName),
		NameLower:              toLower(templateName),
		Html:                   sOut,
		Items:                  items,
		Defs:                   defs,
		Names:                  names,
	}
	v.AppendTemplate(packagePath, temp)
}
func makeMap(data string, seperator1 string, seperator2 string) map[string]string {
	out := make(map[string]string)
	itemsA := strings.Split(data, seperator1)
	for _, item := range itemsA {
		if len(item) > 0 {
			if strings.Contains(item, seperator2) {
				index := strings.Index(item, seperator2)
				out[item[0:index]] = item[index+1:]
			}
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
	v`, sequence, ` := twist.NewTextItem(`, "`", v1, "`", `)`)
	return sequence, fmt.Sprint(`v`, sequence), s
}
func (he *Tag) Definition(sequence int, names map[string]string) (int, string, string) {
	sequence++

	if len(he.Id) > 0 {
		names[he.Id] = fmt.Sprint(`v`, sequence)
	}

	newSequence := sequence
	s := ``

	if len(he.Id) > 0 {
		s += fmt.Sprint(`
	v`, sequence, ` := twist.NewItemId("`, he.Name, `", `)
		s += fmt.Sprint(`t, `)
		s += fmt.Sprint(`c.Writer, `)
		s += fmt.Sprint(`"`, he.Id, `", `)
	} else {
		s += fmt.Sprint(`
	v`, sequence, ` := twist.NewItem("`, he.Name, `", `)
	}

	if len(he.Attributes) > 0 {
		s += `map[string]string{`
		for k, v := range he.Attributes {
			s += fmt.Sprint(`"`, k, `":"`, v, `", `)
		}
		s += `}, `
	} else {
		s += `map[string]string{}, `
	}
	if len(he.Styles) > 0 {
		s += `map[string]string{`
		for k, v := range he.Styles {
			s += fmt.Sprint(`"`, k, `":"`, v, `", `)
		}
		s += `}, `
	} else {
		s += `map[string]string{}, `
	}
	if len(he.Contents) > 0 {
		s += `[]*twist.Item{`
		for _, c := range he.Contents {
			newSequence1, name, def := c.Definition(newSequence, names)
			s += name + `, `
			s = def + s
			newSequence = newSequence1
		}
		s += `}, `
	} else {
		s += `[]*twist.Item{}, `
	}
	s += `)`
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
type Package struct {
	PackageName            string
	PackagePath            string
	PackagePathUnderscores string
}
type Template struct {
	PackageName            string
	PackagePath            string
	PackagePathUnderscores string
	NameUpper              string
	NameLower              string
	Html                   string
	Items                  []Item
	Defs                   string
	Names                  string
}

func getTemplate() string {
	return `
package {{.Package}}

import (
	"twist"
)

{{range .Templates}}
type {{.NameUpper}}_T struct {
	
	name string
	*twist.Template
	{{range .Items}}
	{{.ItemNameUpper}} *twist.Item
	{{end}}

}

func (t *{{.NameUpper}}_T) GetTemplate() *twist.Template {
	return t.Template
}

func {{.NameUpper}}(c *twist.Context, id string) *{{.NameUpper}}_T {
	
	t := twist.GetTemplateByPath("{{.PackagePathUnderscores}}_{{.NameLower}}")
	t.Path = "{{.PackagePathUnderscores}}_{{.NameLower}}"
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(t)

	{{.Defs}}

	t.Contents = []*twist.Item{ {{.Names}} }
	
	return &{{.NameUpper}}_T{
		name : t.Name, 
		Template : t,
		{{range .Items}}
		{{.ItemNameUpper}} : {{.Variable}},
		{{end}}
	}
}
{{end}}`
}

func getTemplateIndex() string {
	return `
package twist

func GetTemplateByPath(path string) *Template {
	switch path {
		{{range .Templates}}
		case "{{.PackagePathUnderscores}}_{{.NameLower}}" : 
			return {{.PackagePathUnderscores}}_{{.NameLower}}_Template()
		{{end}}
	}
	return nil
}

{{range .Templates}}
func {{.PackagePathUnderscores}}_{{.NameLower}}_Template() *Template{
	return &Template {
		Name     : "{{.NameLower}}",
		Html     : ` + "`" + `{{.Html}}` + "`" + `,
	}
}
{{end}}
`
}
