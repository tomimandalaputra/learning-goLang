package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type EmailData struct {
	RecipientName string
	SenderName    string
	Subject       string
	Body          string
	Items         []string
	UnreadCount   int
}

func main() {
	fmt.Println("--- Text Template ---")

	emailTemplate := `
	Subject: {{ .Subject }}

	{{ .Body }}

	{{if .Items}}
		Related items:
		{{range .Items}}
		- {{.}}
		{{end}}
	{{end}}

	{{if gt .UnreadCount 0}}
		You have {{.UnreadCount}} unreads.
	{{else}}
		You have no messages
	{{end}}

	- Thanks
	{{.SenderName}}
	`

	tmpl, err := template.New("email-message").Parse(emailTemplate)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data := EmailData{
		RecipientName: "Tomi",
		SenderName:    "Putra",
		Subject:       "Your weekly Update",
		Body:          "Here is the update you requested. We hope you find it unseful",
		Items:         []string{"Report A", "Report N", "Report D"},
		UnreadCount:   1,
	}

	var output strings.Builder

	err = tmpl.Execute(&output, data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(output.String())
}
