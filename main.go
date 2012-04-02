package main

import "web.go"
import "libxml"
import "io/ioutil"

func serv(ctx *web.Context, message string) (resp string) {
	out, err := ioutil.ReadFile("test.htm")
	if err != nil {
		ctx.Server.Logger.Print("ERROR: " + err.Error())
		ctx.NotFound(err.Error())
		return ""
	}
	ctx.ContentType("html")

	doc := libxml.HtmlParseDoc(string(out))
	doc.AddChild("a", "a link")
	doc.AddChild("b", "a bold")
	doc.AddChild("i", "an italic")
	doc.AddChild("quote", "a quote")
	doc.AddChild("span", "a span")
	doc.AddChild("code", "a code")
	doc.AddChild("font", "a font")
	doc.AddChild("p", "a paragraph")
	doc.AddChild("div", "a div")
	resp = doc.String()
	doc.Free()

	return
}

func main() {
	web.Get("/(.*)", serv)
	web.Run("0.0.0.0:3000")
}
