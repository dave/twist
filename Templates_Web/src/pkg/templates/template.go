package templates


import (
//	"strings"
	//"writer"
)
type Template struct {
	name string
	Html string
	Id string
	parentId string
	Writer *Writer
}

func (t *Template) GetTemplateJavascript() string {
//	s := strings.Replace(t.Html, `"`, `\"`, -1)
//	s1 := strings.Replace(s, ` id=\"`, ` id=\""+id+"_`, -1)
//	s2 := strings.Replace(s1, `
//`, `\n`, -1)
//	s3 := `<script>function template_`+t.name+`(id){return "` + s2 + `"}</script>`
	return t.Html
}
func (t *Template) FullId() string {
	return t.parentId + `_` + t.Id
}

func (t *Template) GetParentId() string {
	return t.parentId
}

type Templater interface {
	GetTemplate() *Template
}