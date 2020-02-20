package widget

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type OneShotEntry struct {
	widget.Entry
	OnSubmitted func(text string) bool
}

func NewOnShotEntry(handler func(string) bool) *OneShotEntry {
	e := &OneShotEntry{OnSubmitted: handler}
	e.ExtendBaseWidget(e)
	return e
}

func (e *OneShotEntry) TypedKey(key *fyne.KeyEvent) {
	if key.Name == fyne.KeyReturn || key.Name == fyne.KeyEnter {
		if e.OnSubmitted(e.Text) {
			e.SetText("")
		}
	} else {
		e.Entry.TypedKey(key)
	}
}
