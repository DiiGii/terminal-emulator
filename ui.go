package main

import (
	"os"
	"os/exec"
	"time"
	
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/creack/pty"
)

func main() {
	a := app.New()
	w := a.NewWindow("germ")

	ui := widget.NewTextGrid()
	ui.SetText("Terminal text!")
	
	c := exec.Command("/bin/bash")
	p, err := pty.Start(c)

	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	defer c.Process.Kill()

	p.Write([]byte("ls\r"))
	time.Sleep(1 * time.Second)
	b := make([]byte, 1024)
	_, err = p.Read(b)
	if err != nil {
		fyne.LogError("Failed to read pty", err)
	}

	ui.SetText(string(b))

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridWrapLayout(fyne.NewSize(420, 200)),
			ui,
		),
	)
	w.ShowAndRun()
}