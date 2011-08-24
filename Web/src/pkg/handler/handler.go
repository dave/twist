package handler

import (
	"http"
	"templates"
	"appengine"
)
func init() {
	http.HandleFunc("/", handler)
}


func handler(wr http.ResponseWriter, r *http.Request) {

	templates.Handler(
		wr, 
		r,
		func(c *appengine.Context, w *templates.Writer, r *http.Request) interface{} { 
			return &Context{Context : c, Writer : w, Request : r }
		},
	)

}

type Context struct{

	*templates.Writer
	Context *appengine.Context
	Request *http.Request

}

func (c Context)Root(root *templates.Item) {
	
	test := templates.Test(c.Writer, "main")
	root.Html(test)
	
	test.Span1.Html("Hello world!")
	test.Text1.Attr("value", "foo")
	
	inner := templates.Inner(c.Writer, "dave")
	test.Para1.Html(inner)
		
	inner.Span1.Html("BAR")
	inner.Img1.Attr("src", "http://pix-eu.dontstayin.com/53812cd7-33c7-44ac-a766-9710e4f14077.jpg")
	inner.Img1.Attr("width", 100)
	inner.Img1.Attr("height", 100)
	
	inner.Img1.Click(Context.MyClickFoo, inner.Span1, test.Text1)
	
	c.Send()
}

func (c Context)MyClick(span1 *templates.Item, text1 *templates.Item) {
	
	span1.Html("WHOOOOOOPPPPPPPPP!!!!! " + text1.Value)
	c.Send()
}

func (c Context)MyClickFoo(span1 *templates.Item, text1 *templates.Item) {

	span1.Html("FOOOOOOOOOOO " + text1.Value)
	c.Send()

}



