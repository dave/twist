package templates

import (
	"fmt"
	"http"
	"strings"
	"appengine"
	"reflect"
	"json"
	//	"strconv"
)

type Context struct {
	*Writer
	Context *appengine.Context
	Request *http.Request
	Root *Item
}

func Server(wr http.ResponseWriter,
r *http.Request,
getFunctionsType func() interface{}) {

	c := appengine.NewContext(r)
	path := r.URL.Path

	if path == "/favicon.ico" {
		
		return

	} else if strings.HasPrefix(path, "/template") {

		name := path[strings.Index(path, "_")+1:]
		template := getTemplateByName(name)
		fmt.Fprint(wr, template.GetTemplateJavascript())

	} else if strings.HasPrefix(path, "/function") {
		w := NewWriter(wr, false)

		r.ParseForm()

		stubs := new(allStubs)
		stubsForCheck := new(allStubs)
		for v, _ := range r.Form {
			json.Unmarshal([]uint8(v), stubs)
			json.Unmarshal([]uint8(v), stubsForCheck)
			c.Infof(v)
			break
		}
		proposedHashFromClient := stubs.Hash

		//clear variant data out of stubsForCheck before checking hash
		stubsForCheck.Hash = ""
		for i, _ := range stubsForCheck.Items {
			stubsForCheck.Items[i].V = ""
		}

		calculatedHash := getHash(*stubsForCheck)

		if proposedHashFromClient != calculatedHash {
			panic("hash mismatch")
		}

		methodValue, method := findMethod(stubs.Func, getFunctionsType)
		context := Context {
			Writer: w, 
			Context: &c, 
			Request: r, 
			Root: Root(w),
		}

		if method.Type.NumIn() != 3 {
			panic("function " + stubs.Func + " should have two fields. It has " + fmt.Sprint(method.Type.NumIn()))
		}
		if method.Type.In(1) != reflect.TypeOf(context) {
			panic("function " + stubs.Func + " first field should be type templates.Context")
		}
		mainParam := reflect.New(method.Type.In(2)).Elem()

		for _, it := range stubs.Items {
			field := mainParam.FieldByName(it.N)
			item := newItemFromAction(it.I, w)
			item.Value = it.V
			field.Set(reflect.ValueOf(item))
		}

		for _, va := range stubs.Values {
			field := mainParam.FieldByName(va.N)
			if va.E {
				panic("to do")
			} else {
				value := newValueFromAction(va.V)
				field.Set(reflect.ValueOf(value))
			}
		}

		functionParams := make([]reflect.Value, 2)
		functionParams[0] = reflect.ValueOf(context)
		functionParams[1] = mainParam
		methodValue.Call(functionParams)

	} else if path == "/" {

		w := NewWriter(wr, true)

		context := Context {
			Writer: w, 
			Context: &c, 
			Request: r, 
			Root: Root(w),
		}

		c := reflect.ValueOf(context)

		params := []reflect.Value{c}

		m, _ := findMethod("Root", getFunctionsType)
		m.Call(params)

	} else if r.Method == "GET" {

		w := NewWriter(wr, true)
		pageName := path[1:]

		context := Context {
			Writer: w, 
			Context: &c, 
			Request: r, 
			Root: Root(w),
		}

		r.ParseForm()

		stubs := allStubs{}
		if len(r.Form) > 0 {
			proposedHashFromClient := ""
			valueStubs := make([]valueStub, 0)
			for n, v := range r.Form {
				if n == "_hash" {
					proposedHashFromClient = v[0]
				} else {
					c.Infof(n, v[0])
					valueStubs = append(valueStubs, valueStub{N:n, V:v[0], E:false})
				}
			}
			stubs = allStubs{Func: pageName, Values: valueStubs}
			calculatedHash := getHash(stubs)
			if proposedHashFromClient != calculatedHash {
				panic("hash mismatch")
			}
		}

		methodValue, method := findMethod(stubs.Func, getFunctionsType)

		if method.Type.NumIn() != 3 {
			panic("function " + stubs.Func + " should have two fields.")
		}
		if method.Type.In(1) != reflect.TypeOf(context) {
			panic("function " + stubs.Func + " first field should be type templates.Context")
		}
		mainParam := reflect.New(method.Type.In(2)).Elem()

		for _, va := range stubs.Values {
			field := mainParam.FieldByName(va.N)
			if va.E {
				panic("to do")
			} else {
				value := newValueFromAction(va.V)
				field.Set(reflect.ValueOf(value))
			}
		}
/*
		field := mainParam.FieldByName("Root")
		item := newItemFromAction("root", w)
		item.Value = ""
		field.Set(reflect.ValueOf(item))
*/
		functionParams := make([]reflect.Value, 2)
		functionParams[0] = reflect.ValueOf(context)
		functionParams[1] = mainParam
		methodValue.Call(functionParams)

		
	}
}

func findMethod(name string, getFunctionsType func() interface{}) (val reflect.Value, met reflect.Method) {

	v := reflect.ValueOf(getFunctionsType())
	t := v.Type()

	if n := t.NumMethod(); n > 0 {
		for i := 0; i < n; i++ {
			m := t.Method(i)
			if m.Name == name {
				return v.Method(i), m
			}
		}
	}
	panic("can't find method " + name)

}
