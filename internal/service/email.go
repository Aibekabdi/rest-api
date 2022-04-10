package service

import (
	"net/mail"
	"strings"
	"unicode"

	"github.com/Aibekabdi/rest-api/internal/models"
)

type EmailServise struct {
}

func NewEmailServise() *EmailServise {
	return &EmailServise{}
}

func (e *EmailServise) FindEmails(emails models.Email) models.Email {
	newEmails := models.Email{}
	for _, v := range emails {
		checkedEmail := findingFromTextEmail(v.Email)
		if checkedEmail != "" {
			v.Email = checkedEmail
			newEmails = append(newEmails, v)
		}
	}
	return newEmails
}

func findingFromTextEmail(str string) string {
	res := strings.TrimFunc(str, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	if err := isValidEmail(res); err != nil {
		return ""
	} else {
		return res
	}
}

func isValidEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}
	return nil
}
