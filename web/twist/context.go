package twist

import (
	"http"
	"appengine"
)

type Context struct {
	*Writer
	Context        *appengine.Context
	Request        *http.Request
	Root           *Item
	itemsInRequest []*Item
}

func (c *Context) Navigate(handlerFunc interface{}, values interface{}) {

	href := c.Root.getLinkUrl(handlerFunc, values)

	c.Root.commands = append(c.Root.commands, func() { c.Root.navigateAtRender(href) })

}
func (i *Item) navigateAtRender(href string) {

	i.writer.Buffer += `
ignoreNextStateChange = true;
History.pushState(null, null, "` + href + `");`

}
