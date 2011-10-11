package twist

import (
//"fmt"
)

type Template struct {
	name     string
	Html     string
	Id       string
	parentId string
	Writer   *Writer
	Contents []*Item
}

func (t *Template) GetTemplateJavascript() string {
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

/*
func (t *Template) GetPlainHtml() string {
	s := ``
	for _, e := range t.Contents {
		s += e.RenderHtml()
	}
	return s
}
*/
