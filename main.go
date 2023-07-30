package main

import (
	// "github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Contact struct {
	firstName string
	lastName  string
	email     string
}

var contacts []Contact

var app = tview.NewApplication()

var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(a) to add a new contact \n(q) to quit")

var form = tview.NewForm()

var pages = tview.NewPages().
	AddPage("Menu", text, true, true).
	AddPage("Add Contact", form, true, false)

func main() {

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 97 {
			addContactForm()
			pages.SwitchToPage("Add contact")
		}
		return event
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func addContactForm() {
	contact := Contact{}

	form.AddInputField("First Name", "", 20, nil, func(firstName string) {
		contact.firstName = firstName
	})

	form.AddInputField("last Name", "", 20, nil, func(lastName string) {
		contact.lastName = lastName
	})

	form.AddInputField("email", "", 20, nil, func(email string) {
		contact.email = email
	})

	form.AddButton("Save", func() {
		contacts = append(contacts, contact)
		pages.SwitchToPage("Menu")
	})

}
