package handler

import (
	"http"
	"templates"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(wr http.ResponseWriter, r *http.Request) {

	templates.Server(
		wr,
		r,
		func() interface{} { return Functions(0) },
	)
}

type Functions int

func (f Functions) Root(c templates.Context, root *templates.Item) {

	

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

	//inner.Img1.Click(Functions.MyClick, &MyClick_T{Span1: inner.Span1, Text1: test.Text1, Img1: inner.Img1})

	inner.Img1.Click(Functions.MyClickNew, MyClickNew_T { Val1 : "testing", Span1: inner.Span1, Img1: inner.Img1, Text1: test.Text1 })
	
	
	

	c.Send()
}

type MyClickNew_T struct {
	Val1  templates.Value
	Span1 *templates.Item
	Img1  *templates.Item
	Text1 *templates.Item
}

func (f Functions) MyClickNew(c templates.Context, v MyClickNew_T) {

	v.Img1.Css("border", "10px solid #ff0000")
	v.Span1.Html("WHOOOOOOPPPPPPPPP!!!!" + v.Val1)

	c.Send()
}

type MyClick_T struct {
	Span1 *templates.Item
	Text1 *templates.Item
	Img1  *templates.Item
}

func (f Functions) MyClick(c templates.Context, v MyClick_T) {

	v.Img1.Css("border", "10px solid #ff0000")
	v.Span1.Html("WHOOOOOOPPPPPPPPP!!!!! " + v.Text1.Value)

	c.Send()
}


type MyClickFoo_T struct {
	Span1 *templates.Item
	Text1 *templates.Item
}

func (f Functions) MyClickFoo(c templates.Context, v MyClickFoo_T) {

	v.Span1.Html("FOOOOOOOOOOO " + v.Text1.Value)
	c.Send()

}
