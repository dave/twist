
package twist

func getTemplateByName(name string) *Template {
	switch name {
		
		case "inner" : 
			return inner_Template()
		
		case "page2" : 
			return page2_Template()
		
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
		name     : "inner",
		Html     : `<script>function template_inner(id){return "<b>This is an inner template, here's some text: <span style=\"border:1px solid #000000;\" id=\""+id+"_span1\" /></b><br />\n...and here's an image: <img id=\""+id+"_img1\" border=\"0\" style=\"border:2px solid #00ff00;\" /><br />\n<a href=\"/\" id=\""+id+"_myLink\">This is a link</a>"}</script>`,
	}
}
func (t *Inner_T) GetTemplate() *Template {
	return t.Template
}

func Inner(c *Context, id string) *Inner_T {
	
	t := inner_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(*t)

	
	v3 := Item{Name: "span", template: t, writer: c.Writer, id: "span1", Styles: map[string]string{"border":"1px solid #000000", }, }
	v2 := Item{Text:`This is an inner template, here's some text: `}
	v1 := Item{Name: "b", Contents: []*Item{&v2, &v3, }, }
	v4 := Item{Name: "br", }
	v5 := Item{Text:``+"\n"+`...and here's an image: `}
	v6 := Item{Name: "img", template: t, writer: c.Writer, id: "img1", Attributes: map[string]string{"border":"0", }, Styles: map[string]string{"border":"2px solid #00ff00", }, }
	v7 := Item{Name: "br", }
	v8 := Item{Text:``+"\n"+``}
	v10 := Item{Text:`This is a link`}
	v9 := Item{Name: "a", template: t, writer: c.Writer, id: "myLink", Attributes: map[string]string{"href":"/", }, Contents: []*Item{&v10, }, }

	t.Contents = []*Item{ &v1, &v4, &v5, &v6, &v7, &v8, &v9,  }
	
	return &Inner_T{
		name : t.name, 
		Template : t,
		
		Span1 : &v3,
		
		Img1 : &v6,
		
		MyLink : &v9,
		
	}
}

type Page2_T struct {
	
	name string
	*Template
	
	Button1 *Item
	

}
func page2_Template() *Template{
	return &Template {
		name     : "page2",
		Html     : `<script>function template_page2(id){return "<h1>Page 2!!!... </h1>\n<p><button id=\""+id+"_button1\">Back</button></p>"}</script>`,
	}
}
func (t *Page2_T) GetTemplate() *Template {
	return t.Template
}

func Page2(c *Context, id string) *Page2_T {
	
	t := page2_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(*t)

	
	v2 := Item{Text:`Page 2!!!... `}
	v1 := Item{Name: "h1", Contents: []*Item{&v2, }, }
	v3 := Item{Text:``+"\n"+``}
	v6 := Item{Text:`Back`}
	v5 := Item{Name: "button", template: t, writer: c.Writer, id: "button1", Contents: []*Item{&v6, }, }
	v4 := Item{Name: "p", Contents: []*Item{&v5, }, }

	t.Contents = []*Item{ &v1, &v3, &v4,  }
	
	return &Page2_T{
		name : t.name, 
		Template : t,
		
		Button1 : &v5,
		
	}
}

type Test_T struct {
	
	name string
	*Template
	
	Span1 *Item
	
	Text1 *Item
	
	Para1 *Item
	

}
func test_Template() *Template{
	return &Template {
		name     : "test",
		Html     : `<script>function template_test(id){return "<p>TEsting... <span id=\""+id+"_span1\">span1</span></p>\n<p><input id=\""+id+"_text1\" type=\"text\" /></p>\n<p></p>\n<p id=\""+id+"_para1\"></p>"}</script>`,
	}
}
func (t *Test_T) GetTemplate() *Template {
	return t.Template
}

func Test(c *Context, id string) *Test_T {
	
	t := test_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(*t)

	
	v4 := Item{Text:`span1`}
	v3 := Item{Name: "span", template: t, writer: c.Writer, id: "span1", Contents: []*Item{&v4, }, }
	v2 := Item{Text:`TEsting... `}
	v1 := Item{Name: "p", Contents: []*Item{&v2, &v3, }, }
	v5 := Item{Text:``+"\n"+``}
	v7 := Item{Name: "input", template: t, writer: c.Writer, id: "text1", Attributes: map[string]string{"type":"text", }, }
	v6 := Item{Name: "p", Contents: []*Item{&v7, }, }
	v8 := Item{Text:``+"\n"+``}
	v9 := Item{Name: "p", }
	v10 := Item{Text:``+"\n"+``}
	v11 := Item{Name: "p", template: t, writer: c.Writer, id: "para1", }

	t.Contents = []*Item{ &v1, &v5, &v6, &v8, &v9, &v10, &v11,  }
	
	return &Test_T{
		name : t.name, 
		Template : t,
		
		Span1 : &v3,
		
		Text1 : &v7,
		
		Para1 : &v11,
		
	}
}
