package twist

func getTemplateByName(name string) *Template {
	switch name {
	case "inner":
		return inner_Template()
	case "test":
		return test_Template()
	}
	return nil
}

type Inner_T struct {
	name string
	*Template
	Span1  *Item
	Img1   *Item
	MyLink *Item
}

func inner_Template() *Template {
	return &Template{
		name: "inner",
		Html: `<script>function template_inner(id){return "<b>This is an inner template, here's some text: <span style=\"border:1px solid #000000;\" id=\""+id+"_span1\" /></b><br />\n...and here's an image: <img id=\""+id+"_img1\" border=\"0\" style=\"border:2px solid #00ff00;\" /><br />\n<a href=\"/\" id=\""+id+"_myLink\">This is a link</a>"}</script>`,
	}
}
func (t *Inner_T) GetTemplate() *Template {
	return t.Template
}

func Inner(w *Writer, id string) *Inner_T {

	t := inner_Template()
	t.Writer = w
	t.Id = id

	w.RegisterTemplate(*t)

	v3 := Item{Name: "span", template: t, writer: w, id: "span1", Styles: map[string]string{"border": "1px solid #000000"}}
	v2 := Item{Text: `This is an inner template, here's some text: `}
	v1 := Item{Name: "b", Contents: []*Item{&v2, &v3}}
	v4 := Item{Name: "br"}
	v5 := Item{Text: `` + "\n" + `...and here's an image: `}
	v6 := Item{Name: "img", template: t, writer: w, id: "img1", Attributes: map[string]string{"border": "0"}, Styles: map[string]string{"border": "2px solid #00ff00"}}
	v7 := Item{Name: "br"}
	v8 := Item{Text: `` + "\n" + ``}
	v10 := Item{Text: `This is a link`}
	v9 := Item{Name: "a", template: t, writer: w, id: "myLink", Attributes: map[string]string{"href": "/"}, Contents: []*Item{&v10}}

	t.Contents = []*Item{&v1, &v4, &v5, &v6, &v7, &v8, &v9}

	return &Inner_T{
		name:     t.name,
		Template: t,
		Span1:    &v3,
		Img1:     &v6,
		MyLink:   &v9,
	}
}

type Test_T struct {
	name string
	*Template
	Span1   *Item
	Text1   *Item
	Para1   *Item
	Button1 *Item
}

func test_Template() *Template {
	return &Template{
		name: "test",
		Html: `<script>function template_test(id){return "<p>TEsting... <span id=\""+id+"_span1\">span1</span></p>\n<p><input id=\""+id+"_text1\" type=\"text\" /></p>\n<p></p>\n<p id=\""+id+"_para1\"></p>\n<p><button id=\""+id+"_button1\" /></p>"}</script>`,
	}
}
func (t *Test_T) GetTemplate() *Template {
	return t.Template
}

func Test(w *Writer, id string) *Test_T {

	t := test_Template()
	t.Writer = w
	t.Id = id

	w.RegisterTemplate(*t)

	v4 := Item{Text: `span1`}
	v3 := Item{Name: "span", template: t, writer: w, id: "span1", Contents: []*Item{&v4}}
	v2 := Item{Text: `TEsting... `}
	v1 := Item{Name: "p", Contents: []*Item{&v2, &v3}}
	v5 := Item{Text: `` + "\n" + ``}
	v7 := Item{Name: "input", template: t, writer: w, id: "text1", Attributes: map[string]string{"type": "text"}}
	v6 := Item{Name: "p", Contents: []*Item{&v7}}
	v8 := Item{Text: `` + "\n" + ``}
	v9 := Item{Name: "p"}
	v10 := Item{Text: `` + "\n" + ``}
	v11 := Item{Name: "p", template: t, writer: w, id: "para1"}
	v12 := Item{Text: `` + "\n" + ``}
	v14 := Item{Name: "button", template: t, writer: w, id: "button1"}
	v13 := Item{Name: "p", Contents: []*Item{&v14}}

	t.Contents = []*Item{&v1, &v5, &v6, &v8, &v9, &v10, &v11, &v12, &v13}

	return &Test_T{
		name:     t.name,
		Template: t,
		Span1:    &v3,
		Text1:    &v7,
		Para1:    &v11,
		Button1:  &v14,
	}
}
