package src

import (
	"io/ioutil"
	"net/http"
	"os"
)

const webPath string = "webpage/"

type Context struct {
	req  *http.Request
	resw http.ResponseWriter
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		req:  r,
		resw: w,
	}
}

func (ctxt *Context) ResHtml(str string) {
	file, err := os.Open(webPath + str)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	ctxt.resw.Write(content)
}
