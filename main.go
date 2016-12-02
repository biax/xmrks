package main

import ui "github.com/gizak/termui"

//import "fmt"

func main() {
	err := ui.Init()
	check(err)
	defer ui.Close()

	ui.Handle("/sys/wnd/resize", refresh)
	ui.Handle("/sys/kbd/q", quit)
	ui.Handle("/sys/kbd/<enter>", enter)
	ui.Handle("/sys/kbd/<up>", up)
	ui.Handle("/sys/kbd/<down>", down)
	ui.Handle("/sys/kbd/<space>", refresh)
	runMenu()
}

func quit(e ui.Event) {
	if actions.stopping {
		return
	}
	ui.StopLoop()
	actions.stopping = true
}

func up(e ui.Event) {
	if actions.up != nil {
		actions.up()
		updateStatus()
	}
}

func down(e ui.Event) {
	if actions.down != nil {
		actions.down()
		updateStatus()
	}
}

var t = false

// Refresh only the status box
func updateStatus() {
	//	status.Text = fmt.Sprintf("Selected: %d:%d", list.sel+1, len(list.items))
	//ui.Render(status)
}

// Refresh entire UI (flickery)
func refresh(e ui.Event) {
	ui.Body.Width = ui.TermWidth()
	ui.Body.Align()
	ui.Render(ui.Body)
}

func enter(e ui.Event) {
	if actions.enter != nil {
		actions.enter()
	}
}
