package controllers

import (
	"bytes"
	"html/template"
)

func ParseTemplate(templateFileName string, data interface{}) (content string, err error) {
	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
