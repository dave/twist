
package master

import (
	"twist"
)


type PlainMaster_T struct {
	
	name string
	*twist.Template
	
	Navigation *twist.Item
	
	Header *twist.Item
	
	Content *twist.Item
	
	Footer *twist.Item
	

}

func (t *PlainMaster_T) GetTemplate() *twist.Template {
	return t.Template
}

func PlainMaster(c *twist.Context, id string) *PlainMaster_T {
	
	t := twist.GetTemplateByPath("html_master_plainMaster")
	t.Path = "html_master_plainMaster"
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(t)

	
	v1 := twist.NewItemId("div", t, c.Writer, "Navigation", map[string]string{}, map[string]string{}, []*twist.Item{}, )
	v2 := twist.NewTextItem(``+"\n"+``)
	v4 := twist.NewTextItem(`Plain page`)
	v3 := twist.NewItemId("h1", t, c.Writer, "Header", map[string]string{}, map[string]string{}, []*twist.Item{v4, }, )
	v5 := twist.NewTextItem(``+"\n"+``)
	v6 := twist.NewItemId("div", t, c.Writer, "Content", map[string]string{}, map[string]string{}, []*twist.Item{}, )
	v7 := twist.NewTextItem(``+"\n"+``)
	v8 := twist.NewItemId("div", t, c.Writer, "Footer", map[string]string{}, map[string]string{" background-color":"#000000", " font-weight":"bold", "color":"#ffffff", " padding":"5px", }, []*twist.Item{}, )

	t.Contents = []*twist.Item{ v1, v2, v3, v5, v6, v7, v8,  }
	
	return &PlainMaster_T{
		name : t.Name, 
		Template : t,
		
		Navigation : v1,
		
		Header : v3,
		
		Content : v6,
		
		Footer : v8,
		
	}
}

type RedMaster_T struct {
	
	name string
	*twist.Template
	
	Navigation *twist.Item
	
	Header *twist.Item
	
	Topper *twist.Item
	
	Location *twist.Item
	
	Date *twist.Item
	
	Content *twist.Item
	
	Footer *twist.Item
	

}

func (t *RedMaster_T) GetTemplate() *twist.Template {
	return t.Template
}

func RedMaster(c *twist.Context, id string) *RedMaster_T {
	
	t := twist.GetTemplateByPath("html_master_redMaster")
	t.Path = "html_master_redMaster"
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(t)

	
	v1 := twist.NewItemId("div", t, c.Writer, "Navigation", map[string]string{}, map[string]string{}, []*twist.Item{}, )
	v2 := twist.NewTextItem(``+"\n"+``)
	v4 := twist.NewTextItem(`Red page`)
	v3 := twist.NewItemId("h1", t, c.Writer, "Header", map[string]string{}, map[string]string{}, []*twist.Item{v4, }, )
	v5 := twist.NewTextItem(``+"\n"+``)
	v11 := twist.NewTextItem(``+"\n"+``)
	v10 := twist.NewItemId("span", t, c.Writer, "Date", map[string]string{}, map[string]string{}, []*twist.Item{}, )
	v9 := twist.NewTextItem(` @ `)
	v8 := twist.NewItemId("span", t, c.Writer, "Location", map[string]string{}, map[string]string{}, []*twist.Item{}, )
	v7 := twist.NewTextItem(``+"\n"+`	`)
	v6 := twist.NewItemId("div", t, c.Writer, "Topper", map[string]string{}, map[string]string{"background-color":"#cc0000", " color":"#ffffff", " font-weight":"bold", " padding":"5px", }, []*twist.Item{v7, v8, v9, v10, v11, }, )
	v12 := twist.NewTextItem(``+"\n"+``)
	v13 := twist.NewItemId("div", t, c.Writer, "Content", map[string]string{}, map[string]string{}, []*twist.Item{}, )
	v14 := twist.NewTextItem(``+"\n"+``)
	v15 := twist.NewItemId("div", t, c.Writer, "Footer", map[string]string{}, map[string]string{" background-color":"#cc0000", " font-weight":"bold", "color":"#ffffff", " padding":"5px", }, []*twist.Item{}, )

	t.Contents = []*twist.Item{ v1, v2, v3, v5, v6, v12, v13, v14, v15,  }
	
	return &RedMaster_T{
		name : t.Name, 
		Template : t,
		
		Navigation : v1,
		
		Header : v3,
		
		Topper : v6,
		
		Location : v8,
		
		Date : v10,
		
		Content : v13,
		
		Footer : v15,
		
	}
}
