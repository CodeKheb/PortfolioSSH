package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


type Contact struct {
	Name     string
	Email    string
	LinkedIn string
	GitHub   string
}

func (model Model) contactContent() string {
	content := []string{
		"Name",
		model.contactName.View(),

		"",
		"Email",
		model.contactEmail.View(),

		"",
		"Message",
		model.contactMessage.View(),

		"",
		"[ Send Message ]",
	}

	if model.contactError != "" {
		content = append(
			content,
			"",
			model.contactError,
		)
	}

	if model.contactSuccess {
		content = append(
			content,
			"",
			"Kheb appreciates the message.",
		)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		content...,
	)
}

func (model Model) ContactView() string {
	header := model.headerView("CONTACT")

	return model.layout(
		header,
		model.viewport.View(),
	)
}

func (model Model) submitMessage() Model {
	name := strings.TrimSpace(model.contactName.Value())
	email := strings.TrimSpace(model.contactEmail.Value())
	message := strings.TrimSpace(model.contactMessage.Value())

	if name == "" {
		model.contactError = "Name is required."
		return model
	}

	if email == "" {
		model.contactError = "Email is required."
		return model
	}

	if message == "" {
		model.contactError = "Message is required."
		return model
	}

	err := model.db.SaveMessage(
		name,
		email,
		message,
	)
	if err != nil {
		model.contactError = err.Error()
		return model
	}

	model.contactSuccess = true
	return model
}

func (model *Model) focusContactField() {
	model.contactName.Blur()
	model.contactEmail.Blur()
	model.contactMessage.Blur()

	switch model.contactField {
	case 0:
		model.contactName.Focus()
	case 1:
		model.contactEmail.Focus()
	case 2:
		model.contactMessage.Focus()
	}
}

func (model Model) contactKeyHandler(message tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch message.String() {
	case "esc":
		model.screen = MenuScreen
		model.contactName.Blur()
		model.contactEmail.Blur()
		model.contactMessage.Blur()
		return model, nil

	case "tab":
		model.contactField++

		if model.contactField > 3 {
			model.contactField = 0
		}

		model.focusContactField()
		return model, nil

	case "shift+tab":
		model.contactField--

		if model.contactField < 0 {
			model.contactField = 3
		}

		model.focusContactField()
		return model, nil

	case "enter":
		if model.contactField == 3 {
			model = model.submitMessage()
			model.viewport.SetContent(model.contactContent())
			return model, nil
		}
	}

	switch model.contactField {
	case 0:
		model.contactName, cmd = model.contactName.Update(message)

	case 1:
		model.contactEmail, cmd = model.contactEmail.Update(message)

	case 2:
		model.contactMessage, cmd = model.contactMessage.Update(message)
	}

	model.viewport.SetContent(model.contactContent())

	return model, cmd
}
