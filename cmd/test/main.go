package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
)

func main() {

	builder, err := gtk.BuilderNewFromFile("/home/per/code/test/assets/main.glade")
	if err != nil {
		fmt.Println(err)
		return
	}
	windowObj, err := builder.GetObject("mainWindow")
	if err != nil {
		fmt.Println(err)
		return
	}
	window, ok := windowObj.(*gtk.ApplicationWindow)
	if !ok {
		fmt.Println("Failed to convert")
		return
	}
	fmt.Println(window.GetName())
}

