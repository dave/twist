package handler

import (
	"http"
	"twist"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(wr http.ResponseWriter, r *http.Request) {

	twist.Server(
		wr,
		r,
		func() interface{} { return Functions(0) },
	)
}

type Functions int

func (f Functions) Root(c *twist.Context) {

	test := twist.Test(c, "main")
	c.Root.Html(test)

	test.Span1.Html("Hello world!")
	test.Text1.Attr("value", "foo")

	inner := twist.Inner(c, "dave")
	test.Para1.Html(inner)

	inner.Span1.Html("BAR")
	inner.Img1.Attr("src", "http://pix-eu.dontstayin.com/53812cd7-33c7-44ac-a766-9710e4f14077.jpg")
	inner.Img1.Attr("width", 100)
	inner.Img1.Attr("height", 100)

	inner.Img1.Click(Functions.MyClickNew, MyClickNew_T{Val1: "testing", Span1: inner.Span1, Img1: inner.Img1, Text1: test.Text1})

	inner.MyLink.Link(Functions.Page1, Page1_T{Val1: "ooooooh!", Val2: 44})

	c.Send()

}

type Page1_T struct {
	Val1 twist.String
	Val2 twist.Int
}

func (f Functions) Page1(c *twist.Context, v Page1_T) {

	//c.Root.Html("Hello World." + v.Val1.Value() + " - " + v.Val2.String())

	page2 := twist.Page2(c, "foo")
	c.Root.Html(page2)

	page2.Button1.Link(Functions.Root, nil)

	c.Send()

}

type MyClickNew_T struct {
	Val1  twist.String
	Span1 *twist.Item
	Img1  *twist.Item
	Text1 *twist.Item
}

func (f Functions) MyClickNew(c *twist.Context, v MyClickNew_T) {

	v.Img1.Css("border", "10px solid #ff0000")
	v.Span1.Html("WHOOOOOOPPPPPPPPP!!!! ", v.Val1.String(), " ", v.Text1.Value)

	c.Send()
}
