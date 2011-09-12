package twist

type Template struct {
	name     string
	Html     string
	Id       string
	parentId string
	Writer   *Writer
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
