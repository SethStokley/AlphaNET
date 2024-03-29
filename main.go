package main

import (
	"github.com/trigex/alphanet/core"
	"github.com/trigex/alphanet/io"
	"github.com/trigex/alphanet/js"
)

func main() {
	// create fs
	fs := io.CreateFilesystem("fs.json")
	//fs.DebugPrint()

	console := io.CreateConsole()

	// create javascript vm
	VM := js.CreateVM(&console, &fs)

	// create computer, which holds these structs
	comp := core.CreateComputer(&fs, &VM)
	// start main loop
	comp.Start()
}
