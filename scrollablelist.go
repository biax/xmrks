package main

import (
	"fmt"

	ui "github.com/gizak/termui"
)

type ScrollableList struct {
	list   *ui.List
	title  string
	items  []string
	items2 []string
	sel    int
}

func CreateScrollableList(title string, w, h int) *ScrollableList {
	l := ScrollableList{list: ui.NewList(), sel: 0}
	l.title = title
	l.list.BorderLabel = title
	l.list.BorderLabelFg = ui.ColorWhite
	l.list.ItemFgColor = ui.ColorGreen
	l.list.Width = w
	l.list.Height = h
	return &l
}

func (l *ScrollableList) Select(t bool) {
	if t {
		l.list.BorderLabelFg = ui.ColorBlack
		l.list.BorderLabelBg = ui.ColorGreen
	} else {
		l.list.BorderLabelFg = ui.ColorWhite
		l.list.BorderLabelBg = ui.ColorBlack
	}
	ui.Render(l.list)
}

func (l *ScrollableList) Render() {
	h := l.list.Height - 3
	start := 0
	if len(l.items2) > h && l.sel > h {
		start = l.sel - h
	}
	l.items[l.sel] = fmt.Sprintf("[%s](fg-black,bg-green)", l.items2[l.sel])
	l.list.Items = l.items[start:]
	if len(l.items2) > h {
		dec := " "
		if l.sel > h {
			dec += "↑"
		}
		if l.sel != len(l.items2)-1 {
			dec += "↓"
		}
		l.list.BorderLabel = l.title + dec
	}
	ui.Render(l.list)
}

func (l *ScrollableList) SetItems(data []string) {
	l.items = make([]string, len(data))
	l.items2 = make([]string, len(l.items))
	copy(l.items, data)
	copy(l.items2, data)
	l.list.Items = l.items
}

func (l *ScrollableList) MoveUp() {
	if l.sel < 1 {
		return
	}
	l.sel--
	l.items[l.sel+1] = l.items2[l.sel+1]
	l.Render()
}

func (l *ScrollableList) MoveDown() {
	if l.sel == len(l.items)-1 {
		return
	}
	l.sel++
	if l.sel > 0 {
		l.items[l.sel-1] = l.items2[l.sel-1]
	}
	l.Render()
}
