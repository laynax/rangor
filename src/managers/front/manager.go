package front

import (
	"github.com/jroimartin/gocui"
	"os/exec"
)

var CurrentPath = "/home/daniel/"

func Manager(g *gocui.Gui) error {
	x, y := g.Size()

	v, err := g.SetView("ls", (x / 5) + 2, 2, (x * 4 / 5) - 2, y - 2)
	if err != gocui.ErrUnknownView {
		return err
	}

	v.Wrap = true
	v.Highlight = true
	v.SelBgColor = gocui.ColorBlack
	v.SelFgColor = gocui.ColorWhite

	_, err = v.Write(getls())
	if err != nil {
		return err
	}

	return nil
}

func getls() []byte {
	cmd := exec.Command("ls", CurrentPath)
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return data
}