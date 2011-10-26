
package misc

import (
	"twist"
)


type Navigation_T struct {
	
	name string
	*twist.Template
	
	Plain1Link *twist.Item
	
	Plain2Link *twist.Item
	
	Plain3Link *twist.Item
	
	Red1Link *twist.Item
	
	Red2Link *twist.Item
	
	Red3Link *twist.Item
	

}

func (t *Navigation_T) GetTemplate() *twist.Template {
	return t.Template
}

func Navigation(c *twist.Context, id string) *Navigation_T {
	
	t := twist.GetTemplateByPath("html_misc_navigation")
	t.Path = "html_misc_navigation"
	t.Writer = c.Writer
	t.Id = id
	
	c.Writer.RegisterTemplate(t)

	
	v20 := twist.NewTextItem(``+"\n"+``)
	v19 := twist.NewTextItem(`Red page 3`)
	v18 := twist.NewItemId("a", t, c.Writer, "red3Link", map[string]string{}, map[string]string{}, []*twist.Item{v19, }, )
	v17 := twist.NewTextItem(` |`+"\n"+`	`)
	v16 := twist.NewTextItem(`Red page 2`)
	v15 := twist.NewItemId("a", t, c.Writer, "red2Link", map[string]string{}, map[string]string{}, []*twist.Item{v16, }, )
	v14 := twist.NewTextItem(` |`+"\n"+`	`)
	v13 := twist.NewTextItem(`Red page 1`)
	v12 := twist.NewItemId("a", t, c.Writer, "red1Link", map[string]string{}, map[string]string{}, []*twist.Item{v13, }, )
	v11 := twist.NewTextItem(` |`+"\n"+`	`)
	v10 := twist.NewTextItem(`Plain page 3`)
	v9 := twist.NewItemId("a", t, c.Writer, "plain3Link", map[string]string{}, map[string]string{}, []*twist.Item{v10, }, )
	v8 := twist.NewTextItem(` |`+"\n"+`	`)
	v7 := twist.NewTextItem(`Plain page 2`)
	v6 := twist.NewItemId("a", t, c.Writer, "plain2Link", map[string]string{}, map[string]string{}, []*twist.Item{v7, }, )
	v5 := twist.NewTextItem(` |`+"\n"+`	`)
	v4 := twist.NewTextItem(`Plain page 1`)
	v3 := twist.NewItemId("a", t, c.Writer, "plain1Link", map[string]string{}, map[string]string{}, []*twist.Item{v4, }, )
	v2 := twist.NewTextItem(``+"\n"+`	`)
	v1 := twist.NewItem("div", map[string]string{}, map[string]string{" padding":" 5px", }, []*twist.Item{v2, v3, v5, v6, v8, v9, v11, v12, v14, v15, v17, v18, v20, }, )

	t.Contents = []*twist.Item{ v1,  }
	
	return &Navigation_T{
		name : t.Name, 
		Template : t,
		
		Plain1Link : v3,
		
		Plain2Link : v6,
		
		Plain3Link : v9,
		
		Red1Link : v12,
		
		Red2Link : v15,
		
		Red3Link : v18,
		
	}
}
