
package content

import (
	"twist"
)


type Plain1_T struct {
	
	name string
	*twist.Template
	
	Output *twist.Item
	
	Plus *twist.Item
	
	Minus *twist.Item
	
	Count *twist.Item
	

}

func (t *Plain1_T) GetTemplate() *twist.Template {
	return t.Template
}

func Plain1(c *twist.Context, id string) *Plain1_T {
	
	t := twist.GetTemplateByPath("html_content_plain1")
	t.Path = "html_content_plain1"
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(t)

	
	v2 := twist.NewTextItem(``+"\n"+`	This is a page using the plain_template.html file.`+"\n"+``)
	v1 := twist.NewItem("p", map[string]string{}, map[string]string{}, []*twist.Item{v2, }, )
	v3 := twist.NewTextItem(``+"\n"+``)
	v5 := twist.NewTextItem(``+"\n"+`	Here's a counter:`+"\n"+``)
	v4 := twist.NewItem("p", map[string]string{}, map[string]string{}, []*twist.Item{v5, }, )
	v6 := twist.NewTextItem(``+"\n"+``)
	v8 := twist.NewTextItem(``+"\n"+`	0`+"\n"+``)
	v7 := twist.NewItemId("h1", t, c.Writer, "Output", map[string]string{}, map[string]string{}, []*twist.Item{v8, }, )
	v9 := twist.NewTextItem(``+"\n"+``)
	v17 := twist.NewTextItem(``+"\n"+``)
	v16 := twist.NewTextItem(`minus`)
	v15 := twist.NewItemId("a", t, c.Writer, "Minus", map[string]string{"href":"#", }, map[string]string{}, []*twist.Item{v16, }, )
	v14 := twist.NewTextItem(` or `)
	v13 := twist.NewTextItem(`plus`)
	v12 := twist.NewItemId("a", t, c.Writer, "Plus", map[string]string{"href":"#", }, map[string]string{}, []*twist.Item{v13, }, )
	v11 := twist.NewTextItem(``+"\n"+`	Click `)
	v10 := twist.NewItem("p", map[string]string{}, map[string]string{}, []*twist.Item{v11, v12, v14, v15, v17, }, )
	v18 := twist.NewTextItem(``+"\n"+``)
	v19 := twist.NewItemId("input", t, c.Writer, "Count", map[string]string{"type":"hidden", }, map[string]string{}, []*twist.Item{}, )

	t.Contents = []*twist.Item{ v1, v3, v4, v6, v7, v9, v10, v18, v19,  }
	
	return &Plain1_T{
		name : t.Name, 
		Template : t,
		
		Output : v7,
		
		Plus : v12,
		
		Minus : v15,
		
		Count : v19,
		
	}
}

type Plain2_T struct {
	
	name string
	*twist.Template
	

}

func (t *Plain2_T) GetTemplate() *twist.Template {
	return t.Template
}

func Plain2(c *twist.Context, id string) *Plain2_T {
	
	t := twist.GetTemplateByPath("html_content_plain2")
	t.Path = "html_content_plain2"
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(t)

	
	v2 := twist.NewTextItem(``+"\n"+`	This is the second plain page.`+"\n"+``)
	v1 := twist.NewItem("p", map[string]string{}, map[string]string{}, []*twist.Item{v2, }, )

	t.Contents = []*twist.Item{ v1,  }
	
	return &Plain2_T{
		name : t.Name, 
		Template : t,
		
	}
}

type Plain3_T struct {
	
	name string
	*twist.Template
	

}

func (t *Plain3_T) GetTemplate() *twist.Template {
	return t.Template
}

func Plain3(c *twist.Context, id string) *Plain3_T {
	
	t := twist.GetTemplateByPath("html_content_plain3")
	t.Path = "html_content_plain3"
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(t)

	
	v2 := twist.NewTextItem(``+"\n"+`	This is another plain page.`+"\n"+``)
	v1 := twist.NewItem("p", map[string]string{}, map[string]string{}, []*twist.Item{v2, }, )

	t.Contents = []*twist.Item{ v1,  }
	
	return &Plain3_T{
		name : t.Name, 
		Template : t,
		
	}
}

type Red1_T struct {
	
	name string
	*twist.Template
	

}

func (t *Red1_T) GetTemplate() *twist.Template {
	return t.Template
}

func Red1(c *twist.Context, id string) *Red1_T {
	
	t := twist.GetTemplateByPath("html_content_red1")
	t.Path = "html_content_red1"
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(t)

	
	v2 := twist.NewTextItem(``+"\n"+`	This is a page using the red_template.html file.`+"\n"+``)
	v1 := twist.NewItem("p", map[string]string{}, map[string]string{}, []*twist.Item{v2, }, )

	t.Contents = []*twist.Item{ v1,  }
	
	return &Red1_T{
		name : t.Name, 
		Template : t,
		
	}
}

type Red2_T struct {
	
	name string
	*twist.Template
	

}

func (t *Red2_T) GetTemplate() *twist.Template {
	return t.Template
}

func Red2(c *twist.Context, id string) *Red2_T {
	
	t := twist.GetTemplateByPath("html_content_red2")
	t.Path = "html_content_red2"
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(t)

	
	v2 := twist.NewTextItem(``+"\n"+`	This is the second red page.`+"\n"+``)
	v1 := twist.NewItem("p", map[string]string{}, map[string]string{}, []*twist.Item{v2, }, )

	t.Contents = []*twist.Item{ v1,  }
	
	return &Red2_T{
		name : t.Name, 
		Template : t,
		
	}
}

type Red3_T struct {
	
	name string
	*twist.Template
	

}

func (t *Red3_T) GetTemplate() *twist.Template {
	return t.Template
}

func Red3(c *twist.Context, id string) *Red3_T {
	
	t := twist.GetTemplateByPath("html_content_red3")
	t.Path = "html_content_red3"
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(t)

	
	v2 := twist.NewTextItem(``+"\n"+`	This is another red page.`+"\n"+``)
	v1 := twist.NewItem("p", map[string]string{}, map[string]string{}, []*twist.Item{v2, }, )

	t.Contents = []*twist.Item{ v1,  }
	
	return &Red3_T{
		name : t.Name, 
		Template : t,
		
	}
}
