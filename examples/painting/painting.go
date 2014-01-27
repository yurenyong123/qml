package main

import (
	"fmt"
	"github.com/niemeyer/qml"
	"github.com/niemeyer/qml/gl"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

type GoRect struct {
	Objects []qml.Object
}

func (r *GoRect) Paint(p *qml.Painter) {
	obj := p.Object()

	width := gl.Float(obj.Int("width"))
	height := gl.Float(obj.Int("height"))

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA);
	gl.Color4f(1.0, 1.0, 1.0, 0.8)
	gl.Begin(gl.QUADS)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(width, 0)
	gl.Vertex2f(width, height)
	gl.Vertex2f(0, height)
	gl.End()

	gl.LineWidth(2.5)
	gl.Color4f(0.0, 0.0, 0.0, 1.0)
	gl.Begin(gl.LINES)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(width, height)
	gl.Vertex2f(width, 0)
	gl.Vertex2f(0, height)
	gl.End()
}

func run() error {
	qml.Init(nil)

	qml.RegisterTypes("GoExtensions", 1, 0, []qml.TypeSpec{{
		Name: "GoRect",
		New: func() interface{} { return &GoRect{} },
	}})

	engine := qml.NewEngine()
	component, err := engine.LoadFile("painting.qml")
	if err != nil {
		return err
	}

	value := component.CreateWindow(nil)
	value.Show()
	value.Wait()

	return nil
}
