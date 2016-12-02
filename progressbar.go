package main

import (
	"fmt"
	"strings"

	ui "github.com/gizak/termui"
)

type ProgressBar struct {
	X       int
	Y       int
	Width   int
	text    string
	blocks  int
	BG      string
	FG      string
	TextCol ui.Attribute
	par     *ui.Par
}

func CreateProgressBar(x, y int, text string) *ProgressBar {
	bar := ProgressBar{
		X:       x,
		Y:       y,
		Width:   ui.TermWidth() - len(text),
		text:    text,
		TextCol: ui.ColorWhite,
		BG:      "░",
		FG:      "█",
	}
	bar.blocks = bar.Width - len(text) - 6
	bar.par = ui.NewPar("")
	bar.par.Border = false
	return &bar
}

func (bar *ProgressBar) Render() {
	bar.par.X = bar.X
	bar.par.Y = bar.Y
	bar.par.Width = bar.Width
	bar.par.TextFgColor = bar.TextCol
	ui.Render(bar.par)
}

func (bar *ProgressBar) Update(progress int) {
	s1 := strings.Repeat(bar.FG, progress*bar.blocks/100)
	s2 := strings.Repeat(bar.BG, bar.blocks-(progress*bar.blocks/100))
	bar.par.Text = fmt.Sprintf("%s[|](fg-yellow)[%s](fg-green,bg-black)%s[|](fg-yellow)%d%%", bar.text, s1, s2, progress)
	bar.Render()
}
