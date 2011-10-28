package handler

import (
	"http"
	"twist"
	"fmt"
	"html/content"
	"html/master"
	"html/misc"
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
	f.Plain1(c, Plain1_V{Start: 0})
}

func getNav(c *twist.Context) *misc.Navigation_T {
	nav := misc.Navigation(c, "Nav")
	nav.Plain1Link.Link(Functions.Plain1, nil)
	nav.Plain2Link.Link(Functions.Plain2, nil)
	nav.Plain3Link.Link(Functions.Plain3, nil)
	nav.Red1Link.Link(Functions.Red1, nil)
	nav.Red2Link.Link(Functions.Red2, nil)
	nav.Red3Link.Link(Functions.Red3, nil)
	return nav
}

func getPlainMaster(c *twist.Context) *master.PlainMaster_T {

	master := master.PlainMaster(c, "Master")
	master.Footer.Html("Here's some HTML in the footer...")
	master.Navigation.Html(getNav(c))
	return master
}

func getRedMaster(c *twist.Context) *master.RedMaster_T {
	master := master.RedMaster(c, "Master")
	master.Footer.Html("Here's some HTML in the red footer...")
	master.Navigation.Html(getNav(c))
	return master
}

type Plain1_V struct {
	Start twist.Int
}

func (f Functions) Plain1(c *twist.Context, v Plain1_V) {

	master := getPlainMaster(c)
	c.Root.Html(master)

	master.Header.Html("Plain 1 heading")

	p := content.Plain1(c, "Plain1")
	master.Content.Html(p)

	p.Plus.Click(Functions.Plain1Add, Plain1Count_V{Count: p.Count, Output: p.Output})
	p.Minus.Click(Functions.Plain1Minus, Plain1Count_V{Count: p.Count, Output: p.Output})
	p.Output.Html(v.Start.String())
	p.Count.Attr("value", v.Start.String())

}

type Plain1Count_V struct {
	Count  *twist.Item
	Output *twist.Item
}

func (f Functions) Plain1Add(c *twist.Context, v Plain1Count_V) {
	i := v.Count.Int()
	i++
	v.Output.Html(fmt.Sprint(i))
	v.Count.Attr("value", fmt.Sprint(i))
	c.Navigate(Functions.Plain1, Plain1_V{Start: twist.Int(i)})
}

func (f Functions) Plain1Minus(c *twist.Context, v Plain1Count_V) {
	i := v.Count.Int()
	i--
	v.Output.Html(fmt.Sprint(i))
	v.Count.Attr("value", fmt.Sprint(i))
	c.Navigate(Functions.Plain1, Plain1_V{Start: twist.Int(i)})
}

func (f Functions) Plain2(c *twist.Context) {

	master := getPlainMaster(c)
	c.Root.Html(master)

	master.Header.Html("Plain 2 heading")

	contents := content.Plain2(c, "Plain2")
	master.Content.Html(contents)

}

func (f Functions) Plain3(c *twist.Context) {
	master := getPlainMaster(c)
	c.Root.Html(master)

	master.Header.Html("Plain 3 heading")

	contents := content.Plain3(c, "Plain3")
	master.Content.Html(contents)

}

func (f Functions) Red1(c *twist.Context) {
	master := getRedMaster(c)
	c.Root.Html(master)

	master.Location.Html("Red 1 location")
	master.Date.Html("red 1 date")

	master.Header.Html("Red 1 heading")

	contents := content.Red1(c, "Red1")
	master.Content.Html(contents)

}

func (f Functions) Red2(c *twist.Context) {
	master := getRedMaster(c)
	c.Root.Html(master)

	master.Location.Html("Red 2 location")
	master.Date.Html("red 2 date")

	master.Header.Html("Red 2 heading")

	contents := content.Red2(c, "Red2")
	master.Content.Html(contents)

}

func (f Functions) Red3(c *twist.Context) {
	master := getRedMaster(c)
	c.Root.Html(master)

	master.Location.Html("Red 3 location")
	master.Date.Html("red 3 date")

	master.Header.Html("Red 3 heading")

	contents := content.Red3(c, "Red3")
	master.Content.Html(contents)

}
