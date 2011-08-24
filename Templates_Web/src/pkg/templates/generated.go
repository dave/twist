package templates

func GetTemplateByName(name string) *Template {
	switch name {
		case "inner" : 
			return inner_Template()
		case "test" :
			return test_Template()
	}
	return nil
}

type Inner_T struct {
	
	name string
	*Template
	Span1 *Item
	Img1 *Item

}
func inner_Template() *Template{
	return &Template {
		name : "inner",
		Html : `
<b>This is an inner template, here's some text: <span id="span1" /></b><br />
...and here's an image: <img id="img1" border="0" /><br />
`,
	}
}
func (t *Inner_T) GetTemplate() *Template {
	return t.Template
}

func Inner(w *Writer, id string) *Inner_T {
	
	
	t := inner_Template()
	t.Writer = w
	t.Id = id
	
	w.RegisterTemplate(t.name)

	
	return &Inner_T{
		name : t.name, 
		Template : t, 
		Span1 : &Item{
			id : "span1",
			template : t, 
			writer : w,
		},
		Img1 : &Item{
			id : "img1",
			template : t, 
			writer : w,
		},
	}
}


type Test_T struct {
	
	name string
	*Template
	Span1 *Item
	Para1 *Item
	Text1 *Item
	Button1 *Item
}
func test_Template() *Template {
	return &Template{
		name : "test",
		Html : `
<p>Template: <span id="span1">span1</span></p>
<p><input id="text1" type="text" /></p>
<p></p>
<p id="para1"></p>
<p><button id="button1" /></p>
`,
	}	
}
func (t *Test_T) GetTemplate() *Template {
	return t.Template
}

func Test(w *Writer, id string) *Test_T {
	
	t := test_Template()
	t.Writer = w
	t.Id = id
	
	w.RegisterTemplate(t.name)
	
	return &Test_T{
		name : t.name, 
		Template : t, 
		Span1 : &Item{
			id:"span1",
			template:t, 
			writer:w,
		},
		Para1 : &Item{
			id:"para1",
			template:t, 
			writer:w,
		},
		Text1 : &Item{
			id:"text1",
			template:t, 
			writer:w,
		},
		Button1 : &Item{
			id:"button1",
			template:t, 
			writer:w,
		},
	}
}


