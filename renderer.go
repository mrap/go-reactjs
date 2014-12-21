package renderer

import (
	"fmt"
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

type renderer struct {
	*otto.Otto
}

type RendererDefaults []string

var SharedDefaults = RendererDefaults{
	"assets/global.js",
	"assets/react.js",
}

func (r *renderer) RenderComponent(name, props string) otto.Value {
	return r.RunCmd(fmt.Sprintf("React.renderToString(React.createElement(%s, %s));", name, props))
}

func NewRenderer(files []string) *renderer {
	r := &renderer{otto.New()}
	for _, f := range SharedDefaults {
		r.runFile(f)
	}
	r.runFiles(files)
	return r
}

func (r *renderer) RunCmd(cmd string) otto.Value {
	v, err := r.Run(cmd)
	if err != nil {
		panic(err)
	}
	return v
}

func (r *renderer) runFiles(files []string) {
	for _, file := range files {
		r.runFile(file)
	}
}

func (r *renderer) runFile(path string) otto.Value {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	result, err := r.Run(data)
	if err != nil {
		panic(err)
	}

	return result
}
