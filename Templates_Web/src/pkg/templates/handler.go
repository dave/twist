package templates

import (
	"fmt"
	"http"
	"strings"
	"appengine"
	"reflect"
)


func Handler(
	wr http.ResponseWriter, 
	r *http.Request,
	getContext func(c *appengine.Context, w *Writer, r *http.Request) interface{}) {

	c := appengine.NewContext(r)
	path := r.URL.Path

	if strings.HasPrefix(path, "/template") {

		name := path[strings.Index(path, "_") + 1:]
		c.Infof(name)
		template := GetTemplateByName(name)
		fmt.Fprint(wr, template.GetTemplateJavascript())
		
	} else if strings.HasPrefix(path, "/function") {
	
		w := NewWriter(wr, false)
		
		r.ParseForm()
		
		params := make([]reflect.Value,0)
		for k, _ := range r.Form {
			if k != "function" {
				item := NewItemStub(k, w)
				item.Value = r.FormValue(k)
				params = append(params, reflect.ValueOf(item))
			}
		}
		
		function := r.FormValue("function")
		c := getContext(&c, w, r)
		v := reflect.ValueOf(c)
		typ := v.Type()
		if n := v.Type().NumMethod(); n > 0 {
			for i := 0; i < n; i++ {
				m := typ.Method(i)
				if m.Name == function {
					v.Method(i).Call(params)
				}
			}
		}
	} else {

		w := NewWriter(wr, true)

		root := Root(w)
		
		context := getContext(&c, w, r)
		
		r := reflect.ValueOf(root)
		
		params := []reflect.Value{r}
		
		v := reflect.ValueOf(context)
		typ := v.Type()
		if n := v.Type().NumMethod(); n > 0 {
			for i := 0; i < n; i++ {
				m := typ.Method(i)
				if m.Name == "Root" {
					v.Method(i).Call(params)
				}
			}
		}

	}

}












