package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"

)

func main() {
	p := bluemonday.NewPolicy()
	// p := bluemonday.UGCPolicy()

	// p := bluemonday.StrictPolicy()

	p.AllowElements("p", "br", "strong", "em", "h1", "h2", "h3", "h4", "h5", "h6", "blockquote", "hr")
	//p.AllowImages()
	p.AllowAttrs("style").OnElements("span")
	p.AllowAttrs("src").OnElements("video")
	p.AllowAttrs("controls").OnElements("video")
	p.AllowAttrs("controlsList").OnElements("video")
	p.AllowAttrs("src").OnElements("audio")
	p.AllowAttrs("controls").OnElements("audio")
	p.AllowAttrs("controlsList").OnElements("audio")
	p.AllowAttrs("src").OnElements("img")
	p.AllowAttrs("'")
	html := p.Sanitize(
		`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
	)

	// Output:
	// <a href="http://www.google.com">Google</a>
	fmt.Println(html)
}
