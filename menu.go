package main

import (
	"strings"

	ui "github.com/gizak/termui"
)

var title = ui.NewPar("")
var summary = ui.NewPar("")
var shortcuts = ui.NewPar("")
var list *ScrollableList
var listlog *ScrollableList

var nextSection int

const (
	nowhere = iota + 1
	mainmenu
	showrecords
)

func runMenu() {
	for {
		ui.ClearArea(ui.TermRect(), ui.ColorBlack)
		if list == nil {
			list = CreateScrollableList(" Select An Option ", ui.TermWidth()/2, ui.TermHeight()-16)
			stuff := []string{
				" [A](fg-blue)dd",
				" [U](fg-blue)pdate",
				" [F](fg-blue)ind",
				" [B](fg-blue)rowse",
				" [V](fg-blue)isit Record Entry",
				" [E](fg-blue)dit Visit Record",
				" [P](fg-blue)rint Day's Visits",
				" [G](fg-blue)enerate Report Summary",
				" [Q](fg-blue)uit",
			}
			list.SetItems(stuff)
			list.sel = 0
			list.list.Bg = ui.ColorBlack // None blacker.
			list.Render()

			w := ui.TermWidth()

			title.Width = w
			title.Height = 1
			title.X = 0
			title.Y = 0
			title.Border = false
			title.Bg = ui.ColorWhite
			title.TextBgColor = ui.ColorWhite
			title.TextFgColor = ui.ColorBlack

			tt := "Dr. Y Klinik"
			s := strings.Repeat(" ", w/2-(len(tt)/2))
			s = s + tt
			s2 := "XMRKS v0.0.1"
			x := w - len(s) - len(s2)
			s = s + strings.Repeat(" ", x) + s2
			title.Text = s

			summary.Width = w
			summary.Height = 1
			summary.Border = false
			summary.Bg = ui.ColorRed
			summary.TextBgColor = ui.ColorRed
			summary.TextFgColor = ui.ColorWhite
			s = "Main Menu"
			x = w - len(s)
			s = strings.Repeat(" ", x/2) + s + strings.Repeat(" ", x/2)
			summary.Text = s

			listlog = CreateScrollableList(" Log ", ui.TermWidth()/2, ui.TermHeight()-24)

			shortcuts.Width = w
			shortcuts.Height = 1
			shortcuts.Border = false
			shortcuts.Bg = ui.ColorWhite
			shortcuts.TextBgColor = ui.ColorWhite
			shortcuts.TextFgColor = ui.ColorBlack
			s = "[Q](fg-blue)uit    [^/v](fg-blue) Navigate    [H](fg-blue)elp   [A](fg-blue)bout"
			x = w - 65
			s = strings.Repeat(" ", x/2) + s + strings.Repeat(" ", x/2)
			shortcuts.Text = s
		} else {
			list.sel = 0
			list.Render()
		}

		// Layout section
		ui.Body.Rows = nil
		ui.Body.AddRows(
			ui.NewRow(
				ui.NewCol(12, 0, title),
			),
			ui.NewRow(
				ui.NewCol(12, 0, summary),
			),
			ui.NewRow(
				ui.NewCol(12, 0, list.list),
			),
			ui.NewRow(
				ui.NewCol(12, 0, listlog.list),
			),
			ui.NewRow(
				ui.NewCol(12, 0, shortcuts),
			),
		)
		ui.Body.BgColor = ui.ColorBlack

		actions.stopping = false
		nextSection = nowhere
		actions.enter = mainEnter
		actions.up = list.MoveUp
		actions.down = list.MoveDown

		updateStatus()
		refresh(ui.Event{})
		ui.Loop()

		switch nextSection {
		case nowhere:
			return
		case showrecords:
			actions.enter = nil
			actions.up = nil
			actions.down = nil
			runShowRecords()
		}
	}
}

func mainEnter() {
	switch list.sel {
	case 0:
		nextSection = showrecords
		ui.StopLoop()
	case 1:
	// Edit record
	case 2:
		// New record
	}
}
