package twist

import (
	"fmt"
	"http"
	"strings"
	"appengine"
	"reflect"
	"json"
	"strconv"
)

type Context struct {
	*Writer
	Context *appengine.Context
	Request *http.Request
	Root *Item
}

func Server(wr http.ResponseWriter, r *http.Request, getFunctionsType func() interface{}) {

	path := r.URL.Path

	if path == "/favicon.ico" {
		
		return

	} else if path == "/" {

		serverRoot(wr, r, getFunctionsType)

	} else if strings.HasPrefix(path, "/template") {

		serverTemplate(wr, r, getFunctionsType)

	} else if strings.HasPrefix(path, "/function") {

		serverFunction(wr, r, getFunctionsType)

	} else if r.Method == "GET" {

		serverPage(wr, r, getFunctionsType)

	}
}

func serverTemplate(wr http.ResponseWriter, r *http.Request, getFunctionsType func() interface{}) {

	path := r.URL.Path
	name := path[strings.Index(path, "_")+1:]
	template := getTemplateByName(name)
	fmt.Fprint(wr, template.GetTemplateJavascript())

}

func serverRoot(wr http.ResponseWriter, r *http.Request, getFunctionsType func() interface{}) {
	
	c := appengine.NewContext(r)
	w := NewWriter(wr, true)

	context := Context {
		Writer: w, 
		Context: &c, 
		Request: r, 
		Root: Root(w),
	}

	contextVal := reflect.ValueOf(context)

	params := []reflect.Value{contextVal}

	m, _ := findMethod("Root", getFunctionsType)
	m.Call(params)

}

func serverFunction(wr http.ResponseWriter, r *http.Request, getFunctionsType func() interface{}) {
	
	c := appengine.NewContext(r)
	w := NewWriter(wr, false)

	r.ParseForm()

	stubs := new(allStubs)
	for v, _ := range r.Form {
		json.Unmarshal([]uint8(v), stubs)
		break
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
	needsHash := false

	val := reflect.New(method.Type.In(2)).Elem()
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		name := typ.Field(i).Name
		field := val.FieldByName(name)
		switch o := field.Interface().(type) {
			case String:
				stub, found := getValueStubByName(stubs.Values, name)
				if found {
					value := String(toString(stub.V))
					field.Set(reflect.ValueOf(value))
				}
			case StringHashed:
				needsHash = true
				stub, found := getValueStubByName(stubs.Values, name)
				if found {
					value := StringHashed(toString(stub.V))
					field.Set(reflect.ValueOf(value))
				}
			case StringEncrypted:
				needsHash = true //???
				panic("TODO")
			case Int:
				stub, found := getValueStubByName(stubs.Values, name)
				if found {
					value := Int(toInt(stub.V))
					field.Set(reflect.ValueOf(value))
				}
			case IntHashed:
				needsHash = true
				stub, found := getValueStubByName(stubs.Values, name)
				if found {
					value := IntHashed(toInt(stub.V))
					field.Set(reflect.ValueOf(value))
				}
			case IntEncrypted:
				needsHash = true //???
				panic("TODO")
			case *Item:
				needsHash = true
				stub, found := getItemStubByName(stubs.Items, name)
				if found {
					item := newItemFromAction(stub.I, w)
					item.Value = stub.V
					field.Set(reflect.ValueOf(item))
				}
			default:
				panic("Incorrect item/value " + name)
		}
	}

	if needsHash {
		proposedHashFromClient := stubs.Hash
		calculatedHash := getHash(*stubs)

		if proposedHashFromClient != calculatedHash {
			panic("hash mismatch")
		}
	}

	functionParams := make([]reflect.Value, 2)
	functionParams[0] = reflect.ValueOf(context)
	functionParams[1] = val
	methodValue.Call(functionParams)

}
func getItemStubByName(items []itemStub, name string) (item itemStub, found bool) {
	for _, i := range items {
		if i.N == name {
			return i, true
		}
	}
	return itemStub{}, false
}
func getValueStubByName(values []valueStub, name string) (val valueStub, found bool) {
	for _, v := range values {
		if v.N == name {
			return v, true
		}
	}
	return valueStub{}, false
}

func serverPage(wr http.ResponseWriter, r *http.Request, getFunctionsType func() interface{}) {

	c := appengine.NewContext(r)
	path := r.URL.Path
	w := NewWriter(wr, true)
	pageName := path[1:]

	context := Context {
		Writer: w, 
		Context: &c, 
		Request: r, 
		Root: Root(w),
	}

	r.ParseForm()

	methodValue, method := findMethod(pageName, getFunctionsType)

	if method.Type.NumIn() != 3 {
		panic("function " + pageName + " should have two fields.")
	}
	if method.Type.In(1) != reflect.TypeOf(context) {
		panic("function " + pageName + " first field should be type templates.Context")
	}
	
	needsHash := false

	valueStubs := make([]valueStub, 0)

	val := reflect.New(method.Type.In(2)).Elem()
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		name := typ.Field(i).Name
		field := val.FieldByName(name)
		switch o := field.Interface().(type) {
			case String:
				v := r.FormValue(name)
				valueStubs = append(valueStubs, valueStub{N:name, V:v, T:1})
				value := String(v)
				field.Set(reflect.ValueOf(value))
			case StringHashed:
				needsHash = true
				v := r.FormValue(name)
				valueStubs = append(valueStubs, valueStub{N:name, V:v, T:2})
				value := StringHashed(v)
				field.Set(reflect.ValueOf(value))
			case StringEncrypted:
				needsHash = true //???
				panic("TODO")
			case Int:
				v := r.FormValue(name)
				valueStubs = append(valueStubs, valueStub{N:name, V:v, T:4})
				vInt, _ := strconv.Atoi(v)
				value := Int(vInt)
				field.Set(reflect.ValueOf(value))
			case IntHashed:
				needsHash = true
				v := r.FormValue(name)
				valueStubs = append(valueStubs, valueStub{N:name, V:v, T:5})
				vInt, _ := strconv.Atoi(v)
				value := IntHashed(vInt)
				field.Set(reflect.ValueOf(value))
			case IntEncrypted:
				needsHash = true //???
				panic("TODO")
			case *Item:
				panic("We can't have Items in a Link - name:" + name)
			default:
				panic("Incorrect value " + name)
		}
	}

	if needsHash {
		proposedHashFromClient := r.FormValue("_hash")
		stubs := allStubs{Func: pageName, Values: valueStubs}

		calculatedHash := getHash(stubs)
		if proposedHashFromClient != calculatedHash {
			panic("hash mismatch")
		}
	}

	functionParams := make([]reflect.Value, 2)
	functionParams[0] = reflect.ValueOf(context)
	functionParams[1] = val
	methodValue.Call(functionParams)

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
