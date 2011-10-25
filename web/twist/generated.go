
package twist

func getTemplateByName(name string) *Template {
	switch name {
		
		case "navigation" : 
			return navigation_Template()
		
		case "plain1" : 
			return plain1_Template()
		
		case "plain2" : 
			return plain2_Template()
		
		case "plain3" : 
			return plain3_Template()
		
		case "plainMaster" : 
			return plainMaster_Template()
		
		case "red1" : 
			return red1_Template()
		
		case "red2" : 
			return red2_Template()
		
		case "red3" : 
			return red3_Template()
		
		case "redMaster" : 
			return redMaster_Template()
		
	}
	return nil
}


type Navigation_T struct {
	
	name string
	*Template
	
	Plain1Link *Item
	
	Plain2Link *Item
	
	Plain3Link *Item
	
	Red1Link *Item
	
	Red2Link *Item
	
	Red3Link *Item
	

}
func navigation_Template() *Template{
	return &Template {
		name     : "navigation",
		Html     : `<script>function template_navigation(id){return "<div style=\"border 1px solid #000000; padding: 5px;\">\n	<a id=\""+id+"_plain1Link\">Plain page 1</a> |\n	<a id=\""+id+"_plain2Link\">Plain page 2</a> |\n	<a id=\""+id+"_plain3Link\">Plain page 3</a> |\n	<a id=\""+id+"_red1Link\">Red page 1</a> |\n	<a id=\""+id+"_red2Link\">Red page 2</a> |\n	<a id=\""+id+"_red3Link\">Red page 3</a>\n</div>"}</script>`,
	}
}
func (t *Navigation_T) GetTemplate() *Template {
	return t.Template
}

func Navigation(c *Context, id string) *Navigation_T {
	
	t := navigation_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.registerTemplate(*t)

	
	v20 := Item{Text:``+"\n"+``}
	v19 := Item{Text:`Red page 3`}
	v18 := Item{Name: "a", template: t, writer: c.Writer, id: "red3Link", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v19, }, }
	v17 := Item{Text:` |`+"\n"+`	`}
	v16 := Item{Text:`Red page 2`}
	v15 := Item{Name: "a", template: t, writer: c.Writer, id: "red2Link", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v16, }, }
	v14 := Item{Text:` |`+"\n"+`	`}
	v13 := Item{Text:`Red page 1`}
	v12 := Item{Name: "a", template: t, writer: c.Writer, id: "red1Link", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v13, }, }
	v11 := Item{Text:` |`+"\n"+`	`}
	v10 := Item{Text:`Plain page 3`}
	v9 := Item{Name: "a", template: t, writer: c.Writer, id: "plain3Link", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v10, }, }
	v8 := Item{Text:` |`+"\n"+`	`}
	v7 := Item{Text:`Plain page 2`}
	v6 := Item{Name: "a", template: t, writer: c.Writer, id: "plain2Link", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v7, }, }
	v5 := Item{Text:` |`+"\n"+`	`}
	v4 := Item{Text:`Plain page 1`}
	v3 := Item{Name: "a", template: t, writer: c.Writer, id: "plain1Link", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v4, }, }
	v2 := Item{Text:``+"\n"+`	`}
	v1 := Item{Name: "div", Attributes: map[string]string{}, Styles: map[string]string{" padding":" 5px", }, Contents: []*Item{&v2, &v3, &v5, &v6, &v8, &v9, &v11, &v12, &v14, &v15, &v17, &v18, &v20, }, }

	t.Contents = []*Item{ &v1,  }
	
	return &Navigation_T{
		name : t.name, 
		Template : t,
		
		Plain1Link : &v3,
		
		Plain2Link : &v6,
		
		Plain3Link : &v9,
		
		Red1Link : &v12,
		
		Red2Link : &v15,
		
		Red3Link : &v18,
		
	}
}

type Plain1_T struct {
	
	name string
	*Template
	
	Output *Item
	
	Plus *Item
	
	Minus *Item
	
	Count *Item
	

}
func plain1_Template() *Template{
	return &Template {
		name     : "plain1",
		Html     : `<script>function template_plain1(id){return "<p>\n	This is a page using the plain_template.html file.\n</p>\n<p>\n	Here's a counter:\n</p>\n<h1 id=\""+id+"_Output\">\n	0\n</h1>\n<p>\n	Click <a href=\"#\" id=\""+id+"_Plus\">plus</a> or <a href=\"#\" id=\""+id+"_Minus\">minus</a>\n</p>\n<input type=\"hidden\" id=\""+id+"_Count\" />"}</script>`,
	}
}
func (t *Plain1_T) GetTemplate() *Template {
	return t.Template
}

func Plain1(c *Context, id string) *Plain1_T {
	
	t := plain1_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.registerTemplate(*t)

	
	v2 := Item{Text:``+"\n"+`	This is a page using the plain_template.html file.`+"\n"+``}
	v1 := Item{Name: "p", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v2, }, }
	v3 := Item{Text:``+"\n"+``}
	v5 := Item{Text:``+"\n"+`	Here's a counter:`+"\n"+``}
	v4 := Item{Name: "p", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v5, }, }
	v6 := Item{Text:``+"\n"+``}
	v8 := Item{Text:``+"\n"+`	0`+"\n"+``}
	v7 := Item{Name: "h1", template: t, writer: c.Writer, id: "Output", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v8, }, }
	v9 := Item{Text:``+"\n"+``}
	v17 := Item{Text:``+"\n"+``}
	v16 := Item{Text:`minus`}
	v15 := Item{Name: "a", template: t, writer: c.Writer, id: "Minus", Attributes: map[string]string{"href":"#", }, Styles: map[string]string{}, Contents: []*Item{&v16, }, }
	v14 := Item{Text:` or `}
	v13 := Item{Text:`plus`}
	v12 := Item{Name: "a", template: t, writer: c.Writer, id: "Plus", Attributes: map[string]string{"href":"#", }, Styles: map[string]string{}, Contents: []*Item{&v13, }, }
	v11 := Item{Text:``+"\n"+`	Click `}
	v10 := Item{Name: "p", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v11, &v12, &v14, &v15, &v17, }, }
	v18 := Item{Text:``+"\n"+``}
	v19 := Item{Name: "input", template: t, writer: c.Writer, id: "Count", Attributes: map[string]string{"type":"hidden", }, Styles: map[string]string{}, }

	t.Contents = []*Item{ &v1, &v3, &v4, &v6, &v7, &v9, &v10, &v18, &v19,  }
	
	return &Plain1_T{
		name : t.name, 
		Template : t,
		
		Output : &v7,
		
		Plus : &v12,
		
		Minus : &v15,
		
		Count : &v19,
		
	}
}

type Plain2_T struct {
	
	name string
	*Template
	

}
func plain2_Template() *Template{
	return &Template {
		name     : "plain2",
		Html     : `<script>function template_plain2(id){return "<p>\n	This is the second plain page.\n</p>"}</script>`,
	}
}
func (t *Plain2_T) GetTemplate() *Template {
	return t.Template
}

func Plain2(c *Context, id string) *Plain2_T {
	
	t := plain2_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.registerTemplate(*t)

	
	v2 := Item{Text:``+"\n"+`	This is the second plain page.`+"\n"+``}
	v1 := Item{Name: "p", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v2, }, }

	t.Contents = []*Item{ &v1,  }
	
	return &Plain2_T{
		name : t.name, 
		Template : t,
		
	}
}

type Plain3_T struct {
	
	name string
	*Template
	

}
func plain3_Template() *Template{
	return &Template {
		name     : "plain3",
		Html     : `<script>function template_plain3(id){return "<p>\n	This is another plain page.\n</p>"}</script>`,
	}
}
func (t *Plain3_T) GetTemplate() *Template {
	return t.Template
}

func Plain3(c *Context, id string) *Plain3_T {
	
	t := plain3_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.registerTemplate(*t)

	
	v2 := Item{Text:``+"\n"+`	This is another plain page.`+"\n"+``}
	v1 := Item{Name: "p", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v2, }, }

	t.Contents = []*Item{ &v1,  }
	
	return &Plain3_T{
		name : t.name, 
		Template : t,
		
	}
}

type PlainMaster_T struct {
	
	name string
	*Template
	
	Navigation *Item
	
	Header *Item
	
	Content *Item
	
	Footer *Item
	

}
func plainMaster_Template() *Template{
	return &Template {
		name     : "plainMaster",
		Html     : `<script>function template_plainMaster(id){return "<div id=\""+id+"_Navigation\" />\n<h1 id=\""+id+"_Header\">Plain page</h1>\n<div id=\""+id+"_Content\" />\n<div id=\""+id+"_Footer\" style=\"color:#ffffff; background-color:#000000; font-weight:bold; padding:5px;\" />"}</script>`,
	}
}
func (t *PlainMaster_T) GetTemplate() *Template {
	return t.Template
}

func PlainMaster(c *Context, id string) *PlainMaster_T {
	
	t := plainMaster_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.registerTemplate(*t)

	
	v1 := Item{Name: "div", template: t, writer: c.Writer, id: "Navigation", Attributes: map[string]string{}, Styles: map[string]string{}, }
	v2 := Item{Text:``+"\n"+``}
	v4 := Item{Text:`Plain page`}
	v3 := Item{Name: "h1", template: t, writer: c.Writer, id: "Header", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v4, }, }
	v5 := Item{Text:``+"\n"+``}
	v6 := Item{Name: "div", template: t, writer: c.Writer, id: "Content", Attributes: map[string]string{}, Styles: map[string]string{}, }
	v7 := Item{Text:``+"\n"+``}
	v8 := Item{Name: "div", template: t, writer: c.Writer, id: "Footer", Attributes: map[string]string{}, Styles: map[string]string{" background-color":"#000000", " font-weight":"bold", "color":"#ffffff", " padding":"5px", }, }

	t.Contents = []*Item{ &v1, &v2, &v3, &v5, &v6, &v7, &v8,  }
	
	return &PlainMaster_T{
		name : t.name, 
		Template : t,
		
		Navigation : &v1,
		
		Header : &v3,
		
		Content : &v6,
		
		Footer : &v8,
		
	}
}

type Red1_T struct {
	
	name string
	*Template
	

}
func red1_Template() *Template{
	return &Template {
		name     : "red1",
		Html     : `<script>function template_red1(id){return "<p>\n	This is a page using the red_template.html file.\n</p>"}</script>`,
	}
}
func (t *Red1_T) GetTemplate() *Template {
	return t.Template
}

func Red1(c *Context, id string) *Red1_T {
	
	t := red1_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.registerTemplate(*t)

	
	v2 := Item{Text:``+"\n"+`	This is a page using the red_template.html file.`+"\n"+``}
	v1 := Item{Name: "p", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v2, }, }

	t.Contents = []*Item{ &v1,  }
	
	return &Red1_T{
		name : t.name, 
		Template : t,
		
	}
}

type Red2_T struct {
	
	name string
	*Template
	

}
func red2_Template() *Template{
	return &Template {
		name     : "red2",
		Html     : `<script>function template_red2(id){return "<p>\n	This is the second red page.\n</p>"}</script>`,
	}
}
func (t *Red2_T) GetTemplate() *Template {
	return t.Template
}

func Red2(c *Context, id string) *Red2_T {
	
	t := red2_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.registerTemplate(*t)

	
	v2 := Item{Text:``+"\n"+`	This is the second red page.`+"\n"+``}
	v1 := Item{Name: "p", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v2, }, }

	t.Contents = []*Item{ &v1,  }
	
	return &Red2_T{
		name : t.name, 
		Template : t,
		
	}
}

type Red3_T struct {
	
	name string
	*Template
	

}
func red3_Template() *Template{
	return &Template {
		name     : "red3",
		Html     : `<script>function template_red3(id){return "<p>\n	This is another red page.\n</p>"}</script>`,
	}
}
func (t *Red3_T) GetTemplate() *Template {
	return t.Template
}

func Red3(c *Context, id string) *Red3_T {
	
	t := red3_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.registerTemplate(*t)

	
	v2 := Item{Text:``+"\n"+`	This is another red page.`+"\n"+``}
	v1 := Item{Name: "p", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v2, }, }

	t.Contents = []*Item{ &v1,  }
	
	return &Red3_T{
		name : t.name, 
		Template : t,
		
	}
}

type RedMaster_T struct {
	
	name string
	*Template
	
	Navigation *Item
	
	Header *Item
	
	Topper *Item
	
	Location *Item
	
	Date *Item
	
	Content *Item
	
	Footer *Item
	

}
func redMaster_Template() *Template{
	return &Template {
		name     : "redMaster",
		Html     : `<script>function template_redMaster(id){return "<div id=\""+id+"_Navigation\" />\n<h1 id=\""+id+"_Header\">Red page</h1>\n<div id=\""+id+"_Topper\" style=\"background-color:#cc0000; color:#ffffff; font-weight:bold; padding:5px;\">\n	<span id=\""+id+"_Location\" /> @ <span id=\""+id+"_Date\" />\n</div>\n<div id=\""+id+"_Content\" />\n<div id=\""+id+"_Footer\" style=\"color:#ffffff; background-color:#cc0000; font-weight:bold; padding:5px;\" />"}</script>`,
	}
}
func (t *RedMaster_T) GetTemplate() *Template {
	return t.Template
}

func RedMaster(c *Context, id string) *RedMaster_T {
	
	t := redMaster_Template()
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.registerTemplate(*t)

	
	v1 := Item{Name: "div", template: t, writer: c.Writer, id: "Navigation", Attributes: map[string]string{}, Styles: map[string]string{}, }
	v2 := Item{Text:``+"\n"+``}
	v4 := Item{Text:`Red page`}
	v3 := Item{Name: "h1", template: t, writer: c.Writer, id: "Header", Attributes: map[string]string{}, Styles: map[string]string{}, Contents: []*Item{&v4, }, }
	v5 := Item{Text:``+"\n"+``}
	v11 := Item{Text:``+"\n"+``}
	v10 := Item{Name: "span", template: t, writer: c.Writer, id: "Date", Attributes: map[string]string{}, Styles: map[string]string{}, }
	v9 := Item{Text:` @ `}
	v8 := Item{Name: "span", template: t, writer: c.Writer, id: "Location", Attributes: map[string]string{}, Styles: map[string]string{}, }
	v7 := Item{Text:``+"\n"+`	`}
	v6 := Item{Name: "div", template: t, writer: c.Writer, id: "Topper", Attributes: map[string]string{}, Styles: map[string]string{"background-color":"#cc0000", " color":"#ffffff", " font-weight":"bold", " padding":"5px", }, Contents: []*Item{&v7, &v8, &v9, &v10, &v11, }, }
	v12 := Item{Text:``+"\n"+``}
	v13 := Item{Name: "div", template: t, writer: c.Writer, id: "Content", Attributes: map[string]string{}, Styles: map[string]string{}, }
	v14 := Item{Text:``+"\n"+``}
	v15 := Item{Name: "div", template: t, writer: c.Writer, id: "Footer", Attributes: map[string]string{}, Styles: map[string]string{" background-color":"#cc0000", " font-weight":"bold", "color":"#ffffff", " padding":"5px", }, }

	t.Contents = []*Item{ &v1, &v2, &v3, &v5, &v6, &v12, &v13, &v14, &v15,  }
	
	return &RedMaster_T{
		name : t.name, 
		Template : t,
		
		Navigation : &v1,
		
		Header : &v3,
		
		Topper : &v6,
		
		Location : &v8,
		
		Date : &v10,
		
		Content : &v13,
		
		Footer : &v15,
		
	}
}
