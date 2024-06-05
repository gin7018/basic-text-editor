package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	fmt.Println("starting text editor gui")
	application := app.New()
	window := application.NewWindow("basic text editor")

	text := canvas.NewText("type here", color.White)
	//str := binding.BindString(&text)
	text.Alignment = fyne.TextAlignLeading
	text.TextStyle = fyne.TextStyle{
		Bold:      false,
		Italic:    false,
		Monospace: false,
		Symbol:    false,
		TabWidth:  0,
	}
	w := widget.NewTextGridFromString("type here")
	window.SetContent(w)

	window.Resize(fyne.NewSize(500, 500))
	window.ShowAndRun()
}

//one way binding with key event capturing
