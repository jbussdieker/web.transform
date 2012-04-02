package main

import "web.go"
import "libxml"
import "io/ioutil"

func serv(ctx *web.Context, message string) (resp string) {
	out, err := ioutil.ReadFile(message)
	if err != nil {
		ctx.Server.Logger.Print("ERROR: " + err.Error())
		ctx.NotFound(err.Error())
		return ""
	}
	ctx.ContentType("html")

	doc := libxml.XmlParseDoc(string(out))
	doc.AddChild("a", "link")
	buf := doc.Dump()
	resp = buf.String()
	buf.Free()
	doc.Free()

	return
}

func main() {
	web.Get("/(.*)", serv)
	web.Run("0.0.0.0:3000")
}
