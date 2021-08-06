package src

import (
	"io/ioutil"
	"net/http"
	"os"
)

const webPath string = "webpage/"

type Context struct {
	Req  *http.Request
	Resw http.ResponseWriter

	allStep []handlerfunc
	index   int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Req:   r,
		Resw:  w,
		index: -1,
	}
}

func (ctxt *Context) AddSteps(handlers ...handlerfunc) {
	ctxt.allStep = append(ctxt.allStep, handlers...)
}

func (ctxt *Context) NextStep() {
	ctxt.index++

	stepLen := len(ctxt.allStep)
	for ; ctxt.index < stepLen; ctxt.index++ {
		ctxt.allStep[ctxt.index](ctxt)
	}
}

func (ctxt *Context) ResHtml(str string) {
	file, err := os.Open(webPath + str)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	ctxt.Resw.Write(content)
}
