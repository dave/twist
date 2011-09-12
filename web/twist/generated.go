package twist

func getTemplateByName(name string) *Template {
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
	MyLink *Item

}
func inner_Template() *Template{
	return &Template {
		name : "inner",
		Html : `<script>function template_inner(id){return "<b>This is an inner template, here's some text: <span style=\"border:1px solid #000000;\" id=\""+id+"_span1\" /></b><br />\n...and here's an image: <img id=\""+id+"_img1\" border=\"0\" style=\"border:2px solid #00ff00;\" /><br />\n<a href=\"/\" id=\""+id+"_myLink\">This is a link</a>"}</script>`,
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
		MyLink : &Item{
			id : "myLink",
			template : t, 
			writer : w,
		},
	}
}
type Test_T struct {
	
	name string
	*Template
	Span1 *Item
	Text1 *Item
	Para1 *Item
	Button1 *Item

}
func test_Template() *Template{
	return &Template {
		name : "test",
		Html : `<script>function template_test(id){return "<p>TEsting... <span id=\""+id+"_span1\">span1</span></p>\n<p><input id=\""+id+"_text1\" type=\"text\" /></p>\n<p></p>\n<p id=\""+id+"_para1\"></p>\n<p><button id=\""+id+"_button1\" /></p>"}</script>`,
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
	
	return &Test_T{
		name : t.name, 
		Template : t,
		Span1 : &Item{
			id : "span1",
			template : t, 
			writer : w,
		},
		Text1 : &Item{
			id : "text1",
			template : t, 
			writer : w,
		},
		Para1 : &Item{
			id : "para1",
			template : t, 
			writer : w,
		},
		Button1 : &Item{
			id : "button1",
			template : t, 
			writer : w,
		},
	}
}
