package main

import ui "github.com/gizak/termui"

func runShowRecords() {
	ui.ClearArea(ui.TermRect(), ui.ColorBlack)
	ui.Body.Rows = make([]*ui.Row, 0)
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, title),
		),
	)
	shortcuts.Y = ui.TermHeight() - 1
	ui.Render(title, shortcuts)
	refresh(ui.Event{})
	ui.Loop()
}
