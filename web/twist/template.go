package twist

import (
//"fmt"
)

type Template struct {
	Name     string
	Path     string
	Html     string
	Id       string
	parentId string
	Writer   *Writer
	Contents []*Item
}

func (t *Template) getTemplateJavascript() string {
	return t.Html
}

func (t *Template) fullId() string {
	return t.parentId + `_` + t.Id
}

func (t *Template) getParentId() string {
	return t.parentId
}

type Templater interface {
	GetTemplate() *Template
}
