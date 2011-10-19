package handler

import (
	"http"
	"twist"
	"fmt"
	"strconv"
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
	f.Plain1(c)
}

func getNav(c *twist.Context) *twist.Navigation_T {
	nav := twist.Navigation(c, "Nav")
	nav.Plain1Link.Link(Functions.Plain1, nil)
	nav.Plain2Link.Link(Functions.Plain2, nil)
	nav.Plain3Link.Link(Functions.Plain3, nil)
	nav.Red1Link.Link(Functions.Red1, nil)
	nav.Red2Link.Link(Functions.Red2, nil)
	nav.Red3Link.Link(Functions.Red3, nil)
	return nav
}

func getPlainMaster(c *twist.Context) *twist.PlainMaster_T {
	master := twist.PlainMaster(c, "Master")
	master.Footer.Html("Here's some HTML in the footer...")
	master.Navigation.Html(getNav(c))
	return master
}

func getRedMaster(c *twist.Context) *twist.RedMaster_T {
	master := twist.RedMaster(c, "Master")
	master.Footer.Html("Here's some HTML in the red footer...")
	master.Navigation.Html(getNav(c))
	return master
}

func (f Functions) Plain1(c *twist.Context) {

	master := getPlainMaster(c)
	c.Root.Html(master)

	master.Header.Html("Plain 1 heading")

	p := twist.Plain1(c, "Plain1")
	master.Content.Html(p)

	p.Plus.Click(Functions.Plain1Add, Plain1Count_T{Count: p.Count, Output: p.Output})
	p.Minus.Click(Functions.Plain1Minus, Plain1Count_T{Count: p.Count, Output: p.Output})
	p.Count.Attr("value", "0")

	c.Send()
}

type Plain1Count_T struct {
	Count  *twist.Item
	Output *twist.Item
}

func (f Functions) Plain1Add(c *twist.Context, v Plain1Count_T) {
	i, _ := strconv.Atoi(v.Count.Value())
	i++
	v.Output.Html(fmt.Sprint(i))
	v.Count.Attr("value", fmt.Sprint(i))
	c.Send()
}
func (f Functions) Plain1Minus(c *twist.Context, v Plain1Count_T) {
	i, _ := strconv.Atoi(v.Count.Value())
	i--
	v.Output.Html(fmt.Sprint(i))
	v.Count.Attr("value", fmt.Sprint(i))
	c.Send()
}

func (f Functions) Plain2(c *twist.Context) {

	master := getPlainMaster(c)
	c.Root.Html(master)

	master.Header.Html("Plain 2 heading")

	contents := twist.Plain2(c, "Plain2")
	master.Content.Html(contents)

	c.Send()
}

func (f Functions) Plain3(c *twist.Context) {
	master := getPlainMaster(c)
	c.Root.Html(master)

	master.Header.Html("Plain 3 heading")

	contents := twist.Plain3(c, "Plain3")
	master.Content.Html(contents)

	c.Send()
}

func (f Functions) Red1(c *twist.Context) {
	master := getRedMaster(c)
	c.Root.Html(master)

	master.Location.Html("Red 1 location")
	master.Date.Html("red 1 date")

	master.Header.Html("Red 1 heading")

	contents := twist.Red1(c, "Red1")
	master.Content.Html(contents)

	c.Send()
}

func (f Functions) Red2(c *twist.Context) {
	master := getRedMaster(c)
	c.Root.Html(master)

	master.Location.Html("Red 2 location")
	master.Date.Html("red 2 date")

	master.Header.Html("Red 2 heading")

	contents := twist.Red2(c, "Red2")
	master.Content.Html(contents)

	c.Send()
}

func (f Functions) Red3(c *twist.Context) {
	master := getRedMaster(c)
	c.Root.Html(master)

	master.Location.Html("Red 3 location")
	master.Date.Html("red 3 date")

	master.Header.Html("Red 3 heading")

	contents := twist.Red3(c, "Red3")
	master.Content.Html(contents)

	c.Send()
}
